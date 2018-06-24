package parsing

import (
	"fmt"
	"reflect"
	"strconv"

	gin "github.com/gin-gonic/gin"
)

func parseInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

// PaginationParam for parse page & limit in URI for offset parameter to limit get rows data in sql
func PaginationParam(c *gin.Context) (page int, limit int) {
	page, limit = parseInt(c.DefaultQuery("page", "1")), parseInt(c.DefaultQuery("limit", "10"))
	// limitSetting := utils.ParseInt(os.Getenv("PAGINATION_DATA_LIMIT"))
	// if limit > limitSetting {
	// 	limit = limitSetting
	// }
	if page < 1 {
		page = 1
	}
	page--
	page *= limit
	return
}

// RequestData for parse request body in HTTP POST, PUT, or DELETE in Content Type application/JSON,
// and binding it to data structure model
func RequestData(c *gin.Context, data interface{}) error {
	ref := reflect.ValueOf(data)
	if ref.Kind() != reflect.Ptr {
		return fmt.Errorf("Error. Must pass pointer")
	}
	if err := c.BindJSON(data); err != nil {
		return fmt.Errorf("Error parsing JSON, wrong data type. %v", err)
	}
	c.Set("bodydata", data)

	ref = reflect.ValueOf(data).Elem()
	typeOfData := ref.Type()
	for i := 0; i < ref.NumField(); i++ {
		key := typeOfData.Field(i).Tag.Get("form")
		val := ref.Field(i).Interface()
		isEmptyValue := reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
		if key == "required" && isEmptyValue {
			property := typeOfData.Field(i).Tag.Get("json")
			return fmt.Errorf("Error. Property '%s' required (must not empty)", property)
		}
	}
	dataType := reflect.TypeOf(data).Elem()
	if reflect.DeepEqual(data, reflect.New(dataType).Interface()) {
		return fmt.Errorf("No data in request body or JSON key undefined. Please read API documentation")
	}
	return nil
}
