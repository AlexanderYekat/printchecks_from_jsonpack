package mercuriy

import (
	"bytes"
	consttypes "clientrabbit/consttypes"
	logsmy "clientrabbit/packetlog"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

var testnomsessii int

func GetSNOByDefault(emulation bool, ipktt string, port int, sessionkey string) (int, error) {
	var resMerc consttypes.TAnswerMercur
	jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"GetRegistrationInfo\"}", sessionkey))
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		return -1, err
	}
	err = json.Unmarshal(buffAnsw, &resMerc)
	if err != nil {
		logsmy.Logsmap[consttypes.LOGERROR].Printf("ошибка (%v) маршалинга результата запроса о резултататах регистрации кассы меркурий\n", err)
		return -1, err
	}
	if resMerc.Result != 0 {
		logsmy.Logsmap[consttypes.LOGERROR].Printf("ошибка (%v) запроса о резултататах регистрации кассы меркурий\n", resMerc.Description)
		err = fmt.Errorf(resMerc.Description)
		if !emulation {
			return -1, err
		} else {
			resMerc.RegistrationInfo = new(consttypes.TMercRegistrationInfo)
			resMerc.RegistrationInfo.TaxSystem = []int{5}
		}
	}
	if len(resMerc.RegistrationInfo.TaxSystem) != 1 {
		err := errors.New("касса зарегистрирована на больше чем одна система налогообложение")
		logsmy.Logsmap[consttypes.LOGERROR].Printf("касса зарегистрирована на больше чем одна система налогообложение")
		return -1, err
	}
	return resMerc.RegistrationInfo.TaxSystem[0], nil
}

