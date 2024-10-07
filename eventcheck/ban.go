package eventcheck

import (
	"errors"

	base "github.com/sesanetwork/go-vassalo/eventcheck"

	"github.com/sesanetwork/go-sesa/eventcheck/epochcheck"
	"github.com/sesanetwork/go-sesa/eventcheck/heavycheck"
)

var (
	ErrAlreadyProcessedBVs   = errors.New("BVs is processed already")
	ErrAlreadyProcessedBR    = errors.New("BR is processed already")
	ErrAlreadyProcessedEV    = errors.New("EV is processed already")
	ErrAlreadyProcessedER    = errors.New("ER is processed already")
	ErrUnknownEpochBVs       = heavycheck.ErrUnknownEpochBVs
	ErrUnknownEpochEV        = heavycheck.ErrUnknownEpochEV
	ErrUndecidedBR           = errors.New("BR is unprocessable yet")
	ErrUndecidedER           = errors.New("ER is unprocessable yet")
	ErrAlreadyConnectedEvent = base.ErrAlreadyConnectedEvent
	ErrSpilledEvent          = base.ErrSpilledEvent
	ErrDuplicateEvent        = base.ErrDuplicateEvent
)

func IsBan(err error) bool {
	if err == epochcheck.ErrNotRelevant ||
		err == ErrAlreadyConnectedEvent ||
		err == ErrAlreadyProcessedBVs ||
		err == ErrAlreadyProcessedBR ||
		err == ErrAlreadyProcessedEV ||
		err == ErrAlreadyProcessedER ||
		err == ErrUnknownEpochBVs ||
		err == ErrUndecidedBR ||
		err == ErrUnknownEpochEV ||
		err == ErrUndecidedER ||
		err == ErrSpilledEvent ||
		err == ErrDuplicateEvent {
		return false
	}
	return err != nil
}
