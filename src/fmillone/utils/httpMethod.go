package utils

const (
	Get HttpMethod = iota
	Post
)

var httpMethods = []string{
	"GET",
	"POST",
}

type HttpMethod int

func (method HttpMethod) GetCode() string {
	return httpMethods[int(method)]
}
