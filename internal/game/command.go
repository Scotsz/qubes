package game

import (
	"context"
	pb "qubes/internal/api"
	"qubes/internal/model"
)

type CommandFactory struct {
	network *NetworkManager
	state   *State
}

func NewCommandFactory(network *NetworkManager, state *State) *CommandFactory {
	return &CommandFactory{
		network: network,
		state:   state,
	}
}
func (c *CommandFactory) RemoveBlock(p model.Point) Command {
	return func(ctx context.Context) {
		deleting := c.state.world.DestroyBlock(ctx, p)
		if len(deleting) > 0 {
			update := &WorldUpdate{
				points:  deleting,
				newType: pb.BlockType_Air,
				tick:    c.state.tick,
			}
			c.network.AddWorldUpdate(update)
		}
	}
}

func (c *CommandFactory) AddPlayer(id model.PlayerID) Command {
	return func(ctx context.Context) {
		c.state.players[id] = model.NewPlayer(id)
		c.network.SendPlayerConnected(id)
	}
}

func (c *CommandFactory) RemovePlayer(id model.PlayerID) Command {
	return func(ctx context.Context) {
		delete(c.state.players, id)
		c.network.SendPlayerDisconnected(id)
	}
}

func (c *CommandFactory) IncTick() Command {
	return func(ctx context.Context) {
		c.state.tick++
	}
}

func (c *CommandFactory) ProcessTickers() Command {
	return func(ctx context.Context) {
		for _, p := range c.state.players {
			p.Tick()
			if p.ShouldUpdate() {
				c.network.AddPlayerUpdate(NewPlayerUpdate(p, c.state.tick))
			}
		}
	}
}
