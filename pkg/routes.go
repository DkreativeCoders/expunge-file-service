package pkg

import (
	"fmt"
	"github.com/joho/godotenv"
)

func NewServer() () {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}



	//pathToFileCleanerJson :=getPathToFileCleanerJson()
	//fileJsonParser := service.NewFileJsonParser()
	//_, _ = fileJsonParser.ParseFileCleanerJson(pathToFileCleanerJson)


}


