package api

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/jepsonwu/inchat_go/customize/response"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	*response.Status
	Data interface{} `json:"data"`
	Time int64 `json:"time"`
}

const (
	RESPONSE_JSON = "json"
	RESPONSE_XML  = "xml"
)

var (
	valid        *validation.Validation = &validation.Validation{}
	responseType string                 = "json"
)

func (c *BaseController) valid() {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Failure(response.GetStatusByMap(err.Message))
		}
	}
}

func (c *BaseController) Success(data interface{}) {
	var responseResult *Response = &Response{}
	responseResult.Status = response.SUCCESS
	responseResult.Data = data
	responseResult.Time = time.Now().Unix()

	switch responseType {
	case RESPONSE_JSON:
		c.Data[RESPONSE_JSON] = responseResult
		c.ServeJSON()
	}
}

func (c *BaseController) Failure(status *response.Status) {
	var responseResult *Response = &Response{}
	responseResult.Status = status
	responseResult.Time = time.Now().Unix()
	responseResult.Data = struct {
	}{}

	switch responseType {
	case RESPONSE_JSON:
		c.Data[RESPONSE_JSON] = responseResult
		c.ServeJSON()
	}
}
