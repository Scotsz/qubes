package game

import (
	pb "qubes/internal/api"
	"qubes/internal/model"
)

type NetUpdate struct {
	blocks  []*WorldUpdate
	players map[model.PlayerID]*PlayerUpdate
}

func NewNetUpdate() *NetUpdate {
	return &NetUpdate{
		blocks:  make([]*WorldUpdate, 0),
		players: make(map[model.PlayerID]*PlayerUpdate, 0),
	}
}

type WorldUpdate struct {
	point   model.Point
	newType pb.BlockType
	tick    model.TickID
}

type PlayerUpdate struct {
	X, Y, Z float32
	id      model.PlayerID
	tick    model.TickID
}

func NewPlayerUpdate(player *model.Player, tick model.TickID) *PlayerUpdate {
	return &PlayerUpdate{
		X:    player.X,
		Y:    player.Y,
		Z:    player.Z,
		id:   player.ID,
		tick: tick,
	}
}
