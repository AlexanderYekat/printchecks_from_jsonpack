//go:generate ./resource/goversioninfo.exe -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"bufio"
	consttypes "clientrabbit/consttypes"
	fptr10 "clientrabbit/fptr"
	merc "clientrabbit/sendtcp"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var kassatype = flag.String("kassatype", "atol", "касса (atol - атол, merc - меркурий)")
var IpMerc = flag.String("ipMerc", "localhost", "ip адрес сервера Меркурия")
var PortMerc = flag.Int("PortMerc", 50009, "порт сервера Меркурия")
var dirOfjsons = flag.String("dirjsons", ".\\jsons\\works\\", "директория json файлов по умолчанию ./jsons/")
var clearLogsProgramm = flag.Bool("clearlogs", true, "очистить логи программы")
var countChecksForPause = flag.Int("countforpause", 0, "число чеков, после которых программа делает небольшую паузу")
var pauseInSecondsMayClose = flag.Int("secpause", 10, "сколько секунд на паузу, если вдруг надо завершить программу")
var LogsDebugs = flag.Int("debug", 0, "уровень логирования всех действий, чем выше тем больше логов")
var comport = flag.Int("com", 0, "ком порт кассы")
var emailforcheck = flag.String("email", "", "email клиента чека")
var PrintCheckOnKKT = flag.String("print", "", "печтать или не печатать чек на ККТ: true - печатать, false - не печетать")
var ipaddresskkt = flag.String("ipkkt", "", "ip адрес ккт")
var portkktatol = flag.Int("portipkkt", 0, "порт ip ккт")
var ipaddressservrkkt = flag.String("ipservkkt", "", "ip адрес сервера ккт")
var emulation = flag.Bool("emul", false, "эмуляция")
var emulatmistakes = flag.Bool("emulmist", false, "эмуляция ошибок")
var emulatmistakesmarks = flag.Bool("emulmistmark", false, "эмуляция ошибок марок")
var countOfCheckingMarks = flag.Int("attempts", 20, "число попыток провекри марки")
var clearTableOfMarks = flag.Bool("clearmarks", true, "очищать таблицу марок перед запуском на ККТ нового чека")
var countOfMistakesCheckForStop = flag.Int("stop_mist", 3, "число ошибочных чеков, после которого останавливать программу")
var pauseOfMarksMistake = flag.Int("pause_mist", 10, "пауза между проблемами с марками")

var countPrintChecks = flag.Int("countchecks", 0, "число успешно распечатнных чеков, после которого остановить программу")
var pauseAfterDay = flag.Int("pauseAfterDay", 0, "число дней, после которого программа делает паузу")
var pauseInSecondsAfterDay = flag.Int("pausefterdaysec", 90, "пауза в секундах после звершение какого-то количества дней напечатнных чеков")

var ExlusionDate = flag.String("exldate", "", "дата исключения из распечатки в формате 2006.01.02")

var filelogmap map[string]*os.File
var logsmap map[string]*log.Logger

const Version_of_program = "2024_04_14_01"

