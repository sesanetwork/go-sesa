package asyncflushproducer

import "github.com/sesanetwork/go-helios/sesadb"

type store struct {
	sesadb.Store
	CloseFn func() error
}

func (s *store) Close() error {
	return s.CloseFn()
}
