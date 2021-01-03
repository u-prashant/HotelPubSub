package main

import (
	"common/messagequeue"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

type configuration struct {
	RMqConfig messagequeue.RabbitMqConfig `yaml:"rMqConfig"`
}

const configFile = "../hotelpublisher/config.yaml"

// Config contains all the configuration for the publisher
// that are read from config yaml
var Config configuration

func loadConfig() error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Errorf("error reading config file - file[%s] err[%s]", configFile, err.Error())
		return err
	}
	return yaml.Unmarshal(data, &Config)
}
