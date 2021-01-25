package ProcessFile

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg/domain"
	"github.com/kpango/glg"
	"log"
	"os"
	"path/filepath"
)

type UsingRecursiveDepth struct {
}

func NewProcessFileUsingRecursiveDepth() *UsingRecursiveDepth {
	return &UsingRecursiveDepth{}
}

func (p UsingRecursiveDepth) PrepareFile(generalConfig domain.GeneralConfig,
	serviceConfig domain.ServiceConfig,
	fileProcessState *domain.FileProcessState) {

	if serviceConfig.EnableRecursiveDepth{
		//perform recursive check
		setOfExcludedFolder := p.getExcludedFolderAsSet(serviceConfig.ExcludeFolders)

		err := filepath.Walk(serviceConfig.RootPath,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() && setOfExcludedFolder[info.Name()]{
					glg.Log("found and skipping excludedFolder==>",info.Name())
					return filepath.SkipDir
				}

				if !info.IsDir(){
					fileProcessState.SetOfFilesPath[path]=true
				}
				//glg.Log(path, "====>", info.Size(), "date===>", info.ModTime())
				return nil
			})
		if err != nil {
			log.Println(err)
		}

		glg.Log("After ProcessFileUsingRecursiveDepth")

		//fileProcessState.SetOfFilesPath=setOfFilePath
	}
}
func (p UsingRecursiveDepth) getExcludedFolderAsSet(excludeFolders []string) map[string]bool{
	excludedFolderAsSet := make(map[string]bool)
	for _, s := range excludeFolders {
		excludedFolderAsSet[s] = true
	}
	return excludedFolderAsSet
}




