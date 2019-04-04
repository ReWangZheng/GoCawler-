package parser

import (
	"regexp"
	"strings"
	"Main/model"
	"Main/engine"
	"log"
	"fmt"
)
var(
	baseRe=regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)
 	nicknameRe=regexp.MustCompile(`<span class="nickName" data-v-3c42fade>([^<]+)</span>`)
 	weightRe=regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([1-9]+kg)</div>`)
 	infoRe=regexp.MustCompile(`<div class="m-btn [a-z]+" data-v-bff6f798>([^<]+)</div>`)
	requ=regexp.MustCompile(`<div class="m-btn" data-v-bff6f798>([^<]+)</div>`)
)
func UserMes(document []byte,gender string) engine.ParseResult{
	fmt.Println("进入了！")
	unhandleMes := string(baseRe.FindSubmatch(document)[1])
	user:=handleMessage(unhandleMes,document)
	user.Gender=gender
	result:=engine.ParseResult{Requests:nil,}
	result.Items=append(result.Items,user)
	log.Println("\n姓名:"+user.Name+" 性别："+user.Gender+
				"\n教育:"+user.Education+" 收入:"+user.Income+
					"\n身高:"+user.Height+" 位置:"+user.Location)
	return result
}



func handleTag(tag [][][]byte)string{
	var tagmes string
	for _,va:=range tag{
		tagmes=tagmes+" | "+string(va[1])
	}
	return tagmes
}

func handleMessage(unhandleMes string,document []byte) model.UserMes{
	mesSlice := strings.Split(unhandleMes, " | ")
	location:=mesSlice[0]
	age:=mesSlice[1]
	education:=mesSlice[2]
	marriage:=mesSlice[3]
	height:=mesSlice[4]
	income:=mesSlice[5]
	user:=model.UserMes{
		Location:location,
		Age:age,
		Education:education,
		Marriage:marriage,
		Height:height,
		Income:income,
	}
	nickname:=string(nicknameRe.FindSubmatch(document)[1])
	infor:=handleTag(infoRe.FindAllSubmatch(document,-1))
	require:=handleTag(requ.FindAllSubmatch(document,-1))
	user.Name=nickname
	user.Height=height
	user.Infor=infor
	user.Require=require
	result:=engine.ParseResult{Requests:nil,}
	result.Items=append(result.Items,user)
	return user
}
