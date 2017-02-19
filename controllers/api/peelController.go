package api

import (
	"github.com/jepsonwu/inchat_go/models"
)

type PeelController struct {
	BaseController
}

type GetValidation struct {
	PaginateLastId
}

func (c *PeelController) Get() {
	c.Valid()
	//last_id per_page
	//c.Input()
	//c.Ctx.WriteString(beego.AppConfig.String("mysql_host"))

	//response
	peelInfo := models.NewPeel()
	c.Success(peelInfo.GetByPrimary())
}
