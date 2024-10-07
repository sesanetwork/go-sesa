package utils

import (
	"github.com/sesanetwork/go-helios/sesadb"
	"github.com/sesanetwork/go-helios/sesadb/table"
)

func NewTableOrSelf(db sesadb.Store, prefix []byte) sesadb.Store {
	if len(prefix) == 0 {
		return db
	}
	return table.New(db, prefix)
}