func PrintCheck(emulation bool, ipktt string, port int, comport int, checkatol consttypes.TCorrectionCheck, sessionkey string, snoDefault int, dontprintrealfortest bool, userint int, passwuser string, emulatmistakesOpenCheck bool) (string, error) {
	var resMerc, resMercCancel consttypes.TAnswerMercur
	var answer []byte
	var answerclosecheck []byte
	var errclosecheck, errOfOpenCheck error
	if sessionkey == "" {
		answer, err := opensession(ipktt, port, comport, userint, passwuser)
		if err != nil {
			descrError := "ошибка открытия сессии к ккт меркурий"
			err = errors.Join(err, errors.New(descrError))
			return descrError, err
		}
		err = json.Unmarshal(answer, &resMerc)
		if err != nil {
			descrError := "ошибка при разобре ответа при отрытии сессии покдлючения к ККТ меркурий"
			err = errors.Join(err, errors.New(descrError))
			return descrError, err
		}
		if resMerc.Result != 0 || resMerc.SessionKey == "" {
			descrError := "ошибка при подключении к ккт меркурий"
			err = fmt.Errorf(resMerc.Description)
			err = errors.Join(err, errors.New(descrError))
			if !emulation {
				return descrError, err
			} else {
				testnomsessii = testnomsessii + 1
				resMerc.SessionKey = "эмуляция" + strconv.Itoa(testnomsessii)
			}
		}
		sessionkey := resMerc.SessionKey
		defer Closesession(ipktt, port, &sessionkey)
	}
	//logsmy.LogginInFile(fmt.Sprintln()
	checheaderkmerc, err := convertAtolToMercHeader(checkatol, snoDefault)
	checheaderkmerc.SessionKey = sessionkey
	if err != nil {
		descrError := "ошибка конвертации шапки чека атол в шапку чека меркурия"
		err = errors.Join(err, errors.New(descrError))
		return descrError, err
	}
	headercheckmerc, err := json.Marshal(checheaderkmerc)
	if err != nil {
		descrError := fmt.Sprintf("ошибка формирования шапки чека для кассы меркурий из структуры (%v)", checheaderkmerc)
		err = errors.Join(err, errors.New(descrError))
		return descrError, err
	}
	answer, err = opencheck(ipktt, port, headercheckmerc)
	if err != nil {
		descrError := "ошибка открытия чека для кассы меркурий"
		err = errors.Join(err, errors.New(descrError))
		return descrError, err
	}
	err = json.Unmarshal(answer, &resMerc)
	if err != nil {
		descrError := "ошибка разбора ответа при открытии чека для кассы меркурий"
		err = errors.Join(err, errors.New(descrError))
		return descrError, err
	}
	if resMerc.Result != 0 { //если не получилось открыть чек, отменяем его и пробуем отрыть заново
		descrError := fmt.Sprintf("ошибка (%v) открытия чека для кассы меркурий (попытка 1)", resMerc.Description)
		errOfOpenCheck = fmt.Errorf(descrError)
		logsmy.LogginInFile(fmt.Sprintf("ошибка (%v) 1-ой попытки открытия чека. Пробуемотменить чек и открыть заново \n", errOfOpenCheck))
		logsmy.LogginInFile("отменяем предыдущий чек \n")
		answerCancel, errCancel := cancelcheck(ipktt, port, &sessionkey) //отменяем предыдущий чек
		if errCancel != nil {
			logsmy.LogginInFile(fmt.Sprintf("ошибка (%v) отмены предыдущего чека \n", errCancel))
			errOfOpenCheck = errors.Join(errOfOpenCheck, errCancel)
		} else {
			errUnMarshCancel := json.Unmarshal(answerCancel, &resMercCancel)
			if errUnMarshCancel != nil {
				logsmy.LogginInFile(fmt.Sprintf("ошибка (%v) рапзборап ответа отмены предыдущего чека \n", errUnMarshCancel))
				errOfOpenCheck = errors.Join(errOfOpenCheck, errUnMarshCancel)
			} else {
				logsmy.LogginInFile(fmt.Sprintf("результат (%v) отмены предыдущего чека \n", resMercCancel.Description))
				if resMercCancel.Result != 0 {
					descrError := fmt.Sprintf("ошибка (%v) отмены чека для кассы меркурий", resMercCancel.Description)
					errOfOpenCheck = errors.Join(errOfOpenCheck, fmt.Errorf(descrError))
				} else {
					answer, err = opencheck(ipktt, port, headercheckmerc) //открываем заново чек
					if err != nil {
						err = json.Unmarshal(answer, &resMerc) //разбираем ответ
						if err != nil {
							logsmy.LogginInFile(fmt.Sprintf("ошибка (%v) разбора ответа отмены чека\n", err))
						}
					}
				}
			}
		}
	}
	if resMerc.Result != 0 { //если не получилось открыть чек
		descrError := fmt.Sprintf("ошибка (%v) открытия чека для кассы меркурий", resMerc.Description)
		err = errors.Join(errOfOpenCheck, errors.New(descrError))
		if !emulation {
			return descrError, err
		}
	}
	for _, pos := range checkatol.Items {
		var currPosType consttypes.TGenearaPosAndTag11921191
		mapstructure.Decode(pos, &currPosType)
		if currPosType.Type != "position" {
			continue
		}
		var currPos consttypes.TPosition
		mapstructure.Decode(pos, &currPos)
		mercPos, err := convertAtolPosToMercPos(currPos)
		mercPos.SessionKey = sessionkey
		if err != nil {
			descrError := fmt.Sprintf("ошибка формирования структуры позиции для кассы меркурий из позиции json-задания (%v)", pos.(consttypes.TPosition))
			err = errors.Join(err, errors.New(descrError))
			return descrError, err
		}
		mercPosJsonBytes, err := json.Marshal(mercPos)
		if err != nil {
			descrError := fmt.Sprintf("ошибка маршалинга структуры позиции для кассы меркурий из (%v)", mercPos)
			err = errors.Join(err, errors.New(descrError))
			return descrError, err
		}
		answer, err = addpos(ipktt, port, mercPosJsonBytes)
		if err != nil {
			descrError := fmt.Sprintf("ошибка добавления позиции %v в чек для кассы меркурий", mercPosJsonBytes)
			err = errors.Join(err, errors.New(descrError))
			if !emulation {
				return descrError, err
			}
		}
		err = json.Unmarshal(answer, &resMerc)
		if err != nil {
			descrError := fmt.Sprintf("ошибка маршалинга результата %v добавления позиции в чек для кассы меркурий", resMerc)
			err = errors.Join(err, errors.New(descrError))
			return descrError, err
		}
		if resMerc.Result != 0 {
			descrError := fmt.Sprintf("ошибка добавления позиции %v в чек для кассы меркурий", mercPosJsonBytes)
			err = fmt.Errorf(resMerc.Description)
			err = errors.Join(err, errors.New(descrError))
			if !emulation {
				return descrError, err
			}
		}
	}
	checkclosekmerc := convertAtolToMercCloseCheck(checkatol)
	checkclosekmerc.SessionKey = sessionkey
	checkclosekmercbytes, err := json.Marshal(checkclosekmerc)
	if err != nil {
		descrError := "ошибка формирования данных для закрытия чек кассы меркурий"
		err = errors.Join(err, errors.New(descrError))
		return descrError, err
	}
	if !dontprintrealfortest {
		answerclosecheck, errclosecheck = closecheck(ipktt, port, checkclosekmercbytes)
	} else {
		answerclosecheck, errclosecheck = cancelcheck(ipktt, port, &sessionkey)
	}
	if errclosecheck != nil {
		descrError := "ошибка закрытия чека для кассы меркурий"
		err = errors.Join(err, errors.New(descrError))
		return descrError, errclosecheck
	}
	errclosecheck = json.Unmarshal(answerclosecheck, &resMerc)
	if errclosecheck != nil {
		descrError := "ошибка разбора резульата закрытия чека для кассы меркурий"
		err = errors.Join(err, errors.New(descrError))
		return descrError, errclosecheck
	}
	if resMerc.Result != 0 {
		descrError := "ошибка закрытия чека для кассы меркурий"
		errclosecheck = fmt.Errorf(resMerc.Description)
		err = errors.Join(err, errors.New(descrError))
		if !emulation {
			return descrError, errclosecheck
		} else {
			errclosecheck = nil
		}
	}
	return string(answerclosecheck), errclosecheck
} //PrintCheck

