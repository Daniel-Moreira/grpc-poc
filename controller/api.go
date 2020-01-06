package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"grpc-poc/controller/domain/audio"
	"grpc-poc/rpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const grpcPort = ":8081"
const httpPort = ":8080"

func main() {
	log.Println("Starting application")

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	rpc.RegisterRecordServiceServer(server, new(RecordService))

	log.Println("Starting grpc server on port " + grpcPort)

	go server.Serve(listen)

	log.Println("Starting HTTP server on port " + httpPort)
	run()
}

func serveSwagger(mux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir("./third_party/swagger-ui"))
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := rpc.RegisterRecordServiceHandlerFromEndpoint(ctx, gw, "localhost"+grpcPort, opts)

	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	curdir, _ := os.Getwd()
	fmt.Println("cur dir", curdir)
	mux.Handle("/api/", gw)
	serveSwagger(mux)

	return http.ListenAndServe(httpPort, mux)
}

type RecordService struct{}

func (r *RecordService) GetRecord(ctx context.Context, id *rpc.RecordID) (*rpc.Response, error) {
	ok, err := audio.GetRecord(id.Id)

	Response := new(rpc.Response)
	Response.Ok = ok
	Response.Record = nil
	Error := new(rpc.Error)
	Error.Message = err.Error()
	Response.Error = Error

	return Response, err
}

func (r *RecordService) BackupRecord(ctx context.Context, record *rpc.Record) (*rpc.Response, error) {
	fmt.Println("Call")
	fmt.Println(record)
	ok, err := audio.BackupRecord(record)

	Response := new(rpc.Response)
	Response.Ok = ok
	Response.Record = nil
	Error := new(rpc.Error)
	Error.Message = err.Error()
	Response.Error = Error

	return Response, err
}
