package ws

import (
	"context"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
	"qubes/internal/model"
	"sync"
)

type GameHandler interface {
	OnConnect(c model.ClientID)
	OnDisconnect(client model.ClientID)
	OnMessage(client model.ClientID, msg []byte)
}

type Server struct {
	clients map[model.ClientID]*Client
	mu      sync.Mutex
	logger  *zap.SugaredLogger
	game    GameHandler
}

func NewServer(logger *zap.SugaredLogger) *Server {
	return &Server{
		clients: make(map[model.ClientID]*Client),
		mu:      sync.Mutex{},
		logger:  logger,
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

	s.game.OnConnect(client.id)

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

func (s *Server) HandleMessage(client *Client, data []byte) error {
	s.game.OnMessage(client.id, data)
	return nil
}

func (s *Server) RemoveClient(client *Client) {
	s.logger.Infof("removing client %v", client.id)
	s.game.OnDisconnect(client.id)

	s.mu.Lock()
	delete(s.clients, client.id)
	s.mu.Unlock()
}

func (s *Server) Send(id model.ClientID, msg []byte) {
	s.clients[id].Send(msg)
}
