package game

import (
	"context"
	"go.uber.org/zap"
	pb "qubes/internal/api"
	"qubes/internal/model"
)

type RequestHandler struct {
	logger       *zap.SugaredLogger
	requestQueue chan *PlayerRequest

	worldManager *stateManager
	network      *NetworkManager

	commandFactory CommandFactory
}

type PlayerRequest struct {
	id model.ClientID
	*pb.Request
}

func (r *RequestHandler) AddRequest(id model.ClientID, req *pb.Request) {
	r.requestQueue <- &PlayerRequest{id, req}
}

func NewRequestHandler(
	logger *zap.SugaredLogger,
	worldManager *stateManager,
	network *NetworkManager,

) *RequestHandler {
	return &RequestHandler{
		logger:       logger,
		requestQueue: make(chan *PlayerRequest, 20),
		network:      network,
		worldManager: worldManager,
	}
}

func (r *RequestHandler) Run(ctx context.Context) {
	r.logger.Debug("RequestHandler running")

	for {
		select {
		case cmd := <-r.requestQueue:
			r.handleRequest(ctx, cmd)
		case <-ctx.Done():
			return
		}
	}
}

func (r *RequestHandler) handleRequest(ctx context.Context, cmd *PlayerRequest) {
	switch cmd.Command.(type) {
	case *pb.Request_Move:
		move := cmd.GetMove()
		r.handleMove(ctx, cmd.id, move)

	case *pb.Request_Shoot:
		r.handleShoot(ctx, cmd.id, cmd.GetShoot())

	case *pb.Request_WorldDiff:
		r.handleWorldDiffRequest(ctx, cmd.id, cmd.GetWorldDiff())

	case *pb.Request_Connect:
	}

	r.logger.Infof("ID[%v] [%T]", cmd.id[:8], cmd.Command)

}

func (r *RequestHandler) handleMove(ctx context.Context, id model.ClientID, m *pb.Move) {
	//player.SetDest(m.Point.X, m.Point.Y, m.Point.Z)

}

func (r *RequestHandler) handleShoot(ctx context.Context, id model.ClientID, m *pb.Shoot) {
	point := model.Point{X: int(m.Point.X), Y: int(m.Point.Y), Z: int(m.Point.Z)}
	r.worldManager.AddCommand(r.commandFactory.RemoveBlock(point))
}

func (r *RequestHandler) handleWorldDiffRequest(ctx context.Context, id model.ClientID, m *pb.WorldDiff) {

	r.logger.Info("diff requested")
	//g.network.SendWorldDiff(id)

}
