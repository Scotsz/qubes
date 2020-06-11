package game

import (
	pb "qubes/api"
	"testing"
)

func TestWorld_SetBlock(t *testing.T) {
	world := NewWorld(256, 256, 64) //4194304

	valid := []Point{
		{X: 34, Y: 256, Z: 64},
		{X: 256, Y: 256, Z: 64},
		{X: 0, Y: 0, Z: 0},
	}
	//invalid := []Point{
	//	{X: -1, Y: 0, Z: 0},
	//	{X: 0, Y: -1, Z: 0},
	//	{X: 0, Y: 0, Z: -1},
	//	{X: 257, Y: 0, Z: 0},
	//	{X: 0, Y: 257, Z: 0},
	//	{X: 0, Y: 0, Z: 65},
	//	{X: -2134, Y: 52543, Z: -5182},
	//}

	for _, p := range valid {
		world.SetBlock(p, &Block{blockType: pb.BlockType_Root})
		block := world.GetBlock(p)
		if block == nil {
			t.Errorf("block %v:%v:%v is nil", p.X, p.Y, p.Z)
		} else if block.blockType != pb.BlockType_Root {
			t.Errorf("wanted %v got %v", pb.BlockType_Root, block.blockType)
		}
	}

}

func TestWorld_DestroyBlock(t *testing.T) {
	w := NewWorld(5, 5, 5)
	w.SetFloor(pb.BlockType_Root)

	points := []Point{
		{1, 1, 0},
		{1, 1, 2},
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 2},
		{4, 1, 2},
	}
	for _, p := range points {
		w.SetBlock(p, &Block{
			blockType: pb.BlockType_Root})
	}

	blocks := w.DestroyBlock(Point{1, 1, 2})
	if len(blocks) != 5 {
		t.Errorf("got %v want %v", len(blocks), 5)
	}
}
