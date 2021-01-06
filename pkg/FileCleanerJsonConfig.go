package pkg

type FileCleanerJsonConfig struct {
	GeneralConfigFileAgeLastModifiedInDays        int
	GeneralConfigEnableFileMovementToBackupFolder string
	GeneralPathToBackupFile                       string `json:"age"`
	Services                                      []string
	ServiceConfigs                                []ServiceConfig
}
