package test__123

import (
	"fmt"
	"net"
	"time"
)

func server_start() {
	l, err := net.Listen("tcp", ":61000")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	fmt.Println("server_start listen")
	for {
		time.Sleep((time.Second * 10))
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			break
		}
		conn_work(conn)
	}
}

func conn_work(c net.Conn) {

}

func client_start() {
	service := ":61000"
	tcp := "tcp4"
	tcpAddr, err := net.ResolveTCPAddr(tcp, service)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("ResolveTCPAddr err:", err)
	}
	conn, err := net.DialTCP(tcp, nil, tcpAddr)
	if err != nil {
		fmt.Println("net.DialTCP err:", err)
		return
	}
	conn_work(conn)

	/*
		func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
			net 参数是 "tcp4"、"tcp6"、"tcp" 中的任意一个，分别表示 TCP(IPv4-only)、TCP(IPv6-only) 或者 TCP(IPv4,IPv6) 的任意一个；
			laddr 表示本机地址，一般设置为 nil；
			raddr 表示远程的服务地址。
	*/

}

func conn_work_bd(c net.Conn) {

}
func test_baidu() {
	service := "80.101.49.12:80"
	tcp := "tcp4"
	tcpAddr, err := net.ResolveTCPAddr(tcp, service)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("ResolveTCPAddr err:", err)
	}
	conn, err := net.DialTCP(tcp, nil, tcpAddr)
	if err != nil {
		fmt.Println("net.DialTCP err:", err)
		return
	}
	conn_work_bd(conn)
}

func TestNet() {
	//server_start()
	test_baidu()
}
