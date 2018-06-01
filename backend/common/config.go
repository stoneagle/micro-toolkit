package common

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type DBConf struct {
	Name     string
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Target   string
	MaxIdle  int
	MaxOpen  int
	Showsql  bool
	Location string
}

type Conf struct {
	App struct {
		Mode      string
		BasicAuth string `yaml:"BasicAuth"`
	}
	Storybox struct {
		Cms struct {
			Url        string
			Supertoken string
		}
		Toolkit struct {
			Database DBConf
		}
		Mqtt struct {
			Database DBConf
			Params   string
		}
		Callback struct {
			Database DBConf
			Config   string
		}
		Upgrade struct {
			Database DBConf
		}
		Album struct {
			Database DBConf
		}
	}
}

var onceConfig *Conf = &Conf{}

func GetConfig() *Conf {
	if (Conf{}) == *onceConfig {
		configPath := os.Getenv("ConfigPath")
		if configPath == "" {
			configPath = "./config/config.yaml"
		}
		yamlFile, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(err)
		}
		config := &Conf{}
		err = yaml.Unmarshal(yamlFile, config)
		if err != nil {
			panic(err)
		}
		onceConfig = config
	}
	return onceConfig
}
