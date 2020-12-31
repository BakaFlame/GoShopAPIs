package CartController

import (
	"GoShop/tool"
	"github.com/gin-gonic/gin"
)

func (banner *BannerStruct) Route(EngineRouter *gin.Engine)  {
	bannerRoutes := EngineRouter.Group("/api/banner")
	bannerRoutes.Use(tool.Cors())
	{

	}
}