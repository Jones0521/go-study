package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Param struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	r := gin.Default()
	r.POST("/test", testHandler)
	err := r.Run()
	if err != nil {
		return
	}
}
func testHandler(c *gin.Context) {
	var p Param
	if err := c.ShouldBindUri(&p); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusOK, "参数错误:"+err.Error())
		return
	}
	fmt.Printf("p: %#v/n", p)
	c.JSON(http.StatusOK, p)

}
