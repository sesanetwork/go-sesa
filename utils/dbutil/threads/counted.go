package threads

import (
	"github.com/sesanetwork/go-vassalo/sesadb"

	"github.com/sesanetwork/go-sesa/logger"
)

type (
	countedDbProducer struct {
		sesadb.DBProducer
	}

	countedFullDbProducer struct {
		sesadb.FullDBProducer
	}

	countedStore struct {
		sesadb.Store
	}

	countedIterator struct {
		sesadb.Iterator
		release func(count int)
	}
)

func CountedDBProducer(dbs sesadb.DBProducer) sesadb.DBProducer {
	return &countedDbProducer{dbs}
}

func CountedFullDBProducer(dbs sesadb.FullDBProducer) sesadb.FullDBProducer {
	return &countedFullDbProducer{dbs}
}

func (p *countedDbProducer) OpenDB(name string) (sesadb.Store, error) {
	s, err := p.DBProducer.OpenDB(name)
	return &countedStore{s}, err
}

func (p *countedFullDbProducer) OpenDB(name string) (sesadb.Store, error) {
	s, err := p.FullDBProducer.OpenDB(name)
	return &countedStore{s}, err
}

var notifier = logger.New("threads-pool")

func (s *countedStore) NewIterator(prefix []byte, start []byte) sesadb.Iterator {
	got, release := GlobalPool.Lock(1)
	if got < 1 {
		notifier.Log.Warn("Too many DB iterators")
	}

	return &countedIterator{
		Iterator: s.Store.NewIterator(prefix, start),
		release:  release,
	}
}

func (it *countedIterator) Release() {
	it.Iterator.Release()
	it.release(1)
}
