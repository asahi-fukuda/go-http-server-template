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
		Name    string
		Message string
	}
	// SaveOutput
	SaveOutput struct{}
)

// Execute
func (u *Save) Execute(ctx context.Context, input *SaveInput) (*SaveOutput, error) {
	err := u.MessageRepository.Save(ctx, &model.Message{
		Name:      input.Name,
		Message:   input.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return &SaveOutput{}, err
}
