package grpc

import (
	"context"
	"log"

	"github.com/underthetreee/fsync/internal/model"
	fs "github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EventConsumer interface {
	ConsumeFileEvent(ctx context.Context) (*fs.FileEvent, error)
}

type FileManager interface {
	StoreFile(file *model.File) error
	GetFile(filename string) (*model.File, error)
	DeleteFile(filename string) error
}

type Client struct {
	client   fs.FileSyncServiceClient
	consumer EventConsumer
	mng      FileManager
}

func NewClient(listenAddr string, consumer EventConsumer, manager FileManager) (*Client, error) {
	conn, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := fs.NewFileSyncServiceClient(conn)
	return &Client{
		client:   client,
		consumer: consumer,
		mng:      manager,
	}, nil
}

func (c *Client) UpdateFile(ctx context.Context, filename string) error {
	resp, err := c.client.DownloadFile(ctx, &fs.DownloadFileRequest{Filename: filename})
	if err != nil {
		log.Println(err)
	}
	protoFile := resp.GetFile()
	file := model.ToModel(protoFile)
	if err := c.mng.StoreFile(file); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteFile(ctx context.Context, filename string) error {
	return c.mng.DeleteFile(filename)
}

func (c *Client) SyncLoop(ctx context.Context) error {
	for {
		event, err := c.consumer.ConsumeFileEvent(ctx)
		if err != nil {
			return err
		}

		switch event.Action {
		case fs.FileEvent_UPLOAD:
			if err := c.UpdateFile(ctx, event.GetFilename()); err != nil {
				return err
			}
		case fs.FileEvent_DELETE:
			if err := c.DeleteFile(ctx, event.GetFilename()); err != nil {
				return err
			}
		}

	}
}
