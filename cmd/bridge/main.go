package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aura-nw/btc-bridge/bridge"
	"github.com/aura-nw/btc-bridge/config"
)

const (
	defaultConfigPath = "./bridge.toml"
)

func main() {
	config, err := config.LoadBridgeConfig(defaultConfigPath)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	bridgeControl, err := bridge.NewControl(ctx, slog.Default(), &config)
	if err != nil {
		panic(err)
	}
	fmt.Println("starting bridge control")
	bridgeControl.Start()
	<-ctx.Done()
	fmt.Println("context done, stopping bridge control")
	bridgeControl.Stop()
}
