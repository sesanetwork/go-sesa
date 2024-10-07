package bvprocessor

import (
	"time"

	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/sesanetwork/go-vassalo/native/dag"
	"github.com/sesanetwork/go-vassalo/utils/cachescale"
)

type Config struct {
	BufferLimit dag.Metric

	SemaphoreTimeout time.Duration

	MaxTasks int
}

func DefaultConfig(scale cachescale.Func) Config {
	return Config{
		BufferLimit: dag.Metric{
			Num:  3000,
			Size: scale.U64(15 * opt.MiB),
		},
		SemaphoreTimeout: 10 * time.Second,
		MaxTasks:         512,
	}
}
