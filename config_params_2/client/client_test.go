package client

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sdeoras/design-patterns/config_params_2/grpc"
)

func TestSet(t *testing.T) {
	global := new(config.Global)
	global.Local = make([]*config.Local, 3)
	global.A = "this is a"
	global.B = "this is b"
	global.C = true

	for i := range global.Local {
		global.Local[i] = new(config.Local)
	}

	for i := range global.Local {
		global.Local[i].D = "this is d"
		global.Local[i].E = 10
	}

	if ack, err := Set(global); err != nil {
		t.Fatal(err)
	} else {
		t.Log("bytes written:", ack.N)
	}
}

func TestGet(t *testing.T) {
	global, err := Get()
	if err != nil {
		t.Fatal(err)
	}

	jb, err := json.MarshalIndent(global, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(jb))
}
