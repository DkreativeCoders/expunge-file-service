package JsonParser

import (
	"encoding/json"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/kpango/glg"
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

	glg.Info("Data read", fileCleanerJsonConfig)
	glg.Info(fileCleanerJsonConfig.ServiceConfigs)

	return fileCleanerJsonConfig, err
}


