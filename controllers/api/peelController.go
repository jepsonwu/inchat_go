package api

import (
	"github.com/jepsonwu/inchat_go/models"
	"os"
)

type PeelController struct {
	BaseController
}

type GetValidation struct {
	LastId  int
	PerPage int
}

func (c *PeelController) Get() {
	getParams := &GetValidation{}
	c.ParseForm(getParams)

	println(getParams.LastId)
	os.Exit(0)
	//getParams.LastId = 1
	getParams.PerPage = 1

	valid.Required(getParams.LastId, "last_id").Message("system_error")
	c.valid()
	//last_id per_page
	//c.Input()
	//c.Ctx.WriteString(beego.AppConfig.String("mysql_host"))

	//response
	peelInfo := models.NewPeel()
	c.Success(peelInfo.GetByPrimary())
}
