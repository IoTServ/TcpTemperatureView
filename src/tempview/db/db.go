package db

import ("runtime"
	"net"
	"fmt"
	"time"
	"tempview/models"
	"github.com/astaxie/beego/orm"
)

func storagehandler(conn net.Conn) {
	o := orm.NewOrm()
	for {
		var recv []byte = make([]byte, 10240)
		if conn.SetReadDeadline(time.Now().Add(time.Minute*60))!=nil{
			conn.Close()
			runtime.Goexit()
		}
		n,e:=conn.Read(recv)
		if e!=nil{
			conn.Close()
			runtime.Goexit()
		}
		if conn.SetReadDeadline(time.Time{}) !=nil{
			conn.Close()
			runtime.Goexit()
		}
		msg:=recv[:n]
		if string(msg)!="" {
			user := models.DEVICE{VALUE:string(msg) , DATE:time.Now()}

			// insert
			id, _ := o.Insert(&user)
			fmt.Printf("Insert "+string(id)+" VALUE: "+string(msg))
		}
	}
}

func init()  {
	go func() {
		l, _ := net.Listen("tcp", ":8002")
		for {
			clientconnn, _ := l.Accept()
			go storagehandler(clientconnn)
		}
	}()
}

