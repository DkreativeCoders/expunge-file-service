package ProcessFile

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"log"
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
		fmt.Println("After ProcessFileEradicate")
		fmt.Println("SetOfFilesPathToBeDeleted ", fileProcessState.SetOfFilesPath)
	}

}

func (p Eradicate) deleteFile(fileProcessState *domain.FileProcessState) {

	for systemFilePath := range fileProcessState.SetOfFilesPath {
		if _, err := os.Stat(systemFilePath); os.IsNotExist(err) {
			// path/to/whatever does not exist
			fmt.Println(systemFilePath, "File does not exist")
			continue
		}
		fmt.Println("Deleting file...", systemFilePath)
		err := os.Remove(systemFilePath)
		fmt.Println("Deleted file", systemFilePath)

		if err != nil {
			log.Fatal(err)
		}

	}
}
