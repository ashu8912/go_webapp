package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router gin.IRouter) {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
