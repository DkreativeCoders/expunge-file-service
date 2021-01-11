package pkg

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/Expunge"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/JsonParser"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/ProcessFile"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/RunnableCrons"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"time"
)

func NewServer() () {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}



	//pathToFileCleanerJson :=utils.GetPathToFileCleanerJson()
	fileJsonParser := JsonParser.NewFileJsonParser()
	factoryProcessFile := ProcessFile.NewFactoryProcessFile()
	fileExpunge := Expunge.NewFileExpunge(fileJsonParser,*factoryProcessFile)
	//fileExpunge.ExecuteDeleteTask()
	//_, _ = fileJsonParser.ParseFileCleanerJson(pathToFileCleanerJson)
	cronFactory := RunnableCrons.NewCronFactory(fileExpunge)

	c := cron.New()
	for _, cronService := range cronFactory.CronService {
		_, _ = c.AddFunc(cronService.GetCronTime(), func() { cronService.Execute() })

		// Added time to see output
	}
	//_, _ = c.AddFunc("@every 0h0m10s", func() { fmt.Println("Every second") })
	_, _ = c.AddFunc("* * * * *", func() { fmt.Println("Every second") })

	c.Start()
	time.Sleep(20 * time.Minute)
	c.Stop()
}


