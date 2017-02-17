package models

type Peel struct {
	id       uint64
	user_id  uint64
	location string
}

func (peel *Peel) GetByPrimary() *Peel {
	peel.id = 1
	peel.user_id = 10098
	peel.location = "杭州"

	return peel
}
