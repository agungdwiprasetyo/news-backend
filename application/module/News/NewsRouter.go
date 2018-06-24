package item

import (
	gin "github.com/gin-gonic/gin"
)

// Router -> [host]/news/[route]
func Router(router *gin.RouterGroup) {
	router.GET("/list", getAllNews)
	router.GET("/find/:id", getNewsByID)
	router.GET("/filter/topic/:id", filterNewsByTopic)
	router.GET("/filter/status/:status", filterNewsByStatus)
	router.POST("/add", addNews)
	router.PUT("/update/:id", updateNews)
	router.DELETE("/topic", deleteTopic)
	router.DELETE("/archive/:id", archiveNews)
	router.DELETE("/remove/:id", removeNews)
}
