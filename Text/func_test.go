package Text

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func Text(t *testing.T){

	datas, e := ioutil.ReadFile("docu.txt")
	if e != nil {
		panic(e)
	}

	fmt.Println("okokok?"+string(datas))




}


