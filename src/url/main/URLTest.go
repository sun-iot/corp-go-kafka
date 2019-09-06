package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"url/urlutil"
)

func main() {
	// 设置访问的路由
	http.HandleFunc("/", urlutil.HttpRequest)
	// 设置监听的端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
