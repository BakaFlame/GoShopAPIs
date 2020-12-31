package TestModel

import (
	"GoShop/model"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

func SelectData(params string) interface{}{
	fmt.Println("查询啦")
	println("参数:",params)
	var user []User //查询多行需要切片声明, 单行first则用普通的 var user User
	if err:=model.DB.Find(&user).Error;err!=nil{
		log.Fatal(err.Error())
	} //select * from users
	return user //返回结构体给控制器
}

//插入语句测试 db.create
func InsertData(params string) bool{
	tx := model.DB.Begin()
	fmt.Println("插入啦")
	println("参数:",params)
	var user =  User{
		Name:       "charu",
		Age:        24,
		Content:    "charucontent",
		CreateTime: model.BetterTime{Time: time.Now()},
		UpdateTime: model.BetterTime{Time: time.Now()},
		Tag:        1,
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
	}
	tx.Commit()
	success:=model.DB.NewRecord(&user) //如果插入的主键没有找到就返回true，找到主键说明刚刚插入的数据成功返回false
	return !success //取反值好表达返回状态给控制器
}

func UpdateData(params string) bool{
	fmt.Println("更新啦")
	println("参数:",params)
	var user = User{ //可以直接指定结构体id进行更新
		ID: 13,
	}
	//更新时默认用主键id来做判断 sql: ... id == ???
	//model.DB.Model(&user).Update("name","update") 更新单列数据的时候用update
	model.DB.Model(&user).Updates(map[string]interface{}{"name":"takeitboy","age":10}) //更新多列数据的时候用updates 但要注意使用结构体的参数来进行更新 值如果为0 false ""  这种都算做为空值并且不对此进行更改 所以建议用map[string]interface{}{"键名":键值}

	success:=model.DB.NewRecord(&user) //如果找到已经被删除的数据那就说明没删到返回了true，如果删除了应该返回false
	return !success //取反值好表达返回状态给控制器
}

func DeleteData(params string) bool{
	fmt.Println("删除啦")
	println("参数:",params)
	var user = User{ //可以直接指定结构体id进行删除
		ID: 13,
	}
	model.DB.Delete(&user)
	success:=model.DB.NewRecord(&user) //如果找到已经被删除的数据那就说明没删到返回了true，如果删除了应该返回false
	return !success //取反值好表达返回状态给控制器
}

func RawSqlTest() string{
	data,err:=model.QuerySql("select * from test")
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func TestStructthing(testid string) string{
	data,_:=model.QuerySql("select * from test where id = "+testid)
	return data
}

func TestReturnRows() {
	test:=TestStruct{}
	record:=model.DB.Raw("select * from tests where id = 0").Scan(&test).RecordNotFound()
	fmt.Println(record)
	fmt.Println(test.ID)
}






