package netinit

import (
	"github.com/sesanetwork/go-sesa/common"
	"github.com/sesanetwork/go-sesa/common/hexutil"
	"github.com/sesanetwork/go-sesa/gossip/contract/netinit100"
)

// GetContractBin is NetworkInitializer contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from sesa-sfc 76c17565a891e241b09de0a9c1693d0ab3689c17, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode(netinit100.ContractBinRuntime)
}

// ContractAddress is the NetworkInitializer contract address
var ContractAddress = common.HexToAddress("0xd1005eed00000000000000000000000000000000")
