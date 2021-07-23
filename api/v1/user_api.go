/**
 * @Description 用户相关表
 **/
package v1

import (
	"52lu/go-import-template/global"
	"52lu/go-import-template/middleware"
	"52lu/go-import-template/model"
	"52lu/go-import-template/model/request"
	"52lu/go-import-template/model/response"
	"52lu/go-import-template/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
 * @description: TODO 用户注册
 * @param ctx
 */
func Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam request.RegisterParam
	_ = ctx.ShouldBindJSON(&registerParam)
	// todo 参数校验
	// 调用注册
	register, err := service.Register(registerParam)
	if err != nil {
		response.Error(ctx, "注册失败!")
		return
	}
	response.OkWithData(ctx, register)
}

/**
 * @description: TODO 用户账号密码登录
 * @param ctx
 */
func Login(ctx *gin.Context) {
	// 绑定参数
	var loginParam request.LoginParam
	_ = ctx.ShouldBindJSON(&loginParam)
	if loginParam.Password == "" || loginParam.Phone == "" {
		response.Error(ctx, "手机号和密码不能为空！")
		return
	}
	// 调用登录服务
	userRecord := model.User{Phone: loginParam.Phone, Password: loginParam.Password}
	if err := service.LoginPwd(&userRecord); err != nil {
		global.GvaLogger.Error("登录失败:", zap.Any("user", userRecord))
		response.Error(ctx, "登录失败,账号或者密码错误!")
		return
	}
	// 生成token
	token, err := middleware.CreateToken(userRecord.ID)
	if err != nil {
		global.GvaLogger.Sugar().Errorf("登录失败,Token生成异常:%s", err)
		response.Error(ctx, "登录失败,账号或者密码错误!")
		return
	}
	userRecord.Token = token
	response.OkWithData(ctx, userRecord)
}

// 查询用户信息
func GetUser(ctx *gin.Context) {
	// 从上下文中获取用户信息，(经过中间件逻辑时，已经设置到上下文)
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
	return
}
