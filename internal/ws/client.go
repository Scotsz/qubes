package ws

import (
	"context"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
	"qubes/internal/model"
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
			{
				err := c.conn.Write(ctx, websocket.MessageText, msg)
				if err != nil {
					c.logger.Error(err)
					continue
				}
			}
		case <-ctx.Done():
			{
				return
			}
		case <-pingTicker.C:
			{
				err := c.conn.Ping(ctx)
				if err != nil {
					c.logger.Info(err.Error())
					return
				}
			}

		}
	}
}
