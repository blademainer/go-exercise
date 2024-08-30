package backend

import (
	"github.com/blademainer/go-exercise/demos/generic/factorymode/types"
)

func init() {
	// init
	types.RegisterBackend("file", &FileFactory{})
}

type FileFactory struct {
}

func (f *FileFactory) Create() (interface{}, error) {
	return &fileBackend{}, nil
}

type fileBackend struct {
}
