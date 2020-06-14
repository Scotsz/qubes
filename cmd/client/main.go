package main

import (
	"context"
	"google.golang.org/protobuf/proto"
	"log"
	"nhooyr.io/websocket"
	pb "qubes/internal/api"
	"time"
)

func main() {

	ctx := context.Background()
	c, _, err := websocket.Dial(ctx, "ws://localhost:8000/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	err = send(ctx, makeMove(), c)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(time.Second * 2)
	err = send(ctx, makeShoot(), c)

	time.Sleep(time.Second * 10)
	c.Close(websocket.StatusNormalClosure, "")
}

func send(ctx context.Context, request *pb.Request, c *websocket.Conn) error {
	data, err := proto.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	return c.Write(ctx, websocket.MessageBinary, data)
}

func makeReq() *pb.Request {
	return &pb.Request{
		Tick:    14,
		Command: &pb.Command{},
	}
}
func makeMove() *pb.Request {
	req := makeReq()
	return req
}

func makeShoot() *pb.Request {
	req := makeReq()
	req.Command.Type = &pb.Command_Shoot{Shoot: &pb.Shoot{}}
	return req
}
