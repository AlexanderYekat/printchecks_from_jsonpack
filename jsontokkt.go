package main

import (
	"bufio"
	fptr10 "clientrabbit/fptr"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var DIROFJSONS = "./jsons/works/"

var dirOfjsons = flag.String("dirjsons", "./jsons/works/", "директория json файлов по умолчанию ./jsons/")
var clearLogsProgramm = flag.Bool("clearlogs", true, "очистить логи программы")
var commandForAction = flag.String("command", "update", "команда")
var debugpr = flag.Bool("debug", false, "дебажим программу")
var emulation = flag.Bool("emul", false, "эмуляция")

var LOGSDIR = "./logs/"
var filelogmap map[string]*os.File
var logsmap map[string]*log.Logger

const LOGINFO = "info"
const LOGINFO_WITHSTD = "info_std"
const LOGERROR = "error"
const LOGSKIP_LINES = "skip_line"
const LOGOTHER = "other"
const LOG_PREFIX = "TASKS"
const Version_of_program = "2023_12_29_05"

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
	clearLogsDescr := fmt.Sprintf("Очистить логи программы %v", *clearLogsProgramm)
	fmt.Println(clearLogsDescr)
	fmt.Println("инициализация лог файлов программы")
	if foundedLogDir, _ := doesFileExist(LOGSDIR); !foundedLogDir {
		os.Mkdir(LOGSDIR, 0777)
	}
	//if foundedLogDir, _ := doesFileExist(RESULTSDIR); !foundedLogDir {
	//	os.Mkdir(RESULTSDIR, 0777)
	//}
	filelogmap, logsmap, descrError, err = initializationLogs(*clearLogsProgramm, LOGINFO, LOGERROR, LOGSKIP_LINES, LOGOTHER)
	defer func() {
		fmt.Println("закрытие дескрипторов лог файлов программы")
		for _, v := range filelogmap {
			if v != nil {
				//fmt.Println("close", k, v)
				v.Close()
			}
		}
	}()
	if err != nil {
		descrMistake := fmt.Sprintf("ошибка инициализации лог файлов %v", descrError)
		fmt.Fprint(os.Stderr, descrMistake)
		log.Panic(descrMistake)
	}
	fmt.Println("лог файлы инициализированы в папке " + LOGSDIR)
	multwriterLocLoc := io.MultiWriter(logsmap[LOGINFO].Writer(), os.Stdout)
	logsmap[LOGINFO_WITHSTD] = log.New(multwriterLocLoc, LOG_PREFIX+"_"+strings.ToUpper(LOGINFO)+" ", log.LstdFlags)
	logsmap[LOGINFO].Println(runDescription)
	logsmap[LOGINFO].Println(clearLogsDescr)

	DIROFJSONS = *dirOfjsons
	if foundedLogDir, _ := doesFileExist(DIROFJSONS); !foundedLogDir {
		err := os.Mkdir(DIROFJSONS, 0777)
		descrError := fmt.Sprintf("ошибка (%v) чтения директории %v с json заданиямию", err, DIROFJSONS)
		logsmap[LOGERROR].Println(descrError)
		log.Fatal(descrError)
	}

	listOfFilesTempr, err := listDirByReadDir(DIROFJSONS)
	if err != nil {
		logsmap[LOGERROR].Printf("ошибка поиска json заданий в директории %v c ошибкой %v", DIROFJSONS, err)
	}

	var listOfFiles []string
	countOfFiles := len(listOfFilesTempr)
	logsmap[LOGINFO_WITHSTD].Println("Всего json файлов", countOfFiles)
	//перебор всех файлов
	for k, v := range listOfFilesTempr {
		currFullFileName := DIROFJSONS + "\\" + v
		numChecka := getFDFromFileName(v)
		printedThisCheck := false
		if numChecka != "" {
			printedThisCheck, _ = printedCheck(DIROFJSONS, numChecka)
		}
		if printedThisCheck {
			logsmap[LOGINFO_WITHSTD].Printf("чек с номером %v уже был распечатан", numChecka)
			continue
		}
		listOfFiles = append(listOfFiles, currFullFileName)
		logsmap[LOGINFO_WITHSTD].Printf("%v = %v\n", k+1, currFullFileName)
	}
	countOfFiles = len(listOfFiles)
	logsmap[LOGINFO_WITHSTD].Println("инициализация драйвера атол")
	fptr, err := fptr10.NewSafe()
	if err != nil {
		descrError := fmt.Sprintf("Ошибка (%v) инициализации драйвера ККТ атол", err)
		logsmap[LOGERROR].Println(descrError)
		log.Fatal(descrError)
	}
	defer fptr.Destroy()
	fmt.Println(fptr.Version())
	//сединение с кассой
	logsmap[LOGINFO_WITHSTD].Println("соединение с кассой")
	comPort, err := getCurrentPortOfKass(DIROFJSONS)
	if err != nil {
		desrErr := fmt.Sprintf("ошибка (%v) чтения параметра com порт соединения с кассой", err)
		logsmap[LOGERROR].Println(desrErr)
		log.Fatal(desrErr)
	}
	if !connectWithKassa(fptr, comPort) {
		descrErr := fmt.Sprintf("ошибка сокдинения с кассовым аппаратом на ком порт %v", comPort)
		logsmap[LOGERROR].Println(descrErr)
		if !*debugpr {
			log.Fatal(descrErr)
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
	file_printed_checks, err := os.OpenFile(DIROFJSONS+"\\"+FILE_NAME_PRINTED_CHECKS, flagsTempOpen, 0644)
	if err != nil {
		descrError := fmt.Sprintf("ошибка создания файла напечатанных чеков %v", err)
		logsmap[LOGERROR].Println(descrError)
		log.Panic("ошибка инициализации напечтанных файла чеков", descrError)
	}
	defer file_printed_checks.Close()

	//перебор json занаий и обработка
	countPrintedChecks := 0
	logsmap[LOGINFO_WITHSTD].Println("начинаем выполнять json чеков", countOfFiles)
	logsmap[LOGINFO_WITHSTD].Println("всего json заданий для печати чека", countOfFiles)
	for k, currFullFileName := range listOfFiles {
		currNumIsprChecka := getFDFromFileName(currFullFileName)
		logsmap[LOGINFO_WITHSTD].Printf("обработка задания %v из %v %v", k+1, countOfFiles, currFullFileName)
		logsmap[LOGINFO].Printf("начинаем читать json файл %v", currFullFileName)
		jsonCorrection, err := readJsonFromFile(currFullFileName)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) чтения json задания чека %v атол", err, currFullFileName)
			logsmap[LOGERROR].Println(errorDescr)
			continue
		}
		logsmap[LOGINFO].Printf("прочитали json файл %v", currFullFileName)
		logsmap[LOGINFO].Printf("послыем команду печати чека кассу json файл %v", jsonCorrection)
		resulOfCommand, err := sendComandeAndGetAnswerFromKKT(fptr, jsonCorrection)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) печати чека %v атол", err, currFullFileName)
			logsmap[LOGERROR].Println(errorDescr)
			continue
		}
		logsmap[LOGINFO].Print("послали команду печати чека кассу json файл")
		if successCommand(resulOfCommand) {
			//при успешной печати чека, записываем данные о номере напечатнного чека
			countPrintedChecks++
			if countPrintedChecks == 1 {
				file_printed_checks.WriteString("\n")
			}
			file_printed_checks.WriteString(currNumIsprChecka + "\n")
		}
	}
	logsmap[LOGINFO_WITHSTD].Printf("распечатно %v из %v чеков", countPrintedChecks, countOfFiles)
	//обработка лог файла
	//log.Fatal("штатный выход")
}