func CheckStatsuConnectionKKT(emulation bool, ipktt string, port int, comport int, sessionkey string, userint int, passwuser string) (string, string, error) {
	var resMerc consttypes.TAnswerMercur
	answerbytesserver, errStatusServer := getStatusServerKKT(ipktt, port)
	if errStatusServer != nil {
		descrError := "ошибка получения статуса сервера ккт меркурий"
		return "", descrError, errStatusServer
	}
	errUnmarshServer := json.Unmarshal(answerbytesserver, &resMerc)
	if errUnmarshServer != nil {
		descrError := fmt.Sprintf("ошибка распаковки ответа %v сервера ккт меркурий", string(answerbytesserver))
		return "", descrError, errUnmarshServer
	}
	if resMerc.Result != 0 {
		descrError := fmt.Sprintf("сервер ККТ меркурий не работает по причине %v", resMerc.Description)
		err := errors.New(resMerc.Description)
		return "", descrError, err
	}
	if sessionkey == "" {
		answer, err := opensession(ipktt, port, comport, userint, passwuser)
		if err != nil {
			descrError := "ошибка при подключении к ккт меркурий"
			return "", descrError, err
		}
		err = json.Unmarshal(answer, &resMerc)
		if err != nil {
			descrError := fmt.Sprintf("ошибка при разборе ответа %v от ккт меркурий", answer)
			return "", descrError, err
		}

		if resMerc.Result != 0 || resMerc.SessionKey == "" {
			descrError := "ошибка при подключении к ккт меркурий"
			err = fmt.Errorf(resMerc.Description)
			if !emulation {
				return "", descrError, err
			} else {
				testnomsessii = testnomsessii + 1
				resMerc.SessionKey = "эмуляция" + strconv.Itoa(testnomsessii)
			}
		}
		sessionkey = resMerc.SessionKey
		//defer closesession(ipktt, port, sessionkey, loginfo)
	}
	answerbyteKKT, errStatusKKT := getStatusKKT(ipktt, port, sessionkey)
	if errStatusKKT != nil {
		descrError := "ошибка получения статуса ккт меркурий"
		Closesession(ipktt, port, &sessionkey)
		return "", descrError, errStatusKKT
	}
	errUnmarshKKT := json.Unmarshal(answerbyteKKT, &resMerc)
	if errUnmarshKKT != nil {
		descrError := fmt.Sprintf("ошибка распаковки ответа %v ккт меркурий", string(answerbyteKKT))
		Closesession(ipktt, port, &sessionkey)
		return "", descrError, errUnmarshKKT
	}
	if resMerc.Result != 0 {
		descrError := fmt.Sprintf("ккт меркурий не работает по причине %v", resMerc.Description)
		if !emulation {
			Closesession(ipktt, port, &sessionkey)
			err := errors.New(resMerc.Description)
			return "", descrError, err
		}
	}
	return sessionkey, "", nil
} //CheckStatsuConnectionKKT

func DissconnectMeruriy(ipktt string, port int, sessionkey string) (string, error) {
	var resMerc consttypes.TAnswerMercur
	if sessionkey != "" {
		Closesession(ipktt, port, &sessionkey)
	}
	jsonmerc := []byte("{\"command\":\"ClosePorts\"}")
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) закрытия всех не активных портов для меркурия", err)
		return descrError, err
	}
	err = json.Unmarshal(buffAnsw, &resMerc)
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) маршалинга результата закрытия закрытия всех не активных портов для меркурия", err)
		return descrError, err
	}
	if resMerc.Result != 0 {
		descrError := fmt.Sprintf("ошибка (%v) закрытия всех не активных портов для меркурий", resMerc.Description)
		err = fmt.Errorf(resMerc.Description)
		return descrError, err
	}
	return "", nil
}

