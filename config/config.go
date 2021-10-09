package config

import (
	"bufio"
	"encoding/json"
	"os"
)

type GlobalConfig struct {
	Domain 		string 		`json:"domain,omitempty"`
	Port   		string 		`json:"port,omitempty"`
	RedisConfig RedisConfig `json:"redis_config"`
}

type RedisConfig struct {
	Addr     string 	`json:"addr,omitempty"`
	Password string 	`json:"password,omitempty"`
	DB       int  		`json:"db,omitempty"`
}

var globalConfig *GlobalConfig = nil

func GetGlobalConfig() *GlobalConfig {
	// 不重复加载文件
	if globalConfig != nil {
		return globalConfig
	}

	file, err := os.Open("config/config-dev.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&globalConfig); err != nil {
		panic(err)
	}
	return globalConfig
}