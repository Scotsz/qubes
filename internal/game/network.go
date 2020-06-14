package game

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	pb "qubes/internal/api"
	"qubes/internal/model"
	"qubes/internal/protocol"
	"qubes/internal/store"
	"time"
)

type Sender interface {
	Send(id model.ClientID, msg []byte)
}

type NetworkManager struct {
	currentUpdates   []*WorldUpdate
	updateRepository store.WorldUpdateRepository

	sender Sender
	logger *zap.SugaredLogger

	response *ResponseBuilder
	protocol protocol.Protocol
	players  *PlayerStore
	tp       *TickProvider
}

func NewNetworkManager(
	worldUpdateRepo store.WorldUpdateRepository,
	sender Sender,
	logger *zap.SugaredLogger,
	protocol protocol.Protocol,
	players *PlayerStore,
	tp *TickProvider,

) *NetworkManager {
	return &NetworkManager{
		currentUpdates:   make([]*WorldUpdate, 0),
		updateRepository: worldUpdateRepo,
		sender:           sender,
		logger:           logger,
		response:         nil,
		protocol:         protocol,
		players:          players,
		tp:               tp,
	}
}
func (n *NetworkManager) Run(ctx context.Context) {
	netTicker := time.NewTicker(time.Millisecond * 100)

	for {
		select {
		case <-ctx.Done():
			return
		case <-netTicker.C:
			gameTick := n.tp.Get()

			n.Broadcast(n.response.AllPlayers(n.players.All(), gameTick))
			if len(n.currentUpdates) > 0 {
				n.logger.Info("BROACASTING UPDATES")
				changes := n.response.WorldUpdates(n.currentUpdates, gameTick)
				bytes, err := n.protocol.Marshal(changes)
				if err != nil {
					n.logger.Error(err)
				}
				n.BroadcastRaw(bytes)
				n.moveChangesToHistory(bytes)
			}
		}
	}
}

func (n *NetworkManager) SendUpdateRange(ctx context.Context, id model.ClientID, start model.TickID, end model.TickID) {
	strings, err := n.updateRepository.GetByRangeRaw(ctx, start, end)
	if err != nil {
		n.logger.Info(err)
		return
	}
	for _, s := range strings {
		n.sender.Send(id, []byte(s))
	}
}

func (n *NetworkManager) SendWorldUpdate(points []Point, bt pb.BlockType) {
	if points != nil {
		n.logger.Infof("sending updates %v", points)
		wu := &WorldUpdate{
			points:  points,
			newType: bt,
		}
		n.currentUpdates = append(n.currentUpdates, wu)
	} else {
		n.logger.Info("updates are nil wtf")
	}
}

func (n *NetworkManager) SendPlayerConnected(id model.ClientID, tick model.TickID) {
	n.Broadcast(n.response.PlayerConnected(string(id), tick))
}
func (n *NetworkManager) SendPlayerDisconnected(id model.ClientID, tick model.TickID) {
	n.Broadcast(n.response.PlayerDisconnected(string(id), tick))
}

func (n *NetworkManager) moveChangesToHistory(bytes []byte) {
	n.updateRepository.StoreRaw(context.Background(), n.tp.Get(), bytes)
	n.currentUpdates = nil
}

func (n *NetworkManager) BroadcastRaw(bytes []byte) {
	for i := range n.players.All() {
		n.sender.Send(i, bytes)
	}
}

func (n *NetworkManager) Broadcast(resp proto.Message) {
	//g.logger.Infof("broadcasting to %v", len(g.players))
	bytes, err := n.protocol.Marshal(resp)
	if err != nil {
		n.logger.Error(err)
	}
	for i := range n.players.All() {
		n.sender.Send(i, bytes)
	}
}
