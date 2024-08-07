//go:generate ./resource/goversioninfo.exe -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"bufio"
	consttypes "clientrabbit/consttypes"
	fptr10 "clientrabbit/fptr"
	logsmy "clientrabbit/packetlog"
	merc "clientrabbit/sendtcp"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
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
var userint = flag.Int("UserMerc", 0, "номер пользователя в кассе меркрий")
var passwuser = flag.String("PaswUserMerc", "", "пароль пользователя меркурий")
var dirOfjsons = flag.String("dirjsons", ".\\jsons\\works\\", "директория json файлов по умолчанию ./jsons/")
var clearLogsProgramm = flag.Bool("clearlogs", true, "очистить логи программы")
var countChecksForPause = flag.Int("countforpause", 0, "число чеков, после которых программа делает небольшую паузу")
var pauseInSecondsMayClose = flag.Int("secpause", 10, "сколько секунд на паузу, если вдруг надо завершить программу")
var LogsDebugs = flag.Int("debug", 0, "уровень логирования всех действий, чем выше тем больше логов")
var comport = flag.Int("com", 0, "ком порт кассы")
var emailforcheck = flag.String("email", "", "email клиента чека")
var PrintCheckOnKKT = flag.String("print", "", "печтать или не печатать чек на ККТ: true - печатать, false - не печетать")
var CassirName = flag.String("cassir", "", "имя кассира")
var ipaddresskkt = flag.String("ipkkt", "", "ip адрес ккт")
var portkktatol = flag.Int("portipkkt", 0, "порт ip ккт")
var ipaddressservrkkt = flag.String("ipservkkt", "", "ip адрес сервера ккт")
var emulation = flag.Bool("emul", false, "эмуляция")
var dontprintrealfortest = flag.Bool("test", false, "тест - не печатать реальный чек")
var emulatmistakes = flag.Bool("emulmist", false, "эмуляция ошибок")
var emulatmistakesOpenCheck = flag.Bool("emulmistopencheck", false, "эмуляция ошибок открытия чека")
var emulatmistakesmarks = flag.Bool("emulmistmark", false, "эмуляция ошибок марок")
var countOfCheckingMarks = flag.Int("attempts", 20, "число попыток провекри марки")
var clearTableOfMarks = flag.Bool("clearmarks", true, "очищать таблицу марок перед запуском на ККТ нового чека")
var countOfMistakesCheckForStop = flag.Int("stop_mist", 3, "число ошибочных чеков, после которого останавливать программу")
var pauseOfMarksMistake = flag.Int("pause_mist", 10, "пауза между проблемами с марками")
var conversChekcCorrectionsType = flag.Bool("converse", false, "для всех чеков бить чеки коррекции сторнирующий")
var changeCashOnBeznal = flag.Bool("cashtobeznal", false, "поменять нал на безнал")

var countPrintChecks = flag.Int("countchecks", 0, "число успешно распечатнных чеков, после которого остановить программу")
var pauseAfterDay = flag.Int("pauseAfterDay", 0, "число дней, после которого программа делает паузу")
var pauseInSecondsAfterDay = flag.Int("pausefterdaysec", 90, "пауза в секундах после звершение какого-то количества дней напечатнных чеков")

var ExlusionDate = flag.String("exldate", "", "дата исключения из распечатки в формате 2006.01.02")

const Version_of_program = "2024_07_30_02"

