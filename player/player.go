package player

import (
	"mahjong-table-go/actions"
	"mahjong-table-go/cards"
)

type Player struct {
	Id   int
	Name string

	SeatId int

	cards.Cards

	PlayerActions
}

func (p *Player) Init() {
	p.Cards = cards.NewCards()
	p.PlayerActions.Init()
}

func NewPlayer(seatId int) *Player {

	p := &Player{SeatId: seatId}
	p.Init()
	return p
}

func (p *Player) GetInfo() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *Player) GetShadowInfo() {

}

func (p *Player) DoAction(action actions.BaseAction) {

}
