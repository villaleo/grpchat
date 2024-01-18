# gRP Chat

gRP Chat is a real-time messaging service, where many clients can connect to the same server and send messages similar to a public chat forum. I built this service to learn and solidify my understanding of gRPC by using protocol buffers and server steaming.

## Dependencies

### go

This service is built using the Go programming language. To run the service, make sure you have Go installed first. Check out the official [Go installation page](https://go.dev/doc/install) for the instructions.

### protoc

`protoc` is the protocol buffer compiler which I used to generate the Go code for my service. There are more detailed instructions here on the [gRPC website](https://grpc.io/docs/languages/go/quickstart/), but I provided a high-level summary below.

Install the protocol compiler plugins:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Update your path:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Running the project

First, run the `Makefile` to assert the protocol compiler is working. Use the following command to run the Makefile:

```sh
make gen
```

Next, we need to start the server. The server takes no command line arguments. It runs on port `50051`, so make sure it is free before running the command below.

```sh
go run server/server.go
```

Finally, we can start an instance of the client. Let's send a message to the server. We need to specify the `-m` mode argument to let the client know we want to chat.

```sh
go run client/client.go -m chat
```

Running the command above will ask for a username and the message body. After sending the information, you will see an id was returned. This code identifies your message. The server will also indicate that it received your message.

You can listen to all the messages received from the server by running the following command. Notice that we didn't need to specify the mode, since the default option is set to 'listen'.

```sh
go run client/client.go
```

Whoops! We didn't mean to send that last message. Let's delete it using its id which we received earlier. Run the command below to start a deletion. It will ask for the code of the message to delete.

```sh
go run client/client.go -m delete
```

Nice! You're all caught up. gRPC enables us to communicate with different services using a schema. gRPC is fast, compact, and uses HTTP 2. Learn more about gRPC [here](https://grpc.io/docs/what-is-grpc/core-concepts/).
