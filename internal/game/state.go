package game

import (
	"context"
	"go.uber.org/zap"
	pb "qubes/internal/api"
	"qubes/internal/model"
)

type state struct {
	logger *zap.SugaredLogger

	tickers map[model.ClientID]Ticker

	worldDiff map[Point]pb.BlockType
	world     *World

	destroy chan Point
	place   chan *WorldUpdate

	worldUpdates chan *WorldUpdate
	network      *NetworkManager
}

type Ticker interface {
	Tick()
}

type WorldManager interface {
	Run(ctx context.Context, tick <-chan model.TickID)
	PlaceBlocks(update *WorldUpdate)
	RemoveBlock(p Point)
	AddTicker(id model.ClientID, ticker Ticker)
	RemoveTicker(id model.ClientID)
}

func NewWorldManager(logger *zap.SugaredLogger, world *World, manager *NetworkManager) *state {
	return &state{
		destroy: make(chan Point, 10),
		logger:  logger,
		tickers: make(map[model.ClientID]Ticker),
		world:   world,
		network: manager,
	}
}

func (s *state) AddTicker(id model.ClientID, ticker Ticker) {
	//s.players.Add(id, NewPlayer())
	s.tickers[id] = ticker
}
func (s *state) RemoveTicker(id model.ClientID) {
	//s.players.Remove(id)
	delete(s.tickers, id)
}

func (s *state) processTickers() {
	for _, p := range s.tickers {
		p.Tick()
	}
}

func (s *state) Run(ctx context.Context, tick <-chan model.TickID) { //<-chan worldupdates
	s.logger.Info("WorldManager running")
	for {
		select {
		case <-ctx.Done():
			return
		case point := <-s.destroy:
			//		deleting := s.world.DestroyBlock(ctx, point)
			//update := &WorldUpdate{
			//	points:  deleting,
			//	newType: pb.BlockType_Air,
			//}
			s.network.AddWorldUpdate(&WorldUpdate{
				points:  []Point{point},
				newType: pb.BlockType_Air,
				tick:    1,
			})

		case update := <-s.place:
			//w.world.placeblocks
			s.network.AddWorldUpdate(update)

		case <-tick:
			s.processTickers()
		}

	}
}
func (s *state) PlaceBlocks(update *WorldUpdate) {
	s.place <- update
}
func (s *state) RemoveBlock(p Point) {
	s.destroy <- p
}

type WorldUpdate struct {
	points  []Point
	newType pb.BlockType
	tick    model.TickID
}
