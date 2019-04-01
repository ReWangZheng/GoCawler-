package engine

import (
	"Main/fetcher"
	"log"
)

func Run(seed ...Request){

	var taskQueen []Request

	for _,va:=range seed{
		taskQueen=append(taskQueen,va)
	}

	for;len(taskQueen)>0;{
		for _,task:=range taskQueen{
			document,err:=fetcher.Getdocument(task.Url)
			log.Printf("访问的URL：%s",task.Url)
			if err!=nil{
				log.Printf("出现错误%v",err)
				continue
			}
			result:=task.ParseFunc(document)
			taskQueen=append(taskQueen,result.Requests...)

			for _,i:=range result.Items{
				log.Printf("Item:%s",i)
			}

		}


	}



}
