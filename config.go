package veego

import (
	"io/ioutil"

	"encoding/json"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"

	"os"
)

type AppConfig struct {
	Host        string `yaml:"host" json:"host"`
	Port        string `yaml:"port" json:"port"`
	DatabaseURL string `yaml:"database_url" json:"database_url"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

// LoadFromFile : loads config details from file
func (a *AppConfig) LoadYML(file string) (*AppConfig, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &AppConfig{}
	if err := yaml.Unmarshal(yamlFile, config); err != nil {
		return nil, err
	}
	return config, nil
}

func (a *AppConfig) LoadEnv(file string) (*AppConfig, error) {
	err := godotenv.Load(file)
	if err != nil {
		return nil, err
	}
	return &AppConfig{
		Host:        os.Getenv("host"),
		Port:        os.Getenv("port"),
		DatabaseURL: os.Getenv("database_url"),
	}, nil
}

func (a *AppConfig) LoadJSON(file string) (*AppConfig, error) {
	jsonFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &AppConfig{}
	if err := json.Unmarshal(jsonFile, config); err != nil {
		return nil, err
	}
	return config, nil
}
