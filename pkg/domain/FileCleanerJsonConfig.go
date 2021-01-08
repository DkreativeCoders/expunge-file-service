package domain

type FileCleanerJsonConfig struct {
	GeneralConfig GeneralConfig
	Services                                      []string
	ServiceConfigs                                []ServiceConfig
}
