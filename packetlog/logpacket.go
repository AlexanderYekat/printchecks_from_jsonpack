package logmodules

import (
	consttypes "clientrabbit/consttypes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var filelogmap map[string]*os.File
var Logsmap map[string]*log.Logger
var descrError string
var err error

var DeegreOfDebug int

func LogginInFile(loggin string) {
	if DeegreOfDebug > 0 {
		Logsmap[consttypes.LOGINFO].Println(loggin)
	}
}

func CloseDescrptorsLogs() {
	fmt.Println("закрытие дескрипторов лог файлов программы")
	for _, v := range filelogmap {
		if v != nil {
			v.Close()
		}
	}
}

func InitializationsLogs(clearLogsProgramm bool, LogsDebugs int) (string, error) {
	startTime := time.Now()
	fmt.Printf("начало инициализации логов: %v\n", startTime.Format("15:04:05"))

	DeegreOfDebug = LogsDebugs

	// Проверяем существование директории логов с таймаутом
	fmt.Println("проверка директории логов...")
	if foundedLogDir, _ := consttypes.DoesFileExist(consttypes.LOGSDIR); !foundedLogDir {
		fmt.Println("создание директории логов...")
		os.Mkdir(consttypes.LOGSDIR, 0777)
		fmt.Println("директория логов создана")
	} else {
		fmt.Println("директория логов уже существует")
	}

	fmt.Println("инициализация лог файлов...")
	filelogmap, Logsmap, descrError, err = initializationLogsLoc(clearLogsProgramm, consttypes.LOGINFO, consttypes.LOGERROR, consttypes.LOGSKIP_LINES, consttypes.LOGOTHER)
	if err != nil {
		descrMistake := fmt.Sprintf("ошибка инициализации лог файлов %v", descrError)
		return descrMistake, err
	}

	fmt.Println("лог файлы инициализированы в папке " + consttypes.LOGSDIR)
	multwriterLocLoc := io.MultiWriter(Logsmap[consttypes.LOGINFO].Writer(), os.Stdout)
	Logsmap[consttypes.LOGINFO_WITHSTD] = log.New(multwriterLocLoc, consttypes.LOG_PREFIX+"_"+strings.ToUpper(consttypes.LOGINFO)+" ", log.LstdFlags)

	elapsed := time.Since(startTime)
	fmt.Printf("инициализация логов завершена за %v\n", elapsed)

	return "OK", nil
}

func initializationLogsLoc(clearLogs bool, logstrs ...string) (map[string]*os.File, map[string]*log.Logger, string, error) {
	var reserr, err error
	reserr = nil
	filelogmapLoc := make(map[string]*os.File)
	logsmapLoc := make(map[string]*log.Logger)
	descrError := ""

	fmt.Printf("инициализация %v лог файлов...\n", len(logstrs))

	for i, logstr := range logstrs {
		fmt.Printf("инициализация лог файла %v/%v: %v\n", i+1, len(logstrs), logstr)

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
		fmt.Printf("лог файл %v инициализирован успешно\n", logstr)
	}

	return filelogmapLoc, logsmapLoc, descrError, reserr
}

func intitLog(logFile string, pref string, clearLogs bool) (*os.File, *log.Logger, error) {
	startTime := time.Now()
	fmt.Printf("открытие лог файла: %v\n", logFile)

	flagsTempOpen := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	if clearLogs {
		flagsTempOpen = os.O_TRUNC | os.O_CREATE | os.O_WRONLY
	}

	// Добавляем таймаут для операции открытия файла
	done := make(chan bool, 1)
	var f *os.File
	var openErr error

	go func() {
		f, openErr = os.OpenFile(logFile, flagsTempOpen, 0644)
		done <- true
	}()

	// Ждем максимум 5 секунд
	select {
	case <-done:
		// Файл открыт
	case <-time.After(5 * time.Second):
		return nil, nil, fmt.Errorf("таймаут открытия файла %v", logFile)
	}

	if openErr != nil {
		fmt.Printf("ошибка открытия файла %v: %v\n", logFile, openErr)
		return nil, nil, openErr
	}

	elapsed := time.Since(startTime)
	fmt.Printf("файл %v открыт за %v\n", logFile, elapsed)

	multwr := io.MultiWriter(f)
	//if pref == LOG_PREFIX+"_INFO" {
	//	multwr = io.MultiWriter(f, os.Stdout)
	//}
	flagsLogs := log.LstdFlags
	if pref == consttypes.LOG_PREFIX+"_ERROR" {
		multwr = io.MultiWriter(f, os.Stderr)
		flagsLogs = log.LstdFlags | log.Lshortfile
	}

	loger := log.New(multwr, pref+" ", flagsLogs)
	return f, loger, nil
}
