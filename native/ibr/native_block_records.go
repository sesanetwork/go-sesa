package ibr

import (
	"github.com/sesanetwork/go-vassalo/common/bigendian"
	"github.com/sesanetwork/go-vassalo/hash"
	"github.com/sesanetwork/go-vassalo/native/idx"
	"github.com/sesanetwork/go-sesa/core/types"

	"github.com/sesanetwork/go-sesa/native"
)

type LlrBlockVote struct {
	Atropos      hash.Event
	Root         hash.Hash
	TxHash       hash.Hash
	ReceiptsHash hash.Hash
	Time         native.Timestamp
	GasUsed      uint64
}

type LlrFullBlockRecord struct {
	Atropos  hash.Event
	Root     hash.Hash
	Txs      types.Transactions
	Receipts []*types.ReceiptForStorage
	Time     native.Timestamp
	GasUsed  uint64
}

type LlrIdxFullBlockRecord struct {
	LlrFullBlockRecord
	Idx idx.Block
}

func (bv LlrBlockVote) Hash() hash.Hash {
	return hash.Of(bv.Atropos.Bytes(), bv.Root.Bytes(), bv.TxHash.Bytes(), bv.ReceiptsHash.Bytes(), bv.Time.Bytes(), bigendian.Uint64ToBytes(bv.GasUsed))
}

func (br LlrFullBlockRecord) Hash() hash.Hash {
	return LlrBlockVote{
		Atropos:      br.Atropos,
		Root:         br.Root,
		TxHash:       native.CalcTxHash(br.Txs),
		ReceiptsHash: native.CalcReceiptsHash(br.Receipts),
		Time:         br.Time,
		GasUsed:      br.GasUsed,
	}.Hash()
}
