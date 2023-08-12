package config

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
	ServerPort  int    `json:"server_port"`
	ServerHost  string `json:"server_host"`
	JWTSecret   string `json:"jwt_secret"`
}

var (
	defaultConfig = Config{
		DatabaseURL: "mysql_database_url",
		ServerPort:  8080,
		ServerHost:  "0.0.0.0",
		JWTSecret:   "Cd18jPdmI6aUWEkmvZomzltcjhetm9MMn64HDEwFBYDpMlks2fZpZ0nGvWp0G9TuMcQeySuGi_P7jwQA3gh7zQ",
	}
	config     Config
	configLock sync.RWMutex
	loaded     bool
	filePath   = "./config/config.json"
)

func init() {
	_ = ReadConfigFromFile()
}

func ReadConfigFromFile() error {
	file, err := os.Open(filePath)
	if err != nil {
		config = defaultConfig // 使用默认配置
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configLock.Lock()
	defer configLock.Unlock()

	err = decoder.Decode(&config)
	if err != nil {
		config = defaultConfig // 使用默认配置
	}
	return nil
}

func GetConfig() Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func SaveConfigToFile() error {
	configLock.RLock()
	defer configLock.RUnlock()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(config)
}
