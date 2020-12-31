package IndexController

import (
	"GoShop/model/IndexModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IndexStruct struct {

}

////随机种类4件物品展示
//func (index *IndexStruct) GetIndexRandTypeItem(context *gin.Context) {
//	rand.Seed(time.Now().Unix())
//	randType := rand.Intn(9) + 1
//	fmt.Println(randType)
//	data := IndexModel.GetIndexRandTypeItem(randType)
//	// fmt.Println(data)
//	context.JSON(200, data)
//}

//随机6件物品展示 //2020.12.30更改为8个随机商品
func (index *IndexStruct) GetIndexRandItem(context *gin.Context) {
	data := IndexModel.GetIndexRandItem()
	context.JSON(200, data)
}
//查询最新5个物品
func (index *IndexStruct) GetIndexNewestItem(context *gin.Context) {
	data := IndexModel.GetIndexNewestItem()
	context.JSON(200, data)
}

//下方商品列表 4个一行
func (index *IndexStruct) GetIndexItemLIst(context *gin.Context) {
	pageData := context.PostForm("page") //page从0开始
	page, err := strconv.Atoi(pageData)
	if err != nil {
		fmt.Println(err)
	}
	context.JSON(200, IndexModel.GetIndexItemLIst(page))
}

//首页banner
func (index *IndexStruct) GetIndexBanner(context *gin.Context) {
	context.JSON(200,IndexModel.GetIndexBanner())
}