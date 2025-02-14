package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ilnar20/storage/v2/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	f, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}
	s.files[f.Id] = f
	return f, nil
}

func (s *Storage) GetById(id uuid.UUID) (*file.File, error) {
	if v, ok := s.files[id]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("file not found for id %v", id)
}
