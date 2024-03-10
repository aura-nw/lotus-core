package operator

import (
	"context"

	"github.com/aura-nw/bitcoin-bridge/protos"
)

type envelopeServiceImpl struct {
	protos.UnimplementedEnvelopeServiceServer
}

// VerifyBitcoinDeposits implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyBitcoinDeposits(context.Context, *protos.VerifyBitcoinDepositsRequest) (*protos.VerifyBitcoinDepositsResponse, error) {
	panic("unimplemented")
}

// VerifyBitcoinWithdrawals implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyBitcoinWithdrawals(context.Context, *protos.VerifyBitcoinWithdrawalsRequest) (*protos.VerifyBitcoinWithdrawalsResponse, error) {
	panic("unimplemented")
}

// VerifyTokenDeposits implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyTokenDeposits(context.Context, *protos.VerifyTokenDepositsRequest) (*protos.VerifyTokenDepositsResponse, error) {
	panic("unimplemented")
}

// VerifyTokenWithdrawals implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyTokenWithdrawals(context.Context, *protos.VerifyTokenWithdrawalsRequest) (*protos.VerifyTokenWithdrawalsResponse, error) {
	panic("unimplemented")
}

var _ protos.EnvelopeServiceServer = &envelopeServiceImpl{}
