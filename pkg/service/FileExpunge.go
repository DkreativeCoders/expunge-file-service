package service

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/utils"
	"log"
)

type FileExpunge struct {
	fileJsonParser IFileJsonParse
}

func NewFileExpunge(fileJsonParser IFileJsonParse) IFileExpunger {
	return &FileExpunge{fileJsonParser: fileJsonParser}
}



func (f FileExpunge) executeDeleteTask() {
	pathToFileCleanerJson :=utils.GetPathToFileCleanerJson()
	fileCleanerJsonConfig,err:= f.fileJsonParser.ParseFileCleanerJson(pathToFileCleanerJson)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("fileCleanerJsonConfig serviceConfigs", fileCleanerJsonConfig.ServiceConfigs)

	panic("implement me")
}