func main() {
	var err error
	var ExlusionDateDate time.Time
	var lastNameOfKassir string
	var fptr *fptr10.IFptr
	var sessionkey string
	mercSNODefault := -1
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
	//закрывес соединение с кассой меркурий если было установлено
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
	input := bufio.NewScanner(os.Stdin)
	descrMistake, err := logsmy.InitializationsLogs(*clearLogsProgramm, *LogsDebugs)
	defer logsmy.CloseDescrptorsLogs()
	if err != nil {
		fmt.Fprint(os.Stderr, descrMistake)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic(descrMistake)
	}
	logsmy.LogginInFile(runDescription)
	logsmy.LogginInFile(clearLogsDescr)
	//ищем все файлы заданий в директории json - заданий
	consttypes.DIROFJSONS = *dirOfjsons
	if foundedLogDir, _ := consttypes.DoesFileExist(consttypes.DIROFJSONS); !foundedLogDir {
		err := os.Mkdir(consttypes.DIROFJSONS, 0777)
		descrError := fmt.Sprintf("ошибка (%v) чтения директории %v с json заданиямию", err, consttypes.DIROFJSONS)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
		println("Нажмите любую клавишу...")
		input.Scan()
		log.Panic(descrError)
	}
	listOfFilesTempr, err := listDirByReadDir(consttypes.DIROFJSONS)
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) поиска json заданий в директории %v", err, consttypes.DIROFJSONS)
		logsmy.Logsmap[consttypes.LOGERROR].Printf(descrError)
		log.Panic(descrError)
	}
	logsmy.LogginInFile(fmt.Sprintln("listOfFilesTempr=", listOfFilesTempr))
	var listOfFiles []string
	countOfFiles := len(listOfFilesTempr)
	logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("Всего json файлов", countOfFiles)
	//убираем json-задания, которые уже были распечатаны
	for _, v := range listOfFilesTempr {
		currFullFileName := consttypes.DIROFJSONS + v
		numChecka := getFDFromFileName(v)
		printedThisCheck := false
		if numChecka == "" {
			logsmy.Logsmap[consttypes.LOGERROR].Printf("пропущен файл %v", currFullFileName)
			continue
		}
		printedThisCheck, _ = printedCheck(consttypes.DIROFJSONS, numChecka)
		if printedThisCheck {
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("чек с номером %v уже был распечатан", numChecka)
			continue
		}
		listOfFiles = append(listOfFiles, currFullFileName)
		//logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("%v = %v\n", k+1, currFullFileName)
	}
	countOfFiles = len(listOfFiles)
	//читаем настроку com - порта в директории json - заданий
	*comport, _ = getCurrentPortOfKass(consttypes.DIROFJSONS)
	logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("порт кассы", *comport)
	//подключаемся к кассовому аппарату
	if *kassatype == "atol" {
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("Тип кассы atol")
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("инициализация драйвера атол")
		fptr, err = fptr10.NewSafe()
		if err != nil {
			descrError := fmt.Sprintf("Ошибка (%v) инициализации драйвера ККТ атол", err)
			logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
			println("Нажмите любую клавишу...")
			input.Scan()
			log.Panic(descrError)
		}
		defer fptr.Destroy()
		fmt.Println(fptr.Version())
		//сединение с кассой
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("соединение с кассой")
		if ok, typepodkluch := connectWithKassa(fptr, *comport, *ipaddresskkt, *portkktatol, *ipaddressservrkkt); !ok {
			descrErr := fmt.Sprintf("ошибка соединения с кассовым аппаратом %v", typepodkluch)
			logsmy.Logsmap[consttypes.LOGERROR].Println(descrErr)
			if !*emulation {
				println("Нажмите любую клавишу...")
				input.Scan()
				log.Panic(descrErr)
			}
		} else {
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", *comport)
		}
		defer fptr.Close()
	} else {
		var err error
		descrError := ""
		sessionkey, descrError, err = merc.CheckStatsuConnectionKKT(*emulation, *IpMerc, *PortMerc, *comport, "", *userint, *passwuser)
		if err != nil {
			logsmy.Logsmap[consttypes.LOGERROR].Printf("ошибка (%v) подлкючение к ККТ меркурий", descrError)
			if !*emulation {
				println("Нажмите любую клавишу...")
				input.Scan()
				log.Panic(descrError)
			} else {
				mercSNODefault, err = merc.GetSNOByDefault(*emulation, *IpMerc, *PortMerc, sessionkey)
				if err != nil {
					mercSNODefault = -1
				}
			}
		} else {
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе меркурий на порт %v прошло успешно. Ключ сессии: %v", *comport, sessionkey)
			mercSNODefault, err = merc.GetSNOByDefault(*emulation, *IpMerc, *PortMerc, sessionkey)
			if err != nil {
				mercSNODefault = -1
			}
		}
		merc.Closesession(*IpMerc, *PortMerc, &sessionkey)
	}
	//открытие для запиписи файла напечатанных чеков
	logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("отрытие для записи таблицы напечатанных чеков")
	flagsTempOpen := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file_printed_checks, err := os.OpenFile(consttypes.DIROFJSONS+consttypes.FILE_NAME_PRINTED_CHECKS, flagsTempOpen, 0644)
	if err != nil {
		descrError := fmt.Sprintf("ошибка создания файла напечатанных чеков %v", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
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
		logsmy.Logsmap[consttypes.LOGERROR].Printf("ошибка (%v) инициализации начальной даты", err)
	}
	prevDateOfCheck := initDate
	logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("начинаем выполнять json чеков", countOfFiles)
	logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("всего json заданий для печати чека", countOfFiles)
	previusWasMarks := false
	if *ExlusionDate != "" {
		ExlusionDateDate, err = time.Parse("2006.01.02", *ExlusionDate)
		if err != nil {
			logsmy.Logsmap[consttypes.LOGERROR].Printf("ошибка (%v) инициализации даты исключения", err)
			*ExlusionDate = ""
		}
	}
	//цикл перебора json-заданий
	for k, currFullFileName := range listOfFiles {
		var receipt consttypes.TCorrectionCheck
		//для кассы меркурий полчаем sessionkey
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
		//
		//для кассы меркурий полчаем sessionkey
		if *kassatype == "merc" {
			var descrError string
			var err error
			if sessionkey != "" {
				logsmy.LogginInFile(fmt.Sprintln("session key before", sessionkey))
				merc.Closesession(*IpMerc, *PortMerc, &sessionkey)
				logsmy.LogginInFile(fmt.Sprintln("session key after", sessionkey))
			}
			sessionkey, descrError, err = merc.CheckStatsuConnectionKKT(*emulation, *IpMerc, *PortMerc, *comport, "", *userint, *passwuser)
			if err != nil {
				logsmy.Logsmap[consttypes.LOGERROR].Printf("ошибка (%v) подлкючение к ККТ меркурий", descrError)
				if !*emulation {
					descrError := "ну удалось полчить ключ сессии ккт меркурий"
					logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
					break
				}
			}
		}
		//инициализируем переменные шага цикла
		globalMistake := false
		globalErrorStr := ""
		command := ""
		//проверяем условия выхода из цикла
		if amountOfMistakesChecks >= *countOfMistakesCheckForStop {
			descrError := "превышено количество ошибок чеков, остановка работы программы"
			logsmy.LogginInFile(descrError)
			resDial := false
			resDial, command = dialogContinuePrintChecks()
			if !resDial && (command != "off/on") {
				descrError := "работы программы прервана пользователем"
				logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
				break
			}
			amountOfMistakesChecks = 0
			amountOfMistakesMarks = 0
		} else {
			if *countPrintChecks > 0 {
				if countPrintedChecks >= *countPrintChecks {
					desctriptionExit := fmt.Sprintf("произошло завершение работы программы, так как число напечатнных чеков %v равно параметру countchecks, переданному при запуске программы", countPrintedChecks)
					logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println(desctriptionExit)
					break //прерываем печать чека
				}
			}
			if *countChecksForPause > 0 {
				if ((countPrintedChecks + 1) % *countChecksForPause) == 0 {
					logsmy.LogginInFile(fmt.Sprintf("делаем паузу в программе через каждые %v чеков для возможной безопасной её остановки на %v секунд...", *countChecksForPause, *pauseInSecondsMayClose))
					logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("если процесс печати чеков нужно прервать, то это можно сделать сейчас")
					duration := time.Second * time.Duration((*pauseInSecondsMayClose))
					time.Sleep(duration)
				}
			}
		}
		if command != "" {
			logsmy.LogginInFile(fmt.Sprintln("command", command))
		}
		if command == "off/on" {
			command = ""
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("переподключение к кассовому аппарату...")
			if *kassatype == "atol" {
				err := reconnectToKKT(fptr)
				if err != nil {
					logsmy.Logsmap[consttypes.LOGERROR].Printf("ошибка переподключения к ККТ %v", err)
					break
				}
			}
		}
		//читаем json - задание
		currNumIsprChecka := getFDFromFileName(currFullFileName)
		logsmy.LogginInFile(fmt.Sprintf("обработка задания %v из %v %v", k+1, countOfFiles, currFullFileName))
		logstr := fmt.Sprintf("начинаем читать json файл %v", currFullFileName)
		logsmy.LogginInFile(logstr)
		jsonCorrection, err := readJsonFromFile(currFullFileName)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) чтения json задания чека %v атол", err, currFullFileName)
			logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			continue
		}
		logstr = fmt.Sprintf("прочитали json файл %v", currFullFileName)
		logsmy.LogginInFile(logstr)
		logsmy.LogginInFile("парсим json задание")
		existMarksInCheck := false
		err = json.Unmarshal([]byte(jsonCorrection), &receipt)
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) json задания чека %v атол", err, jsonCorrection, currFullFileName)
			logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			continue
		}
		lastNameOfKassir = receipt.Operator.Name
		//проеверяем услвоия выхода из цикла по дате чека
		currDateOfCheck, err := time.Parse("2006.01.02", receipt.CorrectionBaseDate) //yyyy.mm.dd
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) парсинга даты %v для чека %v", err, receipt.CorrectionBaseDate, currFullFileName)
			logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		}
		if *ExlusionDate != "" {
			if ExlusionDateDate == currDateOfCheck {
				desrExit := fmt.Sprintf("достигли даты %v исключения - завершаем работы программы ", *ExlusionDate)
				logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println(desrExit)
				break
			}
		}
		if prevDateOfCheck != currDateOfCheck {
			logsmy.LogginInFile(fmt.Sprintf("переходим на новый день %v", currDateOfCheck))
			if prevDateOfCheck != initDate {
				countPrintedDays++
			}
			prevDateOfCheck = currDateOfCheck
		}
		//если надо делаем паузу в работе программы
		if *pauseAfterDay > 0 {
			if countPrintedDays >= *pauseAfterDay {
				logsmy.LogginInFile(fmt.Sprintf("произошло завершение дня, работы программы поставлена на паузу на %v секунд", *pauseInSecondsAfterDay))
				logsmy.LogginInFile(fmt.Sprintf("делаем паузу в программе через каждые %v дней для возможной остановки на %v секунд...", *pauseAfterDay, *pauseInSecondsAfterDay))
				logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("если процесс печати чеков нужно прервать, то это можно сделать сейчас - так как сейчас программа перешла на следующий %v день", prevDateOfCheck)
				countPrintedDays = 0
				duration := time.Second * time.Duration((*pauseInSecondsAfterDay))
				time.Sleep(duration)
			}
		}
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("%v: обработка задания %v из %v %v", receipt.CorrectionBaseDate, k+1, countOfFiles, currFullFileName)
		logsmy.LogginInFile("ищем марки в чеке")
		//очищаем таблицу марок
		if (*clearTableOfMarks) && (previusWasMarks) {
			if *kassatype == "atol" {
				breakProcCheckOfMark(fptr)
				clearTanlesOfMarks(fptr)
			} else {
				merc.BreakAndClearProccessOfMarks(*IpMerc, *PortMerc, *comport, sessionkey, *userint, *passwuser)
			}
		}
		//ищем все марки в json-задании и запускаем по каждой из них проверку
		previusWasMarks = false
		mistakeCheckingMark := false
		markErroDescr := ""
		receipt, existMarksInCheck, mistakeCheckingMark, markErroDescr, globalMistake, globalErrorStr = CheckAndRunsCheckingMarksByCheck(fptr, *kassatype, receipt, currFullFileName, true, *IpMerc, *PortMerc, *comport, sessionkey, *pauseOfMarksMistake, *userint, *passwuser)
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
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("перезапускаем процесс провекри марок для всего чека, так как была ошибкаЖ %v...", markErroDescr)
			if *kassatype == "atol" {
				breakProcCheckOfMark(fptr)
				clearTanlesOfMarks(fptr)
			} else {
				merc.BreakAndClearProccessOfMarks(*IpMerc, *PortMerc, *comport, sessionkey, *userint, *passwuser)
			}
			//отключаемся от ККТ
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("отлючаемся от ККТ")
			if *kassatype == "atol" {
				disconnectWithKKT(fptr, true)
			} else {
				merc.DissconnectMeruriy(*IpMerc, *PortMerc, sessionkey)
			}
			//делаем паузу
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("делаем паузу в %v секунд...", *pauseOfMarksMistake)
			duration := time.Second * time.Duration(*pauseOfMarksMistake)
			time.Sleep(duration)
			//подключаемся к ККТ
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("подключаемся к ККТ")
			if *kassatype == "atol" {
				_, err := connectToKKT(fptr, true)
				if err != nil {
					descrError := fmt.Sprintf("ошибка (%v) подключения к ККТ атол", err)
					logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
					globalErrorStr = descrError
					globalMistake = true
					break
				}
			} else {
				var err error
				descrError := ""
				sessionkey, descrError, err = merc.CheckStatsuConnectionKKT(*emulation, *IpMerc, *PortMerc, *comport, "", *userint, *passwuser)
				if err != nil {
					logsmy.Logsmap[consttypes.LOGERROR].Printf("ошибка (%v) подлкючение к ККТ меркурий", descrError)
					if !*emulation {
						println("Нажмите любую клавишу...")
						input.Scan()
						log.Panic(descrError)
					}
				} else {
					logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе меркурий на порт %v прошло успешно. Ключ сессии: %v", *comport, sessionkey)
				}
				//merc.Closesession(*IpMerc, *PortMerc, &sessionkey)
			}
			//запускаем проверку марки заново
			logsmy.LogginInFile("снова запускаем проверку марки")
			receipt, existMarksInCheck, mistakeCheckingMark, markErroDescr, globalMistake, globalErrorStr = CheckAndRunsCheckingMarksByCheck(fptr, *kassatype, receipt, currFullFileName, true, *IpMerc, *PortMerc, *comport, sessionkey, *pauseOfMarksMistake, *userint, *passwuser)
		}
		//если были серьёзные ошибки при проверки прерываем программу
		if globalMistake {
			errorDescr := fmt.Sprintf("ошибка %v", globalErrorStr)
			logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			//amountOfMistakesMarks++
			break
		}
		//пропускаем чек, если были ошибки при проверке марки
		if mistakeCheckingMark {
			errorDescr := fmt.Sprintf("ошибка (%v) проверки марки для чека %v атол", markErroDescr, currFullFileName)
			logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
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
		if *CassirName != "" {
			receipt.Operator.Name = *CassirName
			wasChangeParametersOfCheck = true
		}
		//пробиваем чек коррекции возврата на неправильный чек
		if *conversChekcCorrectionsType {
			logsmy.LogginInFile("конвертируем тим чека коррекции")
			if receipt.Type == "sellCorrection" {
				logsmy.LogginInFile("меняем тип чека коррекции с продажи на возврат")
				receipt.Type = "sellReturnCorrection"
			} else if receipt.Type == "sellReturnCorrection" {
				logsmy.LogginInFile("меняем тип чека коррекции с возврата на продажу")
				receipt.Type = "sellCorrection"
			} else if receipt.Type == "buyCorrection" {
				logsmy.LogginInFile("меняем тип чека коррекции с покупки на возрат покупки")
				receipt.Type = "buyReturnCorrection"
			} else if receipt.Type == "buyReturnCorrection" {
				logsmy.LogginInFile("меняем тип чека коррекции с возрата покупки на покупку")
				receipt.Type = "buyCorrection"
			}
			wasChangeParametersOfCheck = true
		}
		//меняем тип оплаты с наличной на безнал
		if *changeCashOnBeznal {
			for ind := range receipt.Payments {
				if receipt.Payments[ind].Type == "cash" {
					logsmy.LogginInFile(fmt.Sprintf("меняем тип оплаты с налички на безнал на сумму %v", receipt.Payments[ind].Sum))
					receipt.Payments[ind].Type = "electronically"
					wasChangeParametersOfCheck = true
				}
			}
		}
		if (existMarksInCheck) || (wasChangeParametersOfCheck) {
			jsonCorrWithMarkBytes, err := json.MarshalIndent(receipt, "", "\t")
			if err != nil {
				errorDescr := fmt.Sprintf("ошибка (%v) формирования json-а марками для задания чека %v атол", err, currFullFileName)
				logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
				amountOfMistakesChecks++
				continue
			}
			jsonCorrection = string(jsonCorrWithMarkBytes)
		}
		//печатаем чек
		logstr = fmt.Sprintf("посылаем команду печати чека кассу json файл %v", jsonCorrection)
		logsmy.LogginInFile(logstr)
		resulOfCommand := ""
		if *kassatype == "atol" {
			resulOfCommand, err = sendComandeAndGetAnswerFromKKT(fptr, jsonCorrection)
		} else {
			//mercuriy //меркурий
			resulOfCommand, err = merc.PrintCheck(*emulation, *IpMerc, *PortMerc, *comport, receipt, sessionkey, mercSNODefault, *dontprintrealfortest, *userint, *passwuser, *emulatmistakesOpenCheck)
		}
		//если были ошибку при печати чека, то переходим к следующему заданию
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) печати чека %v", err, currFullFileName)
			logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
			amountOfMistakesChecks++
			continue
		}
		logsmy.LogginInFile("послали команду печати чека кассу json файл")
		//эмулируем ошибку, если режим эмуляции ошибки включен
		if *emulatmistakes {
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("countPrintedChecks", countPrintedChecks)
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("countPrintedChecks%10", countPrintedChecks%10)
			if countPrintedChecks%10 == 0 {
				logsmy.LogginInFile("производим ошибку печати чека")
				logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("производим ошибку печати чека")
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
			logsmy.Logsmap[consttypes.LOGERROR].Printf(descrError)
			logsmy.LogginInFile(descrError)
			amountOfMistakesChecks++
		}
		if amountOfMistakesChecks > 0 {
			logsmy.LogginInFile(fmt.Sprintln("количество не напечатанных чеков", amountOfMistakesChecks))
		}
	} //перебор json заданий
	//закрывес соединение с кассой меркурий если было установлено
	if *kassatype == "merc" {
		if sessionkey != "" {
			merc.Closesession(*IpMerc, *PortMerc, &sessionkey)
		}
	}
	//выводим информацию об количестве напечтатнных чеков
	logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("распечатно %v из %v чеков", countPrintedChecks, countOfFiles)
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

