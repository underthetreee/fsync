package model

import (
	"time"

	fs "github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventAction int32

const (
	UNKNOWN EventAction = iota
	UPLOAD
	DELETE
)

type FileEvent struct {
	Filename  string
	Action    EventAction
	Timestamp time.Time
}

func NewFileEvent(filename string, action EventAction) *FileEvent {
	return &FileEvent{
		Filename:  filename,
		Action:    action,
		Timestamp: time.Now(),
	}
}

func ToProtoEvent(event *FileEvent) *fs.FileEvent {
	return &fs.FileEvent{
		Filename:  event.Filename,
		Action:    fs.FileEvent_EventAction(event.Action),
		Timestamp: timestamppb.New(event.Timestamp),
	}
}
