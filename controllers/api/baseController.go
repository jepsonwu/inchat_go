package api

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/jepsonwu/inchat_go/customize/response"
	"github.com/astaxie/beego/validation"
	"fmt"
	"os"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	*response.Status
	Data interface{} `json:"data"`
	Time int64 `json:"time"`
}

type Paginate struct {
	Page    uint64
	PerPage uint64
}

type PaginateLastId struct {
	LastId  string `valid:"Required"`
	PerPage uint64 `valid:"Numeric"`
}

const (
	RESPONSE_JSON = "json"
	RESPONSE_XML  = "xml"
)

var (
	valid *validation.Validation = &validation.Validation{}

	responseResult *Response = &Response{}
	responseType   string    = "json"
)

func (c *BaseController) Valid() {
	getParams := &GetValidation{}
	//getParams.LastId = ""
	getParams.PerPage = 1

	r, err := valid.Valid(getParams)
	fmt.Println(r)
	os.Exit(0)
	if err != nil {
		c.Failure(response.SYSTEM_ERROR)
	}

	if !r {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
		}
	}
}

func (c *BaseController) Success(data interface{}) {
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
	responseResult.Status = status
	responseResult.Time = time.Now().Unix()

	switch responseType {
	case RESPONSE_JSON:
		c.Data[RESPONSE_JSON] = responseResult
		c.ServeJSON()
	}
}
