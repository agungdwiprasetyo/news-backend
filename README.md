# Backend-API for management News and Topic

Using Golang and MongoDB
Deploy to http://kumparan-api.agungdwiprasetyo.com.

### List API
* GET    /news/list                     **-->** ./application/module/News.getAllNews
* GET    /news/find/:id                 **-->** ./application/module/News.getNewsByID
* GET    /news/filter/topic/:id         **-->** ./application/module/News.filterNewsByTopic
* GET    /news/filter/status/:status    **-->** ./application/module/News.filterNewsByStatus
* POST   /news/add                      **-->** ./application/module/News.addNews
* PUT    /news/update/:id               **-->** ./application/module/News.updateNews
* DELETE /news/topic                    **-->** ./application/module/News.deleteTopic
* DELETE /news/archive/:id              **-->** ./application/module/News.archiveNews
* DELETE /news/remove/:id               **-->** ./application/module/News.removeNews
* GET    /topic/all                     **-->** ./application/module/Topic.getAllTopic
* GET    /topic/list                    **-->** ./application/module/Topic.getListTopic
* GET    /topic/find/:id                **-->** ./application/module/Topic.getTopicByStringID
* POST   /topic/add                     **-->** ./application/module/Topic.addNewTopic
* PUT    /topic/update/:id              **-->** ./application/module/Topic.updateTopic
* DELETE /topic/remove/:id              **-->** ./application/module/Topic.removeTopic

### Data Structure
#### NEWS
```go
News struct {
	ID           bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Title        string        `bson:"title" json:"title,omitempty"`
	Value        string        `bson:"value" json:"value,omitempty"`
	PostedAt     time.Time     `bson:"posted_at" json:"posted_at,omitempty"`
	Status       string        `bson:"status" json:"status,omitempty"`
	TotalComment int           `bson:"total_comment" json:"total_comment"`
	TotalLike    int           `bson:"total_like" json:"total_like"`
	Topics       []Topic       `bson:"topics" json:"topics,omitempty"`
}
```

#### TOPIC
```go
Topic struct {
	ID       bson.ObjectId `bson:"_id" json:"id,omitempty"`
	StringID string        `bson:"string_id" json:"string_id,omitempty"`
	Name     string        `bson:"name" json:"name,omitempty"`
	News     []News        `bson:"-" json:"news,omitempty"`
}
```