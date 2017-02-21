package api

import (
	"github.com/jepsonwu/inchat_go/models"
)

type PeelController struct {
	BaseController
}

type GetRequest struct {
	PeelId uint64 `form:"peel_id"`
}

func (c *PeelController) Get() {
	requestParams := &GetRequest{}
	c.ParseForm(requestParams)//如果类型解析失败 则后面都解析失败
	valid.Required(requestParams.PeelId, "peel_id").Message("peel_id_invalid")
	c.valid()

	//response
	peelModel := models.NewPeel()
	err, peelDetail := peelModel.GetByPrimary(requestParams.PeelId)
	if err != nil {
		c.success(struct {
		}{})
	} else {
		c.success(peelDetail)
	}
}
