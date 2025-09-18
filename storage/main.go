package storage

import "encoding/json"

type Storage interface {
	Save(data any)
}

type FileStorage struct {
	savePath string
}

func (s *FileStorage) Save(data any) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		
	}
}

func NewFileStorage(path string) Storage {
	storage := &FileStorage{}
	storage.savePath = path

	return storage
}
