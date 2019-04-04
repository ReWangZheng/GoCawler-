package engine

import (
	"Main/remote"
)



func (e *CurrentEngine)Run(seeds ...Request){
	e.Scheduler.Run()
	out:=make(chan ParseResult)
	for i:=0;i<e.WorkeCount;i++{
		CreateWorker(out,e.Scheduler)
	}
	for _,seed:=range seeds{
		e.Scheduler.Submit(seed)
	}

	for {
		result:=<-out
		tasks:=result.Requests
		items:=result.Items
		remote.Insert(items)
		for _,task:=range tasks{
			e.Scheduler.Submit(task)
		}
	}

}

func CreateWorker(resulchan chan ParseResult,s Notifier){
	go func(){
		for {
			inch:=s.GetChan()
			s.ReadyWorker(inch)
			request := <-inch
			result, err :=worker(request)
			if  err != nil {
				continue
			}
			resulchan <- *result
		}
	}()

}



