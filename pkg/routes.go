package pkg

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/service"
	"github.com/joho/godotenv"
)

func NewServer() () {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}



	//pathToFileCleanerJson :=getPathToFileCleanerJson()
	fileJsonParser := service.NewFileJsonParser()
	fileExpunge := service.NewFileExpunge(fileJsonParser)
	fileExpunge.ExecuteDeleteTask()
	//_, _ = fileJsonParser.ParseFileCleanerJson(pathToFileCleanerJson)


}


