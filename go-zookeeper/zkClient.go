package main

import(
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var(
	path = "/test"
)

// 增加一个节点（zNode）
func add(conn *zk.Conn)  {
	var data = []byte("test value")
	// flags有4种取值：
	// 0:永久，除非手动删除
	// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
	// zk.FlagSequence  = 2:会自动在节点后面添加序号
	// 3:Ephemeral和Sequence，即，短暂且自动添加序号
	var flags int32 = 0
	acls := zk.WorldACL(zk.PermAll)
	s, err := conn.Create(path, data, flags, acls)
	if err != nil{
		fmt.Printf("创建失败：%v\n", err)
		return
	}
	fmt.Printf("创建节点成功：%s\n", s)
}

// 查
func get(conn *zk.Conn)  {
	data, _, err := conn.Get(path)
	if err != nil {
		fmt.Printf("查询%s失败，err：%v\n", path, err)
		return
	}
	fmt.Printf("%s 的值为 %s\n", path, string(data))
}
// 删改与增不同在于其函数中的version参数,其中version是用于 CAS支持
// 可以通过此种方式保证原子性
// 改
func modify(conn *zk.Conn)  {
	newData := []byte("hello zookeeper")
	_, sate, _ := conn.Get(path)
	_, err := conn.Set(path, newData, sate.Version)
	if err != nil{
		fmt.Printf("数据修改失败：%v\n", err)
		return
	}
	fmt.Printf("数据修改成功：%s\n",newData)
}

// 删
func del(conn *zk.Conn)  {
	delData, sate, _ := conn.Get(path)
	err := conn.Delete(path, sate.Version)
	if err != nil {
		fmt.Printf("数据删除失败：%v\n", err)
		return
	}
	fmt.Printf("数据删除成功:%s\n", delData)
}

func main() {
	// 创建zk连接地址
	hosts := []string{"127.0.0.1:2182"}
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	defer conn.Close()

	if err != nil{
		fmt.Println(err)
		return
	}

	//println(conn.Server())
	add(conn)
	//get(conn)
	//modify(conn)
	//del(conn)
}

