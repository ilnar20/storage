package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	Id   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &File{
		Id:   newId,
		Name: filename,
		Data: blob,
	}, nil
}

func (f File) String() string {
	return fmt.Sprintf("File: (%v [%v], data: %v)", f.Name, f.Id, f.Data)
}
