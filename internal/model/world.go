package model

import (
	"context"
	pb "qubes/internal/api"
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

func (w *World) DestroyBlock(ctx context.Context, p Point) []Point {
	var c int

	if !(w.isValid(p) && w.isSolid(p)) {
		return nil
	}
	if len(w.solidNeighbors(p)) > 2 {
		return []Point{p}
	}

	w.SetBlock(p, pb.BlockType_Air)

	deleting := make(map[Point]bool)
	grounded := make(map[Point]bool)
	deleting[p] = true

	for _, p := range w.solidNeighbors(p) {
		c += w.getConnected(ctx, p, deleting, grounded)
	}
	points := make([]Point, 0)
	for p := range deleting {
		points = append(points, p)
	}
	return points
}

func (w *World) getConnected(ctx context.Context, point Point, deleting map[Point]bool, grounded map[Point]bool) int {
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
			if w.isValid(p) && w.isSolid(p) && !marked[p] {

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

func GetTestWorld(n int) *World {
	w := NewWorld(8, 8, 8)
	var points []Point
	switch n {
	case 0:
		points = w0()
	case 1:
		points = w1()
	}
	w.FillPoints(points, pb.BlockType_Root)
	return w
}

func w1() []Point {
	return []Point{
		{1, 1, 0}, {3, 1, 0}, {5, 1, 0}, {7, 1, 0},
		{1, 1, 1}, {3, 1, 1}, {5, 1, 1}, {7, 1, 1},
		{1, 1, 2}, {3, 1, 2}, {5, 1, 2}, {7, 1, 2},
		{1, 1, 3}, {3, 1, 3}, {5, 1, 3}, {7, 1, 3},
		{1, 1, 4}, {3, 1, 4}, {5, 1, 4}, {7, 1, 4},
		{1, 1, 5}, {3, 1, 5}, {5, 1, 5}, {7, 1, 5},
		{1, 1, 6}, {3, 1, 6}, {5, 1, 6}, {7, 1, 6},
		{1, 1, 7}, {3, 1, 7}, {5, 1, 7}, {7, 1, 7},
	}
}
func w0() []Point {
	return []Point{
		{1, 1, 0},
		{2, 1, 0},
		{1, 2, 0},
		{2, 2, 0},
		{6, 2, 0},
		{1, 3, 0},
		{2, 3, 0},
		{2, 2, 1}, //5 out=6
		{6, 2, 1},
		{2, 2, 2},
		{6, 2, 2}, //3 out=1
		{2, 2, 3},
		{3, 2, 3},
		{4, 2, 3},
		{5, 2, 3}, //4 out=3
		{6, 2, 3},
		{2, 2, 4},
		{6, 2, 4},
		{2, 2, 5}, //2 out=8
		{6, 2, 5}, //1 out=1
		{2, 2, 6},
		{6, 2, 6},
		{2, 2, 7},
		{3, 2, 7},
		{4, 2, 7},
		{5, 2, 7},
		{6, 2, 7},
	}
}
