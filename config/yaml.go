package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	App struct {
		HttpsPort string `yaml:"httpsPort"`
		CertPath  string `yaml:"certPath"`
		KeyPath   string `yaml:"keyPath"`
	} `yaml:"app"`
}

func (c *YamlConfig) ReadConf() *YamlConfig {

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
	return ":" + c.App.HttpsPort
}
func (c *YamlConfig) GetKeyPath() string {
	return c.App.KeyPath
}
func (c *YamlConfig) GetCertPath() string {
	return c.App.CertPath
}
