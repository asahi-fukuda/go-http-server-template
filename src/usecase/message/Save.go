package message

import (
	"context"
	"time"

	"github.com/myuoncorp/go-http-server-template/domain/model"
	"github.com/myuoncorp/go-http-server-template/domain/repository"
)

// Save
type Save struct {
	MessageRepository repository.MessageRepository
}

type (
	// SaveInput
	SaveInput struct {
		Name      string
		Message   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	// SaveOutput
	SaveOutput struct{}
)

// Execute
func (u *Save) Execute(ctx context.Context, input *SaveInput) (*SaveOutput, error) {
	message := &model.Message{
		Name:      input.Name,
		Message:   input.Message,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}

	err := u.MessageRepository.Save(ctx, message)
	if err != nil {
		return nil, err
	}
	return &SaveOutput{}, nil
}
