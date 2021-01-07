package tool

func ReturnData(message string, status bool, redirect bool, args ...interface{}) interface{} {
	data := make(map[string]interface{})
	data["message"] = message
	data["status"] = status
	data["redirect"] = redirect //重定向 true为需要重新回到某个地方，false为不需要，默认为false
	data["extra"] = args
	return data
}
