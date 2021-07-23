/**
 * @Description JWT中间件
 **/
package middleware

import (
	"52lu/go-import-template/global"
	"52lu/go-import-template/model/dao"
	"52lu/go-import-template/model/request"
	"52lu/go-import-template/model/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/**
 * @description: JWT中间件
 * @return func(ctx *gin.Context)
 */
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 获取参数中的token
		token := getToken(ctx)
		global.GvaLogger.Sugar().Infof("token: %s", token)
		if token == "" {
			response.Error(ctx, "Token不能为空!")
			// 中断请求
			ctx.Abort()
			return
		}
		// 验证Token
		userClaim, err := ParseToken(token)
		if err != nil {
			response.ErrorWithToken(ctx, "Token error :"+err.Error())
			// 中断请求
			ctx.Abort()
			return
		}
		// 设置到上下文中
		setContextData(ctx, userClaim, token)
		// 继续请求后续流程
		ctx.Next()
	}
}

// 设置数据到上下文
func setContextData(ctx *gin.Context, userClaim *request.UserClaims, token string) {
	userDao := &dao.UserDao{
		Uid: userClaim.Uid,
	}
	user, err := userDao.FindUser()
	if err != nil {
		response.Error(ctx, "用户不存在!")
		// 中断请求
		ctx.Abort()
		return
	}
	user.Token = token
	ctx.Set("userClaim", userClaim)
	ctx.Set("user", user)
}

// 从请求中获取Token
func getToken(ctx *gin.Context) string {
	var token string
	// 从header中获取
	token = ctx.Request.Header.Get("TOKEN")
	if token != "" {
		return token
	}
	// 获取当前请求方法
	if ctx.Request.Method == http.MethodGet {
		// 从Get请求中获取Token
		token, ok := ctx.GetQuery("token")
		if ok {
			return token
		}
	}
	// 从POST中和获取
	if ctx.Request.Method == http.MethodPost {
		// 从Get请求中获取Token
		postParam := make(map[string]interface{})
		_ = ctx.ShouldBindJSON(&postParam)
		token, ok := postParam["token"]
		if ok {
			return token.(string)
		}
	}
	return ""
}

// 创建Jwt
func CreateToken(uid uint) (string, error) {
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &request.UserClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(global.GvaConfig.Jwt.Expire).Unix(), // 有效期
			Issuer:    global.GvaConfig.Jwt.Issuer,                        // 签发人
			IssuedAt:  time.Now().Unix(),                                  // 签发时间
		},
		Uid: uid,
	})
	return newWithClaims.SignedString([]byte(global.GvaConfig.Jwt.Secret))
}

// 验证JWT
func ParseToken(tokenString string) (*request.UserClaims, error) {
	var err error
	var token *jwt.Token
	token, err = jwt.ParseWithClaims(tokenString, &request.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GvaConfig.Jwt.Secret), nil
	})
	if err != nil {
		global.GvaLogger.Error("解析JWT失败", zap.String("error", err.Error()))
		return nil, err
	}
	// 断言
	userClaims, ok := token.Claims.(*request.UserClaims)
	// 验证
	if !ok || !token.Valid {
		return nil, errors.New("JWT验证失败")
	}
	return userClaims, nil
}
