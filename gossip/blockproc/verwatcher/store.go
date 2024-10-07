package verwatcher

import (
	"sync/atomic"

	"github.com/sesanetwork/go-vassalo/sesadb"

	"github.com/sesanetwork/go-sesa/logger"
)

// Store is a node persistent storage working over physical key-value database.
type Store struct {
	mainDB sesadb.Store

	cache struct {
		networkVersion atomic.Value
		missedVersion  atomic.Value
	}

	logger.Instance
}

// NewStore creates store over key-value db.
func NewStore(mainDB sesadb.Store) *Store {
	s := &Store{
		mainDB:   mainDB,
		Instance: logger.New("verwatcher-store"),
	}

	return s
}
