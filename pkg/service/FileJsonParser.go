package service

import (
	"encoding/json"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"io/ioutil"
	"log"
)

type FileJsonParser struct {
}

func NewFileJsonParser() IFileJsonParse {
	return &FileJsonParser{}
}

func (f *FileJsonParser) ParseFileCleanerJson(filePath string) (*domain.FileCleanerJsonConfig, error) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var fileCleanerJsonConfig *domain.FileCleanerJsonConfig
	if err := json.Unmarshal(data, &fileCleanerJsonConfig); err != nil {
		log.Fatal(err)
		return nil,err
	}

	log.Print("Data read", fileCleanerJsonConfig)
	log.Print(fileCleanerJsonConfig.ServiceConfigs)

	return fileCleanerJsonConfig, err
}


