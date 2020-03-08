package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	HttpsPort string `yaml:"app.httpsPort"`
}

func (c *YamlConfig) GetConf() *YamlConfig {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (c *YamlConfig) GetServerPort() string {
	return ":" + c.HttpsPort
}