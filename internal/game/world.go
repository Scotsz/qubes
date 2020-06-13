package game

import (
	pb "qubes/api"
)

type Point struct {
	X, Y, Z int
}

type World struct {
	width, height, depth int
	blocks               []pb.BlockType
}

func NewWorld(w, d, h int) *World {
	max := w * h * d
	world := &World{
		width:  w,
		height: h,
		depth:  d,
		blocks: make([]pb.BlockType, max, max),
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
	return pb.BlockType_Debug
}

func (w *World) SetBlock(p Point, block pb.BlockType) {
	if w.isValid(p) {
		w.blocks[w.getPos(p.X, p.Y, p.Z)] = block
	}
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

func (w *World) solidNeighbors(point Point) []Point {
	n := make([]Point, 0)
	for _, p := range neighbors(point) {
		if w.isValid(p) && w.isSolid(p) {
			n = append(n, p)
		}
	}
	return n
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
