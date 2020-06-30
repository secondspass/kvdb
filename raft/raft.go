package raft

import (
	"github.com/secondspass/kvdb/db"
)

type DB interface {
	db.DB
}

type raftd struct {
	Name string
	// refers to the local boltdb that stores kv data
	localstore db.DB
}

// TODO: probably want to try using contexts for deferred closing the db and stuff
func Open(name string) (DB, error) {
	r := &raftd{Name: "r"}
	l, err := db.Open(name)
	if err != nil {
		return nil, err
	}
	r.localstore = l
	return r, nil

}

func (r *raftd) Get(key string) string {
	return r.localstore.Get(key)
}

func (r *raftd) Put(key string, val string) error {
	return r.localstore.Put(key, val)
}

func (r *raftd) Close() {
	r.localstore.Close()
}
