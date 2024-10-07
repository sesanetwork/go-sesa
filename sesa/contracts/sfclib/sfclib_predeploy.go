package sfclib

import (
	"github.com/sesanetwork/go-sesa/common"
	"github.com/sesanetwork/go-sesa/common/hexutil"
	"github.com/sesanetwork/go-sesa/gossip/contract/sfclib100"
)

// GetContractBin is SFCLib contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
func GetContractBin() []byte {
	return hexutil.MustDecode(sfclib100.ContractBinRuntime)
}

// ContractAddress is the SFCLib contract address
var ContractAddress = common.HexToAddress("0xfc01face00000000000000000000000000000000")
