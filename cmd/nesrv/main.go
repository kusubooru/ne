package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kusubooru/ne/rpc"
	"github.com/kusubooru/ne/rpc/pb"
	"github.com/kusubooru/ne/shimmie2"
	"github.com/kusubooru/shimmie/store"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

var (
	driverName     = flag.String("driver", "mysql", "database driver")
	dataSourceName = flag.String("datasource", "", "database data source")
	port           = flag.Int("port", 10000, "The server port")
	certFile       = flag.String("tlscert", "", "TLS public key in PEM format.  Must be used together with -tlskey")
	keyFile        = flag.String("tlskey", "", "TLS private key in PEM format.  Must be used together with -tlscert")
	secret         = flag.String("secret", "", "secret used to sign JWT tokens")
	// Set after flag parsing based on certFile & keyFile.
	useTLS bool
)

func main() {
	flag.Parse()
	if *secret == "" {
		log.Println("No secret specified, exiting...")
		return
	}

	useTLS = *certFile != "" && *keyFile != ""
	if *dataSourceName == "" {
		log.Println("No database datasource specified, exiting...")
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if useTLS {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	s := store.Open(*driverName, *dataSourceName)
	userService := shimmie2.NewUserService(s)
	pb.RegisterUsersServer(grpcServer, rpc.NewUsersServer(userService))
	pb.RegisterAuthServer(grpcServer, rpc.NewAuthServer(userService, []byte(*secret)))
	grpcServer.Serve(lis)
}
