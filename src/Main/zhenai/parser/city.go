package parser

import (
	"Main/engine"
	"regexp"
)

//<a href="http://album.zhenai.com/u/1066286717" target="_blank">张坚强吖</a>
//<a href="http://album.zhenai.com/u/1211262040" target="_blank">在路上</a>
//<a href="http://album.zhenai.com/u/1755367918" target="_blank">西贝子</a>
//<td width="180"><span class="grayL">性别：</span>女士</td>
//<a href="http://www.zhenai.com/zhenghun/shizhu/6">下一页</a>
func CityMes(document []byte) engine.ParseResult{

	re:=regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^</]+)</a>`)
	genderRe:=regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
	next:=regexp.MustCompile(`<a href=("http://www.zhenai.com/zhenghun/[a-zA-Z]+/[0-9]+)">下一页</a>`)
	data:=re.FindAllSubmatch(document,-1)
	nexturl:=next.FindSubmatch(document)
	genderData := genderRe.FindAllSubmatch(document, -1)
	result:=engine.ParseResult{}
	for i:=0;i<len(data);i++{
		url:=string(data[i][1])
		gender:=string(genderData[i][1])
		result.Requests=append(result.Requests,engine.Request{
			Url:url,
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return UserMes(bytes,gender)
			},
		})
	}

	if len(nexturl) > 0 {
		url:=string(nexturl[1])
		result.Requests=append(result.Requests, engine.Request{
			Url:url,
			ParseFunc:CityMes,
		})
	}

	return result
}
