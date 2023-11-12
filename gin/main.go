package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Item struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/item", Get)
	r.POST("/item", Post)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Get(c *gin.Context) {
	var item Item

	err := c.BindQuery(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	item.Name = "hello"
	c.JSON(http.StatusOK, item)
}

func Post(c *gin.Context) {
	item := new(Item)
	err := c.Bind(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, item)
}
