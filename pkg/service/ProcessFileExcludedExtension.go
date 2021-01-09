package service

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/DkreativeCoders/expunge-file-service/pkg/utils"
	"path/filepath"
)

type ProcessFileExcludedExtension struct {
}

func NewProcessFileExcludedExtension() *ProcessFileExcludedExtension {
	return &ProcessFileExcludedExtension{}
}



func (p ProcessFileExcludedExtension) prepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	mapOfExcludedExtensionToPaths := make(map[string]map[string] bool) //==> java Map<String, Set<String>>
	excludeExtensionsAsSet := utils.GetListOfStringAsSet(serviceConfig.ExcludeExtensions)
	p.setMapOfExcludedExtensions(fileProcessState, excludeExtensionsAsSet, mapOfExcludedExtensionToPaths)

	for extensions, setOfFilePathOfFilePath := range mapOfExcludedExtensionToPaths {
		for fileToBeDeleted, _ := range setOfFilePathOfFilePath {
			delete(fileProcessState.SetOfFilesPath, fileToBeDeleted)    // Delete
			fmt.Println("Removed file ", fileToBeDeleted, " for this ext =>> ",extensions)
		}
	}

	fmt.Println("After ProcessFileExcludedExtension")
	fmt.Println("SetOfFilesPathToBeDeleted ", fileProcessState.SetOfFilesPath)

}

func (p ProcessFileExcludedExtension) setMapOfExcludedExtensions(fileProcessState *domain.FileProcessState,
	excludeExtensionsAsSet map[string]bool,
	mapOfExcludedExtensionToPaths map[string]map[string]bool) {

	for systemFilePath, _ := range fileProcessState.SetOfFilesPath {
		fileExtension := filepath.Ext(systemFilePath)
		if p.excludedExtensionsContainsFilePath(excludeExtensionsAsSet, systemFilePath) {
			if mapOfExcludedExtensionToPaths[fileExtension] == nil {
				mapOfExcludedExtensionToPaths[fileExtension] = make(map[string]bool)
			}

			setOfFilePath := mapOfExcludedExtensionToPaths[fileExtension]
			setOfFilePath[systemFilePath] = true
		}

	}
	fmt.Println(mapOfExcludedExtensionToPaths)
}

func (p ProcessFileExcludedExtension) excludedExtensionsContainsFilePath(excludeExtensionsAsSet map[string]bool, path string) bool{
	if excludeExtensionsAsSet[filepath.Ext(path)]{
		return true
	}
	return false
}




