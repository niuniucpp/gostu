package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello golang?????", r.URL.Path)
}
func main() {
	http.HandleFunc("/", header)
	fmt.Printf("正在运行中……")
	//创建路由
	http.ListenAndServe("localhost:8080", nil)
	fmt.Printf("程序结束")

}
