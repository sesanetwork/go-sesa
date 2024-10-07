package heavycheck

import (
	"github.com/sesanetwork/go-vassalo/native/dag"

	"github.com/sesanetwork/go-sesa/native"
)

type EventsOnly struct {
	*Checker
}

func (c *EventsOnly) Enqueue(e dag.Event, onValidated func(error)) error {
	return c.Checker.EnqueueEvent(e.(native.EventPayloadI), onValidated)
}

type BVsOnly struct {
	*Checker
}

func (c *BVsOnly) Enqueue(bvs native.LlrSignedBlockVotes, onValidated func(error)) error {
	return c.Checker.EnqueueBVs(bvs, onValidated)
}

type EVOnly struct {
	*Checker
}

func (c *EVOnly) Enqueue(ers native.LlrSignedEpochVote, onValidated func(error)) error {
	return c.Checker.EnqueueEV(ers, onValidated)
}
