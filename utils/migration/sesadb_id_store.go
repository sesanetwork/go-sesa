package migration

import (
	"github.com/sesanetwork/go-helios/sesadb"
	"github.com/sesanetwork/go-sesa/log"
)

// sesadbIDStore stores id
type sesadbIDStore struct {
	table sesadb.Store
	key   []byte
}

// NewsesadbIDStore constructor
func NewsesadbIDStore(table sesadb.Store) *sesadbIDStore {
	return &sesadbIDStore{
		table: table,
		key:   []byte("id"),
	}
}

// GetID is a getter
func (p *sesadbIDStore) GetID() string {
	id, err := p.table.Get(p.key)
	if err != nil {
		log.Crit("Failed to get key-value", "err", err)
	}

	if id == nil {
		return ""
	}
	return string(id)
}

// SetID is a setter
func (p *sesadbIDStore) SetID(id string) {
	err := p.table.Put(p.key, []byte(id))
	if err != nil {
		log.Crit("Failed to put key-value", "err", err)
	}
}
