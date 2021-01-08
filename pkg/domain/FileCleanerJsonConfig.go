package domain

type FileCleanerJsonConfig struct {
	GeneralConfigFileAgeLastModifiedInDays        int
	GeneralConfigEnableFileMovementToBackupFolder bool
	GeneralPathToBackupFile                       string
	Services                                      []string
	ServiceConfigs                                []ServiceConfig
}
