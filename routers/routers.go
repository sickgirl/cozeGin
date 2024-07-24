package routers

import (
	"coze_gin/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user_name":      "张三",
			"user_phone":     "13000001111",
			"user_hosipital": "天津师范医院",
			"user_team":      "李四主任团队",
			"context":        "餐后血糖2.8已超出警戒值请及时干预",
		})
	})

	r.GET("/get_user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user_name":      "张三",
			"user_phone":     "13000001111",
			"user_hosipital": "天津师范医院",
			"user_team":      "李四主任团队",
			"status":         "已入组",
			"context":        "规范用户",
		})
	})

	return r
}
