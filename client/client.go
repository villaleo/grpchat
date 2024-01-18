package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	pb "grpchat/api"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var mode = flag.String("m", "listen",
	"The mode to use when connecting to the server. Modes available are `listen` "+
		"to listen to new messages, `chat` to send messages.",
)

type client struct {
	conn            *grpc.ClientConn
	messengerClient pb.MessengerClient
	ctx             context.Context
	cancel          context.CancelFunc
}

func main() {
	flag.Parse()
	client := newClient()
	defer client.closeConnection()
	switch *mode {
	case "listen":
		client.listenToMessages()
	case "chat":
		client.sendChatMessages()
	case "delete":
		client.deleteChatMessage()
	default:
		log.Fatalf("unknown method: expected 'listen', 'chat', or 'delete'. got %q", *mode)
	}
}

func newClient() client {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return client{
		conn:            conn,
		messengerClient: pb.NewMessengerClient(conn),
		ctx:             ctx,
		cancel:          cancel,
	}
}

func (c client) closeConnection() {
	if err := c.conn.Close(); err != nil {
		log.Fatalf("failed to close connection: %v", err)
	}
	c.cancel()
}

func (c client) listenToMessages() {
	stream, err := c.messengerClient.ListMessages(c.ctx, &pb.Void{})
	if err != nil {
		log.Fatalf("failed to list messages: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive messages from stream: %v", err)
		}
		if msg != nil {
			fmt.Printf("[@%s] %s\n  %s\n",
				msg.GetBody().GetSenderUsername(),
				time.Unix(int64(msg.GetSentAt().GetSeconds()), int64(msg.GetSentAt().GetNanos())).Format(time.DateTime),
				msg.GetBody().GetBody(),
			)
		}
	}
}

func (c client) sendChatMessages() {
	var username string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("enter your username: ")
	scanner.Scan()
	username = scanner.Text()
	var msgBody string
	fmt.Print("enter your message: ")
	scanner.Scan()
	msgBody = scanner.Text()
	body := pb.MessageBody{
		SenderUsername: username,
		Body:           msgBody,
	}
	id, err := c.messengerClient.PublishMessage(c.ctx, &body)
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}
	fmt.Printf("sent message %s\n", id.GetId()[:8])
}

func (c client) deleteChatMessage() {
	var msgId string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("enter message id (first 4 - 8 characters): ")
	scanner.Scan()
	msgId = scanner.Text()
	id := pb.UUID{Id: msgId}
	_, err := c.messengerClient.DeleteMessage(c.ctx, &id)
	if err != nil {
		log.Fatalf("failed to delete message: %v", err)
	}
	fmt.Printf("deleted %s\n", msgId)
}
