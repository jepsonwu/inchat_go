package models

type Peel struct {
	Id       uint64 `json:"id"`
	UserId   uint64 `json:"user_id"`
	Location string `json:"location"`
}

func NewPeel() *Peel {
	return &Peel{}
}

func (peel *Peel) GetByPrimary() *Peel {
	peel.Id = 1
	peel.UserId = 10098
	peel.Location = "杭州"

	return peel
}
