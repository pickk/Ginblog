package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			// 允许所有来源的请求
			// AllowAllOrigins:  true, // 这个配置项允许所有源的请求，但需要注释掉上面的 AllowOrigins 配置
			AllowOrigins: []string{"*"}, // 等同于允许所有域名 #AllowAllOrigins:  true
			// 允许的 HTTP 请求方法
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			// 允许的请求头
			AllowHeaders: []string{"*", "Authorization"}, // 允许所有请求头以及 Authorization 请求头
			// 允许暴露的响应头
			ExposeHeaders: []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			// 是否允许携带凭证（如 Cookies）
			AllowCredentials: true,
			// 预检请求的缓存时间
			MaxAge: 12 * time.Hour, // 预检请求的结果在 12 小时内有效
		},
	)
}
