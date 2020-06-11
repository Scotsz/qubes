package game

import pb "qubes/api"

type Block struct {
	blockType pb.BlockType
}

type World struct {
	width, height, depth int
	blocks               []*Block
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

func (w *World) SetFloor(block *Block) {
	for i := 0; i < w.width; i++ {
		w.SetBlock(Point{i, w.height - 1, 0}, block)
	}
}

func (w *World) GetBlock(p Point) *Block {
	if !(p.X > w.width || p.Y > w.height || p.Z > w.depth) {
		return w.blocks[p.X+w.width*p.Y+w.depth*w.height*p.Z]
	}
	return nil
}

func (w *World) SetBlock(p Point, block *Block) {
	if !(p.X > w.width || p.Y > w.height || p.Z > w.depth) {
		w.blocks[p.X+w.width*p.Y+w.depth*w.height*p.Z] = block
	}
}

func (w *World) ApplyChanges(changes []*Change) {
	for _, c := range changes {
		for _, p := range c.points {
			w.SetBlock(p, &Block{blockType: c.newType})
		}
	}
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

func (w *World) CalculateDestroyChange(p Point) *Change {
	return &Change{
		points:  []Point{p},
		newType: pb.BlockType_Air,
	}
}

type Point struct {
	X, Y, Z int
}
