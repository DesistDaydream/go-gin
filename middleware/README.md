# 目录结构

practice 是各种中间件练习的入口

# Gin 框架实现 Middleware(中间件)

![](https://raw.githubusercontent.com/DesistDaydream/PictureHosting/main/GoWeb/middleware.png)

假如我现在用 gin 框架实现了两个路由

```go
	r.GET("/no_middleware", func(c *gin.Context) {
		c.String(http.StatusOK, "No Use Middleware")
    })

    r.GET("/middleware",  func(c *gin.Context) {
        c.String(http.StatusOK, "Middleware Example")
    })
```

如果想让客户端在登录之后才可以访问 /middleware 页面，那么就需要添加一个处理逻辑，比如：

```go
	r.GET("/no_middleware", func(c *gin.Context) {
		c.String(http.StatusOK, "No Use Middleware")
    })

    // 注册中间件，每次访问中间件之后的路由，都会执行一次中间件的逻辑
    r.Use(MiddleWare())

    r.GET("/middleware",  func(c *gin.Context) {
        c.String(http.StatusOK, "Middleware Example")
    })
```

这样，每次访问 /middleware 的时候，都会执行 `MiddleWare()` 中的逻辑，这个逻辑里可以进行判断，看看本次请求是否带有 TOKEN 或者 Session 之类的认证信息，并与本身数据库中的数据进行匹配，匹配成功才能继续访问 /middleware

**注意：c.Abort() 是很重要的一个方法，用于中止中间件挂起的 Handler。比如验证中间件，当验证不通过时，需要阻止 Handler 继续运行，这时就需要添加一个 c.Abort()**

# 自己动手写一个实现 Session 的中间件

[README.md](./session/README.md)

参考：[老男孩老师在 GitHub 上的代码](https://github.com/Q1mi/ginsession)
