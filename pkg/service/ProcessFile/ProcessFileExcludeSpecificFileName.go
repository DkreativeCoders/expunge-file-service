package ProcessFile

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/DkreativeCoders/expunge-file-service/pkg/utils"
	"github.com/kpango/glg"
	"os"
	"strings"
)

type ExcludeSpecificFileName struct {
}

func NewProcessFileExcludeSpecificFileName() *ExcludeSpecificFileName {
	return &ExcludeSpecificFileName{}
}

func (p ExcludeSpecificFileName) PrepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	mapOfSpecificFileNameToPaths := make(map[string]map[string] bool) //==> java Map<String, Set<String>>
	excludeSpecificFileNameAsSet := utils.GetListOfStringAsSet(serviceConfig.ExcludeSpecificFileNames)
	//glg.Log("excludeSpecificFileNameAsSet ", excludeSpecificFileNameAsSet)

	p.setMapOfExcludedSpecificFileName(fileProcessState, excludeSpecificFileNameAsSet, mapOfSpecificFileNameToPaths)

	for fileName, setOfFilePath := range mapOfSpecificFileNameToPaths {
		for fileToBeDeleted, _ := range setOfFilePath {
			delete(fileProcessState.SetOfFilesPath, fileToBeDeleted)    // Delete
			glg.Log("Removed file ", fileToBeDeleted, " for this fileName =>> ",fileName)
		}
	}

	glg.Log("After ProcessFileExcludeSpecificFileName")
	glg.Log("SetOfFilesPathToBeDeleted ", fileProcessState.SetOfFilesPath)

}

func (p ExcludeSpecificFileName) setMapOfExcludedSpecificFileName(fileProcessState *domain.FileProcessState,
	excludeSpecificFileNameAsSet map[string]bool,
	mapOfSpecificFileNameToPaths map[string]map[string]bool) {

	for systemFilePath, _ := range fileProcessState.SetOfFilesPath {
		fileInfo, _ := os.Stat(systemFilePath)

		fileName := fileInfo.Name()
		if p.excludedSpecificFileNameContainsFileName(excludeSpecificFileNameAsSet, fileName) {
			if mapOfSpecificFileNameToPaths[fileName] == nil {
				mapOfSpecificFileNameToPaths[fileName] = make(map[string]bool)
			}

			setOfFilePath := mapOfSpecificFileNameToPaths[fileName]
			setOfFilePath[systemFilePath] = true
		}

	}
	glg.Log(mapOfSpecificFileNameToPaths)
}

func (p ExcludeSpecificFileName) excludedSpecificFileNameContainsFileName(
	excludeSpecificFileNameAsSet map[string]bool,
	name string) bool{

	splitName := strings.Split(name,".")
	if excludeSpecificFileNameAsSet[splitName[0]]{
		return true
	}
	return false
}




