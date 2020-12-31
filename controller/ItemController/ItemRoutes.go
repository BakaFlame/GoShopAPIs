package ItemController

import (
	"GoShop/tool"
	"github.com/gin-gonic/gin"
)

func (item *ItemStruct) Route(EngineRouter *gin.Engine)  {
	itemRoute := EngineRouter.Group("/api/item")
	itemRoute.Use(tool.Cors())
	{
		itemRoute.GET("/itemlisttest",item.JsonItemPage)
		itemRoute.POST("getbrandlist",item.GetBrandList)//获取品牌列表
		itemRoute.POST("gettypelist",item.GetTypeList)//获取种类列表
		itemRoute.POST("getitemlistbytypeandbrand",item.GetItemListByTypeAndBrand)//按照种类和品牌查询商品列表
		itemRoute.POST("getitemlistbyitemname",item.GetItemListByItemName)//按照名字来查询商品列表
		itemRoute.POST("getiteminfobyid",item.GetItemInfoById)//获取商品详情信息
	}
}
