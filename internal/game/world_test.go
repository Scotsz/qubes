package game

import (
	pb "qubes/api"
	"testing"
)

func TestWorld_SetBlock(t *testing.T) {
	world := NewWorld(256, 256, 64) //4194304
	p := Point{
		X: 34,
		Y: 257,
		Z: 64,
	}
	world.SetBlock(p, &Block{blockType: pb.BlockType_Root})

	block := world.GetBlock(p)
	if block == nil {
		t.Error()
	} else if block.blockType != pb.BlockType_Root {
		t.Errorf("wanted %v got %v", pb.BlockType_Root, block.blockType)
	}
}

func TestWorld_SetFloor(t *testing.T) {

}
