package remote

import (
	"net"
	"net/rpc/jsonrpc"
	"fmt"
	"Main/rpcdemo"
)

func Insert(items []interface{}){
	conn, e := net.Dial("tcp", ":3333")
	if e != nil {
		panic(e)
	}
	var b bool
	client:=jsonrpc.NewClient(conn)
	call := client.Call("DbDao.Insert", rpcdemo.Insertargs{items}, &b)
	if call != nil {
		panic(call)
	}
	if b {
		fmt.Println("插入成功！")
	}


}


