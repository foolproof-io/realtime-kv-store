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

package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "foolproof.io/kvstore/kvstore"
)

const (
	address    = "localhost:50051"
  usage      = "read the source, i'm lazy"
)

func main() {
  
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKvStoreClient(conn)

	// Contact the server and print out its response.
  if len(os.Args) < 2 {
    log.Fatal(usage)
  }
  switch os.Args[1] {
    case "get":
      if len(os.Args) != 3 {
        log.Fatal(usage)
      }
      key := os.Args[2]
      r, err := c.GetValue(context.Background(), &pb.GetValueRequest{Key: key})
      if err != nil {
        log.Fatalf("could not get value: %v", err)
      }
      log.Println(r)
    case "set":
      if len(os.Args) != 4 {
        log.Fatal(usage)
      }
      key := os.Args[2]
      value := os.Args[3]
      r, err := c.SetValue(context.Background(), &pb.SetValueRequest{Key: key, Value: value})
      if err != nil {
        log.Fatalf("could not set value: %v", err)
      }
      log.Println(r)
    default:
      log.Fatal(usage)
  }
}
