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
	"flag"
	"fmt"
	"log"

	pb "foolproof.io/kvstore/kvstore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	usage = "read the source, i'm lazy"
)

func main() {
	address := flag.String("address", "localhost", "the server address")
	port := flag.Int("port", 8080, "the port the server is listening to")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *address, *port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKvStoreClient(conn)

	// Contact the server and print out its response.
	switch flag.Arg(0) {
	case "get":
		key := flag.Arg(1)
		r, err := c.GetValue(context.Background(), &pb.GetValueRequest{Key: key})
		if err != nil {
			log.Fatalf("could not get value: %v", err)
		}
		log.Println(r)
	case "set":
		key := flag.Arg(1)
		value := flag.Arg(2)
		r, err := c.SetValue(context.Background(), &pb.SetValueRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("could not set value: %v", err)
		}
		log.Println(r)
	default:
		log.Fatal(usage)
	}
}
