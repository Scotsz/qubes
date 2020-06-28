package game

import (
	pb "qubes/internal/api"
	"qubes/internal/model"
)

type NetUpdate struct {
	blocks  []*WorldUpdate
	players []*PlayerUpdate
}

func NewNetUpdate() *NetUpdate {
	return &NetUpdate{
		blocks:  make([]*WorldUpdate, 0),
		players: make([]*PlayerUpdate, 0),
	}
}

type WorldUpdate struct {
	points  []model.Point
	newType pb.BlockType
	tick    model.TickID
}

type PlayerUpdate struct {
	X, Y, Z float32
	Name    string
	tick    model.TickID
}

func NewPlayerUpdate(player *model.Player, tick model.TickID) *PlayerUpdate {
	return &PlayerUpdate{
		X:    player.X,
		Y:    player.Y,
		Z:    player.Z,
		Name: player.Name,
		tick: tick,
	}
}
