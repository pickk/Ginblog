package routes

// 引入多模板渲染器包
import "github.com/gin-contrib/multitemplate"

// 引入Gin框架
import "github.com/gin-gonic/gin"

// 引入API版本1的包
import "github.com/wejectchen/ginblog/api/v1"

// 引入中间件包
import "github.com/wejectchen/ginblog/middleware"

// 引入工具函数包
import "github.com/wejectchen/ginblog/utils"

// createMyRender 创建一个多模板渲染器
func createMyRender() multitemplate.Renderer {
	// 初始化多模板渲染器
	p := multitemplate.NewRenderer()
	// 添加管理员模板文件
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	// 添加前端模板文件
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}
func InitRouter() {
	// 设置 Gin 运行模式
	gin.SetMode(utils.AppMode)
	r := gin.New() // 创建一个新的 Gin 实例

	// 设置信任网络 []string
	// nil 表示不计算信任代理，避免性能消耗。上线时应当配置为具体的可信代理 IP 列表
	_ = r.SetTrustedProxies(nil)

	// 设置 HTML 渲染器
	r.HTMLRender = createMyRender()

	// 使用自定义的日志中间件
	r.Use(middleware.Logger())

	// 使用 Gin 的恢复中间件，用于捕获恐慌 (panic) 并恢复，防止程序崩溃
	r.Use(gin.Recovery())

	// 使用自定义的跨域中间件，允许跨域请求
	r.Use(middleware.Cors())

	// 设置静态资源路径
	// 将请求路径 "/static" 映射到本地 "./web/front/dist/static" 目录，提供前端静态资源
	r.Static("/static", "./web/front/dist/static")

	// 将请求路径 "/admin" 映射到本地 "./web/admin/dist" 目录，提供后台管理的静态页面
	r.Static("/admin", "./web/admin/dist")

	// 设置 favicon.ico 文件路径
	// 将 "/favicon.ico" 映射到本地路径 "/web/front/dist/favicon.ico" 文件
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//修改密码
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)
		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArtInfo)
		auth.GET("admin/article", v1.GetArt)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)
		// 评论模块
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章分类信息模块
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)

		// 文章模块
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)

		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)
	}

	_ = r.Run(utils.HttpPort)

}
