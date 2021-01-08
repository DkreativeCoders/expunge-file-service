package service

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/utils"
	"log"
)

type FileExpunge struct {
	fileJsonParser IFileJsonParse
}

func NewFileExpunge(fileJsonParser IFileJsonParse) IFileExpunge {
	return &FileExpunge{fileJsonParser: fileJsonParser}
}



func (f FileExpunge) ExecuteDeleteTask() {
	pathToFileCleanerJson :=utils.GetPathToFileCleanerJson()
	fileCleanerJsonConfig,err:= f.fileJsonParser.ParseFileCleanerJson(pathToFileCleanerJson)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("fileCleanerJsonConfig serviceConfigs", fileCleanerJsonConfig.ServiceConfigs)

	//-Walk through the whole file path and get all the file paths needed if
	//  ---"enableRecursiveDepth" : true,

	//	-here you exclude directories that are not meant to be searched
	//-remove excluded extensions
	//-remove excluded fileNames
	//-remove excludeFileNamesContaining
	//-check last modified date of the file ==  "fileAgeToBeDeleted": 120,
	//-check if "enableFileMovementToBackupFolder": true,
	//   --if true  get  "pathToBackUpFolder":"C:\\Users\\dell\\Documents\\daniel\\backup"
	//-move file to backupFolder
	//delete all the files


	//      "excludeExtensions": [".jar",".bat"],
	//      "excludeFileNames" : ["do-not-delete","delete-not"],
	//      "excludeFileNamesContaining" : ["tomcat","tomcat"],
	//panic("implement me")
}




