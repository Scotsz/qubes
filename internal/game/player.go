package game

import (
	"log"
	"math"
)

type Player struct {
	X, Y         float64
	xDest, yDest int32
	dx, dy       float64
	speed        float64
}

func NewPlayer() *Player {
	return &Player{
		speed: 2,
		X:     1,
		Y:     1,
	}
}
func (p *Player) SetDest(x, y int32) {
	p.xDest = x
	p.yDest = y

	vectorX := float64(p.xDest) - p.X
	vectorY := float64(p.yDest) - p.Y

	length := math.Sqrt(vectorX*vectorX + vectorY*vectorY)

	p.dx = vectorX / length
	p.dy = vectorY / length
	log.Printf("xy:[%v:%v] dest[%v:%v] len[%v]", p.X, p.Y, p.xDest, p.yDest, length)
}
func (p *Player) move() {
	p.X += p.dx * p.speed
	p.Y += p.dy * p.speed

	vectorX := float64(p.xDest) - p.X
	vectorY := float64(p.yDest) - p.Y
	length := math.Sqrt(vectorX*vectorX + vectorY*vectorY)
	if length < 5 {
		p.dx = 0
		p.dy = 0
	}

}
func (p *Player) Tick() {
	defer p.move()
}