func BreakAndClearProccessOfMarks(ipktt string, port int, comport int, sessionkey string, userint int, passwuser string) (string, error) {
	desckErrorBreak, errBreek := BreakProcCheckOfMark(ipktt, port, comport, sessionkey, userint, passwuser)
	desckErrorBreakClear, errClear := ClearTablesOfMarks(ipktt, port, comport, sessionkey, userint, passwuser)
	err := errors.Join(errBreek, errClear)
	return desckErrorBreak + desckErrorBreakClear, err
}

func BreakProcCheckOfMark(ipktt string, port int, comport int, sessionkey string, userint int, passwuser string) (string, error) {
	var resMerc consttypes.TAnswerMercur
	if sessionkey == "" {
		answer, err := opensession(ipktt, port, comport, userint, passwuser)
		if err != nil {
			descrError := "ошибка при подключении к ккт меркурий"
			return descrError, err
		}
		err = json.Unmarshal(answer, &resMerc)
		if err != nil {
			descrError := "ошибка при подключении к ккт меркурий"
			return descrError, err
		}
		if resMerc.Result != 0 || resMerc.SessionKey == "" {
			descrError := "ошибка при подключении к ккт меркурий"
			err = fmt.Errorf(resMerc.Description)
			return descrError, err
		}
		sessionkey = resMerc.SessionKey
		defer Closesession(ipktt, port, &sessionkey)
	}
	jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"AbortMarkingCodeChecking\"}", sessionkey))
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		descrError := "ошибка прерывания проверки марок"
		return descrError, err
	}
	err = json.Unmarshal(buffAnsw, &resMerc)
	if err != nil {
		descrError := "ошибка прерывания проверки марок"
		return descrError, err
	}
	descrError := resMerc.Description
	return descrError, nil
} //breakProcCheckOfMark

func ClearTablesOfMarks(ipktt string, port int, comport int, sessionkey string, userint int, passwuser string) (string, error) {
	var resMerc consttypes.TAnswerMercur
	if sessionkey == "" {
		answer, err := opensession(ipktt, port, comport, userint, passwuser)
		if err != nil {
			descrError := "ошибка при подключении к ккт меркурий"
			return descrError, err
		}
		err = json.Unmarshal(answer, &resMerc)
		if err != nil {
			descrError := "ошибка при подключении к ккт меркурий"
			return descrError, err
		}
		if resMerc.Result != 0 || resMerc.SessionKey == "" {
			descrError := "ошибка при подключении к ккт меркурий"
			err = fmt.Errorf(resMerc.Description)
			return descrError, err
		}
		sessionkey = resMerc.SessionKey
		defer Closesession(ipktt, port, &sessionkey)
	}
	jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"ClearMarkingCodeValidationTable\"}", sessionkey))
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		descrError := "ошибка очистки таблицы марок"
		return descrError, err
	}
	err = json.Unmarshal(buffAnsw, &resMerc)
	if err != nil {
		descrError := "ошибка очистки таблицы марок"
		return descrError, err
	}
	descrError := resMerc.Description
	return descrError, nil
} //BreakProccessMarkAndClearTablesOfMarks

// ////////////////////
func getStatusServerKKT(ipktt string, port int) ([]byte, error) {
	jsonmerc := []byte("{\"command\":\"GetDriverInfo\"}")
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		return nil, err
	}
	return buffAnsw, nil
} //getStatusServerKKT

/*func getInfoKKT(ipktt string, port int, sessionkey string) ([]byte, error) {
	jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"GetCommonInfo\"}", sessionkey))
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		return nil, err
	}
	return buffAnsw, nil
} //getStatusKKT*/

func getJSONBeginProcessMarkCheck(mark string, measureunit int, sessionkey string) ([]byte, error) {
	jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"CheckMarkingCode\", \"mc\":\"%v\", \"plannedStatus\": 255, \"qty\": 10000, \"measureUnit\": %v}", sessionkey, mark, measureunit))
	return jsonmerc, nil
}

func SendCheckOfMark(ipktt string, port int, sessionkey, mark string, measureunit int) ([]byte, error) {
	jsonBeginProcMark, err := getJSONBeginProcessMarkCheck(mark, measureunit, sessionkey)
	if err != nil {
		return nil, err
	}
	return sendCommandTCPMerc(jsonBeginProcMark, ipktt, port)
}

func GetStatusOfChecking(ipktt string, port int, sessionkey string) ([]byte, error) {
	jsonOfStatusProcMark := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"GetMarkingCodeCheckResult\"}", sessionkey))
	return sendCommandTCPMerc(jsonOfStatusProcMark, ipktt, port)
}

