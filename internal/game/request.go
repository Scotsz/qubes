package game

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	pb "qubes/internal/api"
	"qubes/internal/model"
)

type RequestHandler struct {
	logger       *zap.SugaredLogger
	commandQueue chan *PlayerCommand

	worldManager *state
	network      *NetworkManager
}

type PlayerCommand struct {
	id model.ClientID
	*pb.Request
}

func (r *RequestHandler) AddRequest(id model.ClientID, req *pb.Request) {
	r.commandQueue <- &PlayerCommand{id, req}
}
func NewRequestHandler(
	logger *zap.SugaredLogger,
	worldManager *state,
	network *NetworkManager,

) *RequestHandler {
	return &RequestHandler{
		logger:       logger,
		commandQueue: make(chan *PlayerCommand, 20),
		network:      network,
		worldManager: worldManager,
	}
}

func (r *RequestHandler) Run(ctx context.Context) {
	r.logger.Info("RequestHandler running")

	for {
		select {
		case cmd := <-r.commandQueue:
			r.handleCommand(ctx, cmd)
		case <-ctx.Done():
			return
		}
	}
}

func (r *RequestHandler) handleCommand(ctx context.Context, cmd *PlayerCommand) {
	var payload string
	switch cmd.Command.(type) {

	case *pb.Request_Move:
		r.handleMove(ctx, cmd.id, cmd.GetMove())
		payload = fmt.Sprintf(cmd.GetMove().String())

	case *pb.Request_Shoot:
		r.handleShoot(ctx, cmd.id, cmd.GetShoot())
		payload = fmt.Sprintf(cmd.GetShoot().String())

	case *pb.Request_WorldDiff:
		r.handleWorldDiffRequest(ctx, cmd.id, cmd.GetWorldDiff())
		payload = fmt.Sprintf(cmd.GetWorldDiff().String())

	case *pb.Request_Connect:
		payload = fmt.Sprintf(cmd.GetConnect().String())
	}

	r.logger.Infof("ID[%v] [%T] [%v]", cmd.id[:8], cmd.Command, payload)

}

func (r *RequestHandler) handleMove(ctx context.Context, id model.ClientID, m *pb.Move) {
	//player.SetDest(m.Point.X, m.Point.Y, m.Point.Z)

}

func (r *RequestHandler) handleShoot(ctx context.Context, id model.ClientID, m *pb.Shoot) {
	x, y, z := m.Point.X, m.Point.Y, m.Point.Z
	r.worldManager.RemoveBlock(Point{int(x), int(y), int(z)})
}

func (r *RequestHandler) handleWorldDiffRequest(ctx context.Context, id model.ClientID, m *pb.WorldDiff) {

	r.logger.Info("diff requested")
	//g.network.SendWorldDiff(id)

}
