package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//router:=gin.Default()：这是默认的服务器。使用gin的Default方法创建一个路由Handler；
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	router := gin.Default()
	//
	//创建不带中间件的路由:
	//r := gin.New()

	//启动路由的Run方法监听端口。还可以用http.ListenAndServe(":8080", router)，或者自定义Http服务器配置。
	router.Run(":8000")

	//http.ListenAndServe(":8000", router)

	//自定义HTTP服务器的配置：
	//s := &http.Server{
	//	Addr: ":8000",
	//	Handler: router,
	//	ReadHeaderTimeout: 10 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()
}
