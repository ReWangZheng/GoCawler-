package main

import (
	_ "Main/zhenai/parser"
	_ "Main/engine"
	//"Main/zhenai/parser"
	//"Main/engine"
	"Main/engine"
	"Main/zhenai/parser"
	"Main/Scheduler"
)


func main() {
	c:=Scheduler.QueuedScheduler{}

	MainControl:=engine.CurrentEngine{&c,150}

	MainControl.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parser.CityList,
	})


}


