package main

import (
	"github.com/gin-gonic/gin"
	"gyc.com/Fund/cmd"
	"strconv"
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

	//获得某fund数据(在已经下载的情况下)
	r.GET("/fund/data/:code", func(c *gin.Context) {
		code := c.Param("code")
		if !cmd.IsCode(code) {
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

	//获取最新的某fund数据
	r.GET("/fund/update/:code", func(c *gin.Context) {
		code := c.Param("code")
		if !cmd.IsCode(code) {
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

	//获取的某fund涨跌数据（间隔1天）
	r.GET("/fund/rate/:code", func(c *gin.Context) {
		code := c.Param("code")
		if !cmd.IsCode(code) {
			c.JSON(405, gin.H{
				"code": 405,
				"msg":  "输入的code有误",
			})
			return
		}
		cmd.DownLoadFundTxt(code)
		cmd.GenFundJson(code)
		_, data := cmd.JsonToRateArray(code)
		c.JSON(200, gin.H{
			"code": 200,
			//"date": date,
			"data": data,
		})
	})

	//获取的某fund涨跌数据（间隔n天）
	r.GET("/fund/rate/:code/:day", func(c *gin.Context) {
		code := c.Param("code")
		if !cmd.IsCode(code) {
			c.JSON(405, gin.H{
				"code": 405,
				"msg":  "输入的code有误",
			})
			return
		}
		day, _ := strconv.ParseInt(c.Param("day"), 10, 32)

		cmd.DownLoadFundTxt(code)
		cmd.GenFundJson(code)
		_, data := cmd.JsonToRateArrayNDay(code, int(day))
		c.JSON(200, gin.H{
			"code": 200,
			//"date": date,
			"data": data,
		})
	})

	r.Run(":8888")

	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
