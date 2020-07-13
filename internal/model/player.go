package model

import (
	"log"
	"math"
)

type Player struct {
	//Name string

	ID           PlayerID
	X, Y, Z      float32
	xDest, yDest int32
	dx, dy       float32
	speed        float32
}

func NewPlayer(id PlayerID) *Player {
	return &Player{
		ID:    id,
		speed: 2,
		X:     1,
		Y:     1,
	}
}

func (p *Player) SetDest(x, y, z int32) {
	p.xDest = x
	p.yDest = y

	vectorX := float32(p.xDest) - p.X
	vectorY := float32(p.yDest) - p.Y
	length := length(vectorX, vectorY, 0)

	p.dx = vectorX / length
	p.dy = vectorY / length
	log.Printf("xy:[%v:%v] dest[%v:%v] len[%v]", p.X, p.Y, p.xDest, p.yDest, length)
}
func (p *Player) move() {
	p.X += p.dx * p.speed
	p.Y += p.dy * p.speed

	vectorX := float32(p.xDest) - p.X
	vectorY := float32(p.yDest) - p.Y

	length := length(vectorX, vectorY, 0)
	if length < 5 {
		p.dx = 0
		p.dy = 0
	}

}
func (p *Player) Tick() {
	p.move()
}
func (p *Player) ShouldUpdate() bool {
	return true
}

func length(x, y, z float32) float32 {
	return float32(math.Sqrt(float64(x*x + y*y + z*z)))
}
