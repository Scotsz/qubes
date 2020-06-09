package game

import (
	pb "qubes/api"
	"qubes/internal/model"
)

type ResponseBuilder struct {
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{}
}
func (r *ResponseBuilder) envelope(tick model.TickID, payload *pb.Payload) *pb.Response {
	return &pb.Response{
		Tick:    uint64(tick),
		Payload: payload,
	}
}

func (r *ResponseBuilder) Changes(cs []*Change, tick model.TickID) *pb.Response {
	ch := make([]*pb.Change, len(cs))
	for i, c := range cs {
		ch[i] = c.ToProto()
	}
	payload := &pb.Payload{
		Type: &pb.Payload_Changes{
			Changes: &pb.Changes{Changes: ch}}}

	return r.envelope(tick, payload)
}

func (r *ResponseBuilder) AllPlayers(players map[model.ClientID]*Player, tick model.TickID) *pb.Response {
	pbplayers := make([]*pb.Player, len(players))
	i := 0
	for id, p := range players {
		pbplayers[i] = p.ToProto()
		pbplayers[i].Id = string(id)
		i++
	}

	payload := &pb.Payload{
		Type: &pb.Payload_Players{
			Players: &pb.AllPlayers{Player: pbplayers}}}

	return r.envelope(tick, payload)
}

func (r *ResponseBuilder) PlayerConnected(id string, tick model.TickID) *pb.Response {
	payload := &pb.Payload{
		Type: &pb.Payload_PlayerConnect{
			PlayerConnect: &pb.Player{
				Id: id}}}

	return r.envelope(tick, payload)
}
func (r *ResponseBuilder) PlayerDisconnected(id string, tick model.TickID) *pb.Response {
	payload := &pb.Payload{
		Type: &pb.Payload_PlayerDisconnect{
			PlayerDisconnect: &pb.Player{
				Id: id,
			}}}
	return r.envelope(tick, payload)
}
