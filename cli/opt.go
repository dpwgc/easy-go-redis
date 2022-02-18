package cli

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Conn 与redis建立连接
func Conn(addr string, port string) (*net.TCPConn, error) {
	return newConn(addr, port)
}

// Do 执行命令
func Do(conn *net.TCPConn, arr ...string) (string, error) {

	//RESP命令字符串
	cmd := ""

	count := strconv.Itoa(len(arr))

	//拼接命令字符串
	cmd = fmt.Sprintf("%s%s%s\r\n", cmd, "*", count)

	for _, a := range arr {
		size := strconv.Itoa(len(a))
		cmd = fmt.Sprintf("%s%s%s\r\n", cmd, "$", size)
		cmd = fmt.Sprintf("%s%s\r\n", cmd, a)
	}

	return send(conn, cmd)
}

// Close 关闭连接
func Close(conn *net.TCPConn) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}

// Analytic 解析Redis返回的数据（将RESP的数据格式转为string字符串）
func Analytic(data string) []string {

	var res []string

	arr := strings.Split(data, "\r\n") //按换行符切割RESP字符串

	size := len(arr) - 1 //去掉末尾的空字符串

	//如果返回的字符串是单行
	//+OK
	if size == 1 {
		res = append(res, arr[0])
		return res
	}

	//如果返回的字符串是双行
	//$4
	//test
	if size == 2 {
		res = append(res, arr[1])
		return res
	}

	//如果返回的字符串是多行（三行及以上）
	//*2
	//$4
	//test
	//$5
	//hello
	for i := 2; i < size; i = i + 2 {
		res = append(res, arr[i])
	}

	return res
}
