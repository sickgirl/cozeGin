package routers

import (
	"coze_gin/parameter"
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

	packages := parameter.PkgList{
		List: []parameter.ServicePackage{
			{
				PackageName:    "稳糖服务包-太原-年包",
				City:           "taiyuan",
				Sales:          1500,
				Price:          299.99,
				OriginalPrice:  399.99,
				DurationMonths: 12,
				Desc:           "太原:服务包描述1",
			},
			{
				PackageName:    "稳糖服务包-太原-季包",
				City:           "taiyuan",
				Sales:          1200,
				Price:          29.99,
				OriginalPrice:  39.99,
				DurationMonths: 3,
				Desc:           "太原:服务包描述2",
			},
			{
				PackageName:    "稳糖服务包-太原-半年包",
				City:           "taiyuan",
				Sales:          100,
				Price:          129.99,
				OriginalPrice:  139.99,
				DurationMonths: 6,
				Desc:           "太原:服务包描述3",
			},
			{
				PackageName:    "稳糖服务包-太原-月包",
				City:           "taiyuan",
				Sales:          100,
				Price:          19.99,
				OriginalPrice:  19.99,
				DurationMonths: 1,
				Desc:           "太原:服务包描述4",
			},
		},
	}

	r.GET("/get_pkg_list", func(c *gin.Context) {
		c.JSON(200, packages)
	})

	return r
}
