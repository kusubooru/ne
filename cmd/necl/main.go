package main

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/kusubooru/ne/rpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

var (
	serverAddr = flag.String("server", ":10000", "The server address in the format of host:port")
	caFile     = flag.String("tlsca", "", "The file containning the CA root cert file")
	useTLS     bool
)

func main() {
	flag.Parse()
	useTLS = *caFile != ""

	var opts []grpc.DialOption
	if useTLS {
		var sn string
		//if *serverHostOverride != "" {
		//	sn = *serverHostOverride
		//}
		var creds credentials.TransportCredentials
		if *caFile != "" {
			var err error
			creds, err = credentials.NewClientTLSFromFile(*caFile, sn)
			if err != nil {
				grpclog.Fatalf("Failed to create TLS credentials %v", err)
			}
		} else {
			creds = credentials.NewClientTLSFromCert(nil, sn)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUsersClient(conn)

	grpclog.Println("Streaming users")
	stream, err := client.StreamAll(context.Background(), &pb.Page{Limit: 10, Offset: 0})
	if err != nil {
		grpclog.Fatalf("%v.StreamAll(_) = _, %v", client, err)
	}
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Fatalf("%v.StreamAll(_) = _, %v", client, err)
		}
		fmt.Println(user)
	}
}
