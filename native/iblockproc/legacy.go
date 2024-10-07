package iblockproc

import (
	"github.com/sesanetwork/go-helios/hash"
	"github.com/sesanetwork/go-helios/native/idx"
	"github.com/sesanetwork/go-helios/native/pos"

	"github.com/sesanetwork/go-sesa/native"
	"github.com/sesanetwork/go-sesa/sesa"
)

type ValidatorEpochStateV0 struct {
	GasRefund      uint64
	PrevEpochEvent hash.Event
}

type EpochStateV0 struct {
	Epoch          idx.Epoch
	EpochStart     native.Timestamp
	PrevEpochStart native.Timestamp

	EpochStateRoot hash.Hash

	Validators        *pos.Validators
	ValidatorStates   []ValidatorEpochStateV0
	ValidatorProfiles ValidatorProfiles

	Rules sesa.Rules
}
