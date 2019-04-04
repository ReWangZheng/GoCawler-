package main

import (
	"Main/dao"
	"Main/model"
)

func main(){
	user:=model.UserMes{Name:"测试数据"}
	dao.InsertUser(user)
}
func checkError(err error){
	if err != nil {
		panic(err)
	}
}
