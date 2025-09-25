package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage[T any] interface {
	Save(data *T) error
	GetLastId() int
}

type FileStorage[T any] struct {
	savePath string
	file     *os.File
}

func (s *FileStorage[T]) Save(data T) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("Error on marshaling: %v", err)
		return err
	}

	_, errorWrite := s.file.Write(jsonData)

	if errorWrite != nil {
		fmt.Printf("File write error: %v", errorWrite)
		return errorWrite
	}
	return nil
}

func (s *FileStorage[T]) GetLastId() int {

	// decoder := json.NewDecoder(s.file)
	// decoder
	return 0
}

func NewFileStorage[T any](path string) (Storage[T], *error) {
	storage := &FileStorage[*T]{}

	file, err := getOrCreateStorageFile(path)

	if err != nil {
		return nil, err
	}

	storage.file = file
	storage.savePath = path

	return storage, nil
}

func createStorageFile(path string) (*os.File, *error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Create storage Error")
		return nil, &err
	}

	// defer file.Close()
	return file, nil
}

// func getStorageFile(path string) (*os.File, *error) {
// 	file, err := os.Open(path)

// 	if err != nil {
// 		fmt.Printf("Open storage file Error")
// 		return nil, &err
// 	}

// 	// defer file.Close()
// 	return file, nil
// }

// func fileExist(path string) (bool, *error) {
// 	info, err := os.Stat(path)

// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return false, nil
// 		}
// 		return false, &err
// 	}
// 	return !info.IsDir(), nil
// }

func getOrCreateStorageFile(path string) (*os.File, *error) {

	file, _ := createStorageFile(path)
	return file, nil

}

func (s *FileStorage[T]) Close() error {
	if s.file != nil {
		return s.file.Close()
	}
	return nil
}
