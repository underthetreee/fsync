package model

import fs "github.com/underthetreee/fsync/pkg/proto"

type File struct {
	Filename string
	Content  []byte
}

func ToModelFile(protoFile *fs.File) *File {
	return &File{
		Filename: protoFile.GetFilename(),
		Content:  protoFile.GetContent(),
	}
}

func ToProtoFile(file *File) *fs.File {
	return &fs.File{
		Filename: file.Filename,
		Content:  file.Content,
	}
}
