package ProcessFile

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/kpango/glg"
	"os"
)

type Eradicate struct {
}

func NewProcessFileEradicate() *Eradicate {
	return &Eradicate{}
}

func (p Eradicate) PrepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	var enableFileMovementToBackupFolder bool
	if serviceConfig.UseGeneralConfig {
		enableFileMovementToBackupFolder = generalConfig.EnableFileMovementToBackupFolder
	} else {
		enableFileMovementToBackupFolder = serviceConfig.EnableFileMovementToBackupFolder
	}

	if !enableFileMovementToBackupFolder {
		p.deleteFile(fileProcessState)
		glg.Log("After ProcessFileEradicate")
		glg.Log("SetOfFilesPathToBeDeleted ", fileProcessState.SetOfFilesPath)
	}

}

func (p Eradicate) deleteFile(fileProcessState *domain.FileProcessState) {

	for systemFilePath := range fileProcessState.SetOfFilesPath {
		if _, err := os.Stat(systemFilePath); os.IsNotExist(err) {
			// path/to/whatever does not exist
			glg.Log(systemFilePath, "File does not exist")
			continue
		}
		glg.Log("Deleting file...", systemFilePath)
		err := os.Remove(systemFilePath)
		glg.Log("Deleted file", systemFilePath)

		if err != nil {
			glg.Error(err)
		}

	}
}