func AcceptMark(ipktt string, port int, sessionkey string) ([]byte, error) {
	jsonOfStatusProcMark := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"AcceptMarkingCode\"}", sessionkey))
	return sendCommandTCPMerc(jsonOfStatusProcMark, ipktt, port)
}

func getStatusKKT(ipktt string, port int, sessionkey string) ([]byte, error) {
	jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"GetStatus\"}", sessionkey))
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		return nil, err
	}
	return buffAnsw, nil
} //getStatusKKT

func convertAtolToMercHeader(checkatol consttypes.TCorrectionCheck, snoDefault int) (consttypes.TMercOpenCheck, error) {
	var checheaderkmerc consttypes.TMercOpenCheck
	checheaderkmerc.Command = "OpenCheck"
	checheaderkmerc.CheckType = 4
	if checkatol.Type == "buyCorrection" {
		checheaderkmerc.CheckType = 5
	} else if checkatol.Type == "sellReturnCorrection" {
		checheaderkmerc.CheckType = 6
	} else if checkatol.Type == "buyReturnCorrection" {
		checheaderkmerc.CheckType = 7
	}
	if checkatol.TaxationType == "" {
		if snoDefault == -1 {
			err := fmt.Errorf("не задан тип налогообложения")
			return checheaderkmerc, err
		}
	}
	checheaderkmerc.TaxSystem = snoDefault
	if checkatol.TaxationType == "osn" {
		checheaderkmerc.TaxSystem = 0
	} else if checkatol.TaxationType == "usnIncome" {
		checheaderkmerc.TaxSystem = 1
	} else if checkatol.TaxationType == "usnIncomeOutcome" {
		checheaderkmerc.TaxSystem = 2
	} else if checkatol.TaxationType == "esn" {
		checheaderkmerc.TaxSystem = 4
	} else if checkatol.TaxationType == "patent" {
		checheaderkmerc.TaxSystem = 5
	}
	checheaderkmerc.PrintDoc = !checkatol.Electronically
	for _, posItems := range checkatol.Items {
		if posItems.(map[string]interface{})["type"].(string) == "additionalAttribute" {
			//if posItems.(consttypes.TGenearaPosAndTag11921191).Type == "additionalAttribute" {
			//checheaderkmerc.AdditionalProps = posItems.(consttypes.TTag1192_91).Value
			checheaderkmerc.AdditionalProps = posItems.(map[string]interface{})["value"].(string)
		}
		//if posItems.(consttypes.TGenearaPosAndTag11921191).Type == "userAttribute" {
		if posItems.(map[string]interface{})["type"].(string) == "userAttribute" {
			//checheaderkmerc.UserAttribute.AttrName = posItems.(consttypes.TTag1192_91).Name
			//checheaderkmerc.UserAttribute.AttrValue = posItems.(consttypes.TTag1192_91).Value
			checheaderkmerc.UserAttribute.AttrName = posItems.(map[string]interface{})["name"].(string)
			checheaderkmerc.UserAttribute.AttrValue = posItems.(map[string]interface{})["value"].(string)
		}
	}
	if checheaderkmerc.AdditionalProps == "" {
		err := fmt.Errorf("не задан ФП чека основания для чека коррекции")
		return checheaderkmerc, err
	}
	checheaderkmerc.CashierInfo.CashierName = checkatol.Operator.Name
	checheaderkmerc.CashierInfo.CashierINN = checkatol.Operator.Vatin
	if checkatol.ClientInfo != nil {
		if checkatol.ClientInfo.Vatin != "" {
			checheaderkmerc.BuyerInfo = new(consttypes.TBuyerInfoMerc)
			checheaderkmerc.BuyerInfo.BuyerName = checkatol.ClientInfo.Name
			checheaderkmerc.BuyerInfo.BuyerINN = checkatol.ClientInfo.Vatin
		}
	}
	checheaderkmerc.CorrectionInfo.CorrectionType = 0
	if checkatol.CorrectionType == "instruction" {
		checheaderkmerc.CorrectionInfo.CorrectionType = 1
		checheaderkmerc.CorrectionInfo.CauseDocNum = checkatol.CorrectionBaseNumber
	}
	checheaderkmerc.CorrectionInfo.CauseDocDate = strings.ReplaceAll(checkatol.CorrectionBaseDate, ".", "-")
	return checheaderkmerc, nil
} //convertAtolToMercHeader

