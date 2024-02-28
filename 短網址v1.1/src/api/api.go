package api

import (
	"fmt"
	"shortner/create"
	"shortner/db/query"

	"github.com/gin-gonic/gin"
)

func AddLongURLToShortURL(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	url := ctx.PostForm("url")

	code := create.ToShortURL(url)
	ctx.JSON(200, gin.H{"code": code})
}

func ToAnotherWebsite(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	code := ctx.Param("code")
	blog := query.SelectCode(code)
	refreshHtml := fmt.Sprintf("<html><meta http-equiv='refresh' content='0;url=%s'/></html>", blog)

	ctx.Data(200, "text/html; charset=utf-8", []byte(refreshHtml))
}
