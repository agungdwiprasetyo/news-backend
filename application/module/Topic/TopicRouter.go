package comment

import (
	gin "github.com/gin-gonic/gin"
)

// Router -> [host]/comment/[route]
func Router(router *gin.RouterGroup) {
	router.GET("/all", getAllTopic)
	router.GET("/list", getListTopic)
	router.GET("/find/:id", getTopicByStringID)
	router.POST("/add", addNewTopic)
	router.PUT("/update/:id", updateTopic)
	router.DELETE("/remove/:id", removeTopic)
}
