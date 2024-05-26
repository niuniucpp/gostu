package main

import "github.com/gin-gonic/gin"
import "fmt"
import _care "src/care"

func main() {
	fmt.Printf("*********************************\n")
	//gin.New()
	gin.SetMode(gin.ReleaseMode)
	// 接下来是你的路由和中间件设置等
	// ...
	// 启动你的 Gin 服务
	// ...
	fmt.Printf("*********************************\n")

	fmt.Printf("Hello\n")
	fmt.Printf("*********************************\n")

	a := 10
	b := 20
	fmt.Println(_care.SumCare(a, b))

	fmt.Printf("*********************************\n")
	_care.ReferenceCare(a, &b)
	fmt.Println(a, b)

	fmt.Printf("*********************************\n")
	fmt.Println(_care.G) // 0
	_care.ScopeCase(a, b)
	fmt.Println(_care.G) // 131

}
