package application

import (
	news "./module/News"
	topic "./module/Topic"

	gin "github.com/gin-gonic/gin"
)

// InitAppModule untuk inisialisasi routers aplikasi
func InitAppModule(app *gin.Engine) {
	newsGroup := app.Group("/news")
	news.Router(newsGroup)

	topicGroup := app.Group("/topic")
	topic.Router(topicGroup)
}
