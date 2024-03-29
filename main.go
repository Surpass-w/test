package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// 设置路由，将路径 "/" 映射到 helloWorldHandler 函数
	http.HandleFunc("/", helloWorldHandler)

	// 启动HTTP服务器并监听8000端口
	fmt.Println("Server listening on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
