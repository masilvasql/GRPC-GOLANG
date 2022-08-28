package server

import (
	"context"
	"grpcserver/grpc/pb"

	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCreateUserServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	// log.Println("MENSAGEM RECEBIDA ", req)

	response := &pb.Response{
		Id:     uuid.New().String(),
		Pessoa: req.GetPessoa(),
		Outros: req.GetOutros(),
	}
	return response, nil

}
func (service *Server) mustEmbedUnimplementedCreateUserServer() {}

func UP() {
	grpcServer := grpc.NewServer(

		grpc.MaxRecvMsgSize(1024*1024*20),
		grpc.MaxSendMsgSize(1024*1024*20),
	)

	pb.RegisterCreateUserServer(grpcServer, &Server{})
	port := ":5001"
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("SUBIU ", port)
	grpc_error := grpcServer.Serve(listener)

	if grpc_error != nil {
		log.Fatal(err)
	}

}
