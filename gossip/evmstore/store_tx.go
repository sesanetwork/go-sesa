package evmstore

import (
	"github.com/sesanetwork/go-vassalo/hash"
	"github.com/sesanetwork/go-vassalo/native/idx"
	"github.com/sesanetwork/go-sesa/common"
	"github.com/sesanetwork/go-sesa/core/types"
	"github.com/sesanetwork/go-sesa/log"

	"github.com/sesanetwork/go-sesa/native"
)

// SetTx stores non-event transaction.
func (s *Store) SetTx(txid common.Hash, tx *types.Transaction) {
	s.rlp.Set(s.table.Txs, txid.Bytes(), tx)
}

// GetTx returns stored non-event transaction.
func (s *Store) GetTx(txid common.Hash) *types.Transaction {
	tx, _ := s.rlp.Get(s.table.Txs, txid.Bytes(), &types.Transaction{}).(*types.Transaction)

	return tx
}

func (s *Store) GetBlockTxs(n idx.Block, block native.Block, getEventPayload func(hash.Event) *native.EventPayload) types.Transactions {
	if cached := s.GetCachedEvmBlock(n); cached != nil {
		return cached.Transactions
	}

	transactions := make(types.Transactions, 0, len(block.Txs)+len(block.InternalTxs)+len(block.Events)*10)
	for _, txid := range block.InternalTxs {
		tx := s.GetTx(txid)
		if tx == nil {
			log.Crit("Internal tx not found", "tx", txid.String())
			continue
		}
		transactions = append(transactions, tx)
	}
	for _, txid := range block.Txs {
		tx := s.GetTx(txid)
		if tx == nil {
			log.Crit("Tx not found", "tx", txid.String())
			continue
		}
		transactions = append(transactions, tx)
	}
	for _, id := range block.Events {
		e := getEventPayload(id)
		if e == nil {
			log.Crit("Block event not found", "event", id.String())
			continue
		}
		transactions = append(transactions, e.Txs()...)
	}

	transactions = native.FilterSkippedTxs(transactions, block.SkippedTxs)

	return transactions
}