func CheckAndRunsCheckingMarksByCheck(fptr *fptr10.IFptr, kassatype string, receipt consttypes.TCorrectionCheck, FullFileName string, perezapuskatproverku bool, ipktt string, port int, comport int, sessionkey string, pausetimesec int, userint int, passwuser string) (consttypes.TCorrectionCheck, bool, bool, string, bool, string) {
	logsmy.LogginInFile("ищем марки в чеке")
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
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println(logstr)
		imcResultCheckin, errproc := runProcessCheckMark(kassatype, fptr, ipktt, port, sessionkey, currMarkBase64)
		if *emulatmistakesmarks {
			errproc = errors.New("симуляция ошибки провекри марки")
		}
		if errproc != nil {
			//прерываем проверку
			if perezapuskatproverku {
				errorDescr = fmt.Sprintf("будет произведен перезпуск провекри марки, так как была ошибка (%v) запуска проверки марки %v для чека %v атол", errproc, currMarkBase64, FullFileName)
				logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println(errorDescr)
				//***********************************
				logsmy.LogginInFile(fmt.Sprintf("перезапускаем проверку марки %v...", currMarkBase64))
				logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("перезапускаем проверку марки...")
				//прерываем предыдущую провекру марки
				if kassatype == "atol" {
					breakProcCheckOfMark(fptr)
				} else {
					merc.BreakProcCheckOfMark(ipktt, port, comport, sessionkey, userint, passwuser)
				}
				//отключаемся от ККТ
				logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("отключаемся от ККТ")
				if kassatype == "atol" {
					disconnectWithKKT(fptr, true)
				} else {
					merc.DissconnectMeruriy(ipktt, port, sessionkey)
				}
				//делаем паузу
				logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("пауза в %v секунд...", pausetimesec)
				duration := time.Second * time.Duration(pausetimesec)
				time.Sleep(duration)
				//подключаемся к ККТ
				logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("подлкючаемся к ККТ")
				if kassatype == "atol" {
					_, err := connectToKKT(fptr, true)
					if err != nil {
						descrError := fmt.Sprintf("ошибка (%v) подключения к ККТ атол", err)
						logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
						globalErrorStr = descrError
						globalMistake = true
						break
					}
				} else {
					var err error
					descrError := ""
					sessionkey, descrError, err = merc.CheckStatsuConnectionKKT(*emulation, *IpMerc, *PortMerc, comport, "", userint, passwuser)
					if err != nil {
						descrError := fmt.Sprintf("ошибка (%v) подлкючение к ККТ меркурий", descrError)
						logsmy.Logsmap[consttypes.LOGERROR].Printf(descrError)
						if !*emulation {
							globalErrorStr = descrError
							globalMistake = true
							break
						}
					}
					//merc.Closesession(*IpMerc, *PortMerc, &sessionkey)
				}
				//запускаем проверку марки заново
				logsmy.LogginInFile("снова запускаем проверку марки")
				imcResultCheckin, errproc = runProcessCheckMark(kassatype, fptr, ipktt, port, sessionkey, currMarkBase64)
			} //перезапуск провекри марки
			if errproc != nil {
				errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v для чека %v атол", errproc, currMarkBase64, FullFileName)
				logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
				mistakeCheckingMark = true
				break
			}
		}
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("марка успешно проверена")
		//заполняем данные о марке
		logsmy.LogginInFile("заполняем данные о марке")
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
	logsmy.LogginInFile("начало процедуры sendComandeAndGetAnswerFromKKT")
	//return "", nil
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	//fptr.ValidateJson()
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*emulation {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды %v на кассе", err, comJson)
			logsmy.Logsmap[consttypes.LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры sendComandeAndGetAnswerFromKKT c ошибкой", err)
			logsmy.LogginInFile(logstr)
			return desrError, err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	if strings.Contains(result, "Нет связи") {
		logsmy.LogginInFile("нет связи: переподключаемся")
		if ok, typepodkluch := connectWithKassa(fptr, *comport, *ipaddresskkt, *portkktatol, *ipaddressservrkkt); !ok {
			descrErr := fmt.Sprintf("ошибка соединения с кассовым аппаратом %v", typepodkluch)
			logsmy.Logsmap[consttypes.LOGERROR].Println(descrErr)
			if !*emulation {
				println("Нажмите любую клавишу...")
				//input.Scan()
				log.Panic(descrErr)
			}
		} else {
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", *comport)
		}
	}
	//logsmy.Logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmy.Logsmap[LOGOTHER].Println("result", result)
	//if *emulation {
	//}
	//result := "{\"result\": \"all ok\"}"

	logsmy.LogginInFile("конец процедуры sendComandeAndGetAnswerFromKKT без ошибки")
	return result, nil
}

func runProcessCheckMark(kassatype string, fptr *fptr10.IFptr, ipktt string, port int, sessionkey string, mark string) (consttypes.TItemInfoCheckResult, error) {
	var countAttempts int
	var imcResultCheckinObj consttypes.TItemInfoCheckResultObject
	var imcResultCheckin consttypes.TItemInfoCheckResult
	logsmy.LogginInFile("начало процедуры runProcessCheckMark")
	//проверяем - открыта ли смена
	//shiftOpenned, err := checkOpenShift(fptr, true, "админ")
	//if err != nil {
	//	errorDescr := fmt.Sprintf("ошибка (%v). Смена не открыта", err)
	//	logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
	//	return TItemInfoCheckResult{}, errors.New(errorDescr)
	//}
	//if !shiftOpenned {
	//	errorDescr := fmt.Sprintf("ошибка (%v) - смена не открыта", err)
	//	logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
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
				if answerMerc.Result != 0 && !*emulation {
					resJson = "error " + resJson
				}
			}
		}
	}
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v", err, mark)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	if !successCommand(resJson) {
		errorDescr := fmt.Sprintf("ошибка (%v) запуска проверки марки %v", resJson, mark)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
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
					if MercurAnswerOfCheckMark.Result != 0 && !*emulation {
						resJson = "error " + resJson
					}
					if MercurAnswerOfCheckMark.Result != 0 && *emulation {
						MercurAnswerOfCheckMark.IsCompleted = true
					}
				}
			}
		}
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) получения статуса проверки марки %v", err, mark)
			logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
			return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
		}
		if !successCommand(resJson) {
			//делаем паузу
			logsmy.LogginInFile(resJson)
			desrAction := fmt.Sprintf("пауза в %v секунд... так сервер провекри марок не успевает.", *pauseOfMarksMistake)
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println(desrAction)
			duration := time.Second * time.Duration(*pauseOfMarksMistake)
			time.Sleep(duration)
			//if strings.Contains(resJson, "421")
			//errorDescr := fmt.Sprintf("ошибка (%v) получения статуса проверки марки %v", resJson, mark)
			//logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
			//return TItemInfoCheckResult{}, errors.New(errorDescr)
		}
		err = nil
		if kassatype == "atol" {
			err = json.Unmarshal([]byte(resJson), &answerOfCheckMark)
		}
		if err != nil {
			errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) ответа проверки марки %v", err, resJson, mark)
			logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
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
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("попытка %v из %v получения статуса марки", countAttempts+2, *countOfCheckingMarks)
		duration := time.Second
		time.Sleep(duration)
	}
	if countAttempts == *countOfCheckingMarks {
		errorDescr := fmt.Sprintf("ошибка проверки марки %v", mark)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
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
				if MercurAnswerOfResultOfCheckMark.Result != 0 && !*emulation {
					resOfChecking = "error " + resJson
				}
			}
		}
	}
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) принятия марки %v", err, mark)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	if !successCommand(resOfChecking) {
		errorDescr := fmt.Sprintf("ошибка (%v) принятия марки %v", resOfChecking, mark)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	err = nil
	if kassatype == "atol" {
		err = json.Unmarshal([]byte(resOfChecking), &imcResultCheckinObj)
	}
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) парсинга (%v) ответа проверки марки %v", err, resOfChecking, mark)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		return consttypes.TItemInfoCheckResult{}, errors.New(errorDescr)
	}
	if kassatype == "atol" {
		imcResultCheckin.EcrStandAloneFlag = imcResultCheckinObj.ItemInfoCheckResult.EcrStandAloneFlag
		imcResultCheckin.ImcCheckFlag = imcResultCheckinObj.ItemInfoCheckResult.ImcCheckFlag
		imcResultCheckin.ImcCheckResult = imcResultCheckinObj.ItemInfoCheckResult.ImcCheckResult
		imcResultCheckin.ImcEstimatedStatusCorrect = imcResultCheckinObj.ItemInfoCheckResult.ImcEstimatedStatusCorrect
		imcResultCheckin.ImcStatusInfo = imcResultCheckinObj.ItemInfoCheckResult.ImcStatusInfo
	}
	logsmy.LogginInFile("конец процедуры runProcessCheckMark без ошибки")
	return imcResultCheckin, nil
} //runProcessCheckMark

