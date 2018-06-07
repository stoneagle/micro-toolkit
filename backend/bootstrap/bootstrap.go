package bootstrap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"toolkit/backend/common"
	"toolkit/backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	sessionKey = "user"
)

func Boot(app *gin.Engine) {
	app.StaticFS("/static", http.Dir("./static"))
	app.StaticFS("/assets", http.Dir("./static/assets"))
	app.StaticFile("/favicon.ico", "./static/favicon.ico")
	app.GET("", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	store := sessions.NewCookieStore([]byte("secret"))
	app.Use(sessions.Sessions("session", store))

	conf := common.GetConfig()
	setProjectEngine(conf.Storybox.Toolkit.Database)
	setProjectEngine(conf.Storybox.Mqtt.Database)
	setProjectEngine(conf.Storybox.Callback.Database)
	setProjectEngine(conf.Storybox.Upgrade.Database)
	setProjectEngine(conf.Storybox.Album.Database)

	if conf.App.Mode == "debug" {
		app.Use(cors.New(cors.Config{
			AllowHeaders:     []string{"Content-Type", "Access-Control-Allow-Origin", "Authorization"},
			AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
			AllowCredentials: true,
			AllowOrigins:     []string{"http://localhost:8080", "http://localhost:6999"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://localhost:6999"
			},
			MaxAge: 12 * time.Hour,
		}))
	}

	BAConf := map[string]string{}
	err := json.Unmarshal([]byte(conf.App.BasicAuth), &BAConf)
	if err != nil {
		panic(err)
	}

	app.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "")
	})

	v1 := app.Group("/v1")
	{
		v1.GET("/toolkit/login", gin.BasicAuth(BAConf), func(c *gin.Context) {
			user := c.MustGet(gin.AuthUserKey).(string)
			session := sessions.Default(c)
			session.Set(sessionKey, user)
			err := session.Save()
			if err != nil {
				common.ResponseErrorBusiness(c, common.ErrorLogin, "session save error", err)
			} else {
				common.ResponseSuccess(c, struct{}{})
			}
		})

		v1.GET("/toolkit/logout", func(c *gin.Context) {
			session := sessions.Default(c)
			user := session.Get(sessionKey)
			if user == nil {
				common.ResponseErrorBusiness(c, common.ErrorLogin, "invalid session token", nil)
			} else {
				session.Delete(sessionKey)
				session.Save()
				common.ResponseSuccess(c, struct{}{})
			}
		})

		toolkit := v1.Group("/toolkit", AuthRequired())
		{
			controllers.NewAutoBuild(common.GetEngine(conf.Storybox.Toolkit.Database.Name)).Router(toolkit)
		}
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": common.ErrorLogin,
				"data": struct{}{},
				"desc": "invalid session token",
			})
		} else {
			c.Next()
		}
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
	if dbConfig.Showsql {
		engine.ShowSQL(true)
	}

	common.SetEngine(dbConfig.Name, engine)
}
