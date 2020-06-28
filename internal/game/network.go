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
	Broadcast(msg proto.Message)
	BroadcastExcept(id model.ClientID, msg proto.Message)
}

type NetworkManager struct {
	worldUpdates  chan *WorldUpdate
	playerUpdates chan *PlayerUpdate

	queue map[model.TickID]*NetUpdate

	sender Sender
	logger *zap.SugaredLogger

	response *ResponseBuilder
}

func NewNetworkManager(sender Sender, logger *zap.SugaredLogger) *NetworkManager {
	return &NetworkManager{
		worldUpdates:  make(chan *WorldUpdate),
		playerUpdates: make(chan *PlayerUpdate),
		sender:        sender,
		logger:        logger,
		queue:         make(map[model.TickID]*NetUpdate),
		response:      NewResponseBuilder(),
	}
}

func (n *NetworkManager) Run(ctx context.Context) {
	n.logger.Debug("Network running")
	ticker := time.Tick(time.Millisecond * 100)
	for {
		select {
		case <-ctx.Done():
		//	return

		case update := <-n.worldUpdates:
			if _, ok := n.queue[update.tick]; !ok {
				n.queue[update.tick] = NewNetUpdate()
			}

			n.queue[update.tick].blocks = append(n.queue[update.tick].blocks, update)

		case update := <-n.playerUpdates:
			if _, ok := n.queue[update.tick]; !ok {
				n.queue[update.tick] = NewNetUpdate()
			}
			n.queue[update.tick].players = append(n.queue[update.tick].players, update)

		case <-ticker:
			for tick, upd := range n.queue {
				changes := n.response.NetUpdate(upd, tick)
				n.sender.Broadcast(changes)
			}
			n.queue = make(map[model.TickID]*NetUpdate)
		}
	}
}

func (n *NetworkManager) AddPlayerUpdate(update *PlayerUpdate) {
	//n.sender.Broadcast(n.response.AllPlayers(players, tick))
	n.playerUpdates <- update
}

func (n *NetworkManager) SendWorldDiff(id model.ClientID, world *pb.World) {
	n.sender.Send(id, world)
}

func (n *NetworkManager) AddWorldUpdate(update *WorldUpdate) {
	if update != nil {
		n.worldUpdates <- update
	} else {
		n.logger.Debug("updates are nil wtf")
	}
}

func (n *NetworkManager) SendPlayerConnected(id model.ClientID) {
	n.sender.Broadcast(n.response.PlayerConnected(string(id)))
}
func (n *NetworkManager) SendPlayerDisconnected(id model.ClientID) {
	n.sender.Broadcast(n.response.PlayerDisconnected(string(id)))
}
