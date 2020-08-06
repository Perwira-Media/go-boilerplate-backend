package config

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config stores the whole configuration for service.
type Config struct {
	ServiceData ServiceData `yaml:"service_data"`
	SourceData  SourceData  `yaml:"source_data"`
}

// ServiceData contains the service data configuration.
type ServiceData struct {
	Address string `yaml:"address"`
}

// SourceData contains the source data configuration.
type SourceData struct {
	PostgresHost     string `yaml:"postgres_server"`
	PostgresDBName   string `yaml:"postgres_dbname"`
	PostgresUsername string `yaml:"postgres_username"`
	PostgresPassword string `yaml:"postgres_password"`
	PostgresTimeout  int    `yaml:"postgres_timeout"`
	PostgresPort     int    `yaml:"postgres_Port"`
}

// GetConfig parse the configuration from YAML file.
func GetConfig(fileLocation string) (*Config, error) {
	config, err := getRawConfig(fileLocation)
	if err != nil {
		return nil, fmt.Errorf("Unable to get raw config content: %v", err)
	}
	return config, nil
}

func getRawConfig(fileLocation string) (*Config, error) {
	configByte, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		logrus.Errorf("Error Read File Raw Config: %v", err)
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(configByte, config)
	if err != nil {
		logrus.Errorf("Error Unmarshal Raw Config: %v", err)
		return nil, err
	}
	return config, nil
}
