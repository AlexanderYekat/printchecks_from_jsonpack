//go:generate ./resource/goversioninfo.exe -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"bufio"
	fptr10 "clientrabbit/fptr"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type TClientInfo struct {
	EmailOrPhone string `json:"emailOrPhone"`
}

type TTaxNDS struct {
	Type string `json:"type,omitempty"`
}
type TProductCodesAtol struct {
	Undefined    string `json:"undefined,omitempty"` //32 символа только
	Code_EAN_8   string `json:"ean8,omitempty"`
	Code_EAN_13  string `json:"ean13,omitempty"`
	Code_ITF_14  string `json:"itf14,omitempty"`
	Code_GS_1    string `json:"gs10,omitempty"`
	Tag1305      string `json:"gs1m,omitempty"`
	Code_KMK     string `json:"short,omitempty"`
	Code_MI      string `json:"furs,omitempty"`
	Code_EGAIS_2 string `json:"egais20,omitempty"`
	Code_EGAIS_3 string `json:"egais30,omitempty"`
	Code_F_1     string `json:"f1,omitempty"`
	Code_F_2     string `json:"f2,omitempty"`
	Code_F_3     string `json:"f3,omitempty"`
	Code_F_4     string `json:"f4,omitempty"`
	Code_F_5     string `json:"f5,omitempty"`
	Code_F_6     string `json:"f6,omitempty"`
}

type TPayment struct {
	Type string  `json:"type"`
	Sum  float64 `json:"sum"`
}

type TGenearaPosAndTag11921191 struct {
	Type string `json:"type"`
}

type TPosition struct {
	Type            string   `json:"type"`
	Name            string   `json:"name"`
	Price           float64  `json:"price"`
	Quantity        float64  `json:"quantity"`
	Amount          float64  `json:"amount"`
	MeasurementUnit string   `json:"measurementUnit"`
	PaymentMethod   string   `json:"paymentMethod"`
	PaymentObject   string   `json:"paymentObject"`
	Tax             *TTaxNDS `json:"tax,omitempty"`
	//fot type tag1192 //AdditionalAttribute
	Value        string             `json:"value,omitempty"`
	Print        bool               `json:"print,omitempty"`
	ProductCodes *TProductCodesAtol `json:"productCodes,omitempty"`
	ImcParams    *TImcParams        `json:"imcParams,omitempty"`
	//Mark         string             `json:"mark,omitempty"`
}

type TAnsweChekcMark struct {
	Ready bool `json:"ready"`
}

type TBeginTaskMarkCheck struct {
	Type   string     `json:"type"`
	Params TImcParams `json:"params"`
}

type TItemInfoCheckResult struct {
	ImcCheckFlag              bool `json:"imcCheckFlag"`
	ImcCheckResult            bool `json:"imcCheckResult"`
	ImcStatusInfo             bool `json:"imcStatusInfo"`
	ImcEstimatedStatusCorrect bool `json:"imcEstimatedStatusCorrect"`
	EcrStandAloneFlag         bool `json:"ecrStandAloneFlag"`
}

type TImcParams struct {
	ImcType             string                `json:"imcType"`
	Imc                 string                `json:"imc"`
	ItemEstimatedStatus string                `json:"itemEstimatedStatus"`
	ImcModeProcessing   int                   `json:"imcModeProcessing"`
	ImcBarcode          string                `json:"imcBarcode,omitempty"`
	ItemInfoCheckResult *TItemInfoCheckResult `json:"itemInfoCheckResult,omitempty"`
	ItemQuantity        float64               `json:"itemQuantity,omitempty"`
	ItemUnits           string                `json:"itemUnits,omitempty"`
	NotSendToServer     bool                  `json:"notSendToServer,omitempty"`
}

type TAnswerGetStatusOfShift struct {
	ShiftStatus TShiftStatus `json:"shiftStatus"`
}
type TShiftStatus struct {
	DocumentsCount int    `json:"documentsCount"`
	ExpiredTime    string `json:"expiredTime"`
	Number         int    `json:"number"`
	State          string `json:"state"`
}

