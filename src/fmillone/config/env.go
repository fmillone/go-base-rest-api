package config

import pg "gopkg.in/pg.v5"

//Env common interface for environment configurations
type Env interface {
	GetDB() *pg.DB
	GetName() string
}

func GetEnv() Env {
	return prodEnv{}
}
