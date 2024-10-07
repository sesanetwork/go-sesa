package integration

import (
	"fmt"

	"github.com/sesanetwork/go-helios/sesadb"
	"github.com/sesanetwork/go-helios/sesadb/cachedproducer"
	"github.com/sesanetwork/go-helios/sesadb/multidb"
	"github.com/sesanetwork/go-helios/sesadb/skipkeys"

	"github.com/sesanetwork/go-sesa/utils/dbutil/threads"
)

type RoutingConfig struct {
	Table map[string]multidb.Route
}

func (a RoutingConfig) Equal(b RoutingConfig) bool {
	if len(a.Table) != len(b.Table) {
		return false
	}
	for k, v := range a.Table {
		if b.Table[k] != v {
			return false
		}
	}
	return true
}

func MakeMultiProducer(rawProducers map[multidb.TypeName]sesadb.IterableDBProducer, scopedProducers map[multidb.TypeName]sesadb.FullDBProducer, cfg RoutingConfig) (sesadb.FullDBProducer, error) {
	cachedProducers := make(map[multidb.TypeName]sesadb.FullDBProducer)
	var flushID []byte
	var err error
	for typ, producer := range scopedProducers {
		flushID, err = producer.Initialize(rawProducers[typ].Names(), flushID)
		if err != nil {
			return nil, fmt.Errorf("failed to open existing databases: %v. Try to use 'db heal' to recover", err)
		}
		cachedProducers[typ] = cachedproducer.WrapAll(producer)
	}

	p, err := makeMultiProducer(cachedProducers, cfg)
	return threads.CountedFullDBProducer(p), err
}

func MakeDirectMultiProducer(rawProducers map[multidb.TypeName]sesadb.IterableDBProducer, cfg RoutingConfig) (sesadb.FullDBProducer, error) {
	dproducers := map[multidb.TypeName]sesadb.FullDBProducer{}
	for typ, producer := range rawProducers {
		dproducers[typ] = &DummyScopedProducer{producer}
	}
	return MakeMultiProducer(rawProducers, dproducers, cfg)
}

func makeMultiProducer(scopedProducers map[multidb.TypeName]sesadb.FullDBProducer, cfg RoutingConfig) (sesadb.FullDBProducer, error) {
	multi, err := multidb.NewProducer(scopedProducers, cfg.Table, TablesKey)
	if err != nil {
		return nil, fmt.Errorf("failed to construct multidb: %v", err)
	}

	err = multi.Verify()
	if err != nil {
		return nil, fmt.Errorf("incompatible chainstore DB layout: %v. Try to use 'db transform' to recover", err)
	}
	return skipkeys.WrapAllProducer(multi, MetadataPrefix), nil
}
