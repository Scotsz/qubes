package game

import (
	"context"
	pb "qubes/internal/api"
	"qubes/internal/model"
)

type DestroyBlockCommand struct {
	world   *World
	network *NetworkManager
	point   Point
	tick    model.TickID
}

func (d DestroyBlockCommand) execute(ctx context.Context) {
	deleting := d.world.DestroyBlock(ctx, d.point)
	if len(deleting) > 0 {
		update := &WorldUpdate{
			points:  deleting,
			newType: pb.BlockType_Air,
			tick:    d.tick,
		}
		d.network.AddWorldUpdate(update)
	}
}

type AddPlayerCommand struct {
	players PlayerStorage
	id      model.ClientID
	name    string
}

func (a AddPlayerCommand) execute(ctx context.Context) {
	a.players[a.id] = NewPlayer(a.name)
}

type RemovePlayerCommand struct {
	players PlayerStorage
	id      model.ClientID
}

func (r RemovePlayerCommand) execute(ctx context.Context) {
	delete(r.players, r.id)
}
