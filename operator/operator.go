package operator

import (
	"context"
	"crypto/ecdsa"
	"log/slog"
	"net"

	"github.com/aura-nw/btc-bridge-core/clients/bitcoin"
	"github.com/aura-nw/btc-bridge-core/clients/evm"
	"github.com/aura-nw/btc-bridge-core/config"
	"github.com/aura-nw/btc-bridge-core/protos"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Operator struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	logger *slog.Logger
	config *config.OperatorConfig

	btcClient bitcoin.Client
	evmClient *evm.Client

	// cache must thread-safe
	cache BlockCache

	// evm signer info for signing transaction
	evmSignerInfo evmSignerInfo
}

type evmSignerInfo struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
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

	if err := op.initEvmSignerInfo(); err != nil {
		return nil, err
	}

	return op, nil
}

// initClients inits all clients for operating
func (op *Operator) initClients() error {
	btcClient, err := bitcoin.NewClient(op.logger, op.config.Bitcoin)
	if err != nil {
		return err
	}
	op.btcClient = btcClient

	evmClient, err := evm.NewClient(op.logger, op.config.Evm)
	if err != nil {
		return err
	}
	op.evmClient = evmClient

	return nil
}

// initEvmSignerInfo init signer information of evm network
func (op *Operator) initEvmSignerInfo() error {
	privateKey, err := crypto.HexToECDSA(op.config.EvmPrivateKey)
	if err != nil {
		op.logger.Error("init evm signer failed", "err", err)
		return err
	}
	var info evmSignerInfo
	info.PrivateKey = privateKey

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key to ECDSA")
	}
	info.Address = crypto.PubkeyToAddress(*publicKeyECDSA)

	op.evmSignerInfo = info
	return nil
}

func (op *Operator) runGrpcServer() {
	grpcServer := grpc.NewServer()

	go func() {
		<-op.ctx.Done()
		grpcServer.GracefulStop()
	}()

	protos.RegisterEnvelopeServiceServer(grpcServer, NewEnvelopeServiceImpl(op))
	grpc_health_v1.RegisterHealthServer(grpcServer, &healthCheckImpl{})

	lis, err := net.Listen("tcp", ":"+op.config.Server.GrpcPort)
	if err != nil {
		panic(err)
	}
	op.logger.Info("Start serving grpc server", "addr", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func (op *Operator) initCache() error {
	return nil
}

func (op *Operator) Start() {
	op.logger.Info("Starting operator service")
	go op.runGrpcServer()
}
func (op *Operator) Stop() {
	op.logger.Info("Stopping operator service")
	op.ctxCancel()
}

func (op *Operator) EvmAddr() string {
	return op.evmSignerInfo.Address.Hex()
}
