package pkg

type ServiceConfig struct {
	serviceName string
	rootPath string
	excludedFolders []string
	excludedExtensions []string
	excludedFileNames []string
	fileAgeToBeDeleted int
	enableRecursiveDepth string
	enableFileMovementToBackupFolder string
	pathToBackUpFolder string
}

