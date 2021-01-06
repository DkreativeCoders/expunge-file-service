package service

import "file-cleaner-service/pkg"

type IFileJsonParse interface {
	ParseFileCleanerJson() (*pkg.FileCleanerJsonConfig, error)
}
