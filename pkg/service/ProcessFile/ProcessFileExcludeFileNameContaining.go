package ProcessFile

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/DkreativeCoders/expunge-file-service/pkg/utils"
	"github.com/kpango/glg"
	"os"
	"strings"
)

type ExcludeFileNameContaining struct {
}

func NewProcessFileExcludeFileNameContaining() *ExcludeFileNameContaining {
	return &ExcludeFileNameContaining{}
}

func (p ExcludeFileNameContaining) PrepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	mapOfSpecificFileNameToPaths := make(map[string]map[string] bool) //==> java Map<String, Set<String>>
	excludeFileNameContainingAsSet := utils.GetListOfStringAsSet(serviceConfig.ExcludeFileNamesContaining)
	//glg.Log("excludeSpecificFileNameAsSet ", excludeSpecificFileNameAsSet)

	p.setMapOfExcludedSpecificFileName(fileProcessState, excludeFileNameContainingAsSet, mapOfSpecificFileNameToPaths)

	for fileName, setOfFilePath := range mapOfSpecificFileNameToPaths {
		for fileToBeDeleted, _ := range setOfFilePath {
			delete(fileProcessState.SetOfFilesPath, fileToBeDeleted)    // Delete
			glg.Log("Exclude file ", fileToBeDeleted, " for this fileName =>> ",fileName)
		}
	}

	glg.Log("After ProcessFileExcludeFileNameContaining")
	glg.Log("SetOfFilesPathToBeDeleted ", fileProcessState.SetOfFilesPath)

}

func (p ExcludeFileNameContaining) setMapOfExcludedSpecificFileName(fileProcessState *domain.FileProcessState,
	excludeFileNameContainingAsSet map[string]bool,
	mapOfSpecificFileNameToPaths map[string]map[string]bool) {

	for systemFilePath, _ := range fileProcessState.SetOfFilesPath {
		fileInfo, _ := os.Stat(systemFilePath)

		fileName := fileInfo.Name()
		if p.excludeFileNameContainsFileNameConfig(excludeFileNameContainingAsSet, fileName) {
			if mapOfSpecificFileNameToPaths[fileName] == nil {
				mapOfSpecificFileNameToPaths[fileName] = make(map[string]bool)
			}

			setOfFilePath := mapOfSpecificFileNameToPaths[fileName]
			setOfFilePath[systemFilePath] = true
		}

	}
	glg.Log(mapOfSpecificFileNameToPaths)
}

func (p ExcludeFileNameContaining) excludeFileNameContainsFileNameConfig(
	excludeFileNameContainingAsSet map[string]bool,
	name string) bool{

	splitName := strings.Split(name,".")
	//if excludeSpecificFileNameAsSet[splitName[0]]{
	//	return true
	//}
	for fileNameContainsConfig, _ := range excludeFileNameContainingAsSet {
		if strings.Contains(splitName[0],fileNameContainsConfig){
			return true
		}
	}
	return false
}




