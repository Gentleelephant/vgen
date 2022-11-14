package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	Filepath string
)

type Config struct {
	V    string `json:"v,omitempty"`
	Ps   string `json:"ps,omitempty""`
	Add  string `json:"add,omitempty"`
	Port string `json:"port,omitempty"`
	Id   string `json:"id,omitempty"`
	Aid  string `json:"aid,omitempty"`
	Net  string `json:"net,omitempty"`
	Type string `json:"type,omitempty"`
	Host string `json:"host,omitempty"`
	Path string `json:"path,omitempty"`
	Tls  string `json:"tls,omitempty"`
}

func loadConfig() *Config {
	if Filepath == "" {
		Filepath = "./config.yaml"
	}
	file, err := os.ReadFile(Filepath)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	return &config
}

func Config2String() string {
	config := loadConfig()
	marshal, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}
