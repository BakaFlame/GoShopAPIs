package AddressController

import (
	"GoShop/model/AddressModel"
	"GoShop/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AddressStruct struct {
}

//获取地址列表接口
func (address *AddressStruct) GetAddressInfoByUserId(context *gin.Context) {
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	addressData, existBool := AddressModel.GetAddressInfoByUserId(userId)
	var data = map[string]interface{}{"addressdata": addressData, "exist": existBool}
	context.JSON(200, data)
}

//获取单条地址信息接口
func (address *AddressStruct) GetAddressInfoById(context *gin.Context) {
	id := context.PostForm("id")
	addressData := AddressModel.GetAddressInfoById(id)
	context.JSON(200, addressData)
}

//新增地址接口
func (address *AddressStruct) InsertAddressInfo(context *gin.Context) {
	realName := context.PostForm("realname")
	addressText := context.PostForm("address")
	phone := context.PostForm("phone")
	//realName:="测试插入"
	//addressText:="测试插入"
	//phone:="123123123"
	fmt.Println(realName)
	fmt.Println(addressText)
	fmt.Println(phone)
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	go AddressModel.InsertAddressInfo(realName, addressText, phone, userId)
	context.JSON(200, tool.ReturnData("新增地址成功", true, false))
}

//删除地址接口
func (address *AddressStruct) DeleteAddressById(context *gin.Context) {
	id := context.PostForm("id")
	if AddressModel.DeleteAddressById(id) {
		context.JSON(200, tool.ReturnData("删除成功", true, false))
	} else {
		context.JSON(200, tool.ReturnData("删除失败，请联系管理员", false, false))
	}
}

//更新地址接口
func (address *AddressStruct) UpdateAddressById(context *gin.Context) {
	realName := context.PostForm("realname")
	addressText := context.PostForm("address")
	phone := context.PostForm("phone")
	id := context.PostForm("id")
	token := context.GetHeader("thisisnotatoken")
	userId, _ := tool.GetUserInfoFromRedis(token)
	if AddressModel.UpdateAddressById(realName, addressText, phone, userId, id) {
		context.JSON(200, tool.ReturnData("修改成功", true, false))
	} else {
		context.JSON(200, tool.ReturnData("修改失败，请联系管理员", false, false))
	}
}
