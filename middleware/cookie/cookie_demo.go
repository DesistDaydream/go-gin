package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		// 检查本次请求是否携带了名为 key_cookie 的 cookie,并返回该 cookie 的 value 属性。
		// 若不存在该 cookie，则 err 会报错 "http: named cookie not present"
		cookie, err := c.Cookie("key_cookie")

		// 如果 key_cookie 不存在，那么调用 c.SetCookie() 设置一个并响应给客户端
		if err != nil {
			// 第一次访问的话 cookie 的 value 不存在，咱手动设置一个，以便可以在后端调试时查看
			cookie = "NotSet"
			// 设置一个 cookie 的属性。这个 cookie 会包含在 Response 里响应给客户端
			// 客户端一般会保存该 cookie，并在下次访问时带上该 cookie。
			// 需要设置的参数依次为 name、value、MaxAge、path、domain、secure、httpOnly
			c.SetCookie("key_cookie", "value_cookie", 30, "/", "datalake.cn", false, true)
		}
		fmt.Printf("cookie 的值为：%v\n", cookie)
	})

	r.Run(":8080")
}
