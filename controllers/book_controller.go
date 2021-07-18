package controllers

import "github.com/gin-gonic/gin"

func ShowBook(ctx *gin.Context) {
	ctx.JSON(
		200,
			gin.H{
				"value": "ok",
			},
		)
}
