package ProcessFile

import (
	"fmt"
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/kpango/glg"
	"os"
	"path/filepath"
)

type NonRecursive struct {
}

func NewProcessFileNonRecursive() *NonRecursive {
	return &NonRecursive{}
}


func (p NonRecursive) PrepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	if ! serviceConfig.EnableRecursiveDepth{
		//perform recursive check
		setOfExcludedFolder := p.getExcludedFolderAsSet(serviceConfig.ExcludeFolders)
		f, err := os.Open(serviceConfig.RootPath)
		if err != nil {
			glg.Log(err)
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

		glg.Log("After ProcessFileNonRecursive")
		glg.Log("setOfFilePath", fileProcessState.SetOfFilesPath)

	}
}
func (p NonRecursive) getExcludedFolderAsSet(excludeFolders []string) map[string]bool{
	excludedFolderAsSet := make(map[string]bool)
	for _, s := range excludeFolders {
		excludedFolderAsSet[s] = true
	}
	return excludedFolderAsSet
}




