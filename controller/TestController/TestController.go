package TestController

import (
	"GoShop/model/TestModel"
	"GoShop/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type TestStruct struct {

}

//解析 /test
func (test *TestStruct) testControl(context *gin.Context) {
	fmt.Println("test_controller func hello")
	context.Writer.Write([]byte("hduiashdi"))
}

func (test *TestStruct) testDb(context *gin.Context) {
	fmt.Println("test_controller func hello")
	data :=TestModel.SelectData("nothing")
	context.JSON(200,data)
}

func (test *TestStruct) testInsert(context *gin.Context) {
	fmt.Println("test_controller func hello")
	success :=TestModel.InsertData("nothing")
	context.JSON(200,success) //返回插入状态
}

func (test *TestStruct) testDelete(context *gin.Context) {
	fmt.Println("test_controller func hello")
	success :=TestModel.DeleteData("nothing")
	context.JSON(200,success) //返回删除状态
}

func (test *TestStruct) testUpdate(context *gin.Context) {
	fmt.Println("test_controller func hello")
	success :=TestModel.UpdateData("nothing")
	context.JSON(200,success) //返回更新状态
}

func (test *TestStruct) rawSelect(context *gin.Context) {
	context.JSON(200,TestModel.RawSqlTest()) //返回数据
}

func (test *TestStruct) testUpload(context *gin.Context) {
	fmt.Println("test_controller func hello")
	file,err:=context.FormFile("thisname") //拿到表单里input file的name
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(file.Filename)
	dst:=fmt.Sprintf("./%s",file.Filename)//本地服务器存放路径
	context.SaveUploadedFile(file,dst)//开始存放，file指网页上传的文件，dst为存放路径
	context.JSON(http.StatusOK,gin.H{ //存放成功后返回状态
		"message":fmt.Sprintf("%s",file.Filename),
	})
}

func (test *TestStruct) uploadPage(context *gin.Context) { //渲染上传页面
	fmt.Println("test_controller func hello")
	context.HTML(http.StatusOK, "upload.html", gin.H{
		"title": "Main website",
	})
}

func (test *TestStruct) jsonPage(context *gin.Context) { //渲染上传页面
	fmt.Println("test_controller func hello")
	context.HTML(http.StatusOK, "json.html", gin.H{
		"title": "Main website",
	})
}

func (test *TestStruct) teststructthing(context *gin.Context) {
	var stringdata = []string{}
	str :=""
	finalstr:=""
	stringdata = append(stringdata, "1")
	for i := 0; i < len(stringdata); i++ {
		str+=TestModel.TestStructthing(stringdata[i])
	}
	finalstr=strings.Replace(str,"][",",",-1)
	fmt.Println(finalstr)
}

func (test *TestStruct) testjson(context *gin.Context) {
	fmt.Println(TestModel.RawSqlTest())
	context.JSON(200,TestModel.RawSqlTest())
}

func (test *TestStruct) nocors(context *gin.Context) {
	fmt.Println(TestModel.RawSqlTest())
	context.JSON(200,TestModel.RawSqlTest())
}

func (test *TestStruct) testget(context *gin.Context) {
	formthing := context.Query("thing")
	fmt.Println(formthing)
}

func (test *TestStruct) testrows(context *gin.Context) {
	TestModel.TestReturnRows()
}

func (test *TestStruct) testfakesession(context *gin.Context) {
	token:=context.GetHeader("thisisnotatoken")
	encodePassword:=tool.SHA256EncodeWithSalt("mima")
	fmt.Println(encodePassword)
	fmt.Println(token+"\theader没加密")
	fmt.Println(tool.SHA256EncodeWithSalt(token)+"\t加密")
	tool.SHA256Decode(token)
}