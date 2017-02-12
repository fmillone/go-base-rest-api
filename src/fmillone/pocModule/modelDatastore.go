package pocModule

import (
	"fmillone/config"
	"fmillone/utils"

	pg "gopkg.in/pg.v5"
)

type modelDatastore struct {
}

func (ds modelDatastore) Save(model interface{}) error {
	return ds.getDB().Insert(model)
}

func (ds modelDatastore) getDB() *pg.DB {
	return config.GetEnv().GetDB()
}

func GetDatastore() utils.Datastore {
	return modelDatastore{}
}
