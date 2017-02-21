package api

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/jepsonwu/inchat_go/customize/response"
	"github.com/astaxie/beego/validation"
)

const (
	//response type
	RESPONSE_JSON = "json"
	RESPONSE_XML = "xml"

	//paginate
	PAGE = 1
	PER_PAGE = 20
)

var (
	valid        *validation.Validation = &validation.Validation{}
	responseType string = "json"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	*response.Status
	Data interface{} `json:"data"`
	Time int64 `json:"time"`
}

/**-----------------request paginate-------------------------**/
type PaginatePerPage interface {
	GetPerPage() int
}

type Paginate interface {
	GetPage() int
	PaginatePerPage
}

type RequestPaginate struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}

func (r *RequestPaginate) GetPage() int {
	if r.Page == 0 {
		r.Page = PAGE
	}
	return r.Page
}

func (r *RequestPaginate) GetPerPage() int {
	if r.PerPage == 0 {
		r.PerPage = PER_PAGE
	}
	return r.PerPage
}

type PaginateLastId interface {
	GetLastId() int
	PaginatePerPage
}

type RequestPaginateLastId struct {
	LastId  int `form:"last_id"`
	PerPage int `form:"per_page"`
}

func (r *RequestPaginateLastId) GetLastId() int {
	return r.LastId
}

func (r *RequestPaginateLastId) GetPerPage() int {
	if r.PerPage == 0 {
		r.PerPage = PER_PAGE
	}
	return r.PerPage
}

/**---------------------validation--------------------------**/
func (c *BaseController) validRequestPaginate(requestParams Paginate) {
	valid.Min(requestParams.GetPage(), 1, "page").Message("page_invalid")
	valid.Min(requestParams.GetPerPage(), 1, "per_page").Message("per_page_invalid")
}

func (c *BaseController) validRequestPaginateLastId(requestParams PaginateLastId) {
	valid.Min(requestParams.GetLastId(), 0, "last_id").Message("last_id_invalid")
	valid.Min(requestParams.GetPerPage(), 1, "per_page").Message("per_page_invalid")
}

func (c *BaseController) valid() {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.failure(response.GetStatusByMap(err.Message))
		}
	}
}

/**------------------response------------------------**/
func (c *BaseController) success(data interface{}) {
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

func (c *BaseController) failure(status *response.Status) {
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
