package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sdeoras/design-patterns/config_params/confgrpc"
	"github.com/sdeoras/design-patterns/config_params/confparams"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	PORT  = "60050"
	CHUNK = 0x100000
)

func main() {
	server := flag.String("host", "", "host")
	flag.Parse()

	if *server == "" {
		log.Fatal("Please enter server IP using --host")
	}

	ctx := context.Background()

	conn, err := grpc.Dial(*server+":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close()

	c := confgrpc.NewSetClient(conn)
	globalStream, err := c.GlobalParams(ctx)
	if err != nil {
		log.Fatal(err)
	}
	localStream, err := c.LocalParams(ctx)
	if err != nil {
		log.Fatal(err)
	}

	M := confparams.NewSerializer()
	M["I"] = int(10)
	M["S"] = "this is a test"

	for key, val := range M {
		fmt.Println(key, val)
	}

	b, err := M.Serialize()
	if err != nil {
		log.Fatal(err)
	}

	n := 0
	for {
		if len(b[n:]) < CHUNK {
			if err := globalStream.Send(&confgrpc.Data{Buffer: b[n:]}); err != nil {
				log.Fatal(err)
			}
			break
		} else {
			if err := globalStream.Send(&confgrpc.Data{Buffer: b[n : n+CHUNK]}); err != nil {
				log.Fatal(err)
			}
			n += CHUNK
		}
	}

	recv, err := globalStream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("written", recv.N, "bytes")

	M = confparams.NewSerializer()
	M["I"] = int32(12)
	M["B"] = []byte{1, 2, 3, 4, 5}

	b, err = M.Serialize()

	n = 0
	for {
		if len(b[n:]) < CHUNK {
			if err := localStream.Send(&confgrpc.Data{Buffer: b[n:], Index: 0}); err != nil {
				log.Fatal(err)
			}
			break
		} else {
			if err := localStream.Send(&confgrpc.Data{Buffer: b[n : n+CHUNK], Index: 0}); err != nil {
				log.Fatal(err)
			}
			n += CHUNK
		}
	}

	recv, err = localStream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("written", recv.N, "bytes")
}
