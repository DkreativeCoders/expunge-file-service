package service

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"os"
	"path/filepath"
)

type ProcessFileNonRecursive struct {
}

func NewProcessFileNonRecursive() *ProcessFileNonRecursive {
	return &ProcessFileNonRecursive{}
}


func (p ProcessFileNonRecursive) prepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	if ! serviceConfig.EnableRecursiveDepth{
		//perform recursive check
		setOfExcludedFolder := p.getExcludedFolderAsSet(serviceConfig.ExcludeFolders)
		f, err := os.Open(serviceConfig.RootPath)
		if err != nil {
			fmt.Println(err)
		}
		files, err := f.Readdir(-1)
		_ = f.Close()


		for _, file := range files {
			if file.IsDir() && setOfExcludedFolder[file.Name()]{
				fmt.Print("found and skipping excludedFolder==>",file.Name())
				continue
			}
			fullPath := filepath.Join(serviceConfig.RootPath, file.Name())
			fileProcessState.SetOfFilesPath[fullPath]=true
		}
		//fileProcessState.SetOfFilesPath=setOfFilePath

		fmt.Println("After ProcessFileNonRecursive")
		fmt.Println("setOfFilePath", fileProcessState.SetOfFilesPath)

	}
}
func (p ProcessFileNonRecursive) getExcludedFolderAsSet(excludeFolders []string) map[string]bool{
	excludedFolderAsSet := make(map[string]bool)
	for _, s := range excludeFolders {
		excludedFolderAsSet[s] = true
	}
	return excludedFolderAsSet
}




