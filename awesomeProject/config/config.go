package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config — структура config.yml
type Config struct {
	Port      int    `yaml:"port"`
	AuthToken string `yaml:"auth_token"`
	APIURLs   struct {
		Hba1c string `yaml:"hba1c"`
		Ldl   string `yaml:"ldl"`
		Tg    string `yaml:"tg"`
		Hdl   string `yaml:"hdl"`
		Ferr  string `yaml:"ferr"`
	} `yaml:"api_urls"`
}

// Cfg — глобальный инстанс конфига
var Cfg Config

// LoadConfig читает YAML по указанному пути и заполняет Cfg
func LoadConfig(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot open config file %q: %v", path, err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("cannot read config file %q: %v", path, err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		log.Fatalf("cannot unmarshal config file %q: %v", path, err)
	}
}
