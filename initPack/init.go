package initPack

import (
	"GoShop/controller"
	"GoShop/middleWare"
	"GoShop/model"
	"GoShop/tool"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

func InitWeb() (*gin.Engine,*tool.Config,*gorm.DB,*sql.DB,*redis.Client) {
	cfg,err := tool.ParseConfig("./config/Config.json") //配置文件加载
	if err != nil {
		panic(err.Error())
	}
	if cfg.AppMode == "release"{
		gin.SetMode(gin.ReleaseMode)
	}
	app:=gin.Default()
	app.Use(tool.Cors())
	app.Use(middleWare.RequesterWare())
	db,db_query,redisDB:=model.RegisterDB()//开启DB
	controller.RegisterRouter(app) //注册路由
	app.LoadHTMLGlob("html/*") //导入html文件夹
	return app,cfg,db,db_query,redisDB //返回引擎，配置，DB
}
