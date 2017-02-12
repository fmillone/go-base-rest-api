package config

import (
	"sync"

	pg "gopkg.in/pg.v5"
)

type prodEnv struct {
	db *pg.DB
}

var once sync.Once

func (env prodEnv) GetDB() *pg.DB {
	once.Do(func() {
		env.db = pg.Connect(&pg.Options{
			User:     "golang",
			Addr:     "192.168.0.105:5432",
			Database: "golangrest",
		})
	})
	return env.db
}

//GetName returns the environment name duh!!
func (env prodEnv) GetName() string {
	return "Production"
}
