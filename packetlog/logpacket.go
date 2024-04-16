package logmodules

import (
	consttypes "clientrabbit/consttypes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	DeegreOfDebug = LogsDebugs
	if foundedLogDir, _ := consttypes.DoesFileExist(consttypes.LOGSDIR); !foundedLogDir {
		os.Mkdir(consttypes.LOGSDIR, 0777)
	}
	filelogmap, Logsmap, descrError, err = initializationLogsLoc(clearLogsProgramm, consttypes.LOGINFO, consttypes.LOGERROR, consttypes.LOGSKIP_LINES, consttypes.LOGOTHER)
	if err != nil {
		descrMistake := fmt.Sprintf("ошибка инициализации лог файлов %v", descrError)
		return descrMistake, err
	}
	fmt.Println("лог файлы инициализированы в папке " + consttypes.LOGSDIR)
	multwriterLocLoc := io.MultiWriter(Logsmap[consttypes.LOGINFO].Writer(), os.Stdout)
	Logsmap[consttypes.LOGINFO_WITHSTD] = log.New(multwriterLocLoc, consttypes.LOG_PREFIX+"_"+strings.ToUpper(consttypes.LOGINFO)+" ", log.LstdFlags)
	return "OK", nil
}

func initializationLogsLoc(clearLogs bool, logstrs ...string) (map[string]*os.File, map[string]*log.Logger, string, error) {
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
