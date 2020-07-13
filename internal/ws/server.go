package ws

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"nhooyr.io/websocket"
	pb "qubes/internal/api"
	"qubes/internal/model"
	"sync"
)

type GameHandler interface {
	Connect(id model.ClientID)
	Disconnect(id model.ClientID)
	HandleRequest(id model.ClientID, req *pb.Request)
}

type Server struct {
	mu       sync.Mutex
	clients  *ClientStore
	logger   *zap.SugaredLogger
	game     GameHandler
	protocol Protocol
}

func NewServer(logger *zap.SugaredLogger, protocol Protocol, clients *ClientStore) *Server {
	return &Server{
		clients:  clients,
		mu:       sync.Mutex{},
		logger:   logger,
		protocol: protocol,
	}
}

func (s *Server) SetGame(g GameHandler) {
	s.game = g
}

func (s *Server) RemoveClient(client *Client) {
	s.logger.Infof("removing client %v", client.id)
	s.game.Disconnect(client.id)
	s.clients.Remove(client)
}

func (s *Server) HandleConn(c *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())

	client := NewClient(c, s, s.logger)

	s.logger.Infof("new client %v", client.id)
	s.clients.Add(client)
	s.game.Connect(client.id)

	client.Run(ctx)

	defer func() {
		s.RemoveClient(client)
		c.Close(websocket.StatusInternalError, "")
		cancel()
	}()

}
func (s *Server) HandleMessage(client *Client, msg []byte) error {
	req := &pb.Request{}
	err := s.protocol.Unmarshal(msg, req)

	if err != nil {
		s.logger.Info(err)
		return err
	}
	s.game.HandleRequest(client.id, req)
	return nil
}

type sender struct {
	clients  *ClientStore
	protocol Protocol
	logger   *zap.SugaredLogger
}

func NewSender(logger *zap.SugaredLogger, proto Protocol, store *ClientStore) *sender {
	return &sender{
		clients:  store,
		protocol: proto,
		logger:   logger,
	}
}

func (s sender) Send(id model.ClientID, msg proto.Message) {
	bytes, err := s.protocol.Marshal(msg)

	if err != nil {
		s.logger.Error(err)
		return
	}
	s.clients.Get(id).Send(bytes)
}

func (s sender) Broadcast(msg proto.Message) {
	for i := range s.clients.clients {
		s.Send(i, msg)
	}
}

func (s sender) BroadcastExcept(id model.ClientID, msg proto.Message) {
	for i := range s.clients.clients {
		if i != id {
			s.Send(i, msg)
		}
	}
}
