package main

import (
	"demo/api"
	"demo/dao"
)

func main() {
	dao.InitMySQL()
	//dao.InitRedis()
	api.InitRouter()
}
