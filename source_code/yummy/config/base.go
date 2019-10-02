package config

import (
	"os"

	"sample-project/helper"
)

type Config struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Address  string `json:"address"`
	} `json:"database"`
}

func SetupConfig() Config {
	var conf Config

	// Đọc file config.local.json
	configFile, err := os.Open("config.local.json")
	if err != nil {
		defer configFile.Close()
		panic(err)
	}
	defer configFile.Close()

	// Parse dữ liệu JSON và bind vào struct Config
	err = helper.DecodeDataFromJsonFile(configFile, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}
