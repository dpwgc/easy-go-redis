package cli

import (
	"fmt"
	"net"
)

//创建TCP连接
func newConn(addr string, port string) (*net.TCPConn, error) {

	tcpAddr := fmt.Sprintf("%s:%s", addr, port)
	server, err := net.ResolveTCPAddr("tcp4", tcpAddr)

	if err != nil {
		return nil, err
	}

	//建立服务器连接
	conn, err := net.DialTCP("tcp", nil, server)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

//从TCP连接中读取数据
func read(conn *net.TCPConn) (string, error) {
	buffer := make([]byte, 1024)
	msg, err := conn.Read(buffer) //接受服务器信息
	if err != nil {
		return "", nil
	}

	//[:msg]去除buffer数据前的无效字节
	return string(buffer[:msg]), nil
}

//向TCP服务端发送数据
func send(conn *net.TCPConn, data string) (string, error) {

	_, err := conn.Write([]byte(data)) //给服务器发信息
	if err != nil {
		return "", err
	}
	//获取服务端的反馈信息
	return read(conn)
}
