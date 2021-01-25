package utils

import (
	"github.com/kpango/glg"
	"log"
	"os"
	"path/filepath"
)

func SetFileLogLogger() {
	workingDirectoryPath, err := os.Getwd()
	if err != nil {
	}
	logPath := os.Getenv("CONFIG.LOG")
	fullPathToLogFile := filepath.Join(workingDirectoryPath,"logs", logPath)

	_, err = os.Stat(fullPathToLogFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File does not exist.")
			_ = os.MkdirAll(filepath.Join(workingDirectoryPath, "logs"), 0700)
			newFile, err := os.Create(fullPathToLogFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(newFile)
			newFile.Close()
		}
	}
	log.Println("Log path:",fullPathToLogFile)

	log.Println("File does exist. File information:")
	//log.Println(fileInfo.Name())

	infoLog := glg.FileWriter(fullPathToLogFile, 0666)

	//defer infoLog.Close()
	glg.Get().
		SetMode(glg.BOTH). // default is STD
		DisableColor().
		// SetMode(glg.NONE).
		// SetMode(glg.WRITER).
		// SetMode(glg.BOTH).
		// InitWriter().
		// AddWriter(customWriter).
		// SetWriter(customWriter).
		// AddLevelWriter(glg.LOG, customWriter).
		// AddLevelWriter(glg.INFO, customWriter).
		// AddLevelWriter(glg.WARN, customWriter).
		// AddLevelWriter(glg.ERR, customWriter).
		// SetLevelWriter(glg.LOG, customWriter).
		// SetLevelWriter(glg.INFO, customWriter).
		// SetLevelWriter(glg.WARN, customWriter).
		// SetLevelWriter(glg.ERR, customWriter).
		AddLevelWriter(glg.INFO, infoLog).                         // add info log file destination
		AddLevelWriter(glg.DEBG, infoLog).                         // add info log file destination
		AddLevelWriter(glg.FAIL, infoLog).                         // add info log file destination
		AddLevelWriter(glg.FATAL, infoLog).                         // add info log file destination
		AddLevelWriter(glg.WARN, infoLog).                         // add info log file destination
		AddLevelWriter(glg.LOG, infoLog).                         // add info log file destination
		AddLevelWriter(glg.OK, infoLog).                         // add info log file destination
		AddLevelWriter(glg.ERR, infoLog)                          // add error log file destination
	//AddStdLevel(customTag, glg.STD, false).                    //user custom log level
	//AddErrLevel(customErrTag, glg.STD, true)                  // user custom error log level
	//SetLevelColor(glg.TagStringToLevel(customTag), glg.Cyan).  // set color output to user custom level
	//SetLevelColor(glg.TagStringToLevel(customErrTag), glg.Red) // set color output to user custom level
	//
	//glg.Info("info")
	//glg.Infof("%s : %s", "info", "formatted")
	//glg.Log("log")
	//glg.Logf("%s : %s", "info", "formatted")
	//glg.Debug("debug")
	//glg.Debugf("%s : %s", "info", "formatted")
	//glg.Warn("warn")
	//glg.Warnf("%s : %s", "info", "formatted")
	//glg.Error("error")
	//glg.Errorf("%s : %s", "info", "formatted")
	//glg.Success("ok")
	//glg.Successf("%s : %s", "info", "formatted")
	//glg.Fail("fail")
	//glg.Failf("%s : %s", "info", "formatted")
	//glg.Print("Print")

}
