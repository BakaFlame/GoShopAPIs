package CartController

import (
	"GoShop/middleWare"
	"GoShop/tool"
	"github.com/gin-gonic/gin"
)

func (cart *CartStruct) Route(EngineRouter *gin.Engine)  {
	cartRoutes := EngineRouter.Group("/api/cart")
	cartRoutes.Use(tool.Cors(),middleWare.IsLogin())
	{
		cartRoutes.POST("/getcartlistbyuserid",middleWare.IsLogin(),cart.GetCartListByUserId)//获取用户得购物车列表
		cartRoutes.POST("/itempageaddtocart",middleWare.IsLogin(),cart.ItemPageAddToCart)//在商品详情加入到购物车
		cartRoutes.POST("/updatecart",middleWare.IsLogin(),cart.UpdateCart)//购物车页面操作
	}
}