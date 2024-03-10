package operator

import (
	"context"
	"log/slog"
	"net"

	"github.com/aura-nw/bitcoin-bridge/clients/bitcoin"
	"github.com/aura-nw/bitcoin-bridge/clients/evm"
	"github.com/aura-nw/bitcoin-bridge/config"
	"github.com/aura-nw/bitcoin-bridge/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Operator struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	logger *slog.Logger
	config *config.OperatorConfig

	grpcServer *grpc.Server

	bitcoinClient *bitcoin.Client
	evmClient     *evm.Client
}

func NewOperator(ctx context.Context, logger *slog.Logger, config *config.OperatorConfig) (*Operator, error) {
	ctx, ctxCancel := context.WithCancel(ctx)
	op := &Operator{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		logger:    logger,
		config:    config,
	}
	if err := op.initClients(); err != nil {
		return nil, err
	}

	if err := op.initServer(); err != nil {
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

func (op *Operator) initServer() error {
	grpcServer := grpc.NewServer()
	protos.RegisterEnvelopeServiceServer(grpcServer, &envelopeServiceImpl{})

	grpc_health_v1.RegisterHealthServer(grpcServer, &healthCheckImpl{})

	op.grpcServer = grpcServer

	return nil
}

func (op *Operator) Start() {
	lis, err := net.Listen("tcp", ":"+op.config.Server.GrpcPort)
	if err != nil {
		op.logger.Error("Start: bind listener falied", "err", err)
		panic(err)
	}
	op.grpcServer.Serve(lis)
}
func (op *Operator) Stop() {
	op.grpcServer.GracefulStop()
	op.ctxCancel()
}
