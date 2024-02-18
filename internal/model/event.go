package model

import (
	fs "github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventAction int32

const (
	UNKNOWN EventAction = iota
	UPLOAD
	DELETE
)

func NewFileEvent(filename string, action EventAction) *fs.FileEvent {
	return &fs.FileEvent{
		Filename:  filename,
		Action:    fs.FileEvent_EventAction(action),
		Timestamp: timestamppb.Now(),
	}
}
