package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

type Person struct {
	Age int `form:"age" validate:"required,gt=18"`
	Name string `form:"name" validate:"required"`
}

var (
	Uni *ut.UniversalTranslator
	Validate *validator.Validate
)
func main()  {
	Validate := validator.New()
	zh := zh2.New()
	en := en2.New()
	Uni := ut.New(zh,en)

	r:=gin.Default()
	r.GET("/mult_language", func(c *gin.Context) {
		local := c.DefaultQuery("local", "zh")
		trans,_:= Uni.GetTranslator(local)
		switch local {
		case "zh":
			zh_translation.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translation.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translation.RegisterDefaultTranslations(Validate, trans)
		}
		var p Person
		if err:=c.ShouldBind(&p); err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"err":err.Error(),
			})
			c.Abort()
			return
		}
		if err:=Validate.Struct(p);err != nil{
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _,e:= range errs{
				sliceErrs = append(sliceErrs,e.Translate(trans))
			}
			c.JSON(500, gin.H{
				"errs":sliceErrs,
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"person":p,
		})
	})
	r.Run()
}
