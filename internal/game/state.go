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
type PlayerStorage map[model.ClientID]*Player

type state struct {
	logger *zap.SugaredLogger

	players PlayerStorage

	worldDiff map[Point]pb.BlockType
	world     *World

	commandQueue chan Command

	network *NetworkManager
	tick    model.TickID
}

func NewWorldManager(logger *zap.SugaredLogger, world *World, manager *NetworkManager) *state {
	return &state{
		commandQueue: make(chan Command),
		logger:       logger,
		players:      make(map[model.ClientID]*Player),
		world:        world,
		network:      manager,
	}
}

func (s *state) processTickers() {
	for _, p := range s.players {
		p.Tick()
	}
}

func (s *state) Run(ctx context.Context) { //<-chan worldupdates
	s.logger.Info("WorldManager running")
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
					upd := p.GetUpdate()
					upd.tick = s.tick
					s.network.AddPlayerUpdate(upd)
				}
			}

			s.tick++
		}
	}

}
func (s *state) RemoveBlock(p Point) {
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

type WorldUpdate struct {
	points  []Point
	newType pb.BlockType
	tick    model.TickID
}
type PlayerUpdate struct {
	X, Y, Z float32
	name    string
	tick    model.TickID
}
