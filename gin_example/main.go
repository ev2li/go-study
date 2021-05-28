package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Passwd string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

//gin的HelloWorld
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	/*r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})

	r.POST("/xxxPost", getting)
	r.PUT("/xxxPut")*/
/*	r.GET("/user/:name/*action", func(ctx *gin.Context){
		name := ctx.Param("name")
		action := ctx.Param("action")
		ctx.String(http.StatusOK, name + action)
	})

	r.GET("/welcome", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name","Jack")
		ctx.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})


	r.POST("/form", func(context *gin.Context) {
		type1 := context.DefaultPostForm("type", "alter")
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.String(http.StatusOK, fmt.Sprintf("type:%v, username:%v, password:%v"), type1, username, password)
	})*/

	/*v1 := r.Group("/v1") //{} 书写规范
	{
		v1.GET("/login", func(context *gin.Context) {

		})

		v1.GET("/submit", func(context *gin.Context) {

		})
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(context *gin.Context) {

		})

		v2.POST("/submit", func(context *gin.Context) {

		})
	}*/

	//json绑定
	r.POST("loginJSON", func(ctx *gin.Context) {
		login := &Login{}
		if err := ctx.ShouldBindJSON(login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		if login.User != "root" || login.Passwd != "admin" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status":"304"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status":200})
	})


	r.GET("/index", func(context *gin.Context) {
		context.HTML(200,"index.html", gin.H{"title":"我的标题"})
	})

	r.Run(":7071") //默认在8080
}

/*func getting(ctx *gin.Context)  {

}*/
