package logger

import (
	"github.com/sesanetwork/go-sesa/log"
)

type Instance struct {
	Log log.Logger
}

func New(name ...string) Instance {
	if len(name) == 0 {
		return Instance{
			Log: log.New(),
		}
	}
	return Instance{
		Log: log.New("module", name[0]),
	}
}
