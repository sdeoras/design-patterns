package client

import (
	"context"
	"os"

	"github.com/sdeoras/design-patterns/config_params_2/grpc"
)

func Get() (*config.Global, error) {
	file, err := os.Open(CONFIG_FILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	client := config.NewMessengerClientIO(file)
	return client.Get(context.Background(), &config.Empty{})
}

func Set(global *config.Global) (*config.Ack, error) {
	file, err := os.OpenFile(CONFIG_FILE, os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	client := config.NewMessengerClientIO(file)
	return client.Set(context.Background(), global)
}
