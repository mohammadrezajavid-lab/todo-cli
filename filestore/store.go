package filestore

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Store[T interface{}] struct {
	filePath    string
	permFile    os.FileMode
	entityStore []*T
}

func (s *Store[T]) Save(t *T) {
	s.writeToFile(s.serializedData(t))
}

func (s *Store[T]) Load(*T) []*T {

	dataByte := s.readFile()
	dataStr := strings.Split(string(dataByte), "\n")
	var objects []*T = nil
	object := new(T)

	for _, obj := range dataStr {

		if obj == "" {

			continue
		}

		if err := json.Unmarshal([]byte(obj), object); err != nil {
			panic(err)
		}

		objects = append(objects, object)

		object = new(T)
	}

	for _, obj := range objects {
		fmt.Println(obj)
	}

	return objects
}

// GetObjectsStore Getter method
func (s *Store[T]) GetObjectsStore() []*T {
	return s.entityStore
}

// SetObjectsStore Getter method
func (s *Store[T]) SetObjectsStore(entityStore []*T) {
	s.entityStore = entityStore
}

// NewStore constructor method
func NewStore[T any](filePath string, permFile os.FileMode) *Store[T] {
	return &Store[T]{
		filePath:    filePath,
		permFile:    permFile,
		entityStore: make([]*T, 0),
	}
}

// GetFilePath Getter method
func (s *Store[T]) GetFilePath() string {
	return s.filePath
}

// GetPermFile Getter method
func (s *Store[T]) GetPermFile() os.FileMode {
	return s.permFile
}

// SetFilePath Setter method
func (s *Store[T]) SetFilePath(path string) {
	s.filePath = path
}

// SetPermFile Setter method
func (s *Store[T]) SetPermFile(perm os.FileMode) {
	s.permFile = perm
}

func (s *Store[T]) writeToFile(object []byte) {

	// create object of file
	file, _ := os.OpenFile(s.GetFilePath(), os.O_CREATE|os.O_APPEND|os.O_RDWR, s.GetPermFile())

	// defer file close
	defer func(f *os.File) {
		cErr := f.Close()
		if cErr != nil {
			panic(cErr)
		}
	}(file)

	object = append(object, '\n')

	func(f *os.File) {
		_, wErr := file.Write(object)
		if wErr != nil {
			panic(wErr)
		}
	}(file)
}

func (s *Store[T]) readFile() []byte {

	file, _ := os.OpenFile(s.GetFilePath(), os.O_RDONLY, s.GetPermFile())

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	bs, rErr := io.ReadAll(file)
	if rErr != nil {
		panic(rErr)
	}

	return bs
}

func (s *Store[T]) serializedData(t *T) []byte {

	var data, jErr = json.Marshal(t)

	if jErr != nil {
		fmt.Printf("can't marshal user struct to json %v\n", jErr)
		return nil
	}

	return data
}
