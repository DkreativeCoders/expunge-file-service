package domain

type GeneralConfig struct {
	FileAgeLastModifiedInDays        int
	EnableFileMovementToBackupFolder bool
	PathToBackupFile                       string
}
