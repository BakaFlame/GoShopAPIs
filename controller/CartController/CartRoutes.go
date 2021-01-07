package CartController

import (
	"GoShop/middleWare"
	"github.com/gin-gonic/gin"
)

func (cart *CartStruct) Route(EngineRouter *gin.Engine) {
	cartRoutes := EngineRouter.Group("/api/cart")
	cartRoutes.Use(middleWare.IsLogin())
	{
		cartRoutes.POST("/getcartlistbyuserid", cart.GetCartListByUserId) //获取用户得购物车列表
		cartRoutes.POST("/itempageaddtocart", cart.ItemPageAddToCart)     //在商品详情加入到购物车
		cartRoutes.POST("/updatecart", cart.UpdateCart)                   //购物车页面操作
	}
}
