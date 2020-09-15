package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main()  {
	r:= gin.Default()

	//获取post参数 body
	r.POST("/submit", func(c *gin.Context) {
		body,err := ioutil.ReadAll(c.Request.Body) //ReadAll之后，c.Request.Body就不存在了，PostForm取不到值

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		//firstName := c.PostForm("first_name") //
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}

		c.String(http.StatusOK, string(body))
	})

	//获取get参数
	r.GET("/user_info", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "l_name")
		//c.String(200, "%s %s", firstName, lastName)
		c.JSON(http.StatusOK, gin.H{
			"first_name":firstName,
			"last_name":lastName,
		})
	})
	//泛绑定
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(http.StatusOK, "this is user detail")
	})

	r.GET("/param/:name/:id", func(c *gin.Context){
		c.JSON(200, gin.H{
			"name":c.Param("name"),
			"id":c.Param("id"),
		})
	})

	r.Run()
}
