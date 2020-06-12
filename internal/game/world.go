package game

import (
	"log"
	pb "qubes/api"
)

type Point struct {
	X, Y, Z int
}

//type Block struct {
//	blockType pb.BlockType
//}
type World struct {
	width, height, depth int
	blocks               []pb.BlockType
	deleting             map[Point]bool //TODO: convert to []Pos
}
type Change struct {
	points  []Point
	newType pb.BlockType
}

func (c *Change) ToProto() *pb.Change {
	points := make([]*pb.WorldPoint, 0)
	for _, c := range c.points {
		points = append(points, &pb.WorldPoint{X: int32(c.X), Y: int32(c.Y), Z: int32(c.Y)})
	}
	return &pb.Change{Point: points, BlockType: c.newType}
}

func NewWorld(w, d, h int) *World {
	max := w * h * d
	world := &World{
		width:    w,
		height:   h,
		depth:    d,
		blocks:   make([]pb.BlockType, max, max),
		deleting: make(map[Point]bool),
	}
	world.Fill(Point{0, 0, 0}, Point{w - 1, d - 1, h - 1}, pb.BlockType_Air)
	return world
}

func (w *World) Fill(start, end Point, btype pb.BlockType) {
	for i := start.X; i <= end.X; i++ {
		for j := start.Y; j <= end.Y; j++ {
			for k := start.Z; k <= end.Z; k++ {
				w.SetBlock(Point{i, j, k}, btype)
			}
		}
	}
}

func (w *World) FillPoints(points []Point, block pb.BlockType) {
	for _, p := range points {
		w.SetBlock(p, block)
	}
}

func (w *World) SetFloor(btype pb.BlockType) {
	w.Fill(
		Point{X: 0, Y: 0, Z: 0},
		Point{X: w.width - 1, Y: w.depth - 1, Z: 0},
		btype)
}

func (w *World) GetBlock(p Point) pb.BlockType {
	if w.isValid(p) {
		return w.blocks[w.getPos(p.X, p.Y, p.Z)]
	}
	log.Fatalf("trying to get %v", p)
	return pb.BlockType_Debug
}

func (w *World) SetBlock(p Point, block pb.BlockType) {
	if w.isValid(p) {
		w.blocks[w.getPos(p.X, p.Y, p.Z)] = block
	} else {
		log.Fatalf("trying to set %v", p)
	}

}

func (w *World) ApplyChanges(changes []*Change) {
	for _, c := range changes {
		for _, p := range c.points {
			w.SetBlock(p, c.newType)
		}
	}
}

func (w *World) CalculateDestroyChange(p Point) *Change {
	return &Change{
		points:  w.DestroyBlock(p),
		newType: pb.BlockType_Air,
	}
}

func (w *World) isValid(p Point) bool {
	return !(p.X >= w.width || p.Y >= w.depth || p.Z >= w.height) && !(p.X < 0 || p.Y < 0 || p.Z < 0)
}

func (w *World) getPos(x, y, z int) int {
	return x + y*w.width + z*w.depth*w.width
}

func (w *World) isSolid(p Point) bool {
	switch w.GetBlock(p) {
	case pb.BlockType_Air:
		return false
	}
	return true
}

func (w *World) DestroyBlock(p Point) []Point {
	var c int
	var points []Point

	if !(w.isValid(p) && w.isSolid(p)) {
		return nil
	}
	w.SetBlock(p, pb.BlockType_Air)
	w.deleting[p] = true

	for _, p := range neighbors(p) {
		if w.isValid(p) && w.isSolid(p) {
			c += w.getConnected(p)
		}
	}

	for p := range w.deleting {
		points = append(points, p)
	}

	w.deleting = make(map[Point]bool)
	return points
}

func neighbors(p Point) []Point {
	return []Point{
		{p.X, p.Y, p.Z - 1},
		{p.X - 1, p.Y, p.Z},
		{p.X, p.Y - 1, p.Z},
		{p.X + 1, p.Y, p.Z},
		{p.X, p.Y + 1, p.Z},
		{p.X, p.Y, p.Z + 1},
	}
}

func (w *World) getConnected(point Point) int {
	queue := []Point{point}
	marked := make(map[Point]bool)
	marked[point] = true

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if c.Z == 0 {
			return 0
		}

		for _, point := range neighbors(c) {
			if w.isValid(point) && w.isSolid(point) && !marked[point] && !w.deleting[point] {
				marked[point] = true
				queue = append(queue, point)
			}
		}
	}

	for p := range marked {
		w.deleting[p] = true
	}
	return len(marked)
}
