package ProcessFile

import "github.com/DkreativeCoders/expunge-file-service/pkg/domain"

type IProcessFile interface {
	PrepareFile(generalConfig domain.GeneralConfig,
		serviceConfig domain.ServiceConfig,
		fileProcessState *domain.FileProcessState) ()
}