func main() {
	var err error
	var descrError string
	var ExlusionDateDate time.Time
	var lastNameOfKassir string
	var fptr *fptr10.IFptr
	var sessionkey string
	//выводим информацию о программе
	//читаем параметры запуска программы
	//открываем лог файлы
	//ищем все файлы заданий в директории json - заданий
	//убираем json-задания, которые уже были распечатаны
	//читаем настроку com - порта в директории json - заданий
	//подключаемся к кассовому аппарату
	//открытие для запиписи файла напечатанных чеков
	//инициализация переменных для цикла перебора json-заданий
	//цикл перебора json-заданий
	//////инициализируем переменные шага цикла
	//////проверяем условия выхода из цикла
	//////читаем json - задание
	//////проеверяем услвоия выхода из цикла по дате чека
	//////если надо делаем паузу в работе программы
	//////очищаем таблицу марок
	//////ищем все марки в json-задании и запускаем по каждой из них проверку
	//////эмулируем ошибкук провекри марки, если режим эмуляции ошибки включен
	//////перезапускаем полногстью процесс проверки марок, если были ошибки
	//////если были ошибки при печати чека прерываем программу
	//////пропускаем чек, если были ошибки при проверке марки
	//////пересобираем json-задание, если необходимо (вставляем результаты проверки марок, изменяем параметры печати/не печати и email)
	//////печатаем чек
	//////если были ошибку при печати чека, то переходим к следующему заданию
	//////эмулируем ошибку, если режим эмуляции ошибки включен
	//////читаем информацию об результате выполнения команды
	//////если команда выполнена успешно, то записываем в таблицу напечатанных чеков
	//////если команда выполнена неуспешно, то проверяем не превышен ли количество чеков в смену,
	//////и если превышено, то закрываем и открываем смену
	//выводим информацию об количестве напечтатнных чеков
	//
	/////////////////**************************///////////////////////
	//
	//выводим информацию о программе
	runDescription := "программа версии " + Version_of_program + " распечатка чеков из json заданий запущена"
	fmt.Println(runDescription)
	defer fmt.Println("программа версии " + Version_of_program + " распечатка чеков из json заданий остановлена")
	//читаем параметры запуска программы
	fmt.Println("парсинг параметров запуска программы")
	flag.Parse()
	fmt.Println("Эмулирование ККТ", *emulation)
	fmt.Println("Уровень логирования: ", *LogsDebugs)
	clearLogsDescr := fmt.Sprintf("Очистить логи программы %v", *clearLogsProgramm)
	fmt.Println(clearLogsDescr)
	//открываем лог файлы
	fmt.Println("инициализация лог файлов программы")
	if foundedLogDir, _ := doesFileExist(consttypes.LOGSDIR); !foundedLogDir {
		os.Mkdir(consttypes.LOGSDIR, 0777)
	}
	filelogmap, logsmap, descrError, err = initializationLogs(*clearLogsProgramm, consttypes.LOGINFO, consttypes.LOGERROR, consttypes.LOGSKIP_LINES, consttypes.LOGOTHER)
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
	fmt.Println("лог файлы инициализированы в папке " + consttypes.LOGSDIR)
	multwriterLocLoc := io.MultiWriter(logsmap[consttypes.LOGINFO].Writer(), os.Stdout)
	logsmap[consttypes.LOGINFO_WITHSTD] = log.New(multwriterLocLoc, consttypes.LOG_PREFIX+"_"+strings.ToUpper(consttypes.LOGINFO)+" ", log.LstdFlags)
	logginInFile(runDescription)
	logginInFile(clearLogsDescr)
	//ищем все файлы заданий в директории json - заданий
	consttypes.DIROFJSONS = *dirOfjsons
	if foundedLogDir, _ := doesFileExist(consttypes.DIROFJSONS); !foundedLogDir {
		err := os.Mkdir(consttypes.DIROFJSONS, 0777)
		descrError := fmt.Sprintf("ошибка (%v) чтения директории %v с json заданиямию", err, consttypes.DIROFJSONS)
		logsmap[consttypes.LOGERROR].Println(descrError)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic(descrError)
	}
	listOfFilesTempr, err := listDirByReadDir(consttypes.DIROFJSONS)
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) поиска json заданий в директории %v", err, consttypes.DIROFJSONS)
		logsmap[consttypes.LOGERROR].Printf(descrError)
		log.Panic(descrError)
	}
	logginInFile(fmt.Sprintln("listOfFilesTempr=", listOfFilesTempr))
	var listOfFiles []string
	countOfFiles := len(listOfFilesTempr)
	logsmap[consttypes.LOGINFO_WITHSTD].Println("Всего json файлов", countOfFiles)
	//убираем json-задания, которые уже были распечатаны
	for _, v := range listOfFilesTempr {
		currFullFileName := consttypes.DIROFJSONS + v
		numChecka := getFDFromFileName(v)
		printedThisCheck := false
		if numChecka == "" {
			logsmap[consttypes.LOGERROR].Printf("пропущен файл %v", currFullFileName)
			continue
		}
		printedThisCheck, _ = printedCheck(consttypes.DIROFJSONS, numChecka)
		if printedThisCheck {
			logsmap[consttypes.LOGINFO_WITHSTD].Printf("чек с номером %v уже был распечатан", numChecka)
			continue
		}
		listOfFiles = append(listOfFiles, currFullFileName)
		//logsmap[consttypes.LOGINFO_WITHSTD].Printf("%v = %v\n", k+1, currFullFileName)
	}
	countOfFiles = len(listOfFiles)
	//читаем настроку com - порта в директории json - заданий
	*comport, _ = getCurrentPortOfKass(consttypes.DIROFJSONS)
	logsmap[consttypes.LOGINFO_WITHSTD].Println("порт кассы", *comport)
	//подключаемся к кассовому аппарату
	if *kassatype == "atol" {
		logsmap[consttypes.LOGINFO_WITHSTD].Println("Тип кассы atol")
		logsmap[consttypes.LOGINFO_WITHSTD].Println("инициализация драйвера атол")
		fptr, err = fptr10.NewSafe()
		if err != nil {
			descrError := fmt.Sprintf("Ошибка (%v) инициализации драйвера ККТ атол", err)
			logsmap[consttypes.LOGERROR].Println(descrError)
			println("Нажмите любую клавишу...")
			input.Scan()
			log.Panic(descrError)
		}
		defer fptr.Destroy()
		fmt.Println(fptr.Version())
		//сединение с кассой
		logsmap[consttypes.LOGINFO_WITHSTD].Println("соединение с кассой")
		if ok, typepodkluch := connectWithKassa(fptr, *comport, *ipaddresskkt, *portkktatol, *ipaddressservrkkt); !ok {
			descrErr := fmt.Sprintf("ошибка соединения с кассовым аппаратом %v", typepodkluch)
			logsmap[consttypes.LOGERROR].Println(descrErr)
			if !*emulation {
				println("Нажмите любую клавишу...")
				input.Scan()
				log.Panic(descrErr)
			}
		} else {
			logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", *comport)
		}
		defer fptr.Close()
	} else {
		var err error
		logother := new(string)
		descrError := ""
		sessionkey, descrError, err = merc.CheckStatsuConnectionKKT(*IpMerc, *PortMerc, *comport, "", logother)
		if err != nil {
			logsmap[consttypes.LOGERROR].Printf("ошибка (%v) подлкючение к ККТ меркурий", descrError)
			if !*emulation {
				println("Нажмите любую клавишу...")
				input.Scan()
				log.Panic(descrError)
			}
		} else {
			logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе меркурий на порт %v прошло успешно. Ключ сессии: %v", *comport, sessionkey)
		}
		logginInFile(*logother)
		merc.Closesession(*IpMerc, *PortMerc, &sessionkey, logother)
	}
	//открытие для запиписи файла напечатанных чеков
	logsmap[consttypes.LOGINFO_WITHSTD].Println("отрытие для записи таблицы напечатанных чеков")
	flagsTempOpen := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file_printed_checks, err := os.OpenFile(consttypes.DIROFJSONS+consttypes.FILE_NAME_PRINTED_CHECKS, flagsTempOpen, 0644)
	if err != nil {
		descrError := fmt.Sprintf("ошибка создания файла напечатанных чеков %v", err)
		logsmap[consttypes.LOGERROR].Println(descrError)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic("ошибка инициализации напечтанных файла чеков", descrError)
	}
	defer file_printed_checks.Close()
	//инициализация переменных для цикла перебора json-заданий
	countPrintedChecks := 0
	amountOfMistakesChecks := 0
	amountOfMistakesMarks := 0
	countPrintedDays := 0
	initDate, err := time.Parse("2006.01.02", "2006.01.02")
	if err != nil {
		logsmap[consttypes.LOGERROR].Printf("ошибка (%v) инициализации начальной даты", err)
	}
	prevDateOfCheck := initDate
	logsmap[consttypes.LOGINFO_WITHSTD].Println("начинаем выполнять json чеков", countOfFiles)
	logsmap[consttypes.LOGINFO_WITHSTD].Println("всего json заданий для печати чека", countOfFiles)
	previusWasMarks := false
	if *ExlusionDate != "" {
		ExlusionDateDate, err = time.Parse("2006.01.02", *ExlusionDate)
		if err != nil {
			logsmap[consttypes.LOGERROR].Printf("ошибка (%v) инициализации даты исключения", err)
			*ExlusionDate = ""
		}
	}
	//цикл перебора json-заданий
	for k, currFullFileName := range listOfFiles {
		var receipt consttypes.TCorrectionCheck
		//инициализируем переменные шага цикла
		//проверяем условия выхода из цикла
		//читаем json - задание
		//проеверяем услвоия выхода из цикла по дате чека
		//если надо делаем паузу в работе программы
		//очищаем таблицу марок
		//ищем все марки в json-задании и запускаем по каждой из них проверку
		//эмулируем ошибкук провекри марки, если режим эмуляции ошибки включен
		//перезапускаем полногстью процесс проверки марок, если были ошибки
		//если были ошибки при печати чека прерываем программу
		//пропускаем чек, если были ошибки при проверке марки
		//пересобираем json-задание, если необходимо (вставляем результаты проверки марок, изменяем параметры печати/не печати и email)
		//печатаем чек
		//если были ошибку при печати чека, то переходим к следующему заданию
		//эмулируем ошибку, если режим эмуляции ошибки включен
		//читаем информацию об результате выполнения команды
		//если команда выполнена успешно, то записываем в таблицу напечатанных чеков
		//если команда выполнена неуспешно, то проверяем не превышен ли количество чеков в смену,
		//и если превышено, то закрываем и открываем смену
		//////
		//инициализируем переменные шага цикла
		globalMistake := false
		globalErrorStr := ""
		command := ""
		//проверяем условия выхода из цикла
		if amountOfMistakesChecks >= *countOfMistakesCheckForStop {
			descrError := "превышено количество ошибок чеков, остановка работы программы"
			logginInFile(descrError)
			resDial := false
			resDial, command = dialogContinuePrintChecks()
			if !resDial && (command != "off/on") {
				descrError := "работы программы прервана пользователем"
				logsmap[consttypes.LOGERROR].Println(descrError)
				break
			}
			amountOfMistakesChecks = 0
			amountOfMistakesMarks = 0
		} else {
			if *countPrintChecks > 0 {
				if countPrintedChecks >= *countPrintChecks {
					desctriptionExit := fmt.Sprintf("произошло завершение работы программы, так как число напечатнных чеков %v равно параметру countchecks, переданному при запуске программы", countPrintedChecks)
					logsmap[consttypes.LOGINFO_WITHSTD].Println(desctriptionExit)
					break //прерываем печать чека
				}
			}
			if *countChecksForPause > 0 {
				if ((countPrintedChecks + 1) % *countChecksForPause) == 0 {
					logginInFile(fmt.Sprintf("делаем паузу в программе через каждые %v чеков для возможной безопасной её остановки на %v секунд...", *countChecksForPause, *pauseInSecondsMayClose))
					logsmap[consttypes.LOGINFO_WITHSTD].Println("если процесс печати чеков нужно прервать, то это можно сделать сейчас")
					duration := time.Second * time.Duration((*pauseInSecondsMayClose))
					time.Sleep(duration)
				}
			}
		}
		if command != "" {
			logginInFile(fmt.Sprintln("command", command))
		}
		if command == "off/on" {
			command = ""
			logsmap[consttypes.LOGINFO_WITHSTD].Println("переподключение к кассовому аппарату...")
			if *kassatype == "atol" {
				err := reconnectToKKT(fptr)
				if err != nil {
					logsmap[consttypes.LOGERROR].Printf("ошибка переподключения к ККТ %v", err)
					break
				}
			}
		}
		//читаем json - задание
		currNumIsprChecka := getFDFromFileName(currFullFileName)
		logginInFile(fmt.Sprintf("обработка задания %v из %v %v", k+1, countOfFiles, currFullFileName))
		logstr := fmt.Sprintf("начинаем читать json файл %v", currFullFileName)
		logginInFile(logstr)
		jsonCorrection, err := readJsonFromFile(currFullFileName)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) чтения json задания чека %v атол", err, currFullFileName)
			logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			continue
		}
		logstr = fmt.Sprintf("прочитали json файл %v", currFullFileName)
		logginInFile(logstr)
		logginInFile("парсим json задание")
		existMarksInCheck := false
		err = json.Unmarshal([]byte(jsonCorrection), &receipt)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) json задания чека %v атол", err, jsonCorrection, currFullFileName)
			logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			continue
		}
		lastNameOfKassir = receipt.Operator.Name
		//проеверяем услвоия выхода из цикла по дате чека
		currDateOfCheck, err := time.Parse("2006.01.02", receipt.CorrectionBaseDate) //yyyy.mm.dd
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) парсинга даты %v для чека %v", err, receipt.CorrectionBaseDate, currFullFileName)
			logsmap[consttypes.LOGERROR].Println(errorDescr)
		}
		if *ExlusionDate != "" {
			if ExlusionDateDate == currDateOfCheck {
				desrExit := fmt.Sprintf("достигли даты %v исключения - завершаем работы программы ", *ExlusionDate)
				logsmap[consttypes.LOGINFO_WITHSTD].Println(desrExit)
				break
			}
		}
		if prevDateOfCheck != currDateOfCheck {
			logginInFile(fmt.Sprintf("переходим на новый день %v", currDateOfCheck))
			if prevDateOfCheck != initDate {
				countPrintedDays++
			}
			prevDateOfCheck = currDateOfCheck
		}
		//если надо делаем паузу в работе программы
		if *pauseAfterDay > 0 {
			if countPrintedDays >= *pauseAfterDay {
				logginInFile(fmt.Sprintf("произошло завершение дня, работы программы поставлена на паузу на %v секунд", *pauseInSecondsAfterDay))
				logginInFile(fmt.Sprintf("делаем паузу в программе через каждые %v дней для возможной остановки на %v секунд...", *pauseAfterDay, *pauseInSecondsAfterDay))
				logsmap[consttypes.LOGINFO_WITHSTD].Printf("если процесс печати чеков нужно прервать, то это можно сделать сейчас - так как сейчас программа перешла на следующий %v день", prevDateOfCheck)
				countPrintedDays = 0
				duration := time.Second * time.Duration((*pauseInSecondsAfterDay))
				time.Sleep(duration)
			}
		}
		logsmap[consttypes.LOGINFO_WITHSTD].Printf("%v: обработка задания %v из %v %v", receipt.CorrectionBaseDate, k+1, countOfFiles, currFullFileName)
		logginInFile("ищем марки в чеке")
		//очищаем таблицу марок
		if (*clearTableOfMarks) && (previusWasMarks) {
			if *kassatype == "atol" {
				breakProcCheckOfMark(fptr)
				clearTanlesOfMarks(fptr)
			} else {
				logother := new(string)
				merc.BreakAndClearProccessOfMarks(*IpMerc, *PortMerc, *comport, sessionkey, logother)
				logginInFile(*logother)
			}
		}
		//ищем все марки в json-задании и запускаем по каждой из них проверку
		previusWasMarks = false
		mistakeCheckingMark := false
		markErroDescr := ""
		receipt, existMarksInCheck, mistakeCheckingMark, markErroDescr, globalMistake, globalErrorStr = CheckAndRunsCheckingMarksByCheck(fptr, *kassatype, receipt, currFullFileName, true, *IpMerc, *PortMerc, *comport, sessionkey, *pauseOfMarksMistake)
		if existMarksInCheck {
			previusWasMarks = true
		}
		//эмулируем ошибкук провекри марки, если надо
		if (*emulatmistakesmarks) && (existMarksInCheck) {
			markErroDescr = "эмуляция ошибки всего процесса проверки марок чека"
			mistakeCheckingMark = true
		}
		//перезапускаем полногстью процесс проверки марок, если были ошибки
		if mistakeCheckingMark {
			//очищаем все предыдущие провекри
			logsmap[consttypes.LOGINFO_WITHSTD].Printf("перезапускаем процесс провекри марок для всего чека, так как была ошибкаЖ %v...", markErroDescr)
			if *kassatype == "atol" {
				breakProcCheckOfMark(fptr)
				clearTanlesOfMarks(fptr)
			} else {
				logother := new(string)
				merc.BreakAndClearProccessOfMarks(*IpMerc, *PortMerc, *comport, sessionkey, logother)
				logginInFile(*logother)
			}
			//отключаемся от ККТ
			logsmap[consttypes.LOGINFO_WITHSTD].Println("отлючаемся от ККТ")
			if *kassatype == "atol" {
				disconnectWithKKT(fptr, true)
			} else {
				logother := new(string)
				merc.DissconnectMeruriy(*IpMerc, *PortMerc, sessionkey, logother)
				logginInFile(*logother)
			}
			//делаем паузу
			logsmap[consttypes.LOGINFO_WITHSTD].Printf("делаем паузу в %v секунд...", *pauseOfMarksMistake)
			duration := time.Second * time.Duration(*pauseOfMarksMistake)
			time.Sleep(duration)
			//подключаемся к ККТ
			logsmap[consttypes.LOGINFO_WITHSTD].Println("подключаемся к ККТ")
			if *kassatype == "atol" {
				_, err := connectToKKT(fptr, true)
				if err != nil {
					descrError := fmt.Sprintf("ошибка (%v) подключения к ККТ атол", err)
					logsmap[consttypes.LOGERROR].Println(descrError)
					globalErrorStr = descrError
					globalMistake = true
					break
				}
			} else {
				var err error
				logother := new(string)
				descrError := ""
				sessionkey, descrError, err = merc.CheckStatsuConnectionKKT(*IpMerc, *PortMerc, *comport, "", logother)
				if err != nil {
					logsmap[consttypes.LOGERROR].Printf("ошибка (%v) подлкючение к ККТ меркурий", descrError)
					if !*emulation {
						println("Нажмите любую клавишу...")
						input.Scan()
						log.Panic(descrError)
					}
				} else {
					logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе меркурий на порт %v прошло успешно. Ключ сессии: %v", *comport, sessionkey)
				}
				logginInFile(*logother)
				merc.Closesession(*IpMerc, *PortMerc, &sessionkey, logother)
			}
			//запускаем проверку марки заново
			logginInFile("снова запускаем проверку марки")
			receipt, existMarksInCheck, mistakeCheckingMark, markErroDescr, globalMistake, globalErrorStr = CheckAndRunsCheckingMarksByCheck(fptr, *kassatype, receipt, currFullFileName, true, *IpMerc, *PortMerc, *comport, sessionkey, *pauseOfMarksMistake)
		}
		//если были серьёзные ошибки при проверки прерываем программу
		if globalMistake {
			errorDescr := fmt.Sprintf("ошибка %v", globalErrorStr)
			logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			//amountOfMistakesMarks++
			break
		}
		//пропускаем чек, если были ошибки при проверке марки
		if mistakeCheckingMark {
			errorDescr := fmt.Sprintf("ошибка (%v) проверки марки для чека %v атол", markErroDescr, currFullFileName)
			logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			amountOfMistakesMarks++
			continue
		}
		//пересобираем json-задание, если необходимо (вставляем результаты проверки марок, изменяем параметры печати/не печати и email)
		wasChangeParametersOfCheck := false
		if *emailforcheck != "" {
			if receipt.ClientInfo.EmailOrPhone != *emailforcheck {
				receipt.ClientInfo.EmailOrPhone = *emailforcheck
				wasChangeParametersOfCheck = true
			}
		}
		if *PrintCheckOnKKT != "" {
			if printloc, err := getBoolFromString(*PrintCheckOnKKT, !receipt.Electronically); (!receipt.Electronically != printloc) && (err == nil) {
				receipt.Electronically = !printloc
				wasChangeParametersOfCheck = true
			}
		}
		if (existMarksInCheck) || (wasChangeParametersOfCheck) {
			jsonCorrWithMarkBytes, err := json.MarshalIndent(receipt, "", "\t")
			if err != nil {
				errorDescr := fmt.Sprintf("ошибка (%v) формирования json-а марками для задания чека %v атол", err, currFullFileName)
				logsmap[consttypes.LOGERROR].Println(errorDescr)
				amountOfMistakesChecks++
				continue
			}
			jsonCorrection = string(jsonCorrWithMarkBytes)
		}
		//печатаем чек
		logstr = fmt.Sprintf("послыем команду печати чека кассу json файл %v", jsonCorrection)
		logginInFile(logstr)
		resulOfCommand := ""
		if *kassatype == "atol" {
			resulOfCommand, err = sendComandeAndGetAnswerFromKKT(fptr, jsonCorrection)
		} else {
			//mercuriy //меркурий
			loginfoadditinal := new(string)
			resulOfCommand, err = merc.PrintCheck(*IpMerc, *PortMerc, *comport, receipt, loginfoadditinal)
			if *loginfoadditinal != "" {
				logginInFile(*loginfoadditinal)
			}
		}
		//если были ошибку при печати чека, то переходим к следующему заданию
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) печати чека %v атол", err, currFullFileName)
			logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			continue
		}
		logginInFile("послали команду печати чека кассу json файл")
		//эмулируем ошибку, если режим эмуляции ошибки включен
		if *emulatmistakes {
			logsmap[consttypes.LOGINFO_WITHSTD].Println("countPrintedChecks", countPrintedChecks)
			logsmap[consttypes.LOGINFO_WITHSTD].Println("countPrintedChecks%10", countPrintedChecks%10)
			if countPrintedChecks%10 == 0 {
				logginInFile("производим ошибку печати чека")
				logsmap[consttypes.LOGINFO_WITHSTD].Println("производим ошибку печати чека")
				resulOfCommand = "{\"result\": \"error - эмуляция ошибки\"}"
			}
		}
		//читаем информацию об результате выполнения команды
		//если команда выполнена успешно, то записываем в таблицу напечатанных чеков
		//если команда выполнена неуспешно, то проверяем не превышен ли количество чеков в смену,
		//и если превышено, то закрываем и открываем смену
		if successCommand(resulOfCommand) {
			//при успешной печати чека, записываем данные о номере напечатнного чека
			countPrintedChecks++
			if countPrintedChecks == 1 {
				file_printed_checks.WriteString("\n")
			}
			file_printed_checks.WriteString(currNumIsprChecka + "\n")
		} else {
			if strings.Contains(strings.ToUpper(resulOfCommand), strings.ToUpper("исчерпан")) {
				//закрываем, открываем смену
				checkOpenShift(fptr, true, lastNameOfKassir)
			}
			descrError := fmt.Sprintf("ошибка (%v) печати чека %v атол", resulOfCommand, currFullFileName)
			logsmap[consttypes.LOGERROR].Printf(descrError)
			logginInFile(descrError)
			amountOfMistakesChecks++
		}
		if amountOfMistakesChecks > 0 {
			logginInFile(fmt.Sprintln("количество не напечатанных чеков", amountOfMistakesChecks))
		}
	} //перебор json заданий
	//выводим информацию об количестве напечтатнных чеков
	logsmap[consttypes.LOGINFO_WITHSTD].Printf("распечатно %v из %v чеков", countPrintedChecks, countOfFiles)
	println("Нажмите любую клавишу...")
	input.Scan()
}

