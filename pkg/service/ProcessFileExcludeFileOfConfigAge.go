package service

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"os"
	"time"
)

type ProcessFileExcludeFileOfConfigAge struct {
}

func NewProcessFileExcludeFileOfConfigAge() *ProcessFileExcludeFileOfConfigAge {
	return &ProcessFileExcludeFileOfConfigAge{}
}


func (p ProcessFileExcludeFileOfConfigAge) prepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	setOfExcludedFilePathsDueToLastModifiedDate := make(map[string]bool) //==> java Map<String, Set<String>>
	var fileAgeToBeDeleted int
	if serviceConfig.UseGeneralConfig {
		fileAgeToBeDeleted =generalConfig.FileAgeLastModifiedInDays
	}else {
		fileAgeToBeDeleted =serviceConfig.FileAgeToBeDeleted
	}


	p.setMapOfExcludedExtensions(fileProcessState, fileAgeToBeDeleted, setOfExcludedFilePathsDueToLastModifiedDate)

	for fileToBeExcluded, _ := range setOfExcludedFilePathsDueToLastModifiedDate {
		delete(fileProcessState.SetOfFilesPath, fileToBeExcluded)    // Delete
		fmt.Println("Removed file ", fileToBeExcluded )
	}


	fmt.Println("After ProcessFileExcludeFileOfConfigAge")
	fmt.Println("SetOfFilesPathToBeDeleted ", fileProcessState.SetOfFilesPath)

}

func (p ProcessFileExcludeFileOfConfigAge) setMapOfExcludedExtensions(fileProcessState *domain.FileProcessState,
	configuredFileAgeTobeDeleted int,
	setOfExcludedFilePathsDueToLastModifiedDate map[string]bool) {

	for systemFilePath, _ := range fileProcessState.SetOfFilesPath {

		if p.fileLastModifiedDateIsLessThanConfig(configuredFileAgeTobeDeleted, systemFilePath) {
			setOfExcludedFilePathsDueToLastModifiedDate[systemFilePath] = true
		}

	}
	fmt.Println(setOfExcludedFilePathsDueToLastModifiedDate)
}

func (p ProcessFileExcludeFileOfConfigAge) fileLastModifiedDateIsLessThanConfig(
	configuredFileAgeTobeDeleted int,
	path string) bool{
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}

	var configuredFileDurationTobeDeleted time.Duration
	configuredFileAgeTobeDeleted = int(time.Duration(configuredFileAgeTobeDeleted))
	configDate := time.Now().Add(configuredFileDurationTobeDeleted * 24 * time.Hour)

	//fmt.Println(today)
	//fmt.Println(today.Format("02-Jan-2006"))
	fmt.Println(fileInfo.ModTime())
	differenceInDays := int(configDate.Sub(fileInfo.ModTime()).Hours() / 24 )//
	if differenceInDays < configuredFileAgeTobeDeleted{
		return  true
	}

	fmt.Println("file age with respect to configTime ==>",differenceInDays,"file last modifiedTime >>",
		fileInfo.ModTime().Format("02-Jan-2006"))


	return false
}




