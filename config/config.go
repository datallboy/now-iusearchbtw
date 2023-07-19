package config

import (
	"context"

	"github.com/docker/docker/client"
)

var PUBLIC_PATH = "./public"
var LISTENING_ADDRESS = "0.0.0.0"

type Config struct {
	Client           *client.Client
	Context          context.Context
	PublicPath       string
	ListeningAddress string
}

func New() (*Config, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	return &Config{
		Client:           cli,
		Context:          ctx,
		PublicPath:       PUBLIC_PATH,
		ListeningAddress: LISTENING_ADDRESS,
	}, nil
}
