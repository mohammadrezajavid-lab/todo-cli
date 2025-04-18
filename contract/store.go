package contract

type Store[T any] interface {
	Save(t *T)
	Load(t *T) []*T

	GetObjectsStore() []*T
	SetObjectsStore(entityStore []*T)
}
