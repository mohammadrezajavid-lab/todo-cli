package filestore

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type FileStore[T any] struct {
	FilePath string
	PermFile int
}

func (fs *FileStore[T]) Save(t *T) {
	fs.writeToFile(*fs.serializedData(t))
}

func (fs *FileStore[T]) Load(*T) []*T {

	dataByte := fs.readFile()
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

	return objects
}

func (fs *FileStore[T]) writeToFile(object []byte) {

	// create object of file
	file, _ := os.OpenFile(fs.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.FileMode(fs.PermFile))

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

func (fs *FileStore[T]) readFile() []byte {
	file, _ := os.OpenFile(fs.FilePath, os.O_RDONLY, os.FileMode(fs.PermFile))

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

func (fs *FileStore[T]) serializedData(t *T) *[]byte {

	var data, jErr = json.Marshal(t)
	if jErr != nil {
		fmt.Printf("can't marshal user struct to json %v\n", jErr)
		return nil
	}

	return &data
}
