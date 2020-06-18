package game

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	pb "qubes/internal/api"
	"qubes/internal/model"
	"time"
)

type Sender interface {
	Send(id model.ClientID, msg proto.Message)
}

type NetworkManager struct {
	updates chan *WorldUpdate
	queue   []*WorldUpdate

	sender Sender
	logger *zap.SugaredLogger

	response *ResponseBuilder
	players  *PlayerStore
}

func NewNetworkManager(
	sender Sender,
	logger *zap.SugaredLogger,
	players *PlayerStore,

) *NetworkManager {
	return &NetworkManager{
		updates: make(chan *WorldUpdate),
		sender:  sender,
		logger:  logger,
		players: players,
	}
}

func (n *NetworkManager) Run(ctx context.Context) {
	n.logger.Info("Network running")
	ticker := time.Tick(time.Millisecond * 100)
	for {
		select {
		case <-ctx.Done():
		//	return

		case update := <-n.updates:
			//n.Broadcast(n.response.AllPlayers(n.players.All(), gameTick))
			n.queue = append(n.queue, update)

		case <-ticker:
			if len(n.queue) > 0 {
				n.logger.Info("BROACASTING UPDATES")
				changes := n.response.WorldUpdates(n.queue)
				n.Broadcast(changes)
				n.queue = nil
			}

		}
	}
}

func (n *NetworkManager) SendWorldDiff(id model.ClientID, world *pb.World) {
	n.sender.Send(id, world)
}

func (n *NetworkManager) AddWorldUpdate(update *WorldUpdate) {
	if update != nil {
		n.updates <- update
	} else {
		n.logger.Info("updates are nil wtf")
	}
}

func (n *NetworkManager) SendPlayerConnected(id model.ClientID) {
	n.Broadcast(n.response.PlayerConnected(string(id)))
}
func (n *NetworkManager) SendPlayerDisconnected(id model.ClientID) {
	n.Broadcast(n.response.PlayerDisconnected(string(id)))
}

func (n *NetworkManager) Broadcast(resp proto.Message) {
	//g.logger.Infof("broadcasting to %v", len(g.players))

	for i := range n.players.All() {
		n.sender.Send(i, resp)
	}
}
