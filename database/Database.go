package database

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	mongoSession *mgo.Session
	dbname       = os.Getenv("DATABASE")
)

func Init() error {
	var err error
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	mongoHost := fmt.Sprintf("%s:%s", host, port)
	mongoSession, err = mgo.Dial(mongoHost)
	if err != nil {
		return err
	}

	// Init database collection, set unique index
	go func() {
		topics := mongoSession.DB(dbname).C("topics")
		index := mgo.Index{
			Key:    []string{"string_id"},
			Unique: true,
		}
		if err := topics.EnsureIndex(index); err != nil {
			return
		}

		news := mongoSession.DB(dbname).C("news")
		index = mgo.Index{
			Key:    []string{"_id", "topics._id"},
			Unique: true,
		}
		if err := news.EnsureIndex(index); err != nil {
			return
		}
	}()

	return nil
}

func LoadDB() *mgo.Database {
	return mongoSession.DB(dbname)
}
