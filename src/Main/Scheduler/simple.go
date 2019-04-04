package Scheduler


import "Main/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler)Submit(request engine.Request){
	go func() {s.workChan<-request}()
}

func (s *SimpleScheduler)ConfigScheduler(ch chan engine.Request){
		s.workChan=ch

}

func (s *SimpleScheduler)GetChan()(chan engine.Request){
	return s.workChan
}

func (s *SimpleScheduler)ReadyWorker(worker chan engine.Request){
}

func (s *SimpleScheduler)Run(){
	s.workChan=make(chan engine.Request)
}