func convertMeasureUnit(measureUnit string) int {
	resMeasureCode := 0
	if measureUnit == "gram" {
		resMeasureCode = 10
	}
	if measureUnit == "kilogram" {
		resMeasureCode = 11
	}
	if measureUnit == "ton" {
		resMeasureCode = 12
	}
	if measureUnit == "centimeter" {
		resMeasureCode = 20
	}
	if measureUnit == "decimeter" {
		resMeasureCode = 21
	}
	if measureUnit == "meter" {
		resMeasureCode = 22
	}
	if measureUnit == "squareCentimeter" {
		resMeasureCode = 30
	}
	if measureUnit == "squareDecimeter" {
		resMeasureCode = 31
	}
	if measureUnit == "squareMeter" {
		resMeasureCode = 32
	}
	if measureUnit == "milliliter" {
		resMeasureCode = 40
	}
	if measureUnit == "liter" {
		resMeasureCode = 41
	}
	if measureUnit == "cubicMeter" {
		resMeasureCode = 42
	}
	if measureUnit == "kilowattHour" {
		resMeasureCode = 50
	}
	if measureUnit == "gkal" {
		resMeasureCode = 51
	}
	if measureUnit == "day" {
		resMeasureCode = 70
	}
	if measureUnit == "hour" {
		resMeasureCode = 71
	}
	if measureUnit == "minute" {
		resMeasureCode = 72
	}
	if measureUnit == "second" {
		resMeasureCode = 73
	}
	if measureUnit == "kilobyte" {
		resMeasureCode = 80
	}
	if measureUnit == "megabyte" {
		resMeasureCode = 81
	}
	if measureUnit == "gigabyte" {
		resMeasureCode = 82
	}
	if measureUnit == "terabyte" {
		resMeasureCode = 83
	}
	if measureUnit == "otherUnits" {
		resMeasureCode = 255
	}
	return resMeasureCode
} //convertMeasureUnit

func convertTaxNDSCode(taxType string) int {
	resTaxNDS := 6
	if taxType == "vat0" {
		resTaxNDS = 5
	}
	if taxType == "vat10" {
		resTaxNDS = 2
	}
	if taxType == "vat20" {
		resTaxNDS = 1
	}
	if taxType == "vat110" {
		resTaxNDS = 4
	}
	if taxType == "vat120" {
		resTaxNDS = 3
	}
	return resTaxNDS
} //convertTaxNDSCode

func convertSposRasch(sposRash string) int {
	resSposRash := 4
	if sposRash == "fullPrepayment" {
		resSposRash = 1
	}
	if sposRash == "prepayment" {
		resSposRash = 2
	}
	if sposRash == "advance" {
		resSposRash = 3
	}
	if sposRash == "partialPayment" {
		resSposRash = 5
	}
	if sposRash == "credit" {
		resSposRash = 6
	}
	if sposRash == "creditPayment" {
		resSposRash = 7
	}
	return resSposRash
} //convertSposRasch

func convertPredmRash(predmRash string) int {
	resPredmRash := 1 //товар
	if predmRash == "excise" {
		resPredmRash = 2
	}
	if predmRash == "job" {
		resPredmRash = 3
	}
	if predmRash == "service" {
		resPredmRash = 4
	}
	if predmRash == "gamblingBet" {
		resPredmRash = 5
	}
	if predmRash == "gamblingPrize" {
		resPredmRash = 5
	}
	if predmRash == "lottery" {
		resPredmRash = 7
	}
	if predmRash == "lotteryPrize" {
		resPredmRash = 8
	}
	if predmRash == "intellectualActivity" {
		resPredmRash = 9
	}
	if predmRash == "payment" {
		resPredmRash = 10
	}
	if predmRash == "agentCommission" {
		resPredmRash = 11
	}
	if predmRash == "composite" {
		resPredmRash = 12
	}
	if predmRash == "another" {
		resPredmRash = 13
	}
	if predmRash == "proprietaryLaw" {
		resPredmRash = 14
	}
	if predmRash == "nonOperatingIncome" {
		resPredmRash = 15
	}
	if predmRash == "insuranceContributions" {
		resPredmRash = 16
	}
	if predmRash == "otherContributions" {
		resPredmRash = 16
	}
	if predmRash == "merchantTax" {
		resPredmRash = 17
	}
	if predmRash == "resortFee" {
		resPredmRash = 18
	}
	if predmRash == "deposit" {
		resPredmRash = 19
	}
	if predmRash == "consumption" {
		resPredmRash = 20
	}
	if predmRash == "soleProprietorCPIContributions" {
		resPredmRash = 21
	}
	if predmRash == "cpiContributions" {
		resPredmRash = 22
	}
	if predmRash == "soleProprietorCMIContributions" {
		resPredmRash = 23
	}
	if predmRash == "cmiContributions" {
		resPredmRash = 24
	}
	if predmRash == "csiContributions" {
		resPredmRash = 25
	}
	if predmRash == "casinoPayment" {
		resPredmRash = 26
	}
	if predmRash == "fundsIssuance" {
		resPredmRash = 27
	}
	if predmRash == "exciseWithoutMarking" {
		resPredmRash = 30
	}
	if predmRash == "exciseWithMarking" {
		resPredmRash = 31
	}
	if predmRash == "commodityWithoutMarking" {
		resPredmRash = 32
	}
	if predmRash == "commodityWithMarking" {
		resPredmRash = 33
	}
	return resPredmRash
} //convertPredmRash

