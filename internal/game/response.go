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

func (r *ResponseBuilder) WorldUpdates(cs []*WorldUpdate) *pb.Changes {
	ch := make([]*pb.Change, len(cs))
	for i, c := range cs {
		points := make([]*pb.WorldPoint, 0)
		for _, c := range c.points {
			points = append(points, &pb.WorldPoint{X: int32(c.X), Y: int32(c.Y), Z: int32(c.Y)})
		}
		ch[i] = &pb.Change{Point: points, BlockType: c.newType}
	}
	return &pb.Changes{Changes: ch}
}

func (r *ResponseBuilder) AllPlayers(players map[model.ClientID]*Player, tick model.TickID) *pb.AllPlayers {
	pbplayers := make([]*pb.Player, len(players))
	i := 0
	for id, p := range players {
		pbplayers[i] = &pb.Player{
			Id:  string(id),
			Pos: &pb.FloatPoint{X: p.X, Y: p.Y, Z: p.Z},
		}
	}
	return &pb.AllPlayers{Tick: tick.ToUint64(), Players: pbplayers}
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

func (r *ResponseBuilder) Block(p Point, btype pb.BlockType) *pb.Block {
	return &pb.Block{
		Point:     &pb.WorldPoint{X: int32(p.X), Y: int32(p.Y), Z: int32(p.Z)},
		BlockType: btype,
	}
}
func (r *ResponseBuilder) World(blocks []*pb.Block) *pb.World {
	return &pb.World{Blocks: blocks}
}
