package game

import (
	"context"
	"go.uber.org/zap"
	"qubes/internal/model"
	"time"
)

type Command func(ctx context.Context)

type State struct {
	players map[model.PlayerID]*model.Player
	world   *model.World
	tick    model.TickID
}

func NewState(world *model.World) *State {
	return &State{
		players: make(map[model.PlayerID]*model.Player),
		world:   world,
		tick:    0,
	}
}

type stateManager struct {
	logger         *zap.SugaredLogger
	commandQueue   chan Command
	commandFactory *CommandFactory
}

func NewStateManager(logger *zap.SugaredLogger, cf *CommandFactory) *stateManager {
	return &stateManager{
		commandQueue:   make(chan Command),
		logger:         logger,
		commandFactory: cf,
	}
}

func (s *stateManager) Run(ctx context.Context) {
	s.logger.Debug("WorldManager running")
	simTicker := time.NewTicker(time.Millisecond * 50)

	for {
		select {
		case <-ctx.Done():
			return

		case cmd := <-s.commandQueue:
			cmd(ctx)

		case <-simTicker.C:
			s.commandFactory.ProcessTickers()(ctx)
			s.commandFactory.IncTick()(ctx)
		}
	}

}
func (s *stateManager) AddCommand(command Command) {
	s.commandQueue <- command
}
