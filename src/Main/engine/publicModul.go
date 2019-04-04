package engine

import (
	"Main/fetcher"
	"log"
)

type CurrentEngine struct {
	Scheduler
	WorkeCount int
}

type Request struct {
	Url string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items	[]interface{}
}

type SimpleEngine struct{}

//调度器区域
type Scheduler interface {
	Submit(request Request)
	Run()
	Notifier
}

type Notifier interface{
	GetChan()(chan Request)
	ReadyWorker(worker chan Request)
}


//函数组件区域
func worker(task Request)(*ParseResult,error){
	document,err:=fetcher.Getdocument(task.Url)
	log.Printf("访问的URL：%s",task.Url)
	if err!=nil{
		log.Printf("出现错误%v",err)
		return nil,err
	}
	result:=task.ParseFunc(document)
	return &result,nil
}

