package game

import (
	"context"
	"go.uber.org/zap"
	"math/rand"
	pb "qubes/internal/api"
	"qubes/internal/model"
	"strconv"
	"time"
)

type Command interface {
	execute(ctx context.Context)
}
type PlayerStorage map[model.ClientID]*model.Player

type state struct {
	logger *zap.SugaredLogger

	players PlayerStorage

	worldDiff map[model.Point]pb.BlockType
	world     *model.World

	commandQueue chan Command

	network *NetworkManager
	tick    model.TickID
}

func NewWorldManager(logger *zap.SugaredLogger, world *model.World, manager *NetworkManager) *state {
	return &state{
		commandQueue: make(chan Command),
		logger:       logger,
		players:      make(map[model.ClientID]*model.Player),
		world:        world,
		network:      manager,
	}
}

func (s *state) processTickers() {
	for _, p := range s.players {
		p.Tick()
	}
}

func (s *state) Run(ctx context.Context) {
	s.logger.Debug("WorldManager running")
	simTicker := time.NewTicker(time.Millisecond * 50)

	for {
		select {
		case <-ctx.Done():
			return

		case cmd := <-s.commandQueue:
			cmd.execute(ctx)

		case <-simTicker.C:
			s.processTickers()
			for _, p := range s.players {
				if p.ShouldUpdate() {
					s.network.AddPlayerUpdate(NewPlayerUpdate(p, s.tick))
				}
			}

			s.tick++
		}
	}

}
func (s *state) RemoveBlock(p model.Point) {
	s.commandQueue <- DestroyBlockCommand{
		world:   s.world,
		network: s.network,
		point:   p,
		tick:    s.tick,
	}
}

func (s *state) AddCommand(command Command) {
	s.commandQueue <- command
}

func (s *state) AddPlayer(id model.ClientID) {
	s.AddCommand(AddPlayerCommand{players: s.players, id: id, name: strconv.Itoa(rand.Int())})
}

func (s *state) RemovePlayer(id model.ClientID) {
	s.AddCommand(RemovePlayerCommand{players: s.players, id: id})
}
