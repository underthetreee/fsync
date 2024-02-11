package model

import (
	"errors"

	"github.com/underthetreee/fsync/pkg/proto"
)

type File struct {
	Filename string
	Content  []byte
}

func ToModel(file *proto.File) (*File, error) {
	if file == nil {
		return nil, errors.New("file not found")
	}

	f := &File{
		Filename: file.Filename,
		Content:  file.Content,
	}
	return f, nil
}
