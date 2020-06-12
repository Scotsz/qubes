package game

import (
	"github.com/stretchr/testify/assert"
	pb "qubes/api"
	"testing"
)

func TestWorld_SetBlock(t *testing.T) {
	world := NewWorld(256, 256, 64) //4194304

	valid := []Point{
		{X: 34, Y: 255, Z: 63},
		{X: 255, Y: 255, Z: 63},
		{X: 0, Y: 0, Z: 0},
	}

	invalid := []Point{
		{X: -1, Y: 0, Z: 0},
		{X: 0, Y: -1, Z: 0},
		{X: 0, Y: 0, Z: -1},
		{X: 256, Y: 0, Z: 0},
		{X: 0, Y: 256, Z: 0},
		{X: 0, Y: 0, Z: 64},
		{X: -2134, Y: 52543, Z: -5182},
		{X: -2134, Y: 23134342311, Z: -5182},
	}
	for _, p := range valid {
		world.SetBlock(p, &Block{blockType: pb.BlockType_Root})
		block := world.GetBlock(p)
		assert.NotNil(t, block)
		assert.Equal(t, pb.BlockType_Root, block.blockType)
	}

	for _, p := range invalid {
		world.SetBlock(p, &Block{blockType: pb.BlockType_Root})
		block := world.GetBlock(p)
		assert.Nil(t, block)
	}

}
func world1(t *testing.T) *World {
	t.Helper()
	w := NewWorld(8, 8, 8)
	points := []Point{
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
	w.FillPoints(points, &Block{blockType: pb.BlockType_Root})
	return w
}
func world2(t *testing.T) *World {
	t.Helper()
	w := NewWorld(5, 5, 5)
	points := []Point{
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 2},
		//{1, 1, 3},
		//{1, 3, 2},
	}
	w.FillPoints(points, &Block{blockType: pb.BlockType_Root})
	return w
}
func world3(t *testing.T) *World {
	t.Helper()
	w := NewWorld(5, 5, 5)
	points := []Point{
		{0, 1, 1},
		{0, 1, 2},
		{0, 1, 3},
		{0, 2, 3},
		{0, 1, 0},
	}
	w.FillPoints(points, &Block{blockType: pb.BlockType_Root})
	return w
}

func TestWorld_DestroyBlock(t *testing.T) {
	w3 := world3(t)
	blocks := w3.DestroyBlock(Point{0, 1, 1})
	assert.Equal(t, 4, len(blocks), blocks)

	w2 := world2(t)
	blocks = w2.DestroyBlock(Point{1, 1, 1})
	assert.Equal(t, 2, len(blocks), blocks)

	//w1
	w1 := world1(t)
	blocks = w1.DestroyBlock(Point{6, 2, 5})
	assert.Equal(t, 1, len(blocks), blocks)

	blocks = w1.DestroyBlock(Point{2, 2, 5})
	assert.Equal(t, 8, len(blocks), blocks)

	blocks = w1.DestroyBlock(Point{6, 2, 2})
	assert.Equal(t, 1, len(blocks), blocks)

	blocks = w1.DestroyBlock(Point{5, 2, 3})
	assert.Equal(t, 3, len(blocks), blocks)

	blocks = w1.DestroyBlock(Point{2, 2, 1})
	assert.Equal(t, 6, len(blocks), blocks)

	blocks = w1.DestroyBlock(Point{7, 7, 7})
	assert.Equal(t, 0, len(blocks), blocks)
}
