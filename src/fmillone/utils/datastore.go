package utils

type Datastore interface {
	Save(model interface{}) error
}
