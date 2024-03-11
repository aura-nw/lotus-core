package main

import (
	"context"
	"log/slog"

	"github.com/aura-nw/btc-bridge/config"
	"github.com/aura-nw/btc-bridge/operator"
)

const (
	defaultConfigPath = "./operator.toml"
)

func main() {
	config, err := config.LoadOperatorConfig(defaultConfigPath)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	op, err := operator.NewOperator(ctx, slog.Default(), &config, nil)
	if err != nil {
		panic(err)
	}
	op.Start()

	<-ctx.Done()
}
