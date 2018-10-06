package main

import (
	"context"
	"log"
	"net"

	myInternalService "github.com/nigeltiany/micro-istio-internal/proto"
	echo "github.com/nigeltiany/micro_istio/proto"
	"google.golang.org/grpc"
)

type Say struct{
	HiddenService myInternalService.InternalServiceClient
}

func ServiceImplementation (conn *grpc.ClientConn) *Say {
	client := myInternalService.NewInternalServiceClient(conn)

	return &Say{
		HiddenService: client,
	}
}

func (s *Say) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	log.Print("Received Say.Hello request")
	hiddenResponse, err := s.HiddenService.Message(ctx, &myInternalService.InternalRequest{
		Message: "Send this to internal service",
	})
	if err != nil {
		log.Printf("unable to reach internal service %v", err)
	} else {
		log.Println(hiddenResponse.Message)
	}

	rsp := &echo.EchoResponse{}
	rsp.Message = req.Message
	return rsp, nil
}

func main ()  {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	conn, err := grpc.Dial("internal.local:80", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	service := ServiceImplementation(conn)

	echo.RegisterEchoServiceServer(grpcServer, service)

	grpcServer.Serve(lis)
}
