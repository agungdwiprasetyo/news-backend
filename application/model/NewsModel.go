package model

import (
	"reflect"
	"time"

	database "../../database"
	debug "../../utils/debug"

	"gopkg.in/mgo.v2/bson"
)

type (
	News struct {
		ID           bson.ObjectId `bson:"_id" json:"id"`
		Title        string        `bson:"title" json:"title,omitempty"`
		Value        string        `bson:"value" json:"value,omitempty"`
		PostedAt     time.Time     `bson:"posted_at" json:"posted_at,omitempty"`
		Status       string        `bson:"status" json:"status,omitempty"`
		TotalComment int           `bson:"total_comment" json:"total_comment"`
		TotalLike    int           `bson:"total_like" json:"total_like"`
		Topics       []Topic       `bson:"topics" json:"topics,omitempty"`
	}
)

func (this *News) FindDispart(page, limit int) (count int, news []News) {
	db := database.LoadDB()
	query := db.C("news").Find(bson.M{"status": bson.M{"$ne": "dihapus"}}).Sort("-posted_at")
	count, _ = query.Count()
	if err := query.Skip(page).Limit(limit).All(&news); err != nil {
		debug.PrintRed("Error in FindDispart News", err)
	}
	return
}

func (this *News) FindByID(id string) {
	db := database.LoadDB()
	objectID := bson.ObjectIdHex(id)
	if err := db.C("news").FindId(objectID).One(this); err != nil {
		debug.PrintRed("Error in FindByID News", err)
	}
}

func (this *News) FindByTopic(topicID bson.ObjectId) (news []News) {
	db := database.LoadDB()
	query := bson.M{"topics": bson.M{
		"$elemMatch": bson.M{"_id": topicID},
	},
	}
	if err := db.C("news").Find(query).All(&news); err != nil {
		debug.PrintRed("Error in FindByTopic News", err)
	}
	return
}

func (this *News) FindByStatus(status string) (news []News) {
	db := database.LoadDB()
	query := bson.M{"status": status}
	if err := db.C("news").Find(query).All(&news); err != nil {
		debug.PrintRed("Error in FindByStatus News", err)
	}
	return
}

func (this *News) InsertNewData(object News) (bson.ObjectId, error) {
	db := database.LoadDB()
	object.ID = bson.NewObjectId()
	object.PostedAt = time.Now()
	if object.Status == "" {
		object.Status = "konsep"
	}
	for i, topic := range object.Topics {
		// If not found, then insert new topic to database
		if err := topic.FindByName(topic.Name); err != nil {
			topic.InsertNewData(Topic{Name: topic.Name})
		}
		object.Topics[i] = Topic{ID: topic.ID}
	}
	if err := db.C("news").Insert(object); err != nil {
		return object.ID, err
	}

	return object.ID, nil
}

func (this *News) UpdateData(id string, data News) error {
	db := database.LoadDB()
	collections := db.C("news")
	objectID := bson.ObjectIdHex(id)

	topics := new(Topic).FindByNews(id)
	for _, topic := range data.Topics {
		if err := topic.FindByName(topic.Name); err != nil {
			// Topic with name {{topic.Name}} not found, insert new data
			topic.InsertNewData(Topic{Name: topic.Name})
		}
		topics = append(topics, Topic{ID: topic.ID})
	}
	data.Topics = topics
	partialValues := generatePartialValue(data)
	if reflect.DeepEqual(partialValues, bson.M{}) {
		debug.PrintGreen("No data updated")
		return nil
	}

	err := collections.UpdateId(objectID, bson.M{"$set": partialValues})
	return err
}

func (this *News) RemoveTopic(topicID string) error {
	db := database.LoadDB()
	collections := db.C("news")
	query := bson.M{"$pull": bson.M{"topics": bson.M{"_id": bson.ObjectIdHex(topicID)}}}
	err := collections.UpdateId(this.ID, query)
	return err
}

func (this *News) Archive(id string) error {
	db := database.LoadDB()
	err := db.C("news").UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"status": "dihapus"}})
	return err
}

func (this *News) Remove(id string) error {
	var news = News{ID: bson.ObjectIdHex(id)}
	db := database.LoadDB()
	err := db.C("news").Remove(news)
	return err
}
