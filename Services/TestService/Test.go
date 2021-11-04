package testservice

import (
	"github.com/putao520/ourjson"
)

type Test struct {
	Id   int    // public field
	Name string // public field
}

func New() *Test {
	return &Test{
		Id:   0,
		Name: "",
	}
}

// 创建 post
func create(data *ourjson.JsonObject) int {

}

// 更新 put
func update(id int, data *ourjson.JsonObject) int {

}

// 读取 get
func fetchAll() *ourjson.JsonArray {

}

// 读取特定 get
func fetch(id int) *ourjson.JsonObject {

}

// 删除
func delete(id int) int {

}
