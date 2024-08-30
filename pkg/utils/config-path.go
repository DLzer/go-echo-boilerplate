package utils

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "production" {
		return "./production"
	} else if configPath == "development" {
		return "./local"
	}
	return "./local"
}
