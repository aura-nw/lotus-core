package operator

import (
	"context"

	"github.com/aura-nw/btc-bridge/protos"
)

type envelopeServiceImpl struct {
	operator *Operator
	protos.UnimplementedEnvelopeServiceServer
}

func NewEnvelopeServiceImpl(operator *Operator) *envelopeServiceImpl {
	return &envelopeServiceImpl{
		operator: operator,
	}
}

// VerifyBitcoinDeposits implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyBitcoinDeposits(_ context.Context, req *protos.VerifyBitcoinDepositsRequest) (*protos.VerifyBitcoinDepositsResponse, error) {
	return &protos.VerifyBitcoinDepositsResponse{Height: req.Height, Signature: "", Address: e.operator.EvmAddr()}, nil
}

// VerifyBitcoinWithdrawals implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyBitcoinWithdrawals(_ context.Context, req *protos.VerifyBitcoinWithdrawalsRequest) (*protos.VerifyBitcoinWithdrawalsResponse, error) {
	return &protos.VerifyBitcoinWithdrawalsResponse{Height: req.Height, Signature: "", Address: e.operator.EvmAddr()}, nil
}

// VerifyTokenDeposits implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyTokenDeposits(_ context.Context, req *protos.VerifyTokenDepositsRequest) (*protos.VerifyTokenDepositsResponse, error) {
	return &protos.VerifyTokenDepositsResponse{Height: req.Height, Signature: "", Address: e.operator.EvmAddr()}, nil
}

// VerifyTokenWithdrawals implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyTokenWithdrawals(_ context.Context, req *protos.VerifyTokenWithdrawalsRequest) (*protos.VerifyTokenWithdrawalsResponse, error) {
	return &protos.VerifyTokenWithdrawalsResponse{Height: req.Height, Signature: "", Address: e.operator.EvmAddr()}, nil
}

var _ protos.EnvelopeServiceServer = &envelopeServiceImpl{}
