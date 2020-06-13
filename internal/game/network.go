package game

import (
	"context"
	"go.uber.org/zap"
	pb "qubes/api"
	"qubes/internal/model"
	"qubes/internal/protocol"
	"qubes/internal/store"
	"time"
)

type Sender interface {
	Send(id model.ClientID, msg []byte)
}

type NetworkManager struct {
	worldUpdates     []*WorldUpdate
	changeRepository store.WorldUpdateRepository

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
		worldUpdates:     make([]*WorldUpdate, 0),
		changeRepository: worldUpdateRepo,
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
			if len(n.worldUpdates) > 0 {
				n.logger.Info("BROACASTING UPDATES")
				changes := n.response.Changes(n.worldUpdates, gameTick)
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

func (n *NetworkManager) SendUpdateRange(player model.ClientID, start model.TickID, end model.TickID) {
	//TODO
	//strings, err := g.changeRepository.GetByRangeRaw(ctx, start, end)
	//if err != nil {
	//	g.logger.Info(err)
	//}
	//for _, s := range strings {
	//	g.logger.Info(s)
	//	g.sender.Send(id, []byte(s))
	//}
}

func (n *NetworkManager) SendWorldUpdate(points []Point, bt pb.BlockType) {
	if points != nil {
		n.logger.Infof("sending updates %v", points)
		wu := &WorldUpdate{
			points:  points,
			newType: bt,
		}
		n.worldUpdates = append(n.worldUpdates, wu)
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
	n.changeRepository.StoreRaw(context.Background(), n.tp.Get(), bytes)
	n.worldUpdates = nil
}

func (n *NetworkManager) BroadcastRaw(bytes []byte) {
	for i := range n.players.All() {
		n.sender.Send(i, bytes)
	}
}

func (n *NetworkManager) Broadcast(resp *pb.Response) {
	//g.logger.Infof("broadcasting to %v", len(g.players))
	bytes, err := n.protocol.Marshal(resp)
	if err != nil {
		n.logger.Error(err)
	}
	for i := range n.players.All() {
		n.sender.Send(i, bytes)
	}
}
