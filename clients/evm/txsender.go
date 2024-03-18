package evm

import (
	"context"
	"errors"
	"log/slog"

	"github.com/aura-nw/btc-bridge-core/clients/evm/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TxSender struct {
	logger *slog.Logger
	txMgr  txmgr.TxManager
	queue  *txmgr.Queue[int]
}

func NewTxSender(ctx context.Context, logger *slog.Logger, txMgr txmgr.TxManager, maxPending uint64) *TxSender {
	queue := txmgr.NewQueue[int](ctx, txMgr, maxPending)
	return &TxSender{
		logger: logger,
		txMgr:  txMgr,
		queue:  queue,
	}
}

func (s *TxSender) From() common.Address {
	return s.txMgr.From()
}

func (s *TxSender) SendAndWait(txs ...txmgr.TxCandidate) ([]*types.Receipt, error) {
	receiptsCh := make(chan txmgr.TxReceipt[int], len(txs))
	for i, tx := range txs {
		s.queue.Send(i, tx, receiptsCh)
	}
	receipts := make([]*types.Receipt, len(txs))
	completed := 0
	var errs []error
	for completed < len(txs) {
		rcpt := <-receiptsCh
		receipts[rcpt.ID] = rcpt.Receipt
		completed++
		if rcpt.Err != nil {
			errs = append(errs, rcpt.Err)
		} else if rcpt.Receipt != nil {
			if rcpt.Receipt.Status != types.ReceiptStatusSuccessful {
				s.logger.Error("Transaction published but reverted", "tx_hash", rcpt.Receipt.TxHash)
			} else {
				s.logger.Debug("Transaction successfully published", "tx_hash", rcpt.Receipt.TxHash)
			}
		}
	}
	return receipts, errors.Join(errs...)
}
