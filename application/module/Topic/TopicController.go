package comment

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"

	parsing "../../../utils/parsing"
	response "../../../utils/response"
	model "../../model"
)

// METHOD: GET /topic/all
func getAllTopic(c *gin.Context) {
	topics := new(model.Topic).FindAll()
	response.Data(c, topics)
}

// METHOD: GET /topic/list
func getListTopic(c *gin.Context) {
	page, limit := parsing.PaginationParam(c)
	count, topics := new(model.Topic).FindDispart(page, limit)
	for i, topic := range topics {
		topics[i].News = new(model.News).FindByTopic(topic.ID)
		// reset topic id in array of news
		for j := range topics[i].News {
			topics[i].News[j].Topics = nil
		}
	}
	response.DataPagination(c, count, topics)
}

// METHOD: GET /topic/find/:id
func getTopicByStringID(c *gin.Context) {
	stringID := c.Param("id")
	var topic model.Topic
	topic.FindByStringID(stringID)
	news := new(model.News).FindByTopic(topic.ID)
	// Add additional info detail of topic from news object
	for i, tmp := range news {
		var topicID []bson.ObjectId
		for _, topic := range tmp.Topics {
			topicID = append(topicID, topic.ID)
		}
		news[i].Topics = new(model.Topic).GroupByID(topicID...)
	}
	topic.News = news
	response.Data(c, topic)
}

// METHOD: POST /topic/add
func addNewTopic(c *gin.Context) {
	var payload, topic model.Topic
	if err := parsing.RequestData(c, &payload); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := topic.InsertNewData(payload); err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, fmt.Sprintf("Success add new topic. Last ID: %s", topic.ID.Hex()))
}

// METHOD: PUT /topic/update/:id
func updateTopic(c *gin.Context) {
	var payload model.Topic
	if err := parsing.RequestData(c, &payload); err != nil {
		response.BadRequest(c, err)
		return
	}
	id := c.Param("id")
	if err := new(model.Topic).UpdateData(id, payload); err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, "Success update topic")
}

// METHOD: DELETE /topic/remove/:id
func removeTopic(c *gin.Context) {
	id := c.Param("id")
	if err := new(model.Topic).Remove(id); err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "Success remove topic")
}
