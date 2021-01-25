package ProcessFile

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/kpango/glg"
	"os"
	"time"
)

type ExcludeFileOfConfigAge struct {
}

func NewProcessFileExcludeFileOfConfigAge() *ExcludeFileOfConfigAge {
	return &ExcludeFileOfConfigAge{}
}


func (p ExcludeFileOfConfigAge) PrepareFile(generalConfig domain.GeneralConfig,
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
		glg.Log("Removed file ", fileToBeExcluded )
	}


	glg.Log("After ProcessFileExcludeFileOfConfigAge")
	glg.Log("SetOfFilesPathToBeDeleted ", fileProcessState.SetOfFilesPath)

}

func (p ExcludeFileOfConfigAge) setMapOfExcludedExtensions(fileProcessState *domain.FileProcessState,
	configuredFileAgeTobeDeleted int,
	setOfExcludedFilePathsDueToLastModifiedDate map[string]bool) {

	for systemFilePath, _ := range fileProcessState.SetOfFilesPath {

		if p.fileLastModifiedDateIsLessThanConfig(configuredFileAgeTobeDeleted, systemFilePath) {
			setOfExcludedFilePathsDueToLastModifiedDate[systemFilePath] = true
		}

	}
	glg.Log(setOfExcludedFilePathsDueToLastModifiedDate)
}

func (p ExcludeFileOfConfigAge) fileLastModifiedDateIsLessThanConfig(
	configuredFileAgeTobeDeleted int,
	path string) bool{
	fileInfo, err := os.Stat(path)
	if err != nil {
		glg.Log(err)
	}

	var configuredFileDurationTobeDeleted time.Duration
	configuredFileAgeTobeDeleted = int(time.Duration(configuredFileAgeTobeDeleted))
	configDate := time.Now().Add(configuredFileDurationTobeDeleted * 24 * time.Hour)

	//glg.Log(today)
	//glg.Log(today.Format("02-Jan-2006"))
	glg.Log(fileInfo.ModTime())
	differenceInDays := int(configDate.Sub(fileInfo.ModTime()).Hours() / 24 )//
	if differenceInDays < configuredFileAgeTobeDeleted{
		return  true
	}

	glg.Log("file age with respect to configTime ==>",differenceInDays,"file last modifiedTime >>",
		fileInfo.ModTime().Format("02-Jan-2006"))


	return false
}




