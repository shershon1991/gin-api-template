package demo

import (
	"52lu/go-import-template/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Struct(ctx *gin.Context) {
	fmt.Println("study struct.")

	response.Ok(ctx)
}
