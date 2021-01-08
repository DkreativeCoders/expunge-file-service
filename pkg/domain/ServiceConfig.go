package domain

type ServiceConfig struct {
	ServiceName string
	RootPath string
	ExcludeFolders []string
	ExcludeExtensions []string
	ExcludeFileNames []string
	ExcludeFileNamesContaining []string
	FileAgeToBeDeleted int
	EnableRecursiveDepth bool
	EnableFileMovementToBackupFolder bool
	PathToBackUpFolder string
}



