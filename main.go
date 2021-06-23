package main

import (
	"github.com/gin-gonic/gin"
	"gyc.com/Fund/cmd"
	"regexp"
)

func main() {
	//cmd.DownLoadFundTxt("008087")
	//cmd.GenFundJson("006113")
	//_, data := cmd.JsonToDataArray("006113")
	//fmt.Println(data)

	r := gin.Default()
	r.GET("/fund/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/fund/data/:code", func(c *gin.Context) {
		code := c.Param("code")
		isCode, _ := regexp.MatchString("\\b\\d{6}\\b", code)
		if !isCode {
			c.JSON(405, gin.H{
				"code": 405,
				"msg":  "输入的code有误",
			})
			return
		}
		_, data := cmd.JsonToDataArray(code)
		c.JSON(200, gin.H{
			"code": 200,
			"data": data,
		})
	})

	r.GET("/fund/update/:code", func(c *gin.Context) {
		code := c.Param("code")
		isCode, _ := regexp.MatchString("\\b\\d{6}\\b", code)
		if !isCode {
			c.JSON(405, gin.H{
				"code": 405,
				"msg":  "输入的code有误",
			})
			return
		}
		cmd.DownLoadFundTxt(code)
		cmd.GenFundJson(code)
		_, data := cmd.JsonToDataArray(code)
		c.JSON(200, gin.H{
			"code": 200,
			"data": data,
		})
	})

	r.Run(":8888")

	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
