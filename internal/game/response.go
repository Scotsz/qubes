package game

import (
	pb "qubes/internal/api"
	"qubes/internal/model"
)

type ResponseBuilder struct {
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{}
}

func (r *ResponseBuilder) NetUpdate(update *NetUpdate, tick model.TickID) *pb.NetUpdate {
	return &pb.NetUpdate{
		Blocks:   r.WorldUpdates(update.blocks),
		Entities: r.AllPlayers(update.players),
		Tick:     uint64(tick),
	}
}

func (r *ResponseBuilder) WorldUpdates(cs []*WorldUpdate) []*pb.Block {
	ch := make([]*pb.Block, len(cs))
	for i, c := range cs {
		point := &pb.WorldPoint{X: int32(c.point.X), Y: int32(c.point.Y), Z: int32(c.point.Z)}
		ch[i] = &pb.Block{Point: point, BlockType: c.newType}
	}
	return ch
}

func (r *ResponseBuilder) AllPlayers(players map[model.PlayerID]*PlayerUpdate) *pb.EntityUpdates {
	pbplayers := make([]*pb.Player, len(players))
	i := 0
	for id, p := range players {
		pbplayers[i] = &pb.Player{
			Id:  string(id),
			Pos: &pb.FloatPoint{X: p.X, Y: p.Y, Z: p.Z},
		}
		i++
	}
	return &pb.EntityUpdates{Players: pbplayers}
}

func (r *ResponseBuilder) PlayerConnected(id string) *pb.PlayerConnected {
	player := &pb.Player{
		Pos: &pb.FloatPoint{X: 0, Y: 0, Z: 0},
		Id:  id}
	return &pb.PlayerConnected{Player: player}
}

func (r *ResponseBuilder) PlayerDisconnected(id string) *pb.PlayerDisconnected {
	return &pb.PlayerDisconnected{Id: id}
}

func (r *ResponseBuilder) Block(p model.Point, btype pb.BlockType) *pb.Block {
	return &pb.Block{
		Point:     &pb.WorldPoint{X: int32(p.X), Y: int32(p.Y), Z: int32(p.Z)},
		BlockType: btype,
	}
}
func (r *ResponseBuilder) World(blocks []*pb.Block) *pb.World {
	return &pb.World{Blocks: blocks}
}
