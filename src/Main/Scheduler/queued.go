package Scheduler

import (
	"Main/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}


func (s *QueuedScheduler)Submit(request engine.Request){
	s.requestChan<-request
}

//参数是woker的输入接口
func (s *QueuedScheduler)ReadyWorker(worker chan engine.Request){
	s.workerChan<-worker
}

func (s *QueuedScheduler)GetChan() chan engine.Request{
	return make(chan engine.Request)
}

func (s *QueuedScheduler)Run(){
	s.workerChan=make(chan chan engine.Request)
	s.requestChan=make(chan engine.Request)
	var RequestQuene []engine.Request //任务队列
	var WorkerQuene []chan engine.Request  //对外输出的woker接口队列
	go func(){
		for {
			var doworke chan engine.Request
			var request engine.Request
			if len(RequestQuene)>0 && len(WorkerQuene)>0 {
				doworke=WorkerQuene[0]
				request=RequestQuene[0]
			}
			select {
			case task := <-s.requestChan:
				RequestQuene = append(RequestQuene, task)
			case worker := <-s.workerChan:
				WorkerQuene = append(WorkerQuene, worker)
			case doworke <- request:
				WorkerQuene=WorkerQuene[1:]
				RequestQuene=RequestQuene[1:]
			}
		}

	}()
}

