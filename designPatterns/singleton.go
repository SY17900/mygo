package designpatterns

import "sync"

type DB struct {
}

var db *DB
var once sync.Once

func GetDB() *DB {
	once.Do(func() {
		db = &DB{}
	})
	return db
}