func dialogContinuePrintChecks() (bool, string) {
	res := true
	command := ""
	input := bufio.NewScanner(os.Stdin)
	println("Продолжить (да - продолжить печать чеков, нет (по умолчанию) - завершить программу, \"off/on\" - переподключиться к кассе):")
	input.Scan()
	if input.Text() == "off/on" {
		command = "off/on"
	}
	if input.Text() == "" {
		res = false
	} else {
		res, _ = getBoolFromString(input.Text(), res)
	}
	return res, command
}

func CheckAndRunsCheckingMarksByCheck(fptr *fptr10.IFptr, kassatype string, receipt consttypes.TCorrectionCheck, FullFileName string, perezapuskatproverku bool, ipktt string, port int, comport int, sessionkey string, pausetimesec int) (consttypes.TCorrectionCheck, bool, bool, string, bool, string) {
	logginInFile("ищем марки в чеке")
	//читаем данные по маркам
	mistakeCheckingMark := false
	errorDescr := ""
	existMarksInCheck := false
	globalMistake := false
	globalErrorStr := ""
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
		logstr := fmt.Sprintf("запускаем процесс проверки марки %v для чека %v", currMarkBase64, FullFileName)
		logsmap[consttypes.LOGINFO_WITHSTD].Println(logstr)
		imcResultCheckin, errproc := runProcessCheckMark(kassatype, fptr, ipktt, port, sessionkey, currMarkBase64)
		if *emulatmistakesmarks {
			errproc = errors.New("симуляция ошибки провекри марки")
		}
		if errproc != nil {
			//прерываем проверку
			if perezapuskatproverku {
				errorDescr = fmt.Sprintf("будет произведен перезпуск провекри марки, так как была ошибка (%v) запуска проверки марки %v для чека %v атол", errproc, currMarkBase64, FullFileName)
				logsmap[consttypes.LOGINFO_WITHSTD].Println(errorDescr)
				//***********************************
				logginInFile(fmt.Sprintf("перезапускаем проверку марки %v...", currMarkBase64))
				logsmap[consttypes.LOGINFO_WITHSTD].Println("перезапускаем проверку марки...")
				//прерываем предыдущую провекру марки
				if kassatype == "atol" {
					breakProcCheckOfMark(fptr)
				} else {
					loginfoother := new(string)
					merc.BreakProcCheckOfMark(ipktt, port, comport, sessionkey, loginfoother)
					logginInFile(*loginfoother)
				}
				//отключаемся от ККТ
				logsmap[consttypes.LOGINFO_WITHSTD].Println("отключаемся от ККТ")
				if kassatype == "atol" {
					disconnectWithKKT(fptr, true)
				} else {
					loginfoother := new(string)
					merc.DissconnectMeruriy(ipktt, port, sessionkey, loginfoother)
					logginInFile(*loginfoother)
				}
				//делаем паузу
				logsmap[consttypes.LOGINFO_WITHSTD].Printf("пауза в %v секунд...", pausetimesec)
				duration := time.Second * time.Duration(pausetimesec)
				time.Sleep(duration)
				//подключаемся к ККТ
				logsmap[consttypes.LOGINFO_WITHSTD].Println("подлкючаемся к ККТ")
				if kassatype == "atol" {
					_, err := connectToKKT(fptr, true)
					if err != nil {
						descrError := fmt.Sprintf("ошибка (%v) подключения к ККТ атол", err)
						logsmap[consttypes.LOGERROR].Println(descrError)
						globalErrorStr = descrError
						globalMistake = true
						break
					}
				} else {
					var err error
					logother := new(string)
					descrError := ""
					sessionkey, descrError, err = merc.CheckStatsuConnectionKKT(*IpMerc, *PortMerc, comport, "", logother)
					if err != nil {
						descrError := fmt.Sprintf("ошибка (%v) подлкючение к ККТ меркурий", descrError)
						logsmap[consttypes.LOGERROR].Printf(descrError)
						globalErrorStr = descrError
						globalMistake = true
						break
					}
					logginInFile(*logother)
					merc.Closesession(*IpMerc, *PortMerc, &sessionkey, logother)
				}
				//запускаем проверку марки заново
				logginInFile("снова запускаем проверку марки")
				imcResultCheckin, errproc = runProcessCheckMark(kassatype, fptr, ipktt, port, sessionkey, currMarkBase64)
			} //перезапуск провекри марки
			if errproc != nil {
				errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v для чека %v атол", errproc, currMarkBase64, FullFileName)
				logsmap[consttypes.LOGERROR].Println(errorDescr)
				mistakeCheckingMark = true
				break
			}
		}
		logsmap[consttypes.LOGINFO_WITHSTD].Println("марка успешно проверена")
		//заполняем данные о марке
		logginInFile("заполняем данные о марке")
		ImcParams := LocImcParams.(map[string]interface{})
		ImcParams["imc"] = currMarkBase64
		ImcParams["imcModeProcessing"] = 0
		ImcParams["itemEstimatedStatus"] = "itemStatusUnchanged"
		ImcParams["imcType"] = "auto"
		ImcParams["itemInfoCheckResult"] = new(consttypes.TItemInfoCheckResult)
		ItemInfoCheckResult := ImcParams["itemInfoCheckResult"].(*consttypes.TItemInfoCheckResult)
		ItemInfoCheckResult.ImcCheckFlag = imcResultCheckin.ImcCheckFlag
		ItemInfoCheckResult.ImcCheckResult = imcResultCheckin.ImcCheckResult
		ItemInfoCheckResult.ImcStatusInfo = imcResultCheckin.ImcStatusInfo
		ItemInfoCheckResult.ImcEstimatedStatusCorrect = imcResultCheckin.ImcEstimatedStatusCorrect
	}
	return receipt, existMarksInCheck, mistakeCheckingMark, errorDescr, globalMistake, globalErrorStr
} //CheckAndRunsCheckingMarksByCheck

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
		if !*emulation {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды %v на кассе", err, comJson)
			logsmap[consttypes.LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры sendComandeAndGetAnswerFromKKT c ошибкой", err)
			logginInFile(logstr)
			return desrError, err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	if strings.Contains(result, "Нет связи") {
		logginInFile("нет связи: переподключаемся")
		if ok, typepodkluch := connectWithKassa(fptr, *comport, *ipaddresskkt, *portkktatol, *ipaddressservrkkt); !ok {
			descrErr := fmt.Sprintf("ошибка соединения с кассовым аппаратом %v", typepodkluch)
			logsmap[consttypes.LOGERROR].Println(descrErr)
			if !*emulation {
				println("Нажмите любую клавишу...")
				//input.Scan()
				log.Panic(descrErr)
			}
		} else {
			logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", *comport)
		}
	}
	//logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmap[LOGOTHER].Println("result", result)
	//if *emulation {
	//}
	//result := "{\"result\": \"all ok\"}"

	logginInFile("конец процедуры sendComandeAndGetAnswerFromKKT без ошибки")
	return result, nil
}

func runProcessCheckMark(kassatype string, fptr *fptr10.IFptr, ipktt string, port int, sessionkey string, mark string) (consttypes.TItemInfoCheckResult, error) {
	var countAttempts int
	var imcResultCheckinObj consttypes.TItemInfoCheckResultObject
	var imcResultCheckin consttypes.TItemInfoCheckResult
	logginInFile("начало процедуры runProcessCheckMark")
	//проверяем - открыта ли смена
	//shiftOpenned, err := checkOpenShift(fptr, true, "админ")
	//if err != nil {
	//	errorDescr := fmt.Sprintf("ошибка (%v). Смена не открыта", err)
	//	logsmap[consttypes.LOGERROR].Println(errorDescr)
	//	return TItemInfoCheckResult{}, errors.New(errorDescr)
	//}
	//if !shiftOpenned {
	//	errorDescr := fmt.Sprintf("ошибка (%v) - смена не открыта", err)
	//	logsmap[consttypes.LOGERROR].Println(errorDescr)
	//	return TItemInfoCheckResult{}, errors.New(errorDescr)
	//}
	//посылаем запрос на проверку марки
	var resJson string
	var err error
	if kassatype == "atol" {
		resJson, err = sendCheckOfMark(fptr, mark)
	} else {
		var resMercAnswerBytes []byte
		var answerMerc consttypes.TAnswerMercur
		resMercAnswerBytes, err = merc.SendCheckOfMark(ipktt, port, sessionkey, mark, 0)
		if err == nil {
			err = json.Unmarshal(resMercAnswerBytes, &answerMerc)
			if err == nil {
				resJson = answerMerc.Description
				if answerMerc.Result != 0 {
					resJson = "error " + resJson
				}
			}
		}
	}
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v", err, mark)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	if !successCommand(resJson) {
		errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v", resJson, mark)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	for countAttempts = 0; countAttempts < *countOfCheckingMarks; countAttempts++ {
		var answerOfCheckMark consttypes.TAnsweChekcMark
		var MercurAnswerOfCheckMark consttypes.TAnswerMercur
		if kassatype == "atol" {
			resJson, err = getStatusOfChecking(fptr)
		} else {
			var resMercAnswerBytes []byte
			resMercAnswerBytes, err = merc.GetStatusOfChecking(ipktt, port, sessionkey)
			if err == nil {
				err = json.Unmarshal(resMercAnswerBytes, &MercurAnswerOfCheckMark)
				if err == nil {
					resJson = MercurAnswerOfCheckMark.Description
					if MercurAnswerOfCheckMark.Result != 0 {
						resJson = "error " + resJson
					}
				}
			}
		}
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) получения статуса проверки марки %v", err, mark)
			logsmap[consttypes.LOGERROR].Println(errorDescr)
			return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
		}
		if !successCommand(resJson) {
			//делаем паузу
			logginInFile(resJson)
			desrAction := fmt.Sprintf("пауза в %v секунд... так сервер провекри марок не успевает.", *pauseOfMarksMistake)
			logsmap[consttypes.LOGINFO_WITHSTD].Println(desrAction)
			duration := time.Second * time.Duration(*pauseOfMarksMistake)
			time.Sleep(duration)
			//if strings.Contains(resJson, "421")
			//errorDescr := fmt.Sprintf("ошибка (%v) получения статуса проверки марки %v", resJson, mark)
			//logsmap[consttypes.LOGERROR].Println(errorDescr)
			//return TItemInfoCheckResult{}, errors.New(errorDescr)
		}
		err = nil
		if kassatype == "atol" {
			err = json.Unmarshal([]byte(resJson), &answerOfCheckMark)
		}
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) ответа проверки марки %v", err, resJson, mark)
			logsmap[consttypes.LOGERROR].Println(errorDescr)
			return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
		}
		if kassatype != "atol" {
			answerOfCheckMark.Ready = MercurAnswerOfCheckMark.IsCompleted
		}
		if answerOfCheckMark.Ready {
			if (*emulation) && (countAttempts < *countOfCheckingMarks-20) {
				//емулируем задержку полчение марки
			} else {
				break
			}
		}
		//пауза в 1 секунду
		logsmap[consttypes.LOGINFO_WITHSTD].Printf("попытка %v из %v получения статуса марки", countAttempts+2, *countOfCheckingMarks)
		duration := time.Second
		time.Sleep(duration)
	}
	if countAttempts == *countOfCheckingMarks {
		errorDescr := fmt.Sprintf("ошибка проверки марки %v", mark)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	//принимаем марку
	var resOfChecking string
	var MercurAnswerOfResultOfCheckMark consttypes.TAnswerMercur
	if kassatype == "atol" {
		resOfChecking, err = acceptMark(fptr)
	} else {
		var resMercAnswerBytes []byte
		resMercAnswerBytes, err = merc.AcceptMark(ipktt, port, sessionkey)
		if err == nil {
			err = json.Unmarshal(resMercAnswerBytes, &MercurAnswerOfResultOfCheckMark)
			if err == nil {
				resOfChecking = MercurAnswerOfResultOfCheckMark.Description
				if MercurAnswerOfResultOfCheckMark.Result != 0 {
					resOfChecking = "error " + resJson
				}
			}
		}
	}
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) принятия марки %v", err, mark)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	if !successCommand(resOfChecking) {
		errorDescr := fmt.Sprintf("ошибка (%v) принятия марки %v", resOfChecking, mark)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	err = nil
	if kassatype == "atol" {
		err = json.Unmarshal([]byte(resOfChecking), &imcResultCheckinObj)
	}
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) ответа проверки марки %v", err, resOfChecking, mark)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	if kassatype == "atol" {
		imcResultCheckin.EcrStandAloneFlag = imcResultCheckinObj.ItemInfoCheckResult.EcrStandAloneFlag
		imcResultCheckin.ImcCheckFlag = imcResultCheckinObj.ItemInfoCheckResult.ImcCheckFlag
		imcResultCheckin.ImcCheckResult = imcResultCheckinObj.ItemInfoCheckResult.ImcCheckResult
		imcResultCheckin.ImcEstimatedStatusCorrect = imcResultCheckinObj.ItemInfoCheckResult.ImcEstimatedStatusCorrect
		imcResultCheckin.ImcStatusInfo = imcResultCheckinObj.ItemInfoCheckResult.ImcStatusInfo
	}
	logginInFile("конец процедуры runProcessCheckMark без ошибки")
	return imcResultCheckin, nil
} //runProcessCheckMark

