package logger

import (
	"github.com/sirupsen/logrus"
	"sweetcook-backend/config"
	"sweetcook-backend/utils"
	"os"
	//"time"
	"fmt"
)

var log = logrus.New()

func init()  {
	log.Formatter = new(logrus.JSONFormatter)
	
	configJson := config.GetConfigJson()
	logInfo := configJson["log"]
	
	logFileDir, _ := utils.GetStringFromInterfaceMap(logInfo, "fileDir")
	logLevel, _ := utils.GetStringFromInterfaceMap(logInfo, "level")
	logOut, _ := utils.GetStringFromInterfaceMap(logInfo, "out")
	logFilePath := logFileDir + "/log.log"
	
	if logOut == "file" {
		createLogfile(logFileDir, logFilePath)
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err == nil {
			log.Out = file
			log.Info("logger set to logfile.")
		} else {
			log.Out = os.Stdout
			log.Info("Failed to log to file, using default stderr")
		}
	}else {
		log.Out = os.Stdout
		log.Info("logger set to stdout.")
	}
	
	log.Info("current log level is: ", logLevel)
	
	switch logLevel {
	case "debug":
		log.Level = logrus.DebugLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "fatal":
		log.Level = logrus.FatalLevel
	case "panic":
		log.Level = logrus.PanicLevel
	}
}

func createLogfile(fileDir string, filePath string) {
	_, err := os.Stat(filePath)
	if err == nil {
		fmt.Println("file: ", filePath, " existed.")
		return
	}
	if os.IsNotExist(err) {
		//目录不存在则递归创建目录
		_, err = os.Stat(fileDir)
		if err != nil {
			os.MkdirAll(fileDir, 0777)
			fmt.Println("fileDir: ", fileDir, " created.")
		}
	
		os.Create(filePath)
		fmt.Println("file: ", filePath, " created.")
	}
}

func Debug(args ...interface{})  {
	log.WithFields(logrus.Fields{
	}).Debug(args)
}

func Warn(args ...interface{})  {
	log.WithFields(logrus.Fields{
	}).Warn(args)
}

func Info(args ...interface{})  {
	log.WithFields(logrus.Fields{
	}).Info(args)
}

func Error(args ...interface{})  {
	log.WithFields(logrus.Fields{
	}).Error(args)
}

func Fatal(args ...interface{})  {
	log.WithFields(logrus.Fields{
	}).Fatal(args)
}

func Panic(args ...interface{})  {
	log.WithFields(logrus.Fields{
	}).Panic(args)
}