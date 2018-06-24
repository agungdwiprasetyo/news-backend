package model

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"

	database "../../database"
	debug "../../utils/debug"
)

type (
	Topic struct {
		ID       bson.ObjectId `bson:"_id" json:"id"`
		StringID string        `bson:"string_id" json:"string_id,omitempty"`
		Name     string        `bson:"name" json:"name,omitempty"`
		News     []News        `bson:"-" json:"news,omitempty"`
	}
)

func (this *Topic) FindAll() (topics []Topic) {
	db := database.LoadDB()
	if err := db.C("topics").Find(bson.M{}).All(&topics); err != nil {
		debug.PrintRed("Error in FindAll Topic", err)
	}
	return
}

func (this *Topic) FindDispart(page, limit int) (count int, topics []Topic) {
	db := database.LoadDB()
	query := db.C("topics").Find(bson.M{})
	count, _ = query.Count()
	if err := query.Skip(page).Limit(limit).All(&topics); err != nil {
		debug.PrintRed("Error in FindDispart Topic", err)
	}
	return
}

func (this *Topic) FindByID(id string) {
	db := database.LoadDB()
	objectID := bson.ObjectIdHex(id)
	if err := db.C("topics").FindId(objectID).One(this); err != nil {
		debug.PrintRed("Error in FindByID Topic", err)
	}
}

func (this *Topic) FindByStringID(stringID string) {
	db := database.LoadDB()
	query := bson.M{"string_id": stringID}
	if err := db.C("topics").Find(query).One(this); err != nil {
		debug.PrintRed("Error in FindByStringID Topic", err)
	}
}

func (this *Topic) GroupByID(topicID ...bson.ObjectId) (topics []Topic) {
	db := database.LoadDB()
	query := bson.M{"_id": bson.M{"$in": topicID}}
	if err := db.C("topics").Find(query).All(&topics); err != nil {
		debug.PrintRed("Error Topic GroupByID.", err)
	}
	return
}

func (this *Topic) FindByName(name string) error {
	db := database.LoadDB()
	replacer := strings.NewReplacer("/", "-", " ", "-", ",", "-", ".", "-")
	stringID := strings.ToLower(replacer.Replace(name))
	if err := db.C("topics").Find(bson.M{"string_id": stringID}).One(this); err != nil {
		debug.PrintRed(fmt.Sprintf("Name '%s' not found.", name), err)
		return err
	}
	return nil
}

func (this *Topic) FindByNews(newsID string) (topics []Topic) {
	db := database.LoadDB()
	objectID := bson.ObjectIdHex(newsID)
	var news News
	if err := db.C("news").FindId(objectID).One(&news); err != nil {
		debug.PrintRed("Error in FindByNews Topic", err)
	}
	topics = news.Topics
	return
}

func (this *Topic) InsertNewData(object Topic) error {
	if object.Name == "" {
		return fmt.Errorf("Empty name")
	}
	db := database.LoadDB()
	replacer := strings.NewReplacer("/", "-", " ", "-", ",", "-", ".", "-")
	object.ID = bson.NewObjectId()
	object.StringID = strings.ToLower(replacer.Replace(object.Name))
	this.ID = object.ID
	this.StringID = object.StringID
	err := db.C("topics").Insert(object)
	return err
}

func (this *Topic) UpdateData(id string, object Topic) error {
	partialValues := generatePartialValue(object)
	fmt.Println(debug.PrettyJSON(partialValues))
	db := database.LoadDB()
	err := db.C("topics").UpdateId(bson.ObjectIdHex(id), bson.M{"$set": partialValues})
	return err
}

func (this *Topic) Archive(id string) error {
	return nil
}

func (this *Topic) Remove(id string) error {
	var topic = Topic{ID: bson.ObjectIdHex(id)}
	db := database.LoadDB()
	err := db.C("topics").Remove(topic)
	return err
}
