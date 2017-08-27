/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "foolproof.io/kvstore/kvstore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement kvstore.KvStoreServer.
type server struct {
	kvs map[string]string
}

func (s *server) GetValue(ctx context.Context, request *pb.GetValueRequest) (*pb.GetValueReply, error) {
	return &pb.GetValueReply{Value: s.kvs[request.Key]}, nil
}
func (s *server) SetValue(ctx context.Context, request *pb.SetValueRequest) (*pb.SetValueReply, error) {
	s.kvs[request.Key] = request.Value
	return &pb.SetValueReply{}, nil
}

func (s *server) Listen(req *pb.ListenRequest, stream pb.KvStore_ListenServer) error {
	log.Printf("got a stream request, sending %d elements immediately", len(s.kvs))
	for k, v := range s.kvs {
		log.Printf("sending %s -> %s", k, v)
		stream.Send(&pb.ListenReply{
			Key:   k,
			Value: v,
		})
	}
	return nil
}

func main() {
	port := flag.Int("port", 8080, "the port the server will listen to")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	host, _ := os.Hostname()
	pb.RegisterKvStoreServer(s, &server{
		kvs: map[string]string{
			"host": host,
		},
	})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