func breakProcCheckOfMark(fptr *fptr10.IFptr) error {
	var err error
	logsmy.LogginInFile("прерываемм проверку марки")
	turnCheckMarkJson := "{\"type\": \"cancelMarkingCodeValidation\"}"
	resturnCheckMark, _ := sendComandeAndGetAnswerFromKKT(fptr, turnCheckMarkJson)
	logsmy.LogginInFile(fmt.Sprintf("результат прерывания проверки марки: %v", resturnCheckMark))
	return err
}

func clearTanlesOfMarks(fptr *fptr10.IFptr) error {
	var err error
	logsmy.LogginInFile("очищаем таблицу марок")
	clearTableMarksJson := "{\"type\": \"clearMarkingCodeValidationResult\"}"
	resClearTableMarks, _ := sendComandeAndGetAnswerFromKKT(fptr, clearTableMarksJson)
	logsmy.LogginInFile(fmt.Sprintf("результат очистки таблицы марок: %v", resClearTableMarks))
	return err
}

func sendCheckOfMark(fptr *fptr10.IFptr, mark string) (string, error) {
	var err error
	logsmy.LogginInFile("начало процедуры sendCheckOfMark")
	//return "", nil
	comJson, err := getJsonOfBeginCheck(mark)
	if err != nil {
		desrError := fmt.Sprintf("ошибка (%v) фоормирования json задания проверки марки", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(desrError)
		logstr := fmt.Sprint("конец процедуры sendCheckOfMark c ошибкой", err)
		logsmy.LogginInFile(logstr)
		return "", err
	}
	//comJson = "{\"type\": \"reportX\",\"operator\": {\"name\": \"Иванов\",\"vatin\": \"123654789507\"}}"
	logstr := fmt.Sprintf("отправляем запрос (%v) на проверку марки", comJson)
	logsmy.LogginInFile(logstr)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*emulation {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды начать проверку марки на кассовом аппарате", err)
			logsmy.Logsmap[consttypes.LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры sendCheckOfMark c ошибкой", err)
			logsmy.LogginInFile(logstr)
			return "", err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	logstr = fmt.Sprintf("ответ на начало проверки марки: (%v) ", result)
	logsmy.LogginInFile(logstr)
	//logsmy.Logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmy.Logsmap[LOGOTHER].Println("result", result)
	//if *emulation {
	//}
	//result := "{\"result\": \"all ok\"}"
	logsmy.LogginInFile("конец процедуры sendCheckOfMark без ошибки")
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
		logsmy.Logsmap[consttypes.LOGERROR].Println(desrError)
		return "", err
	}
	return string(resstr), err
}

func getStatusOfChecking(fptr *fptr10.IFptr) (string, error) {
	var err error
	logsmy.LogginInFile("начало процедуры getStatusOfChecking")
	//return "", nil
	comJson := "{\"type\": \"getMarkingCodeValidationStatus\"}"
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*emulation {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды получения статуса марки на кассовом аппарате", err)
			logsmy.Logsmap[consttypes.LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры getStatusOfChecking c ошибкой", err)
			logsmy.LogginInFile(logstr)
			return "", err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	if *emulation {
		result = "{\"ready\": true}"
	}
	//logsmy.Logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmy.Logsmap[LOGOTHER].Println("result", result)
	//if *emulation {
	//}
	//result := "{\"result\": \"all ok\"}"
	logsmy.LogginInFile("конец процедуры getStatusOfChecking без ошибки")
	return result, nil
}

func acceptMark(fptr *fptr10.IFptr) (string, error) {
	var err error
	logsmy.LogginInFile("начало процедуры acceptMark")
	//return "", nil
	comJson := "{\"type\": \"acceptMarkingCode\"}"
	fptr.SetParam(fptr10.LIBFPTR_PARAM_JSON_DATA, comJson)
	if !*emulation {
		err = fptr.ProcessJson()
	}
	if err != nil {
		if !*emulation {
			desrError := fmt.Sprintf("ошибка (%v) выполнение команды принятия марки на кассовом аппарате", err)
			logsmy.Logsmap[consttypes.LOGERROR].Println(desrError)
			logstr := fmt.Sprint("конец процедуры acceptMark c ошибкой", err)
			logsmy.LogginInFile(logstr)
			return "", err
		}
	}
	result := fptr.GetParamString(fptr10.LIBFPTR_PARAM_JSON_DATA)
	if *emulation {
		result = "{\"itemInfoCheckResult\": {\"ecrStandAloneFlag\": false,\"imcCheckFlag\": true,\"imcCheckResult\": true,\"imcEstimatedStatusCorrect\": true,\"imcStatusInfo\": true}}"
	}
	//logsmy.Logsmap[LOGOTHER].Println("comJson", comJson)
	//logsmy.Logsmap[LOGOTHER].Println("result", result)
	//if *emulation {
	//}
	logstr := fmt.Sprintf("результат проверки марки (%v) ", result)
	logsmy.LogginInFile(logstr)
	//result := "{\"result\": \"all ok\"}"
	logsmy.LogginInFile("конец процедуры acceptMark без ошибки")
	return result, nil
}

func listDirByReadDir(path string) ([]string, error) {
	var spisFiles []string
	var spisFileFD []int
	logstr := fmt.Sprintf("перебор файлов в директории %v--BEGIN\n", path)
	logsmy.LogginInFile(logstr)
	defer logsmy.LogginInFile(fmt.Sprintf("перебор файлов в директории %v--END\n", path))
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
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("пропускаем файл с ифнормацией о напечатнных чеках")
			continue
		}
		if consttypes.FILE_NAME_CONNECTION == val.Name() {
			logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("пропускаем файл с ифнормацией о настройки связи с ККТ")
			continue
		}
		/*logsmy.Logsmap[consttypes.LOGINFO].Println(val.Name())
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
		//logsmy.LogginInFile(logstr)
		if matched {
			spisFiles = append(spisFiles, val.Name())
		}
	}
	logsmy.LogginInFile(fmt.Sprintln("spisFiles=", spisFiles))
	for _, filename := range spisFiles {
		fdstr := getFDFromFileName(filename)
		fdint, err := strconv.Atoi(fdstr)
		if err != nil {
			logsmy.LogginInFile(fmt.Sprintf("ошибка (%v) получения номера ФД из имени файла %v при сортировке списка файлов по номеру ФД", err, filename))
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
				logsmy.LogginInFile(fmt.Sprintf("ошибка (%v) получения номера ФД из имени файла %v при сортировке списка файлов по номеру ФД", err, filename))
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
	logsmy.LogginInFile(logstr)
	plan, err := ioutil.ReadFile(currFullFileName)
	//logstr = fmt.Sprintln("plan", plan)
	//logsmy.LogginInFile(logstr)
	//logstr = fmt.Sprintln("error", err)
	//logsmy.LogginInFile(logstr)
	//logstr = fmt.Sprintln("конец процедуры readJsonFromFile")
	//logsmy.LogginInFile(logstr)
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
		logsmy.LogginInFile(desrError)
		//logsmy.Logsmap[consttypes.LOGERROR].Println(desrError)
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
	logsmy.LogginInFile("снова создаём объект драйвера...")
	if createComObj {
		fptr, err = fptr10.NewSafe()
	}
	if err != nil {
		descrError := fmt.Sprintf("ошибка (%v) инициализации драйвера ККТ атол", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
		return descrError, errors.New(descrError)
	}
	//сединение с кассой
	logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("соединение с кассой...")
	*comport, _ = getCurrentPortOfKass(consttypes.DIROFJSONS)
	if ok, typepodkluch := connectWithKassa(fptr, *comport, *ipaddresskkt, *portkktatol, *ipaddressservrkkt); !ok {
		descrErr := fmt.Sprintf("ошибка сокдинения с кассовым аппаратом %v", typepodkluch)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrErr)
		if !*emulation {
			return descrErr, errors.New(descrErr)
		}
	} else {
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", *comport)
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
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrError)
		return errors.New(descrError)
		//println("Нажмите любую клавишу...")
		//input.Scan()
		//log.Panic(descrError)
	}
	//defer fptr.Destroy()
	fmt.Println(fptr.Version())
	//сединение с кассой
	logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Println("соединение с кассой")
	*comport, _ = getCurrentPortOfKass(consttypes.DIROFJSONS)
	//if err != nil {
	//	desrErr := fmt.Sprintf("ошибка (%v) чтения параметра com порт соединения с кассой", err)
	//	logsmy.Logsmap[consttypes.LOGERROR].Println(desrErr)
	//	return errors.New(desrErr)
	//	//println("Нажмите любую клавишу...")
	//	//input.Scan()
	//	//log.Panic(desrErr)
	//}
	if ok, typepodkluch := connectWithKassa(fptr, *comport, *ipaddresskkt, *portkktatol, *ipaddressservrkkt); !ok {
		//if !connectWithKassa(fptr, *comport, *ipaddresskkt, *ipaddressservrkkt) {
		descrErr := fmt.Sprintf("ошибка сокдинения с кассовым аппаратом %v", typepodkluch)
		logsmy.Logsmap[consttypes.LOGERROR].Println(descrErr)
		if !*emulation {
			return errors.New(descrErr)
			//println("Нажмите любую клавишу...")
			//input.Scan()
			//log.Panic(descrErr)
		}
	} else {
		logsmy.Logsmap[consttypes.LOGINFO_WITHSTD].Printf("подключение к кассе на порт %v прошло успешно", *comport)
	}
	return nil
	//defer fptr.Close()
}

func checkOpenShift(fptr *fptr10.IFptr, openShiftIfClose bool, kassir string) (bool, error) {
	logsmy.LogginInFile("получаем статус ККТ")
	getStatusKKTJson := "{\"type\": \"getDeviceStatus\"}"
	resgetStatusKKT, err := sendComandeAndGetAnswerFromKKT(fptr, getStatusKKTJson)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) получения статуса кассы", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		return false, err
	}
	if !successCommand(resgetStatusKKT) {
		errorDescr := fmt.Sprintf("ошибка (%v) получения статуса кассы", resgetStatusKKT)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		//logsmy.LogginInFile(errorDescr)
		return false, errors.New(errorDescr)
	}
	logsmy.LogginInFile("получили статус кассы")
	//проверяем - открыта ли смена
	var answerOfGetStatusofShift consttypes.TAnswerGetStatusOfShift
	err = json.Unmarshal([]byte(resgetStatusKKT), &answerOfGetStatusofShift)
	if err != nil {
		errorDescr := fmt.Sprintf("ошибка (%v) распарсивания статуса кассы", err)
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		return false, err
	}
	if answerOfGetStatusofShift.ShiftStatus.State == "expired" {
		errorDescr := "ошибка - смена на кассе уже истекла. Закройте смену"
		logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
		return false, errors.New(errorDescr)
	}
	if answerOfGetStatusofShift.ShiftStatus.State == "closed" {
		if openShiftIfClose {
			if kassir == "" {
				errorDescr := "не указано имя кассира для открытия смены"
				logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
			jsonOpenShift := fmt.Sprintf("{\"type\": \"openShift\",\"operator\": {\"name\": \"%v\"}}", kassir)
			resOpenShift, err := sendComandeAndGetAnswerFromKKT(fptr, jsonOpenShift)
			if err != nil {
				errorDescr := fmt.Sprintf("ошбика (%v) - не удалось открыть смену", err)
				logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
			if !successCommand(resOpenShift) {
				errorDescr := fmt.Sprintf("ошбика (%v) - не удалось открыть смену", resOpenShift)
				logsmy.Logsmap[consttypes.LOGERROR].Println(errorDescr)
				return false, errors.New(errorDescr)
			}
		} else {
			return false, nil
		}
	}
	return true, nil
} //checkOpenShift
