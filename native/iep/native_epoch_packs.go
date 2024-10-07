package iep

import (
	"github.com/sesanetwork/go-sesa/native"
	"github.com/sesanetwork/go-sesa/native/ier"
)

type LlrEpochPack struct {
	Votes  []native.LlrSignedEpochVote
	Record ier.LlrIdxFullEpochRecord
}
