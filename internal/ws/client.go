package ws

import (
	"context"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
	"qubes/internal/model"
	"sync"
	"time"
)

type Client struct {
	id     model.ClientID
	conn   *websocket.Conn
	server *Server
	send   chan []byte
	logger *zap.SugaredLogger
}

func NewClient(conn *websocket.Conn, server *Server, logger *zap.SugaredLogger) *Client {
	return &Client{
		id:     model.NewID(),
		conn:   conn,
		server: server,
		logger: logger,
		send:   make(chan []byte, 100),
	}
}
func (c *Client) Send(msg []byte) {
	c.send <- msg
}

func (c *Client) Run(ctx context.Context) {
	go c.writer(ctx)
	c.reader(ctx)
}

func (c *Client) reader(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, data, err := c.conn.Read(ctx)
			if err != nil {
				c.logger.Info(err)
				return
			}
			err = c.server.HandleMessage(c, data)
			if err != nil {
				c.logger.Info(err)
			}
		}
	}
}

func (c *Client) writer(ctx context.Context) {
	pingTicker := time.NewTicker(time.Second * 5)
	defer pingTicker.Stop()

	for {
		select {
		case msg := <-c.send:
			err := c.conn.Write(ctx, websocket.MessageText, msg)
			if err != nil {
				continue
			}
		case <-ctx.Done():
			return
		case <-pingTicker.C:
			err := c.conn.Ping(ctx)
			if err != nil {
				return
			}

		}
	}
}

type ClientStore struct {
	mu      sync.Mutex
	clients map[model.ClientID]*Client
}

func NewClientStore() *ClientStore {
	return &ClientStore{
		mu:      sync.Mutex{},
		clients: make(map[model.ClientID]*Client),
	}
}
func (c *ClientStore) Add(client *Client) {
	c.mu.Lock()
	c.clients[client.id] = client
	c.mu.Unlock()
}
func (c *ClientStore) Remove(client *Client) {
	c.mu.Lock()
	delete(c.clients, client.id)
	c.mu.Unlock()
}
func (c *ClientStore) Get(id model.ClientID) *Client {
	return c.clients[id]
}
