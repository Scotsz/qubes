package game

import (
	"context"
	"go.uber.org/zap"
	pb "qubes/internal/api"
)

type WorldManager struct {
	world   *World
	network *NetworkManager
	logger  *zap.SugaredLogger

	destroy chan Point
	place   chan Point
}

func NewWorldManager(logger *zap.SugaredLogger, network *NetworkManager) *WorldManager {
	return &WorldManager{
		destroy: make(chan Point, 10),
		network: network,
		logger:  logger,
	}
}

func (w *WorldManager) WithTestWorld() *WorldManager {
	w.world = GetTestWorld()
	return w
}

func (w *WorldManager) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case point := <-w.destroy: //point := <-w.destroy:
			deleting := w.world.destroyBlock(ctx, point)
			w.network.SendWorldUpdate(deleting, pb.BlockType_Air)
		case point := <-w.place:
			w.network.SendWorldUpdate([]Point{point}, pb.BlockType_Root)
		}

	}
}
func (w *WorldManager) TryPlace(p Point, block pb.BlockType) {
	w.place <- p
}
func (w *WorldManager) TryRemove(p Point) {
	w.destroy <- p
}

type WorldUpdate struct {
	points  []Point
	newType pb.BlockType
}

func (c *WorldUpdate) ToProto() *pb.Change {
	points := make([]*pb.WorldPoint, 0)
	for _, c := range c.points {
		points = append(points, &pb.WorldPoint{X: int32(c.X), Y: int32(c.Y), Z: int32(c.Y)})
	}
	return &pb.Change{Point: points, BlockType: c.newType}
}

func (w *WorldManager) ApplyUpdates(updates []*WorldUpdate) {
	for _, c := range updates {
		for _, p := range c.points {
			w.world.SetBlock(p, c.newType)
		}
	}
}
