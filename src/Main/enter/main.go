package main

import (
	"Main/engine"
	"Main/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parser.CityList,
	})

}


