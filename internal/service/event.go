package service

import (
	"time"

	"github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	UNKNOWN proto.FileEvent_EventType = iota
	UPLOAD
	DOWNLOAD
	DELETE
)

func NewFileEvent(filename string, eventType proto.FileEvent_EventType) *proto.FileEvent {
	return &proto.FileEvent{
		Filename:  filename,
		Type:      eventType,
		Timestamp: timestamppb.New(time.Now()),
	}
}
