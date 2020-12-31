package tool

import "sync"

// 协程工具 用于处理用户不重要并且不需要看到或者不及时的数据（例如点击量增加）
// 这样可以不用等待更新数据库时间直接先返回页面

func init(){ // init函数 在调用此包的时候第一优先级调用此函数
	chanList:=getTaskList()//得到任务列表
	go func(){
		for t:=range chanList{
			t.Exec()//执行任务
		}
	}()
}

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor //任务列表
var once sync.Once

func getTaskList() chan *TaskExecutor{
	once.Do(func(){
		taskList=make(chan *TaskExecutor) // 初始化taskList
	})
	return taskList
}



type TaskExecutor struct {
	f TaskFunc
	params []interface{}
}

func (this *TaskExecutor) Exec(){
	this.f(this.params)
}

func NewTaskExecutor(f TaskFunc, params []interface{}) *TaskExecutor {
	return &TaskExecutor{f: f, params: params}
}

func Task(f TaskFunc,params ...interface{}){
	getTaskList()<-NewTaskExecutor(f,params) //增加到队列
}