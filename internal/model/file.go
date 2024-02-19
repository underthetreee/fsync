package model

import fs "github.com/underthetreee/fsync/pkg/proto"

type File struct {
	Filename string
	Content  []byte
}

func ToModel(protoFile *fs.File) *File {
	return &File{
		Filename: protoFile.GetFilename(),
		Content:  protoFile.GetContent(),
	}
}

func ToProto(file *File) *fs.File {
	return &fs.File{
		Filename: file.Filename,
		Content:  file.Content,
	}
}
