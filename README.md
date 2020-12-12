# Gin 介绍
```go
package main
import "github.com/gin-gonic/gin"
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```
这是官方的入门介绍，对于编程新手来说不太友好,可以改成下面这个样子  
```go
package main
import "github.com/gin-gonic/gin"
func main() {
	r := gin.Default()
	r.GET("/ping", PingGet)
	r.Run()
}

func PingGet(c *gin.Context) {
	h := gin.H{
		"message": "pong",
	}

	c.JSON(200, h)
}
```
具体含义，详解[main.go](main.go)

# Features gin的特性
## 绑定
1. Gin 提供了非常方便的数据绑定功能，可以将用户传来的参数自动跟我们定义的结构体绑定在一起。  
1. 模型绑定可以将请求体绑定给一个类型，目前支持绑定的类型有 JSON, XML 和标准表单数据 (foo=bar&boo=baz)。  
1. 绑定时需要给字段设置绑定类型的标签。比如绑定 JSON 数据时，设置 json:"fieldname"。 使用绑定方法时，Gin 会根据请求头中 Content-Type 来自动判断需要解析的类型。如果你明确绑定的类型，你可以不用自动推断，而用 BindWith 方法。  
1. 可以指定某字段是必需的。如果一个字段被 binding:"required" 修饰而值却是空的，请求会失败并返回错误。  

详见[binding.go](./features/binding.go)

# Gin 热更新
为了使代码发生变化时，可以自动编译加载，而不用重新 `go run` 可以通过 fresh 工具实现  
`go get github.com/pilu/fresh`  
获取 fresh 之后，在项目目录直接执行 `fresh` 命令即可。fresh命令会代理 go run 命令来执行程序，并且监控代码文件，当发生变化时，可以自动build  