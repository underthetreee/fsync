package model

import (
	"errors"

	"github.com/underthetreee/fsync/pkg/proto"
)

type File struct {
	Filename string
	Content  []byte
}

func ToModel(protoFile *proto.File) (*File, error) {
	if protoFile == nil {
		return nil, errors.New("file not found")
	}

	file := &File{
		Filename: protoFile.Filename,
		Content:  protoFile.Content,
	}
	return file, nil
}

func ToProto(file *File) (*proto.File, error) {
	if file == nil {
		return nil, errors.New("file not found")
	}

	protoFile := &proto.File{
		Filename: file.Filename,
		Content:  file.Content,
	}
	return protoFile, nil
}
