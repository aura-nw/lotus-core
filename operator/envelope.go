package operator

import (
	"context"

	"github.com/aura-nw/btc-bridge-core/protos"
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

// VerifyBtcDeposits implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyBtcDeposits(_ context.Context, req *protos.VerifyBtcDepositsRequest) (*protos.VerifyBtcDepositsResponse, error) {
	return &protos.VerifyBtcDepositsResponse{Height: req.Height, Signature: "", Address: e.operator.EvmAddr()}, nil
}

// VerifyBtcWithdrawals implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyBtcWithdrawals(_ context.Context, req *protos.VerifyBtcWithdrawalsRequest) (*protos.VerifyBtcWithdrawalsResponse, error) {
	return &protos.VerifyBtcWithdrawalsResponse{Height: req.Height, Signature: "", Address: e.operator.EvmAddr()}, nil
}

// VerifyInscriptionDeposits implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyInscriptionDeposits(_ context.Context, req *protos.VerifyInscriptionDepositsRequest) (*protos.VerifyInscriptionDepositsResponse, error) {
	return &protos.VerifyInscriptionDepositsResponse{Height: req.Height, Signature: "", Address: e.operator.EvmAddr()}, nil
}

// VerifyInscriptionWithdrawals implements protos.EnvelopeServiceServer.
func (e *envelopeServiceImpl) VerifyInscriptionWithdrawals(_ context.Context, req *protos.VerifyInscriptionWithdrawalsRequest) (*protos.VerifyInscriptionWithdrawalsResponse, error) {
	return &protos.VerifyInscriptionWithdrawalsResponse{Height: req.Height, Signature: "", Address: e.operator.EvmAddr()}, nil
}

var _ protos.EnvelopeServiceServer = &envelopeServiceImpl{}
