package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)
//自定义验证器 https://godoc.org/gopkg.in/go-playground/validator.v10
type Booking struct {
	CheckIn time.Time `form:"check_in" binding:"required,gttoday" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

//check_in 大于今天
func gtToday(fl validator.FieldLevel) bool {
	today:= time.Now()
	if fl.Field().Interface().(time.Time).Unix() > today.Unix() {
		return true
	}

	return false
}

func main()  {
	r:= gin.Default()

	if v,ok:=binding.Validator.Engine().(*validator.Validate);ok {
		v.RegisterValidation("gttoday", gtToday)
	}
	r.GET("/valid", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b);err !=nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":err.Error(),
				"aaaaaaaaaa":"aaaaaaaaaa",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"booking":b,
		})
	})
	r.Run()
}



/*//更多验证规则查看 http://godoc.org/gopkg.in/go-playaround/validator.v8
type Person struct {
	Name string `form:"name" binding:"required"`
	Age int `form:"age" binding:"required,gt=10"`
	Address string `form:"address" binding:"required"`
}
func main()  {
	r:= gin.Default()
	r.GET("/valid", func(c *gin.Context) {
		var person Person
		if err:=c.ShouldBind(&person);err !=nil {
			c.String(http.StatusInternalServerError, "%v", err)
			c.Abort()
			return
		}
		c.String(http.StatusOK, "%v",person)
	})
	r.Run()
}*/
