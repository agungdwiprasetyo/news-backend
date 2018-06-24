package item

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	parsing "../../../utils/parsing"
	response "../../../utils/response"
	model "../../model"

	gin "github.com/gin-gonic/gin"
)

// METHOD: GET /news/list?page=&limit=
func getAllNews(c *gin.Context) {
	page, limit := parsing.PaginationParam(c)
	count, news := new(model.News).FindAll(page, limit)
	for i, tmp := range news {
		var topicID []bson.ObjectId
		for _, topic := range tmp.Topics {
			topicID = append(topicID, topic.ID)
		}
		news[i].Topics = new(model.Topic).GroupByID(topicID...)
	}
	response.DataPagination(c, count, news)
}

// METHOD: GET /news/find/:id
func getNewsByID(c *gin.Context) {
	id := c.Param("id")
	var news model.News
	news.FindByID(id)
	var topicID []bson.ObjectId
	for _, topic := range news.Topics {
		topicID = append(topicID, topic.ID)
	}
	news.Topics = new(model.Topic).GroupByID(topicID...)
	response.Data(c, news)
}

// METHOD: GET /news/filter/topic/:id
func filterNewsByTopic(c *gin.Context) {
	topicID := c.Param("id")
	var topic model.Topic
	topic.FindByStringID(topicID)
	news := new(model.News).FindByTopic(topic.ID)
	// Add additional info detail of topic from news object
	for i, tmp := range news {
		var topicID []bson.ObjectId
		for _, topic := range tmp.Topics {
			topicID = append(topicID, topic.ID)
		}
		news[i].Topics = new(model.Topic).GroupByID(topicID...)
	}
	response.Data(c, news)
}

// METHOD: GET /news/filter/status/:status
func filterNewsByStatus(c *gin.Context) {
	status := c.Param("status")
	news := new(model.News).FindByStatus(status)
	// Add additional info detail of topic from news object
	for i, tmp := range news {
		var topicID []bson.ObjectId
		for _, topic := range tmp.Topics {
			topicID = append(topicID, topic.ID)
		}
		news[i].Topics = new(model.Topic).GroupByID(topicID...)
	}
	response.Data(c, news)
}

// METHOD: POST /news/add
func addNews(c *gin.Context) {
	var payload model.News
	if err := parsing.RequestData(c, &payload); err != nil {
		response.Failed(c, err)
		return
	}
	newsID, err := new(model.News).InsertNewData(payload)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, fmt.Sprintf("Success add new news. Last ID: %s", newsID.Hex()))
}

// METHOD: PUT /news/update/:id
func updateNews(c *gin.Context) {
	var payload model.News
	if err := parsing.RequestData(c, &payload); err != nil {
		response.BadRequest(c, err)
		return
	}
	id := c.Param("id")
	if err := new(model.News).UpdateData(id, payload); err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, "Success update news")
}

// METHOD: DELETE /news/topic?news_id=&topic_id=
func deleteTopic(c *gin.Context) {
	var news model.News
	newsID, topicID := c.Query("news_id"), c.Query("topic_id")
	news.ID = bson.ObjectIdHex(newsID)
	if err := news.RemoveTopic(topicID); err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "Success remove topic in news")
}

// METHOD: DELETE /news/archive/:id
// for soft delete
func archiveNews(c *gin.Context) {
	id := c.Param("id")
	if err := new(model.News).Archive(id); err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "Success archive news")
}

// METHOD: DELETE /news/remove/:id
// for hard delete
func removeNews(c *gin.Context) {
	id := c.Param("id")
	if err := new(model.News).Remove(id); err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "Success remove news")
}
