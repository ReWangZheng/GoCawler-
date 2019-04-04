package Text

import (
	"regexp"
	"fmt"
)

func mains(){
	text:="my email is 29370330@qq.com " +
		"" +
			"51543@163.com " +
				"5dasc1@qq.com"
	re:=regexp.MustCompile(`([a-zA-z0-9]+)+@([a-zA-z0-9.]+)+\.([a-zA-z0-9]+)`)
	mubiao:=re.FindAllStringSubmatch(text,-1)
	for _,str:=range mubiao{
		fmt.Println(str)
	}
 }
