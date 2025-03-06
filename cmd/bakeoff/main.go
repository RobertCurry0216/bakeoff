package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertcurry0216/bakeoff/internal/helpers"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/user/:email", func(c *gin.Context) {
		email := c.Params.ByName("email")
		attributes, err := helpers.GetUserAttributesFromEmail(email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"baseURL": helpers.BaseUrl(),
				"err":     err.Error(),
			})
		} else {
			switch c.ContentType() {
			case "application/json":
				c.JSON(http.StatusOK, attributes)
			default:
				c.HTML(http.StatusOK, "basic.tmpl", gin.H{
					"title": "User",
					"email": email,
					"data":  attributes,
				})
			}
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8081")
}
