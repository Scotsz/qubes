package model

import "github.com/google/uuid"

type ClientID string
type PlayerID string
type TickID uint64

func (t TickID) ToUint64() uint64 {
	return uint64(t)
}

func NewID() ClientID {
	str := uuid.New().String()
	return ClientID(str)
}

type Point struct {
	X, Y, Z int
}
type FPoint struct {
	X, Y, Z float64
}