func breakProcCheckOfMark(fptr *fptr10.IFptr) error {
	var err error
	logginInFile("прерываемм проверку марки")
	turnCheckMarkJson := "{\"type\": \"cancelMarkingCodeValidation\"}"
	resturnCheckMark, _ := sendComandeAndGetAnswerFromKKT(fptr, turnCheckMarkJson)
	logginInFile(fmt.Sprintf("результат прерывания проверки марки: %v", resturnCheckMark))
	return err
}

func clearTanlesOfMarks(fptr *fptr10.IFptr) error {
	var err error
	logginInFile("очищаем таблицу марок")
	clearTableMarksJson := "{\"type\": \"clearMarkingCodeValidationResult\"}"
	resClearTableMarks, _ := sendComandeAndGetAnswerFromKKT(fptr, clearTableMarksJson)
	logginInFile(fmt.Sprintf("результат очистки таблицы марок: %v", resClearTableMarks))
	return err
}

func sendCheckOfMark(fptr *fptr10.IFptr, mark string) (string, error) {
	var err error
	logginInFile("начало процедуры sendCheckOfMark")
	//return "", nil
	comJson, err := getJsonOfBeginCheck(mark)
	if err != nil {
		desrError := fmt.Sprintf("ошибка (%v) фоормирования json задания проверки марки", err)
		logsmap[consttypes.LOGERROR].Println(desrError)
		logstr := fmt.Sprint("конец процедуры sendCheckOfMark c ошибкой", err)
		logginInFile(logstr)
		return "", err
	}
	//comJson = "{\"type\": \"reportX\",\"operator\": {\"name\": \"Иванов\",\"vatin\": \"123654789507\"}}"
	logstr := fmt.Sprintf("отправляем запрос (%v) на проверку марки", comJson)
	logginInFile(logstr)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*emulation {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды начать проверку марки на кассовом аппарате", err)
			logsmap[consttypes.LOGERROR].Println(desrError)
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
	//if *emulation {
	//}
	//result := "{\"result\": \"all ok\"}"
	logginInFile("конец процедуры sendCheckOfMark без ошибки")
	return result, nil
}

func getJsonOfBeginCheck(mark string) (string, error) {
	var begMarkStructur consttypes.TBeginTaskMarkCheck
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
		logsmap[consttypes.LOGERROR].Println(desrError)
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
		if !*emulation {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды получения статуса марки на кассовом аппарате", err)
			logsmap[consttypes.LOGERROR].Println(desrError)
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
	//if *emulation {
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
		if !*emulation {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды принятия марки на кассовом аппарате", err)
			logsmap[consttypes.LOGERROR].Println(desrError)
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
	//if *emulation {
	//}
	logstr := fmt.Sprintf("результат проверки марки (%v) ", result)
	logginInFile(logstr)
	//result := "{\"result\": \"all ok\"}"
	logginInFile("конец процедуры acceptMark без ошибки")
	return result, nil
}

func listDirByReadDir(path string) ([]string, error) {
	var spisFiles []string
	var spisFileFD []int
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
		if consttypes.FILE_NAME_PRINTED_CHECKS == val.Name() {
			logsmap[consttypes.LOGINFO_WITHSTD].Println("пропускаем файл с ифнормацией о напечатнных чеках")
			continue
		}
		if consttypes.FILE_NAME_CONNECTION == val.Name() {
			logsmap[consttypes.LOGINFO_WITHSTD].Println("пропускаем файл с ифнормацией о настройки связи с ККТ")
			continue
		}
		/*logsmap[consttypes.LOGINFO].Println(val.Name())
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
		//logstr = fmt.Sprintln("matched=", matched)
		//logginInFile(logstr)
		if matched {
			spisFiles = append(spisFiles, val.Name())
		}
	}
	logginInFile(fmt.Sprintln("spisFiles=", spisFiles))
	for _, filename := range spisFiles {
		fdstr := getFDFromFileName(filename)
		fdint, err := strconv.Atoi(fdstr)
		if err != nil {
			logginInFile(fmt.Sprintf("ошибка (%v) получения номера ФД из имени файла %v при сортировке списка файлов по номеру ФД", err, filename))
			return spisFiles, err
			//continue
		}
		spisFileFD = append(spisFileFD, fdint)
	}
	var spisResOfFiles []string
	sort.Ints(spisFileFD)
	for _, fdint := range spisFileFD {
		for _, filename := range spisFiles {
			fdstr := getFDFromFileName(filename)
			fdintFile, err := strconv.Atoi(fdstr)
			if err != nil {
				logginInFile(fmt.Sprintf("ошибка (%v) получения номера ФД из имени файла %v при сортировке списка файлов по номеру ФД", err, filename))
				return spisFiles, err
				//continue
			}
			if fdint == fdintFile {
				spisResOfFiles = append(spisResOfFiles, filename)
			}
		}
	}
	return spisResOfFiles, nil
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
		preflog := consttypes.LOG_PREFIX + "_" + strings.ToUpper(logstr)
		fullnamelogfile := consttypes.LOGSDIR + filenamelogfile
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
	file_printed_checks, err := os.Open(dirjsons + consttypes.FILE_NAME_PRINTED_CHECKS)
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
	//logstr = fmt.Sprintln("plan", plan)
	//logginInFile(logstr)
	//logstr = fmt.Sprintln("error", err)
	//logginInFile(logstr)
	//logstr = fmt.Sprintln("конец процедуры readJsonFromFile")
	//logginInFile(logstr)
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

func connectWithKassa(fptr *fptr10.IFptr, comportint int, ipaddresskktper string, portkktper int, ipaddresssrvkktper string) (bool, string) {
	//if !strings.Contains(comport, "COM") {
	//	sComPorta = "COM" + comport
	//}
	typeConnect := ""
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	if ipaddresssrvkktper != "" {
		fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_REMOTE_SERVER_ADDR, ipaddresssrvkktper)
		typeConnect = fmt.Sprintf("через сервер ККТ по IP %v", ipaddresssrvkktper)
	}
	if comportint == 0 {
		if ipaddresskktper != "" {
			fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_TCPIP))
			fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_IPADDRESS, ipaddresskktper)
			typeConnect = fmt.Sprintf("%v по IP %v ККТ на порт %v", typeConnect, ipaddresskktper, portkktper)
			if portkktper != 0 {
				fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_IPPORT, strconv.Itoa(portkktper))
			}
		} else {
			fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
			typeConnect = fmt.Sprintf("%v по USB", typeConnect)
		}
	} else {
		sComPorta := "COM" + strconv.Itoa(comportint)
		typeConnect = fmt.Sprintf("%v по COM порту %v", typeConnect, sComPorta)
		fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_COM))
		fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_COM_FILE, sComPorta)
		fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_BAUDRATE, strconv.Itoa(fptr10.LIBFPTR_PORT_BR_115200))
	}
	fptr.ApplySingleSettings()
	fptr.Open()
	return fptr.IsOpened(), typeConnect
}

func getCurrentPortOfKass(dirOfJsons string) (int, error) {
	if *comport > 0 {
		return *comport, nil
	}
	comportb, err := os.ReadFile(dirOfJsons + consttypes.FILE_NAME_CONNECTION)
	if err != nil {
		desrError := fmt.Sprintf("ошибка (%v) октрытия файла с параметрами соедиения кассы", err)
		logginInFile(desrError)
		//logsmap[consttypes.LOGERROR].Println(desrError)
		return 0, nil
	}
	comportstr := string(comportb)
	if strings.Contains(comportstr, "COM") {
		indOfCom := strings.Index(comportstr, "COM")
		comportstr = comportstr[indOfCom:]
	}
	comportint, _ := strconv.Atoi(comportstr)
	return comportint, nil
}

func logginInFile(loggin string) {
	if (*LogsDebugs) > 0 {
		logsmap[consttypes.LOGINFO].Println(loggin)
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
	if pref == consttypes.LOG_PREFIX+"_ERROR" {
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

func connectToKKT(fptr *fptr10.IFptr, createComObj bool) (string, error) {
	var err error
	logginInFile("снова создаём объект драйвера...")
	if createComObj {
		fptr, err = fptr10.NewSafe()
	}
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) инициализации драйвера ККТ атол", err)
		logsmap[consttypes.LOGERROR].Println(descrError)
		return descrError, errors.New(descrError)
	}
	//сединение с кассой
	logsmap[consttypes.LOGINFO_WITHSTD].Println("соединение с кассой...")
	*comport, _ = getCurrentPortOfKass(consttypes.DIROFJSONS)
	if ok, typepodkluch := connectWithKassa(fptr, *comport, *ipaddresskkt, *portkktatol, *ipaddressservrkkt); !ok {
		descrErr := fmt.Sprintf("ошибка сокдинения с кассовым аппаратом %v", typepodkluch)
		logsmap[consttypes.LOGERROR].Println(descrErr)
		if !*emulation {
			return descrErr, errors.New(descrErr)
		}
	} else {
		logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", *comport)
	}
	return "", nil
}

func disconnectWithKKT(fptr *fptr10.IFptr, destroyComObject bool) {
	fptr.Close()
	if destroyComObject {
		fptr.Destroy()
	}
}

func reconnectToKKT(fptr *fptr10.IFptr) error {
	fptr.Close()
	fptr.Destroy()
	fptr, err := fptr10.NewSafe()
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) инициализации драйвера ККТ атол", err)
		logsmap[consttypes.LOGERROR].Println(descrError)
		return errors.New(descrError)
		//println("Нажмите любую клавишу...")
		//input.Scan()
		//log.Panic(descrError)
	}
	//defer fptr.Destroy()
	fmt.Println(fptr.Version())
	//сединение с кассой
	logsmap[consttypes.LOGINFO_WITHSTD].Println("соединение с кассой")
	*comport, _ = getCurrentPortOfKass(consttypes.DIROFJSONS)
	//if err != nil {
	//	desrErr := fmt.Sprintf("ошибка (%v) чтения параметра com порт соединения с кассой", err)
	//	logsmap[consttypes.LOGERROR].Println(desrErr)
	//	return errors.New(desrErr)
	//	//println("Нажмите любую клавишу...")
	//	//input.Scan()
	//	//log.Panic(desrErr)
	//}
	if ok, typepodkluch := connectWithKassa(fptr, *comport, *ipaddresskkt, *portkktatol, *ipaddressservrkkt); !ok {
		//if !connectWithKassa(fptr, *comport, *ipaddresskkt, *ipaddressservrkkt) {
		descrErr := fmt.Sprintf("ошибка сокдинения с кассовым аппаратом %v", typepodkluch)
		logsmap[consttypes.LOGERROR].Println(descrErr)
		if !*emulation {
			return errors.New(descrErr)
			//println("Нажмите любую клавишу...")
			//input.Scan()
			//log.Panic(descrErr)
		}
	} else {
		logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", *comport)
	}
	return nil
	//defer fptr.Close()
}

func checkOpenShift(fptr *fptr10.IFptr, openShiftIfClose bool, kassir string) (bool, error) {
	logginInFile("получаем статус ККТ")
	getStatusKKTJson := "{\"type\": \"getDeviceStatus\"}"
	resgetStatusKKT, err := sendComandeAndGetAnswerFromKKT(fptr, getStatusKKTJson)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) получения статуса кассы", err)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return false, err
	}
	if !successCommand(resgetStatusKKT) {
		errorDescr := fmt.Sprintf("ошибка (%v) получения статуса кассы", resgetStatusKKT)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		//logginInFile(errorDescr)
		return false, errors.New(errorDescr)
	}
	logginInFile("получили статус кассы")
	//проверяем - открыта ли смена
	var answerOfGetStatusofShift consttypes.TAnswerGetStatusOfShift
	err = json.Unmarshal([]byte(resgetStatusKKT), &answerOfGetStatusofShift)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) распарсивания статуса кассы", err)
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return false, err
	}
	if answerOfGetStatusofShift.ShiftStatus.State == "expired" {
		errorDescr := "ошибка - смена на кассе уже истекла. Закройте смену"
		logsmap[consttypes.LOGERROR].Println(errorDescr)
		return false, errors.New(errorDescr)
	}
	if answerOfGetStatusofShift.ShiftStatus.State == "closed" {
		if openShiftIfClose {
			if kassir == "" {
				errorDescr := "не указано имя кассира для открытия смены"
				logsmap[consttypes.LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
			jsonOpenShift := fmt.Sprintf("{\"type\": \"openShift\",\"operator\": {\"name\": \"%v\"}}", kassir)
			resOpenShift, err := sendComandeAndGetAnswerFromKKT(fptr, jsonOpenShift)
			if err != nil {
				errorDescr := fmt.Sprintf("ошбика (%v) - не удалось открыть смену", err)
				logsmap[consttypes.LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
			if !successCommand(resOpenShift) {
				errorDescr := fmt.Sprintf("ошбика (%v) - не удалось открыть смену", resOpenShift)
				logsmap[consttypes.LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
		} else {
			return false, nil
		}
	}
	return true, nil
} //checkOpenShift
