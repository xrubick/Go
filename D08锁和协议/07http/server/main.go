package main

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}

func main() {

	//有参数get请求server端
	http.HandleFunc("/", getHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server failed:", err)
		return
	}

}
