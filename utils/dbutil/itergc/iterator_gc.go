package itergc

import (
	"sync"

	"github.com/sesanetwork/go-vassalo/sesadb"
)

type Snapshot struct {
	sesadb.Snapshot
	nextID uint64
	iters  map[uint64]sesadb.Iterator
	mu     sync.Locker
}

type Iterator struct {
	sesadb.Iterator
	mu    sync.Locker
	id    uint64
	iters map[uint64]sesadb.Iterator
}

// Wrap snapshot to automatically close all pending iterators upon snapshot release
func Wrap(snapshot sesadb.Snapshot, mu sync.Locker) *Snapshot {
	return &Snapshot{
		Snapshot: snapshot,
		iters:    make(map[uint64]sesadb.Iterator),
		mu:       mu,
	}
}

func (s *Snapshot) NewIterator(prefix []byte, start []byte) sesadb.Iterator {
	s.mu.Lock()
	defer s.mu.Unlock()
	it := s.Snapshot.NewIterator(prefix, start)
	id := s.nextID
	s.iters[id] = it
	s.nextID++

	return &Iterator{
		Iterator: it,
		mu:       s.mu,
		id:       id,
		iters:    s.iters,
	}
}

func (s *Iterator) Release() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Iterator.Release()
	delete(s.iters, s.id)
}

func (s *Snapshot) Release() {
	s.mu.Lock()
	defer s.mu.Unlock()
	// release all pending iterators
	for _, it := range s.iters {
		it.Release()
	}
	s.iters = nil
	s.Snapshot.Release()
}
