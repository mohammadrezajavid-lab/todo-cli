package contract

type Store[T any] interface {
	Save(t *T)
	Load(*T) []*T
}
