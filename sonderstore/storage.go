package sonderstore

type Config struct {
	Namespace string
}

type Sonderstore interface {
	path(Doc) string
}
