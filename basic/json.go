package main

// http://cizixs.com/2016/12/19/golang-json-guide

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string          `json:name`
	IsAdmin bool            `json:is_admin`
	Auth    json.RawMessage `json:auth`
}

func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}
	// 尽管 status 字段没有小数点，我们希望它是整数值，解析的时候它还是会被解析成 float64，如果直接把它当做 int 访问，会出现错误：
	//
	// var status = result["status"].(int) //error
	// fmt.Println("status value:", status)

	var userJson = []byte(`{"name": "Wang", "is_admin": false, "auth": { "token": "token1" } }`)
	var user User
	if err := json.Unmarshal(userJson, &user); err != nil {
		fmt.Println("error:", err)
		return
	}
	json.Unmarshal(user.Auth, &result)
	fmt.Println("status value:", result["token"])
}
