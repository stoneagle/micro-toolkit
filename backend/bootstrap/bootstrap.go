package bootstrap

import (
	"fmt"
	"time"
	"toolkit/backend/common"
	"toolkit/backend/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Boot(app *gin.Engine) {
	conf := common.GetConfig()
	setProjectEngine(conf.Storybox.Toolkit)

	toolkit := app.Group("/toolkit")
	{
		controllers.NewAutoBuild(common.GetEngine(conf.Storybox.Toolkit.Name)).Router(toolkit)
	}
}

func setProjectEngine(dbConfig common.DBConf) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Target)

	engine, err := xorm.NewEngine(dbConfig.Type, source)
	if err != nil {
		panic(err)
	}
	engine.SetMaxIdleConns(dbConfig.MaxIdle)
	engine.SetMaxOpenConns(dbConfig.MaxOpen)

	location, err := time.LoadLocation(dbConfig.Location)
	if err != nil {
		panic(err)
	}
	engine.TZLocation = location
	engine.ShowSQL(dbConfig.ShowSQL)

	common.SetEngine(dbConfig.Name, engine)
}
