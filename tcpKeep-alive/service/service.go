package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)
var ConnMap map[string]*net.TCPConn
func main() {
	var tcpAddr *net.TCPAddr
	ConnMap=make(map[string]*net.TCPConn)
	//获得tcp连接的ip，port
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	//启动本地连接，ip端口为tcpAddr
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		//接受下次请求，并生成新的conn
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected : ", tcpConn.RemoteAddr().String())
		ConnMap[tcpConn.RemoteAddr().String()]=tcpConn
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(string(message))
		mes := time.Now().String() + "\n"

		boradcastMessage(mes)

	}
}
func boradcastMessage(mes string){
	b:=[]byte(mes)
	for _,con:=range ConnMap {
		con.Write(b)
	}
}