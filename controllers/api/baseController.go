package api

import (
	"github.com/astaxie/beego"
	"time"
	"fmt"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	code    int32
	message string
	data    interface{}
	time    int64
}

var responseResult = &Response{} //定义变量  var re *dd 错误

func (c *BaseController) success(data interface{}) {
	responseResult.code = 0
	responseResult.message = ""
	responseResult.data = data
	responseResult.time = time.Now().Unix()
	fmt.Println(responseResult)
	return
	c.Data["json"] = responseResult
	c.ServeJSON()
}

//func failure() {
//
//}
