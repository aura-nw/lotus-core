package operator

import (
	"context"
	"log/slog"
	"net"

	"github.com/aura-nw/btc-bridge/clients/bitcoin"
	"github.com/aura-nw/btc-bridge/clients/evm"
	"github.com/aura-nw/btc-bridge/config"
	"github.com/aura-nw/btc-bridge/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Operator struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	logger *slog.Logger
	config *config.OperatorConfig

	bitcoinClient *bitcoin.Client
	evmClient     *evm.Client

	// cache must thread-safe
	cache BlockCache
}

func NewOperator(
	ctx context.Context,
	logger *slog.Logger,
	config *config.OperatorConfig,
	cache BlockCache,
) (*Operator, error) {
	ctx, ctxCancel := context.WithCancel(ctx)
	op := &Operator{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		logger:    logger,
		config:    config,
		cache:     cache,
	}

	if err := op.initCache(); err != nil {
		return nil, err
	}

	if err := op.initClients(); err != nil {
		return nil, err
	}

	return op, nil
}

func (op *Operator) initClients() error {
	bitcoinClient, err := bitcoin.NewClient(op.logger, op.config.Bitcoin)
	if err != nil {
		return err
	}
	op.bitcoinClient = bitcoinClient

	evmClient, err := evm.NewClient(op.logger, op.config.Evm)
	if err != nil {
		return err
	}
	op.evmClient = evmClient

	return nil
}

func (op *Operator) runGrpcServer() {
	grpcServer := grpc.NewServer()

	go func() {
		<-op.ctx.Done()
		grpcServer.GracefulStop()
	}()

	protos.RegisterEnvelopeServiceServer(grpcServer, &envelopeServiceImpl{})
	grpc_health_v1.RegisterHealthServer(grpcServer, &healthCheckImpl{})

	lis, err := net.Listen("tcp", ":"+op.config.Server.GrpcPort)
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func (op *Operator) initCache() error {
	return nil
}

func (op *Operator) Start() {
	go op.runGrpcServer()
}
func (op *Operator) Stop() {
	op.ctxCancel()
}