func convertPlannedStatusOfmc(status string) int {
	resStatus := 255
	if status == "itemPieceSold" {
		resStatus = 1
	}
	if status == "itemDryForSale" {
		resStatus = 2
	}
	if status == "itemPieceReturn" {
		resStatus = 3
	}
	if status == "itemDryReturn" {
		resStatus = 4
	}
	return resStatus
} //convertPlannedStatusOfmc

func convertAtolPosToMercPos(pos consttypes.TPosition) (consttypes.TMercPosition, error) {
	var mercPos consttypes.TMercPosition
	mercPos.Command = "AddGoods"
	mercPos.ProductName = pos.Name
	mercPos.Qty = int(pos.Quantity * 10000)
	mercPos.MeasureUnit = convertMeasureUnit(pos.MeasurementUnit)
	mercPos.TaxCode = 6
	if pos.Tax != nil {
		mercPos.TaxCode = convertTaxNDSCode(pos.Tax.Type)
	}
	mercPos.PaymentFormCode = convertSposRasch(pos.PaymentMethod)
	mercPos.ProductTypeCode = convertPredmRash(pos.PaymentObject)
	mercPos.Price = int(pos.Price * 100)
	if (pos.AgentInfo != nil) && (pos.SupplierInfo != nil) {
		mercPos.Agent = new(consttypes.TAgentMerc)
		mercPos.Agent.Code = 5 //комиссионер
		mercPos.Agent.SupplierPhone = append(mercPos.Agent.SupplierPhone, pos.SupplierInfo.Phones...)
		mercPos.Agent.SupplierName = pos.SupplierInfo.Name
		mercPos.Agent.SupplierINN = pos.SupplierInfo.Vatin
	}
	if pos.ImcParams != nil {
		mercPos.McInfo = new(consttypes.TMcInfoMerc)
		mercPos.McInfo.Mc = pos.ImcParams.Imc
		mercPos.McInfo.PlannedStatus = convertPlannedStatusOfmc(pos.ImcParams.ItemEstimatedStatus)
		mercPos.McInfo.ProcessingMode = pos.ImcParams.ImcModeProcessing
		mercPos.McInfo.Ean = pos.ImcParams.ImcBarcode
	}
	return mercPos, nil
} //convertAtolPosToMercPos

func convertAtolToMercCloseCheck(checkatol consttypes.TCorrectionCheck) consttypes.TMercCloseCheck {
	var checclosekmerc consttypes.TMercCloseCheck
	checclosekmerc.Command = "CloseCheck"
	if checkatol.ClientInfo != nil {
		checclosekmerc.SendCheckTo = checkatol.ClientInfo.EmailOrPhone
	}
	for _, payments := range checkatol.Payments {
		if payments.Type == "cash" {
			checclosekmerc.Payment.Cash = int(payments.Sum * 100)
		}
		if payments.Type == "electronically" {
			checclosekmerc.Payment.Ecash = int(payments.Sum * 100)
		}
		if payments.Type == "prepaid" {
			checclosekmerc.Payment.Prepayment = int(payments.Sum * 100)
		}
		if payments.Type == "credit" {
			checclosekmerc.Payment.Credit = int(payments.Sum * 100)
		}
		if payments.Type == "other" {
			checclosekmerc.Payment.Consideration = int(payments.Sum * 100)
		}
	}
	return checclosekmerc
} //convertAtolToMercCloseCheck

