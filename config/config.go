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
	JWTExp      int64  `json:"jwt_exp"`
}

var (
	defaultConfig = Config{
		DatabaseURL: "mysql_database_url",
		ServerPort:  8080,
		ServerHost:  "0.0.0.0",
		JWTSecret:   "Cd18jPdmI6aUWEkmvZomzltcjhetm9MMn64HDEwFBYDpMlks2fZpZ0nGvWp0G9TuMcQeySuGi_P7jwQA3gh7zQ",
		JWTExp:      60 * 60 * 24, // 一天
	}
	config     Config
	configLock sync.RWMutex
	filePath   = "./config/config.json"
)

func init() {
	_ = ReadConfigFromFile()
}

// ReadConfigFromFile 从文件中读取配置、配置文件不存在就使用默认的配置
func ReadConfigFromFile() error {
	file, err := os.Open(filePath)
	if err != nil {
		config = defaultConfig
		return nil
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	decoder := json.NewDecoder(file)
	configLock.Lock()
	defer configLock.Unlock()

	err = decoder.Decode(&config)
	if err != nil {
		config = defaultConfig
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
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}
