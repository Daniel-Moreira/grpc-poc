package main

import (
	"context"
	"log"
	"net"

	"grpc-poc/controller/domain/audio"
	"grpc-poc/rpc"

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

	rpc.RecordServiceServer(server, new(RecordService))
	rpc.RecordServiceServer(server, new(RecordService))

	log.Println("Starting grpc server on port " + grpcPort)

	go server.Serve(listen)

	log.Println("Starting HTTP server on port " + httpPort)
	// run()
}

// func serveSwagger(w http.ResponseWriter, r *http.Request) {
// 	//swagger := http.FileServer(http.Dir("./3rdparty/swagger-ui"))
// 	fmt.Println("request", r.URL.Path)
// 	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
// 	p = path.Join("3rdparty/swagger-ui/", p)
// 	fmt.Println("request map ", p)
// 	http.ServeFile(w, r, p)
// }

// func run() error {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	gw := runtime.NewServeMux()
// 	opts := []grpc.DialOption{grpc.WithInsecure()}
// 	err := rpc.RegisterDeviceServiceHandlerFromEndpoint(ctx, gw, "localhost"+grpcPort, opts)

// 	if err != nil {
// 		return err
// 	}

// 	mux := http.NewServeMux()
// 	// mux.HandleFunc("/swagger/", serveSwagger)
// 	curdir, _ := os.Getwd()
// 	fmt.Println("cur dir", curdir)
// 	//swagger := http.FileServer(http.Dir(filepath.Join(curdir, "3rdparty", "swagger-ui")))
// 	//mux.Handle("/swagger/", swagger)
// 	mux.Handle("/api/", gw)

// 	return http.ListenAndServe(httpPort, mux)
// }

type RecordService struct{}

func (r *RecordService) GetRecord(ctx context.Context, id *rpc.RecordID) (*rpc.Response, error) {
  ok, err := audio.GetRecord(id.RecordID)
  
	return &new(rpc.Response{
    ok: ok,
    error: "someMessage"
  }), err
}

func (r *RecordService) BackupRecord(ctx context.Context, record *rpc.Record) (*rpc.Response, error) {
  ok, err := rpc.BackupRecord(record)
  
  return &new(rpc.Response{
    ok: ok,
    error: "someMessage"
  }), err
}
