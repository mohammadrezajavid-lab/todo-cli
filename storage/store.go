package storage

type Store[T any] interface {
	Save(t *T)
}
