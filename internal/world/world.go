package world

type BlockType int

const (
	TypeAir BlockType = iota
	TypeRoot
)

type Block struct {
	blockType BlockType
}

type World struct {
	width, height, depth int
	blocks               []*Block
}

func NewWorld(w, h int) *World {
	world := &World{
		width:  w,
		height: h,
		depth:  1,
		blocks: make([]*Block, w*h, w*h),
	}
	world.Fill(Point{0, 0, 0}, Point{w - 1, h - 1, 0}, &Block{blockType: TypeAir})
	return world
}

func (w *World) Fill(start, end Point, block *Block) {
	for i := start.Y; i <= end.Y; i++ {
		for j := start.X; j <= end.X; j++ {
			w.SetBlock(Point{j, i, 0}, block)
		}
	}
}

func (w *World) SetFloor(block *Block) {
	for i := 0; i < w.width; i++ {
		w.SetBlock(Point{i, w.height - 1, 0}, block)
	}
}

func (w *World) GetBlock(p Point) *Block {
	return w.blocks[p.X+w.width*(p.Y+w.depth*p.Z)]
}

func (w *World) SetBlock(p Point, block *Block) {
	w.blocks[p.X+w.width*(p.Y+w.depth*p.Z)] = block
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
	newType BlockType
}

func (w *World) CalculateDestroyChange(p Point) *Change {
	return &Change{
		points:  []Point{p},
		newType: TypeAir,
	}
}

type Point struct {
	X, Y, Z int
}
