package engine



func (e SimpleEngine)Run(seed ...Request){
	var taskQueen []Request

	for _,va:=range seed{
		taskQueen=append(taskQueen,va)
	}

	for;len(taskQueen)>0;{
		for _,task:=range taskQueen{
			taskQueen=taskQueen[1:]
			result,err:=worker(task)
			if err != nil {
				continue
			}
			taskQueen=append(taskQueen,result.Requests...)
		}
	}
}




