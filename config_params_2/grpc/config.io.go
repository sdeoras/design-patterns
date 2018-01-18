package config

import (
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type MessengerClientIO interface {
	Get(ctx context.Context, in *Empty) (*Global, error)
	Set(ctx context.Context, in *Global) (*Ack, error)
}

type messengerClientIO struct {
	cc io.ReadWriter
}

func NewMessengerClientIO(cc io.ReadWriter) MessengerClientIO {
	return &messengerClientIO{cc}
}

func (c *messengerClientIO) Get(ctx context.Context, in *Empty) (*Global, error) {
	out := new(Global)
	b, err := ioutil.ReadAll(c.cc)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(b, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *messengerClientIO) Set(ctx context.Context, in *Global) (*Ack, error) {
	b, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	if n, err := c.cc.Write(b); err != nil {
		return nil, err
	} else {
		return &Ack{N: int64(n)}, nil
	}
}
