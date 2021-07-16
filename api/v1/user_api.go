/**
 * @Description 用户相关表
 **/
package v1

import (
	"52lu/go-import-template/global"
	"52lu/go-import-template/model/entity"
	"52lu/go-import-template/model/request"
	"52lu/go-import-template/model/response"
	"52lu/go-import-template/service"
	"fmt"
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
		response.Fail(ctx,"注册失败!")
		return
	}
	response.OkWithData(ctx,register)
}

/**
 * @description: TODO 用户账号密码登录
 * @param ctx
 */
func Login(ctx *gin.Context) {
	// 绑定参数
	var loginParam request.LoginParam
	_ = ctx.ShouldBindJSON(&loginParam)
	fmt.Println("参数:", loginParam)
	if loginParam.Password == "" || loginParam.Phone == "" {
		response.Fail(ctx, "手机号和密码不能为空！")
		return
	}
	// 调用登录服务
	userRecord := entity.User{Phone: loginParam.Phone, Password: loginParam.Password}
	if err := service.LoginPwd(&userRecord);err != nil {
		global.GvaLogger.Error("登录失败:",zap.Any("user",userRecord))
		response.Fail(ctx,"登录失败,账号或者密码错误!")
		return
	}
	response.OkWithData(ctx, userRecord)
}
