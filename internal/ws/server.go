package ws

import (
	"context"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
	pb "qubes/internal/api"
	"qubes/internal/model"
	"qubes/internal/protocol"
	"sync"
)

type GameHandler interface {
	Connect(id model.ClientID)
	Disconnect(id model.ClientID)
	HandleRequest(id model.ClientID, req *pb.Request)
}

type Server struct {
	clients  map[model.ClientID]*Client
	mu       sync.Mutex
	logger   *zap.SugaredLogger
	game     GameHandler
	protocol protocol.Protocol
}

func NewServer(logger *zap.SugaredLogger, protocol protocol.Protocol) *Server {
	return &Server{
		clients:  make(map[model.ClientID]*Client),
		mu:       sync.Mutex{},
		logger:   logger,
		protocol: protocol,
	}
}

func (s *Server) SetGame(g GameHandler) {
	s.game = g
}

func (s *Server) AddClient(client *Client) {
	s.logger.Infof("new client %v", client.id)

	s.mu.Lock()
	s.clients[client.id] = client
	s.mu.Unlock()

	s.game.Connect(client.id)

}
func (s *Server) HandleConn(c *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())

	client := NewClient(c, s, s.logger)
	s.AddClient(client)
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

	s.logger.Infof("got message %T", req.Command)
	return nil
}

func (s *Server) RemoveClient(client *Client) {
	s.logger.Infof("removing client %v", client.id)
	s.game.Disconnect(client.id)

	s.mu.Lock()
	delete(s.clients, client.id)
	s.mu.Unlock()
}

func (s *Server) Send(id model.ClientID, msg []byte) {
	s.clients[id].Send(msg)
}
