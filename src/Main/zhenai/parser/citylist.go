package parser

import (
	"Main/engine"
	"regexp"
)
//<div class="m-btn purple" data-v-bff6f798="">未婚</div>


func CityList(documentCity []byte)engine.ParseResult{


	//<a href="http://www.zhenai.com/zhenghun/laibin" data-v-5e16505f="">来宾</a>
	match:=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	citys:=match.FindAllSubmatch(documentCity,-1)
	result:=engine.ParseResult{}
	for _,va:=range citys{

		result.Items=append(result.Items,string(va[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(va[1]),
			ParseFunc:CityMes,
		})
	}
	return result
}