func sendCommandTCPMerc(bytesjson []byte, ip string, port int) ([]byte, error) {
	var buffAnsw []byte
	logsmy.LogginInFile(string(bytesjson))
	conn, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), 5*time.Second)
	if err != nil {
		descError := fmt.Sprintf("error: ошибка рукопожатия tcp %v\r\n", err)
		descError = descError + fmt.Sprintln("сервер ККТ не отвечает ККТ")
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	defer conn.Close()
	jsonBytes := bytesjson
	lenTCP := int32(len(jsonBytes))
	bytesLen := make([]byte, 4)
	bytesLen[3] = byte(lenTCP >> 0)
	bytesLen[2] = byte(lenTCP >> (1 * 8))
	bytesLen[1] = byte(lenTCP >> (2 * 8))
	bytesLen[0] = byte(lenTCP >> (3 * 8))
	var bufTCP bytes.Buffer
	_, err = bufTCP.Write(bytesLen)
	if err != nil {
		descError := fmt.Sprintf("error: ошибка записи в буфер данных длины пакета: %v\r\n", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	bufTCP.Write(jsonBytes)
	bufTCPReader := bytes.NewReader(bufTCP.Bytes())
	buffAnsw = make([]byte, 1024)
	var n int
	_, err = mustCopy(conn, bufTCPReader)
	if err != nil {
		descError := fmt.Sprintf("error: ошибка отправка tcp заароса серверу Мекрурия %v\r\n", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	n, err = conn.Read(buffAnsw)
	if err != nil {
		descError := fmt.Sprintf("error: ошибка получения ответа от сервера Меркурия  %v \r\n", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	logsmy.LogginInFile(string(buffAnsw))
	return buffAnsw[4:n], nil
} //sendCommandTCPMerc

func mustCopy(dst io.Writer, src io.Reader) (int64, error) {
	count, err := io.Copy(dst, src)
	if err != nil {
		descError := fmt.Sprintf("ошибка копирования %v\r\n", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
	}
	return count, err
} //mustCopy

func opensession(ipktt string, port int, comport int, userint int, passwuser string) ([]byte, error) {
	var jsonmerc []byte
	//jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"null\", \"command\":\"OpenSession\", \"portName\":\"COM%v\"}", comport))
	//jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":null, \"command\":\"OpenSession\", \"portName\":\"COM%v\"}", comport))
	if (userint != 0) || (passwuser != "") {
		jsonmerc = []byte(fmt.Sprintf("{\"sessionKey\":null, \"command\":\"OpenSession\", \"portName\":\"COM%v\", \"model\":\"185F\", \"userNumber\": %v,\"userPassword\": \"%v\", \"debug\": true, \"logPath\": \"c:\\\\logs\\\\\"}", comport, userint, passwuser))
	} else {
		jsonmerc = []byte(fmt.Sprintf("{\"sessionKey\":null, \"command\":\"OpenSession\", \"portName\":\"COM%v\", \"model\":\"185F\", \"debug\": true, \"logPath\": \"c:\\\\logs\\\\\"}", comport))
	}

	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		descError := fmt.Sprintf("ошибка (%v) открытия сессии для кассы меркурий", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	return buffAnsw, nil
} //opensession

func opencheck(ipktt string, port int, headercheckjson []byte) ([]byte, error) {
	buffAnsw, err := sendCommandTCPMerc(headercheckjson, ipktt, port)
	if err != nil {
		descError := fmt.Sprintf("ошибка (%v) открытия чека для кассы меркурий", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	return buffAnsw, nil
} //opencheck

func addpos(ipktt string, port int, posjson []byte) ([]byte, error) {
	buffAnsw, err := sendCommandTCPMerc(posjson, ipktt, port)
	if err != nil {
		descError := fmt.Sprintf("ошибка (%v) добавления позиции для кассы меркурий", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	return buffAnsw, nil
}

func closecheck(ipktt string, port int, forclosedatamerc []byte) ([]byte, error) {
	buffAnsw, err := sendCommandTCPMerc(forclosedatamerc, ipktt, port)
	if err != nil {
		descError := fmt.Sprintf("ошибка (%v) закрытия чека для кассы меркурий", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	return buffAnsw, nil
} //closecheck

func cancelcheck(ipktt string, port int, sessionkey *string) ([]byte, error) {
	jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"ResetCheck\"}", *sessionkey))
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	if err != nil {
		descError := fmt.Sprintf("ошибка (%v) отмены чека для кассы меркурий", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descError)
		return buffAnsw, err
	}
	return buffAnsw, nil
} //closecheck

func Closesession(ipktt string, port int, sessionkey *string) (string, error) {
	var resMerc consttypes.TAnswerMercur
	jsonmerc := []byte(fmt.Sprintf("{\"sessionKey\":\"%v\", \"command\":\"CloseSession\"}", *sessionkey))
	buffAnsw, err := sendCommandTCPMerc(jsonmerc, ipktt, port)
	*sessionkey = ""
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) закрытия сессии для кассы меркурий", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
		return descrError, err
	}
	err = json.Unmarshal(buffAnsw, &resMerc)
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) маршалинга результата закрытия сессии для кассы меркурий", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
		return descrError, err
	}
	if resMerc.Result != 0 {
		descrError := fmt.Sprintf("ошибка (%v) закрытия сессии для кассы меркурий", resMerc.Description)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
		err = fmt.Errorf(resMerc.Description)
		return descrError, err
	}
	return "", nil
} //closesession
