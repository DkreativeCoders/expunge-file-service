package service

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("Data read", fileCleanerJsonConfig)
	fmt.Println(fileCleanerJsonConfig.ServiceConfigs)

	return fileCleanerJsonConfig, err
}


