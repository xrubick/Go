package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// //Client发送http请求

// func main() {
// 	resp, err := http.Get("https://www.baidu.com")
// 	if err != nil {
// 		fmt.Println("get failed:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("read resp body failed:", err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }

// //带参数的get请求
// func main() {
// 	//get请求
// 	apiurl := "http://127.0.0.1:9000/get"
// 	data := url.Values{}
// 	data.Set("name", "二哈")
// 	data.Set("age", "2")
// 	u, err := url.ParseRequestURI(apiurl)
// 	if err != nil {
// 		fmt.Println("parse uri failed:", err)
// 	}

// 	//url encode
// 	u.RawQuery = data.Encode()
// 	fmt.Println(u.String())
// 	resp, err := http.Get(u.String())
// 	if err != nil {
// 		fmt.Println("get failed:", err)
// 		return
// 	}
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("get resp failed:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	fmt.Println(string(b))
// }

//post请求
func main() {
	url := "http://127.0.0.1:9000/post"
	// // 表单数据
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=二哈&age=2"
	// json
	contentType := "application/json"
	data := `{"name":"二哈","age":2}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed:", err)
		return
	}
	fmt.Println(string(b))
}
