package service

import "github.com/DkreativeCoders/expunge-file-service/pkg/domain"

type IFileJsonParse interface {
	ParseFileCleanerJson(filePath string) (*domain.FileCleanerJsonConfig, error)
}
