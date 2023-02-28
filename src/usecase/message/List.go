package message

import (
	"context"

	"github.com/myuoncorp/go-http-server-template/domain/model"
	"github.com/myuoncorp/go-http-server-template/domain/repository"
)

// List
type List struct {
	MessageRepository repository.MessageRepository
}

type (
	// ListInput
	ListInput struct{}
	// ListOutput
	ListOutput struct {
		Messages []*model.Message
	}
)

// Execute
func (u *List) Execute(ctx context.Context, input *ListInput) (*ListOutput, error) {
	messages, err := u.MessageRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return &ListOutput{Messages: messages}, nil
}
