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

func (r *ResponseBuilder) WorldUpdates(cs []*WorldUpdate, tick model.TickID) *pb.Changes {
	ch := make([]*pb.Change, len(cs))
	for i, c := range cs {
		ch[i] = c.ToProto()
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

func (r *ResponseBuilder) PlayerConnected(id string, tick model.TickID) *pb.PlayerConnected {
	player := &pb.Player{
		Pos: &pb.FloatPoint{X: 0, Y: 0, Z: 0},
		Id:  id}
	return &pb.PlayerConnected{Player: player}
}

func (r *ResponseBuilder) PlayerDisconnected(id string, tick model.TickID) *pb.PlayerDisconnected {
	return &pb.PlayerDisconnected{Id: id}
}
