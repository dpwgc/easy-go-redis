# easy-go-redis

## 一个简易的Redis Golang客户端

***

### 导入客户端

* 引入包：`go get github.com/dpwgc/easy-go-redis`

***

### 函数说明

#### cli/opt.go

##### Conn

```
// Conn 与redis建立连接
func Conn(addr string,port string) (*net.TCPConn,error)

addr: Redis服务器的IP地址
port: Redis的服务端口号
return: TCP连接及报错信息
```

##### Do

```
// Do 执行命令
func Do(conn *net.TCPConn,arr ...string) (string,error)

conn: TCP连接
arr: 命令关键字切片
return: 命令执行后服务器的反馈信息及报错信息
```

##### Close

```
// Close 关闭连接
func Close(conn *net.TCPConn) error

conn: TCP连接
return: 报错信息
```

##### Analytic

```
// Analytic 解析Redis返回的数据（将RESP的数据格式转为string字符串）
func Analytic(data string) []string

data: 命令执行后服务器的反馈信息
return: 解析后的反馈信息数组
```

***

### 使用示范

```go
import (
    "fmt"
    "github.com/dpwgc/easy-go-redis/cli"
    "strconv"
)

func test()  {

	//与redis建立连接
	conn,err := cli.Conn("127.0.0.1","6379")
	if err != nil {
		panic(err)
	}
	
	

	//输入密码
	res,err := cli.Do(conn,"auth","123456")
	if err != nil {
		panic(err)
	}
	//解析并输出服务端的反馈信息
	fmt.Println(cli.Analytic(res))
	
	

	//插入一条key为"test"，value为"hello world"的string类型数据
	res,err = cli.Do(conn,"set","test","hello world")
	if err != nil {
		panic(err)
	}
	//解析并输出服务端的反馈信息
	fmt.Println(cli.Analytic(res))
	
	

	//获取key为"test"的value
	res,err = cli.Do(conn,"get","test")
	if err != nil {
		panic(err)
	}
	//解析并输出服务端的反馈信息
	fmt.Println(cli.Analytic(res))
	
	

	//循环往list中添加数据
	for i:=0;i<3;i++ {
		data := "hi-" + strconv.Itoa(i) //"hi-1","hi-2","hi-3"
		res,err = cli.Do(conn,"lpush","test_list",data)
		if err != nil {
			panic(err)
		}
		//解析并输出服务端的反馈信息
		fmt.Println(cli.Analytic(res))
	}
	
	

	//读取列表中指定区间的元素
	res,err = cli.Do(conn,"lrange","test_list","0","10")
	//解析并输出服务端的反馈信息
	fmt.Println(cli.Analytic(res))
	
	

	//关闭连接
	err = cli.Close(conn)
	if err != nil {
		panic(err)
	}
}
```