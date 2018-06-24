package model

import (
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

type MainModel struct {
	CreatedAt  string `bson:"created_at" db:"created_at" json:"created_at,omitempty"`
	ModifiedAt string `bson:"modified_at" db:"modified_at" json:"modified_at,omitempty"`
	DeletedAt  string `bson:"deleted_at" db:"deleted_at" json:"deleted_at,omitempty"`
}

// for generate dynamic values for partial update
func generatePartialValue(data interface{}) bson.M {
	var result = make(bson.M)
	ref := reflect.ValueOf(data)
	if ref.Kind() == reflect.Ptr {
		ref = reflect.ValueOf(data).Elem()
	}
	typeOfData := ref.Type()
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		key := typeOfData.Field(i).Tag.Get("bson")
		if key == "" || key == "-" {
			continue
		}
		val := field.Interface()
		isEmptyValue := reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
		if !isEmptyValue {
			result[key] = val
		}
	}

	return result
}
