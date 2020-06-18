package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
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

	for i := 1; i < 7; i++ {
		for j := 1; j <= 3; j++ {
			send(ctx, shoot(j*2-1, 1, i), c)
			fmt.Printf("%v:%v:%v\n", j*2+1, 1, i)
			time.Sleep(time.Millisecond * 100)
		}
	}

}

func shoot(x, y, z int) *pb.Request {
	return &pb.Request{
		Tick:    0,
		Command: &pb.Request_Shoot{Shoot: &pb.Shoot{Point: &pb.WorldPoint{X: int32(x), Y: int32(y), Z: int32(z)}}},
	}
}
func send(ctx context.Context, request *pb.Request, c *websocket.Conn) error {
	data, err := protojson.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	return c.Write(ctx, websocket.MessageBinary, data)
}
