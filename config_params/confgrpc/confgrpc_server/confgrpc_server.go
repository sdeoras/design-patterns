package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"errors"

	"github.com/sdeoras/design-patterns/config_params/confgrpc"
	"github.com/sdeoras/design-patterns/config_params/confparams"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	ADDR = ":60050"
)

var (
	globalConf *confparams.Global
)

type server struct{}

func (s *server) GlobalParams(stream confgrpc.Set_GlobalParamsServer) error {
	buffer := make([]byte, 0, 0)
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		buffer = append(buffer, data.GetBuffer()...)
	}

	S, err := confparams.Decode(buffer)
	if err != nil {
		return err
	}

	if err := globalConf.Set(S); err != nil {
		return err
	}

	if jb, err := json.Marshal(globalConf); err != nil {
		return err
	} else {
		fmt.Println(string(jb))
	}

	ack := new(confgrpc.Ack)
	ack.N = int64(len(buffer))
	return stream.SendAndClose(ack)
}

func (s *server) LocalParams(stream confgrpc.Set_LocalParamsServer) error {
	buffer := make([]byte, 0, 0)
	var index int
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		buffer = append(buffer, data.GetBuffer()...)
		index = int(data.Index)
	}

	S, err := confparams.Decode(buffer)
	if err != nil {
		return err
	}

	if index >= len(globalConf.L) {
		return errors.New("Index out of bounds")
	}

	if err := globalConf.L[index].Set(S); err != nil {
		return err
	}

	if jb, err := json.Marshal(globalConf); err != nil {
		return err
	} else {
		fmt.Println(string(jb))
	}

	ack := new(confgrpc.Ack)
	ack.N = int64(len(buffer))
	return stream.SendAndClose(ack)
}

func main() {
	lis, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	confgrpc.RegisterSetServer(s, &server{})
	reflection.Register(s)

	globalConf, err = confparams.NewGlobal(3)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
