package game

import (
	"context"
	pb "qubes/api"
)

type WorldManager struct {
	world   *World
	network *NetworkManager

	destroy chan Point
}

func NewWorldManager(world *World, network *NetworkManager) *WorldManager {
	return &WorldManager{
		destroy: make(chan Point, 10),
		network: network,
		world:   world,
	}
}

func (w *WorldManager) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-w.destroy: //point := <-w.destroy:
			//deleting := w.destroyBlock(ctx, point)
			w.network.SendWorldUpdate([]Point{{1, 2, 3}}, pb.BlockType_Air)
		}
	}
}
func (w *WorldManager) TryPlace(p Point, block pb.BlockType) {
	//TODO
}
func (w *WorldManager) TryRemove(p Point) {
	w.destroy <- p
}

func (w *WorldManager) destroyBlock(ctx context.Context, p Point) []Point {
	var c int

	if !(w.world.isValid(p) && w.world.isSolid(p)) {
		return nil
	}
	if len(w.world.solidNeighbors(p)) > 2 {
		return []Point{p}
	}

	w.world.SetBlock(p, pb.BlockType_Air)

	deleting := make(map[Point]bool)
	grounded := make(map[Point]bool)
	deleting[p] = true

	for _, p := range w.world.solidNeighbors(p) {
		c += w.getConnected(ctx, p, deleting, grounded)
	}

	points := make([]Point, 0)
	for p := range deleting {
		points = append(points, p)
	}

	return points
}

func (w *WorldManager) getConnected(ctx context.Context, point Point, deleting map[Point]bool, grounded map[Point]bool) int {
	queue := []Point{point}

	marked := make(map[Point]bool)
	marked[point] = true
	for len(queue) > 0 {
		select {
		case <-ctx.Done():
			return len(marked)
		default:
		}

		c := queue[0]
		queue = queue[1:]
		marked[c] = true
		if c.Z == 0 {
			for p := range marked {
				grounded[p] = true
			}
			return 0
		}

		for _, p := range neighbors(c) {
			if w.world.isValid(p) && w.world.isSolid(p) && !marked[p] {

				if grounded[p] {
					for m := range marked {
						grounded[m] = true
					}
					return 0
				}
				marked[p] = true
				queue = append(queue, p)
			}
		}
	}

	for p := range marked {
		deleting[p] = true
	}
	return len(marked)
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
