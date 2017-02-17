package api

import (
	"github.com/jepsonwu/inchat_go/models"
)

type PeelController struct {
	BaseController
}

func (c *PeelController) Get() {
	//last_id per_page
	//c.Input()
	//c.Ctx.WriteString(beego.AppConfig.String("mysql_host"))

	//response
	peelInfo := &models.Peel{}
	c.success(peelInfo.GetByPrimary())
}
