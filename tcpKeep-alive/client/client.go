package main

import (
	"bufio"
	"fmt"
	"net"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	//返回tcp端口的地址
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	//拨号，连接到到tcp的ip端口上，获得连接
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected")
	go onMessageRecived(conn)
	for {
		fmt.Println("-------")
		var msg string
		fmt.Scan(&msg)
		if msg == "break" {
			break
		}
		b := []byte(msg + "\n")
		conn.Write(b)
	}

}
func onMessageRecived(con *net.TCPConn) {
	reader := bufio.NewReader(con)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
	}
}
