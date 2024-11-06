package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wejectchen/ginblog/utils"
	"github.com/wejectchen/ginblog/utils/errmsg"
	"net/http"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 定义错误
var (
	TokenExpired     = errors.New("token已过期,请重新登录。")
	TokenNotValidYet = errors.New("token无效,请重新登录。")
	TokenMalformed   = errors.New("token不正确,请重新登录。")
	TokenInvalid     = errors.New("这不是一个token,请重新登录。")
)

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	// 验证token
	if token.Valid {
		return nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return TokenMalformed
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return TokenExpired
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return TokenInvalid
	} else {
		return TokenNotValidYet
	}
}

// JwtToken jwt中间件
// todo 优化此类代码
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义状态码变量
		var code int

		// 从请求头获取 Authorization 字段
		tokenHeader := c.Request.Header.Get("Authorization")

		// 检查 Authorization 字段是否为空
		if tokenHeader == "" {
			// 设置错误状态码，表示缺少 token
			code = errmsg.ERROR_TOKEN_EXIST

			// 返回 JSON 响应，包含状态码和错误信息
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code), // 获取对应的错误信息
			})

			// 中止后续处理
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// 解析token
		err := j.ParserToken(checkToken[1])
		if err != nil {
			if errors.Is(err, TokenExpired) {
				c.JSON(http.StatusOK, gin.H{
					"status":  errmsg.ERROR,
					"message": "token授权已过期,请重新登录",
					"data":    nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		//c.Set("username",)
		c.Next()
	}
}
