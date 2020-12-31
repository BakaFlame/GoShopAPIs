package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

//项目的config文件夹

type Config struct {
	AppName    string `json:"app_name"`
	AppMode    string `json:"app_mode"`
	AppHost    string `json:"app_host"`
	AppPort    string `json:"app_port"`
	AppEmail   string `json:"app_email"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
}

var _cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
