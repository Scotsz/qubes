package game

import (
	pb "qubes/api"
)

type Point struct {
	X, Y, Z int
}

type Block struct {
	blockType pb.BlockType
}

type World struct {
	width, height, depth int
	blocks               []*Block
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

func NewWorld(w, h, d int) *World {
	max := w * h * d
	world := &World{
		width:  w,
		height: h,
		depth:  d,
		blocks: make([]*Block, max, max),
	}
	world.Fill(Point{0, 0, 0}, Point{w - 1, h - 1, d - 1}, pb.BlockType_Air)
	return world
}

func (w *World) Fill(start, end Point, btype pb.BlockType) {
	for i := start.X; i <= end.X; i++ {
		for j := start.Y; j <= end.Y; j++ {
			for k := start.Z; k <= end.Z; k++ {
				w.SetBlock(Point{i, j, k}, &Block{blockType: btype})
			}
		}
	}
}

func (w *World) SetFloor(btype pb.BlockType) {
	w.Fill(
		Point{X: 0, Y: 0, Z: w.depth - 1},
		Point{X: w.width - 1, Y: w.height - 1, Z: w.depth - 1},
		btype)
}

func (w *World) GetBlock(p Point) *Block {
	if w.isValid(p) {
		return w.blocks[w.getPos(p.X, p.Y, p.Z)]
	}
	return nil
}

func (w *World) SetBlock(p Point, block *Block) {
	if w.isValid(p) {
		w.blocks[w.getPos(p.X, p.Y, p.Z)] = block
	}
}

func (w *World) ApplyChanges(changes []*Change) {
	for _, c := range changes {
		for _, p := range c.points {
			w.SetBlock(p, &Block{blockType: c.newType})
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
	return !(p.X > w.width || p.Y > w.height || p.Z > w.depth) && !(p.X < 0 || p.Y < 0 || p.Z < 0)
}

func (w *World) getPos(x, y, z int) int {
	return x + y*w.width + z*w.depth*w.height
}

func (w *World) isSolid(p Point) bool {
	switch w.GetBlock(p).blockType {
	case pb.BlockType_Air:
		return false
	}
	return true
}

func (w *World) DestroyBlock(p Point) []Point {
	var queue []Point
	marked := make(map[Point]bool)

	queue = append(queue, p)

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if c.Z >= w.height-1 {
			return nil
		}

		points := []Point{
			{c.X, c.Y, c.Z - 1},
			{c.X - 1, c.Y, c.Z},
			{c.X, c.Y - 1, c.Z},
			{c.X + 1, c.Y, c.Z},
			{c.X, c.Y + 1, c.Z},
			{c.X, c.Y, c.Z + 1},
		}

		for _, point := range points {
			if w.isValid(point) && w.isSolid(point) && !marked[point] {
				marked[point] = true
				queue = append(queue, point)
			}
		}
	}

	destroyed := make([]Point, 0, len(marked))
	for p := range marked {
		destroyed = append(destroyed, p)
	}
	return destroyed
}
