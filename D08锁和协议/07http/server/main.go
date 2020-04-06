package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//handle 带参数的get请求
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}

//handle post请求
func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// //请求类型为application/x-www-form-urlencoded时解析form数据
	// r.ParseForm()
	// fmt.Println(r.PostForm) // 打印form数据
	// fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	//请求类型为application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request failed:", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {
	//有参数get请求server端
	//http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server failed:", err)
		return
	}
}
