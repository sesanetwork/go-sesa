package snap2udb

import (
	"github.com/sesanetwork/go-vassalo/sesadb"
	"github.com/sesanetwork/go-vassalo/sesadb/devnulldb"
	"github.com/sesanetwork/go-sesa/log"
)

type Adapter struct {
	sesadb.Snapshot
}

var _ sesadb.Store = (*Adapter)(nil)

func Wrap(v sesadb.Snapshot) *Adapter {
	return &Adapter{v}
}

func (db *Adapter) Put(key []byte, value []byte) error {
	log.Warn("called Put on snapshot")
	return nil
}

func (db *Adapter) Delete(key []byte) error {
	log.Warn("called Delete on snapshot")
	return nil
}

func (db *Adapter) GetSnapshot() (sesadb.Snapshot, error) {
	return db.Snapshot, nil
}

func (db *Adapter) NewBatch() sesadb.Batch {
	log.Warn("called NewBatch on snapshot")
	return devnulldb.New().NewBatch()
}

func (db *Adapter) Compact(start []byte, limit []byte) error {
	return nil
}

func (db *Adapter) Close() error {
	return nil
}

func (db *Adapter) Drop() {}

func (db *Adapter) Stat(property string) (string, error) {
	return "", nil
}
