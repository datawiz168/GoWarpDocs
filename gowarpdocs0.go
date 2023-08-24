package main

import (
	"fmt"
	"sync"
)

// DB 表示一个简单的键值数据库
type DB struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewDB 创建一个新的数据库实例
func NewDB() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

// Put 存储一个键值对
func (db *DB) Put(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

// Get 检索一个键的值
func (db *DB) Get(key string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, exists := db.data[key]
	return value, exists
}

func main() {
	db := NewDB()
	db.Put("name", "GoWarpDocs")
	value, _ := db.Get("name")
	fmt.Println("Value:", value) // 输出: Value: GoWarpDocs
}
