package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AsciiJSON() {
	r := gin.Default()

	r.GET("/asciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Gin框架&33&&",
			"tag":  "<br>",
		}

		// 输出: {"lang":"Gin\u6846\u67b6","tag":"\u003cbr\u003e"}
		// AsciiJSON 可以将不是ascii字符的数据转化为u你
		c.JSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
