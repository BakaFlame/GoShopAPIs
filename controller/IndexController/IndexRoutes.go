package IndexController

import (
	"GoShop/tool"
	"github.com/gin-gonic/gin"
)

func (index *IndexStruct) Route(EngineRouter *gin.Engine)  {
	indexRoutes := EngineRouter.Group("/api/index")
	indexRoutes.Use(tool.Cors())
	{
		indexRoutes.POST("/getindexnewestitem",index.GetIndexNewestItem)//获取首页最新的5个商品
		indexRoutes.POST("/getindexranditem",index.GetIndexRandItem)//获取首页随机8个商品
		indexRoutes.POST("getindexitemlist",index.GetIndexItemLIst)//首页商品列表
		indexRoutes.POST("getindexbanner",index.GetIndexBanner)//首页banner
	}
}