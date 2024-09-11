// Package response provides wrapper gin returns json data in the same format.
package response

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/zhufuyi/sponge/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Result output data format
type Result struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	TimeStamp int64       `json:"time_stamp"`
	RequestId string      `json:"request_id"`
}

func newResp(code int, msg string, data interface{}) *Result {
	// 生成随机的 16 位数字
	rand.Seed(time.Now().UnixNano())
	randomPart := fmt.Sprintf("%016d", rand.Intn(10000000000000000))

	// 生成包含随机字符串和时间戳的 requestId
	requestId := fmt.Sprintf("%s%d", randomPart, time.Now().Unix())

	// 对 requestId 进行 MD5 加密
	requestIdHash := md5.Sum([]byte(requestId))
	requestId = fmt.Sprintf("%x", requestIdHash)
	//md5一下
	resp := &Result{
		Code:      code,
		Msg:       msg,
		TimeStamp: time.Now().Unix(),
		RequestId: requestId,
	}

	// ensure that the data field is not nil on return, note that it is not nil when resp.data=[]interface {}, it is serialized to null
	if data == nil {
		resp.Data = &struct{}{}
	} else {
		resp.Data = data
	}

	return resp
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func writeJSON(c *gin.Context, code int, res interface{}) {
	c.Writer.WriteHeader(code)
	writeContentType(c.Writer, jsonContentType)
	err := json.NewEncoder(c.Writer).Encode(res)
	if err != nil {
		fmt.Printf("json encode error, err = %s\n", err.Error())
	}
}

func respJSONWithStatusCode(c *gin.Context, code int, msg string, data ...interface{}) {
	var firstData interface{}
	if len(data) > 0 {
		firstData = data[0]
	}
	resp := newResp(code, msg, firstData)

	writeJSON(c, code, resp)
}

// Output return json data by http status code
func Output(c *gin.Context, code int, msg ...interface{}) {
	switch code {
	case http.StatusOK:
		respJSONWithStatusCode(c, http.StatusOK, "ok", msg...)
	case http.StatusBadRequest:
		respJSONWithStatusCode(c, http.StatusBadRequest, errcode.InvalidParams.Msg(), msg...)
	case http.StatusUnauthorized:
		respJSONWithStatusCode(c, http.StatusUnauthorized, errcode.Unauthorized.Msg(), msg...)
	case http.StatusForbidden:
		respJSONWithStatusCode(c, http.StatusForbidden, errcode.Forbidden.Msg(), msg...)
	case http.StatusNotFound:
		respJSONWithStatusCode(c, http.StatusNotFound, errcode.NotFound.Msg(), msg...)
	case http.StatusRequestTimeout:
		respJSONWithStatusCode(c, http.StatusRequestTimeout, errcode.Timeout.Msg(), msg...)
	case http.StatusConflict:
		respJSONWithStatusCode(c, http.StatusConflict, errcode.AlreadyExists.Msg(), msg...)
	case http.StatusInternalServerError:
		respJSONWithStatusCode(c, http.StatusInternalServerError, errcode.InternalServerError.Msg(), msg...)
	case http.StatusTooManyRequests:
		respJSONWithStatusCode(c, http.StatusTooManyRequests, errcode.LimitExceed.Msg(), msg...)
	case http.StatusServiceUnavailable:
		respJSONWithStatusCode(c, http.StatusServiceUnavailable, errcode.ServiceUnavailable.Msg(), msg...)

	default:
		respJSONWithStatusCode(c, code, http.StatusText(code), msg...)
	}
}

// Out return json data by http status code, converted by errcode
func Out(c *gin.Context, err *errcode.Error, data ...interface{}) {
	code := err.ToHTTPCode()
	switch code {
	case http.StatusOK:
		respJSONWithStatusCode(c, http.StatusOK, "ok", data...)
	case http.StatusBadRequest:
		respJSONWithStatusCode(c, http.StatusBadRequest, err.Msg(), data...)
	case http.StatusUnauthorized:
		respJSONWithStatusCode(c, http.StatusUnauthorized, err.Msg(), data...)
	case http.StatusForbidden:
		respJSONWithStatusCode(c, http.StatusForbidden, err.Msg(), data...)
	case http.StatusNotFound:
		respJSONWithStatusCode(c, http.StatusNotFound, err.Msg(), data...)
	case http.StatusRequestTimeout:
		respJSONWithStatusCode(c, http.StatusRequestTimeout, err.Msg(), data...)
	case http.StatusConflict:
		respJSONWithStatusCode(c, http.StatusConflict, err.Msg(), data...)
	case http.StatusInternalServerError:
		respJSONWithStatusCode(c, http.StatusInternalServerError, err.Msg(), data...)
	case http.StatusTooManyRequests:
		respJSONWithStatusCode(c, http.StatusTooManyRequests, err.Msg(), data...)
	case http.StatusServiceUnavailable:
		respJSONWithStatusCode(c, http.StatusServiceUnavailable, err.Msg(), data...)

	default:
		respJSONWithStatusCode(c, http.StatusNotExtended, err.Msg(), data...)
	}
}

// status code flat 200, custom error codes in data.code
func respJSONWith200(c *gin.Context, code int, msg string, data ...interface{}) {
	var firstData interface{}
	if len(data) > 0 {
		firstData = data[0]
	}
	resp := newResp(code, msg, firstData)

	writeJSON(c, http.StatusOK, resp)
}

// Success return success
func Success(c *gin.Context, data ...interface{}) {
	respJSONWith200(c, 0, "操作成功", data...)
}

// Error return error
func Error(c *gin.Context, err *errcode.Error, data ...interface{}) {
	respJSONWith200(c, err.Code(), err.Msg(), data...)
}
