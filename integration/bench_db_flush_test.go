package integration

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/sesanetwork/go-vassalo/consensus"
	"github.com/sesanetwork/go-vassalo/hash"
	"github.com/sesanetwork/go-vassalo/native/idx"
	"github.com/sesanetwork/go-vassalo/utils/cachescale"
	"github.com/sesanetwork/go-sesa/common"

	"github.com/sesanetwork/go-sesa/gossip"
	"github.com/sesanetwork/go-sesa/integration/makefakegenesis"
	"github.com/sesanetwork/go-sesa/native"
	"github.com/sesanetwork/go-sesa/utils"
	"github.com/sesanetwork/go-sesa/vecmt"
)

func BenchmarkFlushDBs(b *testing.B) {
	dir := tmpDir("flush_bench")
	defer os.RemoveAll(dir)
	genStore := makefakegenesis.FakeGenesisStore(1, utils.Tosesa(1), utils.Tosesa(1))
	g := genStore.Genesis()
	_, _, store, s2, _, closeDBs := MakeEngine(dir, &g, Configs{
		sesa:            gossip.DefaultConfig(cachescale.Identity),
		sesaStore:       gossip.DefaultStoreConfig(cachescale.Identity),
		Hashgraph:      consensus.DefaultConfig(),
		HashgraphStore: consensus.DefaultStoreConfig(cachescale.Identity),
		VectorClock:    vecmt.DefaultConfig(cachescale.Identity),
		DBs:            DefaultDBsConfig(cachescale.Identity.U64, 512),
	})
	defer closeDBs()
	defer store.Close()
	defer s2.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n := idx.Block(0)
		randUint32s := func() []uint32 {
			arr := make([]uint32, 128)
			for i := 0; i < len(arr); i++ {
				arr[i] = uint32(i) ^ (uint32(n) << 16) ^ 0xd0ad884e
			}
			return []uint32{uint32(n), uint32(n) + 1, uint32(n) + 2}
		}
		for !store.IsCommitNeeded() {
			store.SetBlock(n, &native.Block{
				Time:        native.Timestamp(n << 32),
				Atropos:     hash.Event{},
				Events:      hash.Events{},
				Txs:         []common.Hash{},
				InternalTxs: []common.Hash{},
				SkippedTxs:  randUint32s(),
				GasUsed:     uint64(n) << 24,
				Root:        hash.Hash{},
			})
			n++
		}
		b.StartTimer()
		err := store.Commit()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func tmpDir(name string) string {
	dir, err := ioutil.TempDir("", name)
	if err != nil {
		panic(err)
	}
	return dir
}
