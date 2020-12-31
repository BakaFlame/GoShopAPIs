package TestController

import (
	"GoShop/middleWare"
	"github.com/gin-gonic/gin"
)

//解析路由，请手动添加每一个路由的请求相对路径

func (test *TestStruct) Route(EngineRouter *gin.Engine)  {
	//testRoute := EngineRouter.Group("/api/test")
	//testRoute.Use(tool.Cors())
	//{
	EngineRouter.POST("/testjson",test.testjson)
	//}
	//EngineRouter.Use(tool.Cors())
	EngineRouter.POST("/nocors",test.testjson)
	EngineRouter.GET("/test",middleWare.RequesterWare(),test.testControl)
	EngineRouter.GET("/testdb",test.testDb)
	EngineRouter.GET("/testInsert",test.testInsert)
	EngineRouter.GET("/testdelete",test.testDelete)
	EngineRouter.GET("/testupdate",test.testUpdate)
	EngineRouter.POST("/testupload",test.testUpload)
	EngineRouter.GET("/jsonpage",test.jsonPage)
	EngineRouter.GET("/uploadpage",test.uploadPage)
	EngineRouter.GET("/teststructthing",test.teststructthing)
	EngineRouter.GET("/testget",test.testget)
	EngineRouter.GET("/testrows",test.testrows)
	EngineRouter.POST("/testfakesession",test.testfakesession)
}