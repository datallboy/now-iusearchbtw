package config

import (
	"context"

	"github.com/docker/docker/client"
)

type Config struct {
	Client  *client.Client
	Context context.Context
}

func New() (*Config, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	return &Config{
		Client:  cli,
		Context: ctx,
	}, nil
}
