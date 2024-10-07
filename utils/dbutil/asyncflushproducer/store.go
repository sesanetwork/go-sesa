package asyncflushproducer

import "github.com/sesanetwork/go-vassalo/sesadb"

type store struct {
	sesadb.Store
	CloseFn func() error
}

func (s *store) Close() error {
	return s.CloseFn()
}
