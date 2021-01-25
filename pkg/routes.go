package pkg

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/controller"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/Expunge"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/JsonParser"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/ProcessFile"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service/RunnableCrons"
	"github.com/DkreativeCoders/expunge-file-service/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kpango/glg"
	"github.com/robfig/cron/v3"
	"net/http"
	"os"
)

func NewServer() *http.Server {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	utils.SetFileLogLogger()

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



	//pathToFileCleanerJson :=utils.GetPathToFileCleanerJson()
	//----------------------------------------------
	fileJsonParser := JsonParser.NewFileJsonParser()
	factoryProcessFile := ProcessFile.NewFactoryProcessFile()
	fileExpunge := Expunge.NewFileExpunge(fileJsonParser,*factoryProcessFile)
	//fileExpunge.ExecuteDeleteTask()
	//_, _ = fileJsonParser.ParseFileCleanerJson(pathToFileCleanerJson)
	cronFactory := RunnableCrons.NewCronFactory(fileExpunge)

	c := cron.New()

	enableCronJob := os.Getenv("ENABLE.CRONJOB")

	if enableCronJob=="true"{
		for _, cronService := range cronFactory.CronService {
			_, _ = c.AddFunc(cronService.GetCronTime(), func() { cronService.Execute() })
			// Added time to see output
		}
	}


	router := mux.NewRouter()
	controller.NewExpungeController(router, fileExpunge)

	//_, _ = c.AddFunc("@every 0h0m10s", func() { fmt.Println("Every second") })
	//_, _ = c.AddFunc("* * * * *", func() { fmt.Println("Every second") })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8600" //localhost
	}

	glg.Log("service started on port: ",port)
	srv := &http.Server{Handler: router, Addr: ":" + port}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
	c.Start()
	//time.Sleep(20 * time.Minute)
	//c.Stop()

	return srv

}


