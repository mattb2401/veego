package veego

import (
	"io/ioutil"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"

	"os"
)

type AppConfig struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	DatabaseURL string `yaml:"database_url"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

// LoadFromFile : loads config details from file
func (ac *AppConfig) LoadYML(file string) (*AppConfig, error) {
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

func (ac *AppConfig) LoadEnv(file string) (*AppConfig, error) {
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
