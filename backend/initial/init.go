package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"toolkit/backend/common"
	cm "toolkit/backend/models"

	yaml "gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	config := &common.Conf{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(err)
	}
	toolkitConfig := config.Storybox.Toolkit
	initToolkit(toolkitConfig)
}

func initToolkit(dbConfig common.DBConf) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Target)

	engine, err := xorm.NewEngine(dbConfig.Type, source)
	if err != nil {
		panic(err)
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	engine.TZLocation = location
	engine.StoreEngine("InnoDB")
	engine.Charset("utf8")
	err = engine.DropTables(new(cm.AutoBuild))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(cm.AutoBuild))
	if err != nil {
		panic(err)
	}
}
