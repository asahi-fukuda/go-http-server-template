package repository

import (
	"context"

	"github.com/myuoncorp/go-http-server-template/domain/model"
)

type MessageRepository interface {
	Save(ctx context.Context, message *model.Message) error
	List(ctx context.Context) ([]*model.Message, error)
}
