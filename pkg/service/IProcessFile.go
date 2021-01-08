package service

import "github.com/DkreativeCoders/expunge-file-service/pkg/domain"

type IProcessFile interface {
	prepareFile(generalConfig domain.GeneralConfig,
		serviceConfig domain.ServiceConfig,
		fileProcessState *domain.FileProcessState) ()
}

