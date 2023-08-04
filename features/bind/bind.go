// gin 有一个 bind 功能，可以将客户端传来的参数与我们自己定义的结构体中的属性绑定在一起
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoginForm 该结构体属性中的TAG，用来作为bind的依据。
// - form:"user" 表示User与表单中的user绑定
// - binding:"required" 是一个规则修饰符，gin读取到的时候，表明该属性必须绑定，如果字段为空，则会报错。
// Gin 默认规定使用 form 这个 Tag 作为绑定的依据。只有具有 form Tag 的字段才会与请求时传入的各种参数（URL 的 Query、请求体、等等）相关联。
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 声明一个结构体的变量，将form中传递的参数与该变量结构体中的属性绑定
		var form LoginForm
		// 可以使用显式绑定声明绑定 form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定。在这种情况下，将自动选择合适的绑定
		if c.ShouldBind(&form) == nil {
			// 绑定完成后，就可以使用绑定后的变量，来直接调用结构体中的属性。
			// 其中form.User这个结构体中属性的值，就是表单中user的值。
			// 并且两者的类型都相同，比如结构体中如果是int类型，那么表单中的数据到代码中依然是int类型。
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
		// 当然，如果不使用 binding 特性，也可以直接通过 gin 中，作用在 Context 结构体上的 PostForm() 方法来获取表单中的值
		// 下面这个就是输出，名为 username 表单的值
		fmt.Println(c.PostForm("username"))
	})
	router.Run(":8080")
}
