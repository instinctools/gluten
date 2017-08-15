package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "bitbucket.org/instinctools/gluten/cli/proto_service"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func LaunchServer(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProtoServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to the GLuTEN!\n")
}

func LaunchWebServer(address string) {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(address, router))
}
