package AddressController

import (
	"GoShop/middleWare"
	"github.com/gin-gonic/gin"
)

func (address *AddressStruct) Route(EngineRouter *gin.Engine) {
	addressRoutes := EngineRouter.Group("/api/address")
	addressRoutes.Use(middleWare.IsLogin())
	{
		addressRoutes.POST("/getaddressinfobyuserid", address.GetAddressInfoByUserId) //获取地址列表接口
		addressRoutes.POST("/getaddressbyid", address.GetAddressInfoById)             //获取单条地址信息接口
		addressRoutes.POST("/insertaddressinfo", address.InsertAddressInfo)           //新增地址接口
		addressRoutes.POST("/deleteaddressbyid", address.DeleteAddressById)           //删除地址接口
		addressRoutes.POST("/updateaddressbyid", address.UpdateAddressById)           //更新地址接口
	}
}
