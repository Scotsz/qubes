package ws

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"nhooyr.io/websocket"
	pb "qubes/api"
	"qubes/internal/model"
	"sync"
)

type GameHandler interface {
	OnConnect(c model.ClientID)
	OnDisconnect(client model.ClientID)
	OnMessage(client model.ClientID, req *pb.Request)
}

type Server struct {
	clients  map[model.ClientID]*Client
	mu       sync.Mutex
	logger   *zap.SugaredLogger
	game     GameHandler
	protocol Protocol
}

func NewServer(logger *zap.SugaredLogger) *Server {
	return &Server{
		clients:  make(map[model.ClientID]*Client),
		mu:       sync.Mutex{},
		logger:   logger,
		protocol: &jsonProto{},
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
	req := pb.Request{}
	err := s.protocol.Unmarshal(data, &req)
	if err != nil {
		log.Println(err)
		return err
	}
	s.game.OnMessage(client.id, &req)
	return nil
}

func (s *Server) RemoveClient(client *Client) {
	s.logger.Infof("removing client %v", client.id)
	s.game.OnDisconnect(client.id)

	s.mu.Lock()
	delete(s.clients, client.id)
	s.mu.Unlock()
}

func (s *Server) Send(id model.ClientID, msg proto.Message) {
	data, err := s.protocol.Marshal(msg)
	if err != nil {
		s.logger.Info(err)
	}
	s.clients[id].Send(data)
}

type Protocol interface {
	Marshal(msg proto.Message) ([]byte, error)
	Unmarshal(data []byte, msg proto.Message) error
}
type jsonProto struct {
}

func (j jsonProto) Marshal(msg proto.Message) ([]byte, error) {
	return protojson.Marshal(msg)
}

func (j jsonProto) Unmarshal(data []byte, msg proto.Message) error {
	return protojson.Unmarshal(data, msg)
}
