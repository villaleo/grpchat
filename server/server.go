package main

import (
	"context"
	pb "grpchat/api"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessengerServer

	msgsMu sync.Mutex    // Protects the msgs slice.
	msgs   []*pb.Message // The messages sent to the server.
}

func newServer() *server {
	return &server{msgs: make([]*pb.Message, 0)}
}

// PublishMessage creates a new message from the MessageBody input parameter and
// publishes it to the public chatroom. The UUID of the new Message is sent back
// to the client.
func (s *server) PublishMessage(_ context.Context, req *pb.MessageBody) (*pb.UUID, error) {
	now := time.Now().UTC()
	msg := &pb.Message{
		Id: uuid.New().String(),
		Body: &pb.MessageBody{
			SenderUsername: req.GetSenderUsername(),
			Body:           req.GetBody(),
		},
		SentAt: &pb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	s.msgsMu.Lock()
	s.msgs = append(s.msgs, msg)
	s.msgsMu.Unlock()

	log.Printf("SendMessage: received message %s from @%s",
		msg.GetId()[:8],
		req.GetSenderUsername(),
	)
	return &pb.UUID{Id: msg.GetId()}, nil
}

// DeleteMessage removes a saved message by its ID.
func (s *server) DeleteMessage(_ context.Context, req *pb.UUID) (*pb.Void, error) {
	s.msgsMu.Lock()
	defer s.msgsMu.Unlock()

	for i, msg := range s.msgs {
		if id := msg.GetId(); strings.HasPrefix(id, req.GetId()) {
			if i+1 == len(s.msgs) {
				s.msgs = s.msgs[:i]
			} else {
				s.msgs = append(s.msgs[:i], s.msgs[i+1:]...)
			}
			log.Printf("deleted %s", id[:8])
			return &pb.Void{}, nil
		}
	}

	log.Printf("no message found: nothing deleted")
	return nil, nil
}

// ListMessages fetches all the messages sent to the public chat and returns a
// stream of the messages to the client.
func (s *server) ListMessages(_ *pb.Void, in pb.Messenger_ListMessagesServer) error {
	log.Println("listing messages..")
	count := 0
	for _, msg := range s.msgs {
		if err := in.Send(msg); err != nil {
			log.Printf("failed to get: %v\n", err)
			return err
		}
		count++
	}

	log.Printf("done. got %d messages\n", count)
	return nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterMessengerServer(srv, newServer())
	log.Printf("listening at %v", listener.Addr())
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
