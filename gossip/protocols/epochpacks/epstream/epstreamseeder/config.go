package epstreamseeder

import (
	"github.com/sesanetwork/go-helios/gossip/basestream/basestreamseeder"
	"github.com/sesanetwork/go-helios/utils/cachescale"
)

type Config basestreamseeder.Config

func DefaultConfig(scale cachescale.Func) Config {
	return Config{
		SenderThreads:           2,
		MaxSenderTasks:          64,
		MaxPendingResponsesSize: scale.I64(32 * 1024 * 1024),
		MaxResponsePayloadNum:   4096,
		MaxResponsePayloadSize:  8 * 1024 * 1024,
		MaxResponseChunks:       12,
	}
}
