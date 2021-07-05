/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/7/5 3:42 下午
 **/
package v1

import (
	"52lu/go-import-template/model/response"
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	response.Ok(ctx)
}
func Test(ctx *gin.Context) {
	response.Ok(ctx)
}
