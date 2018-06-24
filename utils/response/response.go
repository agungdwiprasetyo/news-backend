package response

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	gin "github.com/gin-gonic/gin"
)

type (
	dataResponse struct {
		Success  bool        `json:"success"`
		Message  string      `json:"message,omitempty"`
		Instance interface{} `json:"instance,omitempty"`
		Total    int         `json:"total,omitempty"`
		Data     interface{} `json:"data,omitempty"`
	}

	paginationResponse struct {
		Success     bool        `json:"success"`
		Message     string      `json:"message,omitempty"`
		Instance    interface{} `json:"instance,omitempty"`
		Page        int         `json:"page"`
		Limit       int         `json:"limit"`
		DataPerPage int         `json:"data_per_page"`
		TotalPage   int         `json:"total_page"`
		TotalData   int         `json:"total_all_data"`
		Data        interface{} `json:"data,omitempty"`
	}

	resp struct {
		Success bool        `json:"success"`
		Status  string      `json:"status,omitempty"`
		Message string      `json:"message,omitempty"`
		LastID  interface{} `json:"last_id,omitempty"`
	}
)

const (
	errMess = "Error when fetch json. Undefined key. Please read API Documentation."
)

func parseInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

// Data is method for return JSON array data or JSON object data in http body, and http status code 200
func Data(c *gin.Context, param ...interface{}) {
	if param == nil || len(param) > 2 {
		Failed(c, fmt.Errorf("Error in parameter"))
		return
	}
	httpCode := http.StatusOK
	var message string
	var data interface{}
	for i, val := range param {
		if mess, ok := val.(string); ok && i == 0 {
			message = mess
		} else {
			data = val
		}
	}
	ref := reflect.ValueOf(data)
	if ref.Kind() == reflect.Invalid {
		Failed(c, fmt.Errorf("Data must object or array"))
		return
	}

	res := dataResponse{
		Success: true,
		Message: message,
	}
	if ref.Kind() == reflect.Slice {
		res.Total = ref.Len()
	}
	res.Data = data
	c.Writer.Header().Set("Copyright", "Create with Golang 1.8 by Agung DP")
	c.JSON(httpCode, res)
}

// DataPagination is JSON response structure for data in dispart on pagination
func DataPagination(c *gin.Context, totalData int, data interface{}) {
	httpCode := http.StatusOK
	var instance interface{}
	if claims, exists := c.Get("claims"); exists {
		instance = fmt.Sprint(claims.(map[string]interface{})["instance"])
	}
	page, limit := parseInt(c.DefaultQuery("page", "1")), parseInt(c.DefaultQuery("limit", "10"))
	c.Writer.Header().Set("Copyright", "Create with Golang 1.8 by Agung DP")
	c.JSON(httpCode, paginationResponse{
		Success:     true,
		Message:     "Success",
		Instance:    instance,
		Page:        page,
		Limit:       limit,
		TotalPage:   int(math.Ceil(float64(totalData) / float64(limit))),
		DataPerPage: reflect.ValueOf(data).Len(),
		TotalData:   totalData,
		Data:        data,
	})
}

func Success(c *gin.Context, message string) {
	httpCode := http.StatusOK
	res := resp{
		Success: true,
		Status:  "OK",
		Message: message,
	}
	c.JSON(httpCode, res)
}

func Updated(c *gin.Context, message string) {
	httpCode := http.StatusOK
	res := resp{
		Success: true,
		Status:  "OK",
		Message: message,
	}
	c.JSON(httpCode, res)
}

func Deleted(c *gin.Context, message string) {
	httpCode := http.StatusOK
	res := resp{
		Success: true,
		Status:  "OK",
		Message: message,
	}
	c.JSON(httpCode, res)
}

// Failed is method for http response 200 if something error
func Failed(c *gin.Context, err error) {
	httpCode := http.StatusOK
	res := resp{
		Success: false,
		Status:  "Failed",
		Message: fmt.Sprint(err),
	}
	c.AbortWithStatusJSON(httpCode, res)
}

// BadRequest return http status code 400 in response header
func BadRequest(c *gin.Context, err error) {
	httpCode := http.StatusBadRequest
	var res resp
	if err != nil {
		res.Message = err.Error()
	} else {
		res.Message = errMess
	}
	res.Status = "400 Bad Request"
	c.AbortWithStatusJSON(httpCode, res)
}

// Unauthorized return http status code 401 in response header
func Unauthorized(c *gin.Context, err error) {
	httpCode := http.StatusUnauthorized
	var res resp
	res.Status = "401 Unauthorized"
	res.Message = err.Error()
	c.AbortWithStatusJSON(httpCode, res)
}

// Forbidden return http status code 403 in response header
func Forbidden(c *gin.Context, err error) {
	httpCode := http.StatusForbidden
	var res resp
	res.Status = "403 Forbidden"
	res.Message = err.Error()
	c.AbortWithStatusJSON(httpCode, res)
}

// AdminForbidden return http status code 403 in response header
func AdminForbidden(c *gin.Context) {
	httpCode := http.StatusForbidden
	err := fmt.Errorf("Only SUPERADMIN can access it")
	var res resp
	res.Status = "403 Forbidden"
	res.Message = err.Error()
	c.AbortWithStatusJSON(httpCode, res)
}

// NotFound return http status code 404 in response header
func NotFound(c *gin.Context, err error) {
	httpCode := http.StatusNotFound
	var res resp
	res.Status = "404 Not Found"
	res.Message = err.Error()
	c.AbortWithStatusJSON(httpCode, res)
}

// Conflict return http status code 409 in response header
func Conflict(c *gin.Context, err error) {
	httpCode := http.StatusConflict
	var res resp
	res.Status = "409 Conflict"
	res.Message = err.Error()
	c.AbortWithStatusJSON(httpCode, res)
}

// InternalServerError return http status code 500 in response header
func InternalServerError(c *gin.Context, err error) {
	httpCode := http.StatusInternalServerError
	var res resp
	res.Status = "500 Internal Server Error"
	res.Message = err.Error()
	c.AbortWithStatusJSON(httpCode, res)
}

// SQLerror return custom http status code in response header suitable with sql error status
func SQLerror(c *gin.Context, err error) {
	mess := strings.Split(err.Error(), ":")
	switch mess[0] {
	case "Error 1052":
		// ambigous column
		InternalServerError(c, err)
	// case "Error 1054":
	// Unknown column
	// 	InternalServerError(c, err)
	case "Error 1062":
		// Duplicate entry
		Conflict(c, err)
		// case "Error 1064":
		// 	InternalServerError(c, err)
	case "Error 1364":
		// field doesn't have a default value
		BadRequest(c, err)
	default:
		InternalServerError(c, err)
	}
}
