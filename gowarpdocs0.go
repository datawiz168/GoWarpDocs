package main // 定义包名为 main，表示可执行程序的入口包

import (
	"log" // 导入日志包，用于记录错误或信息
	"github.com/wiredtiger/go-wiredtiger/wt" // 导入 WiredTiger 的 Golang 绑定
)

// DB 代表我们的数据库实例
type DB struct {
	conn *wt.Conn // WiredTiger 连接对象
}

// NewDB 创建一个新的数据库实例
func NewDB(path string) (*DB, error) {
	conn, err := wt.Open(path, "create") // 打开或创建 WiredTiger 连接
	if err != nil {
		return nil, err // 返回连接错误
	}

	return &DB{conn: conn}, nil // 返回数据库实例
}

// CreateTable 创建一个新的数据表
func (db *DB) CreateTable(name string) error {
	session, err := db.conn.OpenSession("") // 打开 WiredTiger 会话
	if err != nil {
		return err // 返回会话错误
	}
	defer session.Close() // 确保会话在函数结束时关闭

	// 创建数据表
	_, err = session.Create("table:"+name, "key_format=S,value_format=S")
	return err // 返回创建表的错误（如果有）
}

// Insert 插入一条记录到指定的数据表
func (db *DB) Insert(table, key, value string) error {
	session, err := db.conn.OpenSession("") // 打开 WiredTiger 会话
	if err != nil {
		return err // 返回会话错误
	}
	defer session.Close() // 确保会话在函数结束时关闭

	cursor, err := session.OpenCursor("table:"+table, nil, "") // 打开数据表的游标
	if err != nil {
		return err // 返回游标错误
	}
	defer cursor.Close() // 确保游标在函数结束时关闭

	cursor.SetKey(key)   // 设置键
	cursor.SetValue(value) // 设置值
	return cursor.Insert() // 插入记录并返回任何错误
}

func main() {
	// 创建数据库实例
	db, err := NewDB("data")
	if err != nil {
		log.Fatalf("Failed to create DB: %v", err) // 记录并退出错误
	}

	// 创建表
	if err := db.CreateTable("gowarpdocs"); err != nil {
		log.Fatalf("Failed to create table: %v", err) // 记录并退出错误
	}

	// 插入数据
	if err := db.Insert("gowarpdocs", "key1", "value1"); err != nil {
		log.Fatalf("Failed to insert data: %v", err) // 记录并退出错误
	}

	log.Println("Database initialized successfully!") // 记录成功信息
}
