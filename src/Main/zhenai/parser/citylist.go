package parser

import (
	"Main/engine"
	"regexp"
)

func CityList(documentCity []byte)engine.ParseResult{


	//<a href="http://www.zhenai.com/zhenghun/laibin" data-v-5e16505f="">来宾</a>
	match:=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	city:=match.FindAllSubmatch(documentCity,-1)
	result:=engine.ParseResult{}
	for _,va:=range city{

		result.Items=append(result.Items,string(va[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(va[1]),
			ParseFunc:nillParser,
		})
	}
	return result
}

func nillParser(data []byte) engine.ParseResult{
	return engine.ParseResult{}
}
