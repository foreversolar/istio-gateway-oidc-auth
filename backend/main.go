package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 设置HTTP路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend service at path: %s", r.URL.Path)
	})

	// 启动HTTP服务监听在8080端口
	fmt.Println("Starting server at port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