type TTag1192_91 struct {
	Type  string `json:"type"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Print bool   `json:"print,omitempty"`
}

type TOperator struct {
	Name  string `json:"name"`
	Vatin string `json:"vatin,omitempty"`
}

// При работе по ФФД ≥ 1.1 чеки коррекции имеют вид, аналогичный обычным чекам, но с
// добавлением информации о коррекции: тип, описание, дата документа основания и
// номер документа основания.
type TCorrectionCheck struct {
	Type string `json:"type"` //sellCorrection - чек коррекции прихода
	//buyCorrection - чек коррекции расхода
	//sellReturnCorrection - чек коррекции возврата прихода (ФФД ≥ 1.1)
	//buyReturnCorrection - чек коррекции возврата расхода
	Electronically       bool        `json:"electronically"`
	TaxationType         string      `json:"taxationType,omitempty"`
	ClientInfo           TClientInfo `json:"clientInfo"`
	CorrectionType       string      `json:"correctionType"` //
	CorrectionBaseDate   string      `json:"correctionBaseDate"`
	CorrectionBaseNumber string      `json:"correctionBaseNumber"`
	Operator             TOperator   `json:"operator"`
	//Items                []TPosition `json:"items"`
	Items    []interface{} `json:"items"` //либо TTag1192_91, либо TPosition
	Payments []TPayment    `json:"payments"`
	Total    float64       `json:"total,omitempty"`
}

var DIROFJSONS = ".\\jsons\\works\\"

var dirOfjsons = flag.String("dirjsons", ".\\jsons\\works\\", "директория json файлов по умолчанию ./jsons/")
var clearLogsProgramm = flag.Bool("clearlogs", true, "очистить логи программы")
var debugpr = flag.Bool("debug1", false, "дебажим программу")
var debug = flag.Bool("debug2", false, "режим отладки")
var emulation = flag.Bool("emul", false, "эмуляция")
var countOfCheckingMarks = flag.Int("attempts", 20, "число попыток провекри марки")
var clearTableOfMarks = flag.Bool("clearmarks", true, "очищать таблицу марок перед запуском на ККТ нового чека")
var countOfMistakesCheckForStop = flag.Int("stop_mist", 3, "число ошибочных чеков, после которого останавливать программу")

var LOGSDIR = "./logs/"
var filelogmap map[string]*os.File
var logsmap map[string]*log.Logger

const LOGINFO = "info"
const LOGINFO_WITHSTD = "info_std"
const LOGERROR = "error"
const LOGSKIP_LINES = "skip_line"
const LOGOTHER = "other"
const LOG_PREFIX = "TASKS"
const Version_of_program = "2024_02_05_02"

const FILE_NAME_PRINTED_CHECKS = "printed.txt"
const FILE_NAME_CONNECTION = "connection.txt"

func main() {
	var err error
	var descrError string

	runDescription := "программа версии " + Version_of_program + " распечатка чеков из json заданий запущена"
	fmt.Println(runDescription)
	defer fmt.Println("программа версии " + Version_of_program + " распечатка чеков из json заданий остановлена")

	fmt.Println("парсинг параметров запуска программы")
	flag.Parse()
	fmt.Println("emulation", *emulation)
	fmt.Println("debugpr", *debugpr)
	fmt.Println("debug: ", *debug)
	clearLogsDescr := fmt.Sprintf("Очистить логи программы %v", *clearLogsProgramm)
	fmt.Println(clearLogsDescr)
	fmt.Println("инициализация лог файлов программы")
	if foundedLogDir, _ := doesFileExist(LOGSDIR); !foundedLogDir {
		os.Mkdir(LOGSDIR, 0777)
	}
	filelogmap, logsmap, descrError, err = initializationLogs(*clearLogsProgramm, LOGINFO, LOGERROR, LOGSKIP_LINES, LOGOTHER)
	defer func() {
		fmt.Println("закрытие дескрипторов лог файлов программы")
		for _, v := range filelogmap {
			if v != nil {
				v.Close()
			}
		}
	}()
	input := bufio.NewScanner(os.Stdin)
	if err != nil {
		descrMistake := fmt.Sprintf("ошибка инициализации лог файлов %v", descrError)
		fmt.Fprint(os.Stderr, descrMistake)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic(descrMistake)
	}
	fmt.Println("лог файлы инициализированы в папке " + LOGSDIR)
	multwriterLocLoc := io.MultiWriter(logsmap[LOGINFO].Writer(), os.Stdout)
	logsmap[LOGINFO_WITHSTD] = log.New(multwriterLocLoc, LOG_PREFIX+"_"+strings.ToUpper(LOGINFO)+" ", log.LstdFlags)
	logginInFile(runDescription)
	logginInFile(clearLogsDescr)
	//
	DIROFJSONS = *dirOfjsons
	if foundedLogDir, _ := doesFileExist(DIROFJSONS); !foundedLogDir {
		err := os.Mkdir(DIROFJSONS, 0777)
		descrError := fmt.Sprintf("ошибка (%v) чтения директории %v с json заданиямию", err, DIROFJSONS)
		logsmap[LOGERROR].Println(descrError)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic(descrError)
	}

	listOfFilesTempr, err := listDirByReadDir(DIROFJSONS)
	if err != nil {
		logsmap[LOGERROR].Printf("ошибка поиска json заданий в директории %v c ошибкой %v", DIROFJSONS, err)
	}

	var listOfFiles []string
	countOfFiles := len(listOfFilesTempr)
	logsmap[LOGINFO_WITHSTD].Println("Всего json файлов", countOfFiles)
	//перебор всех файлов
	for _, v := range listOfFilesTempr {
		currFullFileName := DIROFJSONS + v
		numChecka := getFDFromFileName(v)
		printedThisCheck := false
		if numChecka == "" {
			logsmap[LOGERROR].Printf("пропущен файл %v", currFullFileName)
			continue
		}
		printedThisCheck, _ = printedCheck(DIROFJSONS, numChecka)
		if printedThisCheck {
			logsmap[LOGINFO_WITHSTD].Printf("чек с номером %v уже был распечатан", numChecka)
			continue
		}
		listOfFiles = append(listOfFiles, currFullFileName)
		//logsmap[LOGINFO_WITHSTD].Printf("%v = %v\n", k+1, currFullFileName)
	}
	countOfFiles = len(listOfFiles)
	logsmap[LOGINFO_WITHSTD].Println("инициализация драйвера атол")
	fptr, err := fptr10.NewSafe()
	if err != nil {
		descrError := fmt.Sprintf("Ошибка (%v) инициализации драйвера ККТ атол", err)
		logsmap[LOGERROR].Println(descrError)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic(descrError)
	}
	defer fptr.Destroy()
	fmt.Println(fptr.Version())
	//сединение с кассой
	logsmap[LOGINFO_WITHSTD].Println("соединение с кассой")
	comPort, err := getCurrentPortOfKass(DIROFJSONS)
	if err != nil {
		desrErr := fmt.Sprintf("ошибка (%v) чтения параметра com порт соединения с кассой", err)
		logsmap[LOGERROR].Println(desrErr)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic(desrErr)
	}
	if !connectWithKassa(fptr, comPort) {
		descrErr := fmt.Sprintf("ошибка сокдинения с кассовым аппаратом на ком порт %v", comPort)
		logsmap[LOGERROR].Println(descrErr)
		if !*debugpr {
			println("Нажмите любую клавишу...")
			input.Scan()
			log.Panic(descrErr)
		}
	} else {
		logsmap[LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", comPort)
	}
	defer fptr.Close()
	//jsonAnswer, err := sendComandeAndGetAnswerFromKKT(fptr, string(d.Body))
	//jsonAnswer, err := sendComandeAndGetAnswerFromKKT(fptr, "{\"type\": \"openShift\"}")
	//fmt.Println(jsonAnswer)
	//инициализация файла напечтанных чеков
	logsmap[LOGINFO_WITHSTD].Println("отрытие для записи таблицы напечатанных чеков")
	flagsTempOpen := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file_printed_checks, err := os.OpenFile(DIROFJSONS+FILE_NAME_PRINTED_CHECKS, flagsTempOpen, 0644)
	if err != nil {
		descrError := fmt.Sprintf("ошибка создания файла напечатанных чеков %v", err)
		logsmap[LOGERROR].Println(descrError)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic("ошибка инициализации напечтанных файла чеков", descrError)
	}
	defer file_printed_checks.Close()

	//перебор json заданаий и обработка
	countPrintedChecks := 0
	amountOfMistakesChecks := 0
	amountOfMistakesMarks := 0
	logsmap[LOGINFO_WITHSTD].Println("начинаем выполнять json чеков", countOfFiles)
	logsmap[LOGINFO_WITHSTD].Println("всего json заданий для печати чека", countOfFiles)
	for k, currFullFileName := range listOfFiles {
		var receipt TCorrectionCheck
		command := ""
		if amountOfMistakesChecks >= *countOfMistakesCheckForStop {
			descrError := "превышено количество ошибок чеков, остановка работы программы"
			logginInFile(descrError)
			resDial, command := dialogContinuePrintChecks()
			if !resDial && (command != "off/on") {
				descrError := "работы программы прервана пользователем"
				logsmap[LOGERROR].Println(descrError)
				//println("Нажмите любую клавишу...")
				//input.Scan()
				//log.Panic(descrError)
				break
			}
			amountOfMistakesChecks = 0
			amountOfMistakesMarks = 0
		}
		if command == "off/on" {
			command = ""
			err := reconnectToKKT(fptr)
			if err != nil {
				logsmap[LOGERROR].Printf("ошибка переподключения к ККТ %v", err)
				break
			}
		}
		currNumIsprChecka := getFDFromFileName(currFullFileName)
		logsmap[LOGINFO_WITHSTD].Printf("обработка задания %v из %v %v", k+1, countOfFiles, currFullFileName)
		logstr := fmt.Sprintf("начинаем читать json файл %v", currFullFileName)
		logginInFile(logstr)
		jsonCorrection, err := readJsonFromFile(currFullFileName)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) чтения json задания чека %v атол", err, currFullFileName)
			logsmap[LOGERROR].Println(errorDescr)
			amountOfMistakesChecks += 1
			continue
		}
		logstr = fmt.Sprintf("прочитали json файл %v", currFullFileName)
		logginInFile(logstr)
		//ищем марки в чеке
		logginInFile("ищем марки в чеке")
		existMarksInCheck := false
		err = json.Unmarshal([]byte(jsonCorrection), &receipt)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) json задания чека %v атол", err, jsonCorrection, currFullFileName)
			logsmap[LOGERROR].Println(errorDescr)
			amountOfMistakesChecks += 1
			continue
		}
		//очищаем таблицу марок
		if *clearTableOfMarks {
			logginInFile("очищаем таблицу марок")
			clearTableMarksJson := "{\"type\": \"clearMarkingCodeValidationResult\"}"
			resClearTableMarks, _ := sendComandeAndGetAnswerFromKKT(fptr, clearTableMarksJson)
			logstr = fmt.Sprintf("результат очистки таблицы марок: %v", resClearTableMarks)
			logginInFile(logstr)
		}
		//читаем данные по маркам
		mistakeChechingMark := false
		for _, v := range receipt.Items {
			typeItem := v.(map[string]interface{})["type"]
			if typeItem != "position" {
				continue

			}
			LocImcParams, ok := v.(map[string]interface{})["imcParams"]
			if !ok {
				continue
			}
			currMarkBase64interface, ok := LocImcParams.(map[string]interface{})["imc"]
			if !ok {
				continue
			}
			currMarkBase64 := currMarkBase64interface.(string)
			if currMarkBase64 == "" {
				continue
			}
			existMarksInCheck = true
			//QuantityPosInterface := v.(map[string]interface{})["quantity"]
			//MeasureUnitInterface := v.(map[string]interface{})["measurementUnit"]
			//QuantityPos := QuantityPosInterface.(float64)
			//MeasureUnit := MeasureUnitInterface.(string)
			logstr := fmt.Sprintf("запускаем процесс проверки марки %v для чека %v", currMarkBase64, currFullFileName)
			logsmap[LOGINFO_WITHSTD].Println(logstr)
			//logginInFile(logstr)
			imcResultCheckin, err := runProcessCheckMark(fptr, currMarkBase64, 0, "")
			if err != nil {
				errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v для чека %v атол", err, currMarkBase64, currFullFileName)
				logsmap[LOGERROR].Println(errorDescr)
				mistakeChechingMark = true
				break
			}
			//v.(map[string]interface{})["mark"] = ""
			logsmap[LOGINFO_WITHSTD].Println("марка успешно проверена")
			//заполняем данные о марке
			logginInFile("заполняем данные о марке")
			//v.(map[string]interface{})["imcParams"] = new(TImcParams)
			//ImcParams := v.(map[string]interface{})["ImcParams"].(*TImcParams)
			ImcParams := LocImcParams.(map[string]interface{})
			ImcParams["imc"] = currMarkBase64
			ImcParams["imcModeProcessing"] = 0
			ImcParams["itemEstimatedStatus"] = "itemStatusUnchanged"
			ImcParams["imcType"] = "auto"
			ImcParams["itemInfoCheckResult"] = new(TItemInfoCheckResult)
			ItemInfoCheckResult := ImcParams["itemInfoCheckResult"].(*TItemInfoCheckResult)
			ItemInfoCheckResult.ImcCheckFlag = imcResultCheckin.ImcCheckFlag
			ItemInfoCheckResult.ImcCheckResult = imcResultCheckin.ImcCheckResult
			ItemInfoCheckResult.ImcStatusInfo = imcResultCheckin.ImcStatusInfo
			ItemInfoCheckResult.ImcEstimatedStatusCorrect = imcResultCheckin.ImcEstimatedStatusCorrect
			//mcParams.Imc = currMarkBase64
			//ImcParams.ImcModeProcessing = 0
			//ImcParams.ItemEstimatedStatus = "itemStatusUnchanged"
			//ImcParams.ImcType = "auto"
			//ImcParams.ItemInfoCheckResult = new(TItemInfoCheckResult)
			//ImcParams.ItemInfoCheckResult.ImcCheckFlag = imcResultCheckin.ImcCheckFlag
			//ImcParams.ItemInfoCheckResult.ImcCheckResult = imcResultCheckin.ImcCheckResult
			//ImcParams.ItemInfoCheckResult.ImcStatusInfo = imcResultCheckin.ImcStatusInfo
			//ImcParams.ItemInfoCheckResult.ImcEstimatedStatusCorrect = imcResultCheckin.ImcEstimatedStatusCorrect
		}
		if mistakeChechingMark {
			errorDescr := fmt.Sprintf("ошибка проверки марки для чека %v атол", currFullFileName)
			logsmap[LOGERROR].Println(errorDescr)
			amountOfMistakesChecks += 1
			amountOfMistakesMarks += 1
			continue
		}
		if existMarksInCheck {
			jsonCorrWithMarkBytes, err := json.MarshalIndent(receipt, "", "\t")
			if err != nil {
				errorDescr := fmt.Sprintf("ошибка (%v) формирования json-а марками для задания чека %v атол", err, currFullFileName)
				logsmap[LOGERROR].Println(errorDescr)
				amountOfMistakesChecks += 1
				continue
			}
			jsonCorrection = string(jsonCorrWithMarkBytes)
		}
		logstr = fmt.Sprintf("послыем команду печати чека кассу json файл %v", jsonCorrection)
		logginInFile(logstr)
		resulOfCommand, err := sendComandeAndGetAnswerFromKKT(fptr, jsonCorrection)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) печати чека %v атол", err, currFullFileName)
			logsmap[LOGERROR].Println(errorDescr)
			amountOfMistakesChecks += 1
			continue
		}
		logginInFile("послали команду печати чека кассу json файл")
		if successCommand(resulOfCommand) {
			//при успешной печати чека, записываем данные о номере напечатнного чека
			countPrintedChecks++
			if countPrintedChecks == 1 {
				file_printed_checks.WriteString("\n")
			}
			file_printed_checks.WriteString(currNumIsprChecka + "\n")
		}
	} ///перебор json заданий
	logsmap[LOGINFO_WITHSTD].Printf("распечатно %v из %v чеков", countPrintedChecks, countOfFiles)
	//обработка лог файла
	//log.Fatal("штатный выход")
	println("Нажмите любую клавишу...")
	input.Scan()
}

func dialogContinuePrintChecks() (bool, string) {
	res := true
	command := ""
	input := bufio.NewScanner(os.Stdin)
	println("Продолжить (да - продолжить печать чеков, нет (по умолчанию) - завершить программу, \"off/on\" - переподключиться к кассе):")
	if input.Text() == "off/on" {
		command = "off/on"
	}
	input.Scan()
	res, _ = getBoolFromString(input.Text(), res)
	return res, command
}

func sendComandeAndGetAnswerFromKKT(fptr *fptr10.IFptr, comJson string) (string, error) {
	var err error
	logginInFile("начало процедуры sendComandeAndGetAnswerFromKKT")
	//return "", nil
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	//fptr.ValidateJson()
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*debugpr {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды %v на кассе", err, comJson)
			logsmap[LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры sendComandeAndGetAnswerFromKKT c ошибкой", err)
			logginInFile(logstr)
			return desrError, err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	//logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmap[LOGOTHER].Println("result", result)
	//if *debugpr {
	//}
	//result := "{\"result\": \"all ok\"}"
	logginInFile("конец процедуры sendComandeAndGetAnswerFromKKT без ошибки")
	return result, nil
}

func runProcessCheckMark(fptr *fptr10.IFptr, mark string, qnt float64, itemUnits string) (TItemInfoCheckResult, error) {
	var countAttempts int
	type TItemInfoCheckResultObject struct {
		ItemInfoCheckResult TItemInfoCheckResult `json:"itemInfoCheckResult"`
	}
	var imcResultCheckinObj TItemInfoCheckResultObject
	var imcResultCheckin TItemInfoCheckResult
	logginInFile("начало процедуры runProcessCheckMark")
	//проверяем - открыта ли смена
	//shiftOpenned, err := checkOpenShift(fptr, true, "админ")
	//if err != nil {
	//	errorDescr := fmt.Sprintf("ошибка (%v). Смена не открыта", err)
	//	logsmap[LOGERROR].Println(errorDescr)
	//	return TItemInfoCheckResult{}, errors.New(errorDescr)
	//}
	//if !shiftOpenned {
	//	errorDescr := fmt.Sprintf("ошибка (%v) - смена не открыта", err)
	//	logsmap[LOGERROR].Println(errorDescr)
	//	return TItemInfoCheckResult{}, errors.New(errorDescr)
	//}
	//посылаем запрос на проверку марки
	resJson, err := sendCheckOfMark(fptr, mark, qnt, itemUnits)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v", err, mark)
		logsmap[LOGERROR].Println(errorDescr)
		return TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	if !successCommand(resJson) {
		errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v", resJson, mark)
		logsmap[LOGERROR].Println(errorDescr)
		return TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	for countAttempts = 0; countAttempts < *countOfCheckingMarks; countAttempts++ {
		var answerOfCheckMark TAnsweChekcMark
		resJson, err := getStatusOfChecking(fptr)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) получения статуса проверки марки %v", err, mark)
			logsmap[LOGERROR].Println(errorDescr)
			return TItemInfoCheckResult{}, errors.New(errorDescr)
		}
		if !successCommand(resJson) {
			errorDescr := fmt.Sprintf("ошибка (%v) получения статуса проверки марки %v", resJson, mark)
			logsmap[LOGERROR].Println(errorDescr)
			return TItemInfoCheckResult{}, errors.New(errorDescr)
		}
		err = json.Unmarshal([]byte(resJson), &answerOfCheckMark)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) ответа проверки марки %v", err, resJson, mark)
			logsmap[LOGERROR].Println(errorDescr)
			return TItemInfoCheckResult{}, errors.New(errorDescr)
		}
		if answerOfCheckMark.Ready {
			if (*emulation) && (countAttempts < *countOfCheckingMarks-18) {

			} else {
				break
			}
		}
		//пауза в 1 секунду
		logsmap[LOGINFO_WITHSTD].Printf("попытка %v из %v получения статуса марки", countAttempts+2, *countOfCheckingMarks)
		duration := time.Second
		time.Sleep(duration)
	}
	if countAttempts == *countOfCheckingMarks {
		errorDescr := fmt.Sprintf("ошибка проверки марки %v", mark)
		logsmap[LOGERROR].Println(errorDescr)
		return TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	//принимаем марку
	resOfChecking, err := acceptMark(fptr)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) принятия марки %v", err, mark)
		logsmap[LOGERROR].Println(errorDescr)
		return TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	if !successCommand(resOfChecking) {
		errorDescr := fmt.Sprintf("ошибка (%v) принятия марки %v", resOfChecking, mark)
		logsmap[LOGERROR].Println(errorDescr)
		return TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	err = json.Unmarshal([]byte(resOfChecking), &imcResultCheckinObj)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) ответа проверки марки %v", err, resOfChecking, mark)
		logsmap[LOGERROR].Println(errorDescr)
		return TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	imcResultCheckin.EcrStandAloneFlag = imcResultCheckinObj.ItemInfoCheckResult.EcrStandAloneFlag
	imcResultCheckin.ImcCheckFlag = imcResultCheckinObj.ItemInfoCheckResult.ImcCheckFlag
	imcResultCheckin.ImcCheckResult = imcResultCheckinObj.ItemInfoCheckResult.ImcCheckResult
	imcResultCheckin.ImcEstimatedStatusCorrect = imcResultCheckinObj.ItemInfoCheckResult.ImcEstimatedStatusCorrect
	imcResultCheckin.ImcStatusInfo = imcResultCheckinObj.ItemInfoCheckResult.ImcStatusInfo
	logginInFile("конец процедуры runProcessCheckMark без ошибки")
	return imcResultCheckin, nil
} //runProcessCheckMark

func sendCheckOfMark(fptr *fptr10.IFptr, mark string, qnt float64, itemUnits string) (string, error) {
	var err error
	logginInFile("начало процедуры sendCheckOfMark")
	//return "", nil
	comJson, err := getJsonOfBeginCheck(mark, qnt, itemUnits)
	if err != nil {
		desrError := fmt.Sprintf("ошибка (%v) фоормирования json задания проверки марки", err)
		logsmap[LOGERROR].Println(desrError)
		logstr := fmt.Sprint("конец процедуры sendCheckOfMark c ошибкой", err)
		logginInFile(logstr)
		return "", err
	}
	logstr := fmt.Sprintf("отправляем запрос (%v) на проверку марки", comJson)
	logginInFile(logstr)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*debugpr {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды начать проверку марки на кассовом аппарате", err)
			logsmap[LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры sendCheckOfMark c ошибкой", err)
			logginInFile(logstr)
			return "", err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	logstr = fmt.Sprintf("ответ на начало проверки марки: (%v) ", result)
	logginInFile(logstr)
	//logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmap[LOGOTHER].Println("result", result)
	//if *debugpr {
	//}
	//result := "{\"result\": \"all ok\"}"
	logginInFile("конец процедуры sendCheckOfMark без ошибки")
	return result, nil
}

func getJsonOfBeginCheck(mark string, qnt float64, itemUnits string) (string, error) {
	var begMarkStructur TBeginTaskMarkCheck
	begMarkStructur.Type = "beginMarkingCodeValidation"
	begMarkStructur.Params.Imc = mark
	//begMarkStructur.Params.ItemQuantity = qnt
	//begMarkStructur.Params.ItemUnits = itemUnits
	begMarkStructur.Params.ImcType = "auto"
	begMarkStructur.Params.ItemEstimatedStatus = "itemStatusUnchanged"
	begMarkStructur.Params.ImcModeProcessing = 0
	//begMarkStructur.Params.NotSendToServer = false
	resstr, err := json.MarshalIndent(begMarkStructur, "", "\t")
	if err != nil {
		desrError := fmt.Sprintf("ошибка (%v) формирования json задания проверки марки", err)
		logsmap[LOGERROR].Println(desrError)
		return "", err
	}
	return string(resstr), err
}

func getStatusOfChecking(fptr *fptr10.IFptr) (string, error) {
	var err error
	logginInFile("начало процедуры getStatusOfChecking")
	//return "", nil
	comJson := "{\"type\": \"getMarkingCodeValidationStatus\"}"
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*debugpr {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды принятия марки на кассовом аппарате", err)
			logsmap[LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры getStatusOfChecking c ошибкой", err)
			logginInFile(logstr)
			return "", err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	if *emulation {
		result = "{\"ready\": true}"
	}
	//logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmap[LOGOTHER].Println("result", result)
	//if *debugpr {
	//}
	//result := "{\"result\": \"all ok\"}"
	logginInFile("конец процедуры getStatusOfChecking без ошибки")
	return result, nil
}

func acceptMark(fptr *fptr10.IFptr) (string, error) {
	var err error
	logginInFile("начало процедуры acceptMark")
	//return "", nil
	comJson := "{\"type\": \"acceptMarkingCode\"}"
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*debugpr {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды принятия марки на кассовом аппарате", err)
			logsmap[LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры acceptMark c ошибкой", err)
			logginInFile(logstr)
			return "", err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	if *emulation {
		result = "{\"itemInfoCheckResult\": {\"ecrStandAloneFlag\": false,\"imcCheckFlag\": true,\"imcCheckResult\": true,\"imcEstimatedStatusCorrect\": true,\"imcStatusInfo\": true}}"
	}
	//logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmap[LOGOTHER].Println("result", result)
	//if *debugpr {
	//}
	logstr := fmt.Sprintf("результат проверки марки (%v) ", result)
	logginInFile(logstr)
	//result := "{\"result\": \"all ok\"}"
	logginInFile("конец процедуры acceptMark без ошибки")
	return result, nil
}

func listDirByReadDir(path string) ([]string, error) {
	var spisFiles []string
	logstr := fmt.Sprintf("перебор файлов в директории %v--BEGIN\n", path)
	logginInFile(logstr)
	defer logginInFile(fmt.Sprintf("перебор файлов в директории %v--END\n", path))
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		return spisFiles, err
	}
	for _, val := range lst {
		if val.IsDir() {
			continue
		}
		matched := true
		if FILE_NAME_PRINTED_CHECKS == val.Name() {
			logsmap[LOGINFO_WITHSTD].Println("пропускаем файл с ифнормацией о напечатнных чеках")
			continue
		}
		if FILE_NAME_CONNECTION == val.Name() {
			logsmap[LOGINFO_WITHSTD].Println("пропускаем файл с ифнормацией о настройки связи с ККТ")
			continue
		}
		/*logsmap[LOGINFO].Println(val.Name())
		matched, err := regexp.MatchString(`fptr10\.log\.(2023\-1(1\-([0]9|[1-3][0-9])|2\-[0-9]{2}))\.gz`, val.Name())
		if !matched {
			matched, err = regexp.MatchString(`fptr10\.log\.(2023\-1(1\-([0]9|[1-3][0-9])|2\-[0-9]{2}))`, val.Name())
		}
		//matched, err := regexp.MatchString(`fptr10\.log\.2023\-(11|12)\-(09|[12][0-9]|)`, val.Name())
		if val.Name() == "fptr10.log" {
			matched = true
		}
		if err != nil {
			return spisFiles, err
		}*/
		logstr = fmt.Sprintln("matched=", matched)
		logginInFile(logstr)
		if matched {
			spisFiles = append(spisFiles, val.Name())
		}
	}
	return spisFiles, nil
} //listDirByReadDir

func doesFileExist(fullFileName string) (found bool, err error) {
	found = false
	if _, err = os.Stat(fullFileName); err == nil {
		// path/to/whatever exists
		found = true
	}
	return
}

func initializationLogs(clearLogs bool, logstrs ...string) (map[string]*os.File, map[string]*log.Logger, string, error) {
	var reserr, err error
	reserr = nil
	filelogmapLoc := make(map[string]*os.File)
	logsmapLoc := make(map[string]*log.Logger)
	descrError := ""
	for _, logstr := range logstrs {
		filenamelogfile := logstr + "logs.txt"
		preflog := LOG_PREFIX + "_" + strings.ToUpper(logstr)
		fullnamelogfile := LOGSDIR + filenamelogfile
		filelogmapLoc[logstr], logsmapLoc[logstr], err = intitLog(fullnamelogfile, preflog, clearLogs)
		if err != nil {
			descrError = fmt.Sprintf("ошибка инициализации лог файла %v с ошибкой %v", fullnamelogfile, err)
			fmt.Fprintln(os.Stderr, descrError)
			reserr = err
			break
		}
	}
	return filelogmapLoc, logsmapLoc, descrError, reserr
}

func printedCheck(dirjsons, numerChecka string) (bool, error) {
	//file_printed_checks, err := os.OpenFile(DIROFJSONS+"\\"+FILE_NAME_PRINTED_CHECKS, flagsTempOpen)
	res := false
	file_printed_checks, err := os.Open(DIROFJSONS + FILE_NAME_PRINTED_CHECKS)
	if err != nil {
		return res, err
	}
	defer file_printed_checks.Close()
	sc := bufio.NewScanner(file_printed_checks)
	for sc.Scan() {
		fd := string(sc.Text())
		if fd == numerChecka {
			res = true
			break
		}
	}
	return res, nil
}

func getFDFromFileName(fileNameOfJson string) string {
	fd := ""
	//7281440500811652_2333.json
	indOfPodch := strings.Index(fileNameOfJson, "_")
	if indOfPodch == -1 {
		return fd
	}
	rightStrObr := fileNameOfJson[indOfPodch+1:]
	indOfPoint := strings.Index(rightStrObr, ".")
	if indOfPodch == -1 {
		return fd
	}
	fd = rightStrObr[:indOfPoint]
	return fd
}

func readJsonFromFile(currFullFileName string) (string, error) {
	logstr := fmt.Sprintln("начало процедуры readJsonFromFile чтения файла", currFullFileName)
	logginInFile(logstr)
	plan, err := ioutil.ReadFile(currFullFileName)
	logstr = fmt.Sprintln("plan", plan)
	logginInFile(logstr)
	logstr = fmt.Sprintln("error", err)
	logginInFile(logstr)
	logstr = fmt.Sprintln("конец процедуры readJsonFromFile")
	logginInFile(logstr)
	return string(plan), err
}

func successCommand(resulJson string) bool {
	res := true
	indOsh := strings.Contains(resulJson, "ошибка")
	indErr := strings.Contains(resulJson, "error")
	if indErr || indOsh {
		res = false
	}
	return res
} //successCommand

func connectWithKassa(fptr *fptr10.IFptr, comport string) bool {
	sComPorta := comport
	if !strings.Contains(comport, "COM") {
		sComPorta = "COM" + comport
	}
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_COM))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_COM_FILE, sComPorta)
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_BAUDRATE, strconv.Itoa(fptr10.LIBFPTR_PORT_BR_115200))
	fptr.ApplySingleSettings()
	fptr.Open()
	return fptr.IsOpened()
}

func getCurrentPortOfKass(dirOfJsons string) (string, error) {
	comportb, err := os.ReadFile(dirOfJsons + FILE_NAME_CONNECTION)
	if err != nil {
		desrError := fmt.Sprintf("ошибка (%v) октрытия файла с параметрами соедиения кассы", err)
		logsmap[LOGERROR].Println(desrError)
		return desrError, err
	}
	return string(comportb), nil
}

func logginInFile(loggin string) {
	if *debug {
		logsmap[LOGINFO].Println(loggin)
	}
}

func intitLog(logFile string, pref string, clearLogs bool) (*os.File, *log.Logger, error) {
	flagsTempOpen := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	if clearLogs {
		flagsTempOpen = os.O_TRUNC | os.O_CREATE | os.O_WRONLY
	}
	f, err := os.OpenFile(logFile, flagsTempOpen, 0644)
	multwr := io.MultiWriter(f)
	//if pref == LOG_PREFIX+"_INFO" {
	//	multwr = io.MultiWriter(f, os.Stdout)
	//}
	flagsLogs := log.LstdFlags
	if pref == LOG_PREFIX+"_ERROR" {
		multwr = io.MultiWriter(f, os.Stderr)
		flagsLogs = log.LstdFlags | log.Lshortfile
	}
	if err != nil {
		fmt.Println("Не удалось создать лог файл ", logFile, err)
		return nil, nil, err
	}
	loger := log.New(multwr, pref+" ", flagsLogs)
	return f, loger, nil
}

func getBoolFromString(val string, onErrorDefault bool) (bool, error) {
	var err error
	res := onErrorDefault
	if (val == "да") || (val == "ДА") || (val == "Да") || (val == "yes") || (val == "Yes") || (val == "YES") {
		res = true
	} else if (val == "НЕТ") || (val == "нет") || (val == "Нет") || (val == "no") || (val == "No") || (val == "NO") {
		res = false
	} else {
		res, err = strconv.ParseBool(val)
		if err != nil {
			res = onErrorDefault
		}
	}
	return res, err
}

func reconnectToKKT(fptr *fptr10.IFptr) error {
	fptr.Close()
	fptr.Destroy()
	fptr, err := fptr10.NewSafe()
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) инициализации драйвера ККТ атол", err)
		logsmap[LOGERROR].Println(descrError)
		return errors.New(descrError)
		//println("Нажмите любую клавишу...")
		//input.Scan()
		//log.Panic(descrError)
	}
	//defer fptr.Destroy()
	fmt.Println(fptr.Version())
	//сединение с кассой
	logsmap[LOGINFO_WITHSTD].Println("соединение с кассой")
	comPort, err := getCurrentPortOfKass(DIROFJSONS)
	if err != nil {
		desrErr := fmt.Sprintf("ошибка (%v) чтения параметра com порт соединения с кассой", err)
		logsmap[LOGERROR].Println(desrErr)
		return errors.New(desrErr)
		//println("Нажмите любую клавишу...")
		//input.Scan()
		//log.Panic(desrErr)
	}
	if !connectWithKassa(fptr, comPort) {
		descrErr := fmt.Sprintf("ошибка сокдинения с кассовым аппаратом на ком порт %v", comPort)
		logsmap[LOGERROR].Println(descrErr)
		if !*debugpr {
			return errors.New(descrErr)
			//println("Нажмите любую клавишу...")
			//input.Scan()
			//log.Panic(descrErr)
		}
	} else {
		logsmap[LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", comPort)
	}
	return nil
	//defer fptr.Close()
}

/* func checkOpenShift(fptr *fptr10.IFptr, openShiftIfClose bool, kassir string) (bool, error) {
	logginInFile("получаем статус ККТ")
	getStatusKKTJson := "{\"type\": \"getDeviceStatus\"}"
	resgetStatusKKT, err := sendComandeAndGetAnswerFromKKT(fptr, getStatusKKTJson)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) получения статуса кассы", err)
		logsmap[LOGERROR].Println(errorDescr)
		return false, err
	}
	if !successCommand(resgetStatusKKT) {
		errorDescr := fmt.Sprintf("ошибка (%v) получения статуса кассы", resgetStatusKKT)
		logsmap[LOGERROR].Println(errorDescr)
		//logginInFile(errorDescr)
		return false, errors.New(errorDescr)
	}
	logginInFile("получили статус кассы")
	//проверяем - открыта ли смена
	var answerOfGetStatusofShift TAnswerGetStatusOfShift
	err = json.Unmarshal([]byte(resgetStatusKKT), &answerOfGetStatusofShift)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) распарсивания статуса кассы", err)
		logsmap[LOGERROR].Println(errorDescr)
		return false, err
	}
	if answerOfGetStatusofShift.ShiftStatus.State == "expired" {
		errorDescr := "ошибка - смена на кассе уже истекла. Закройте смену"
		logsmap[LOGERROR].Println(errorDescr)
		return false, errors.New(errorDescr)
	}
	if answerOfGetStatusofShift.ShiftStatus.State == "closed" {
		if openShiftIfClose {
			if kassir == "" {
				errorDescr := "не указано имя кассира для открытия смены"
				logsmap[LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
			jsonOpenShift := fmt.Sprintf("{\"type\": \"openShift\",\"operator\": {\"name\": \"%v\"}}", kassir)
			resOpenShift, err := sendComandeAndGetAnswerFromKKT(fptr, jsonOpenShift)
			if err != nil {
				errorDescr := fmt.Sprintf("ошбика (%v) - не удалось открыть смену", err)
				logsmap[LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
			if !successCommand(resOpenShift) {
				errorDescr := fmt.Sprintf("ошбика (%v) - не удалось открыть смену", resOpenShift)
				logsmap[LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
		} else {
			return false, nil
		}
	}
	return true, nil
} //checkOpenShift
*/
