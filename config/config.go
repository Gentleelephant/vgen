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
	V    int    `json:"v,omitempty" yaml:"v,omitempty"`
	Ps   string `json:"ps,omitempty" yaml:"ps,omitempty"`
	Add  string `json:"add,omitempty" yaml:"add,omitempty"`
	Port string `json:"port,omitempty" yaml:"port,omitempty"`
	Id   string `json:"id,omitempty" yaml:"id,omitempty"`
	Aid  int    `json:"aid,omitempty" yaml:"aid,omitempty"`
	Net  string `json:"net,omitempty" yaml:"net,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
	Tls  string `json:"tls,omitempty" yaml:"tls,omitempty"`
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
	fmt.Printf("%+v\n", config)
	if err != nil {
		panic(err)
	}
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
