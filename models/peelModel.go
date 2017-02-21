package models

import (
	"github.com/astaxie/beego/orm"
)

type Peel struct {
	Id       uint64 `json:"id"`
	UserId   uint64 `json:"user_id"`
	Location string `json:"location"`
}

func NewPeel() *Peel {
	return &Peel{}
}

func init() {
	orm.RegisterDataBase("peel", "mysql", sprintfDsn("inchat_peel"))
	orm.RegisterModel(NewPeel())
	o = orm.NewOrm()
	o.Using("peel")
}

func (p *Peel) GetByPrimary(peelId uint64) (error, *Peel) {
	p.Id = peelId
	err := o.Read(p)
	if err != nil {
		return err, nil
	}

	return nil, p
}
