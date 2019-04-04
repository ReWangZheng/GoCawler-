package rpcdemo

import (
	"Main/dao"
	"net/rpc"
	"net"
	"fmt"
	"net/rpc/jsonrpc"
	"Main/model"
	"reflect"
	"log"
)

type DbDao struct {}
type Insertargs struct {
	Items []interface{}
}

type Workeages struct{

}


func (*DbDao)Insert(args Insertargs,result *bool) error{

	if len(args.Items)==1 {
		item:=args.Items[0]
		m,ok:=item.(map[string]interface{})
		if ok {

			user:=model.UserMes{
				Age:m["Age"].(string),
				Education:m["Education"].(string),
				Gender:m["Gender"].(string),
				Height:m["Height"].(string),
				Income:m["Income"].(string),
				Infor:m["Infor"].(string),
				Name:m["Name"].(string),
				Require:m["Require"].(string),
				Location:m["Location"].(string),
				Marriage:m["Marriage"].(string),
			}
			isok, err := dao.InsertUser(user)
			if err != nil {
				*result=isok
				return err
			}
			log.Println("插入数据:"+user.Name+"   "+user.Gender+"   "+user.Age)
		}
	}
	return nil

}

func Starservice() error{
	fmt.Println("远程数据存储引擎开始工作！")
	rpc.Register(&DbDao{})
	listener, e := net.Listen("tcp", ":3333")
	if e != nil {
		return e
	}
	for{
		conn, i := listener.Accept()
		if i != nil {
			return i
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}