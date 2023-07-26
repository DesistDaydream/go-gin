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

这是官方的入门介绍，对于编程新手来说不太友好,可以改成下面这个样子:

```go
package main
import "github.com/gin-gonic/gin"
func PingGet(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func main() {
	r := gin.Default()
	r.GET("/ping", PingGet)
	r.Run()
}
```

具体含义，详解[main.go](main.go)

其中有两个概念非常重要：

- `r.GET()` 用来处理 GET 的请求，一般称之为 **Handler(处理器，用来处理 http 请求的处理器)** 。还包括 r.POST、r.PUT 等等对应处理各种不同类型请求的 Handler。
  > `r.GET()` 是 `router.Handle("GET", path, handle)` 的简化版，一般也成为 路由处理器、路由注册器 等等。这种行为也成为注册路由，一般都放在一个单独目录里统一注册，再由 main() 调用
- `gin.Context` 这个结构体是非常重要的，**net/http 基本库** 的 `http.ResponseWriter` 与 `*http.Request` 就包含在这个结构体中，所以

# Features gin 的特性

## 绑定

1. Gin 提供了非常方便的数据绑定功能，可以将用户传来的参数自动跟我们定义的结构体绑定在一起。
1. 模型绑定可以将请求体绑定给一个类型，目前支持绑定的类型有 JSON, XML 和标准表单数据 (foo=bar&boo=baz)。
1. 绑定时需要给字段设置绑定类型的标签。比如绑定 JSON 数据时，设置 json:"fieldname"。 使用绑定方法时，Gin 会根据请求头中 Content-Type 来自动判断需要解析的类型。如果你明确绑定的类型，你可以不用自动推断，而用 BindWith 方法。
1. 可以指定某字段是必需的。如果一个字段被 binding:"required" 修饰而值却是空的，请求会失败并返回错误。

详见[binding.go](./features/binding.go)

# Gin 热更新

为了使代码发生变化时，可以自动编译加载，而不用重新 `go run` 可以通过 fresh 工具实现

`go get github.com/pilu/fresh`

获取 fresh 之后，在项目目录直接执行 `fresh` 命令即可。fresh 命令会代理 go run 命令来执行程序，并且监控代码文件，当发生变化时，可以自动 build