func sendComandeAndGetAnswerFromKKT(fptr *fptr10.IFptr, comJson string) (string, error) {
	var err error
	logsmap[LOGINFO].Print("начало процедуры sendComandeAndGetAnswerFromKKT")
	//return "", nil
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	//fptr.ValidateJson()
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*debugpr {
			desrError := fmt.Sprintf("ошибка (%v) печати чека коррекции на кассовом аппарате", err)
			logsmap[LOGERROR].Println(desrError)
			logsmap[LOGINFO].Print("начало процедуры sendComandeAndGetAnswerFromKKT c ошибкой", err)
			return desrError, err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	logsmap[LOGOTHER].Println("comJson", comJson)
	logsmap[LOGOTHER].Println("result", result)
	//if *debugpr {
	//}
	//result := "{\"result\": \"all ok\"}"
	logsmap[LOGINFO].Print("начало процедуры sendComandeAndGetAnswerFromKKT без ошибки")
	return result, nil
}

func listDirByReadDir(path string) ([]string, error) {
	var spisFiles []string
	logsmap[LOGINFO].Printf("перебор файлов в директории %v--BEGIN\n", path)
	defer logsmap[LOGINFO].Printf("перебор файлов в директории %v--END\n", path)
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
		logsmap[LOGINFO].Println("matched=", matched)
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
	file_printed_checks, err := os.Open(DIROFJSONS + "\\" + FILE_NAME_PRINTED_CHECKS)
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
	logsmap[LOGINFO].Println("начало процедуры readJsonFromFile чтения файла", currFullFileName)
	plan, err := ioutil.ReadFile(currFullFileName)
	logsmap[LOGINFO].Println("plan", plan)
	logsmap[LOGINFO].Println("error", err)
	logsmap[LOGINFO].Println("конец процедуры readJsonFromFile")
	return string(plan), err
}

func successCommand(resulJson string) bool {
	res := true
	indOsh := strings.Index(resulJson, "ошибка")
	indErr := strings.Index(resulJson, "error")
	if indErr != -1 || indOsh != -1 {
		res = false
	}
	return res
}

func connectWithKassa(fptr *fptr10.IFptr, comport string) bool {
	sComPorta := comport
	if strings.Index(comport, "COM") == -1 {
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
	comportb, err := os.ReadFile(dirOfJsons + "\\" + FILE_NAME_CONNECTION)
	if err != nil {
		desrError := fmt.Sprintf("ошибка (%v) октрытия файла с параметрами соедиения кассы", err)
		logsmap[LOGERROR].Println(desrError)
		return desrError, err
	}
	return string(comportb), nil
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
