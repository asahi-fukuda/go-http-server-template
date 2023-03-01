package infra

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/myuoncorp/go-http-server-template/domain/model"
)

type MessageRepository struct {
	DB *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (m *MessageRepository) Save(ctx context.Context, msg *model.Message) error {
	message := &model.Message{
		Name:      msg.Name,
		Message:   msg.Message,
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
	}

	_, err := m.DB.NamedExecContext(ctx, `
		INSERT INTO messages(name, message, created_at, updated_at)
		VALUES(:name, :message, :created_at, :updated_at)
		`, &message)

	return err
}

func (m *MessageRepository) List(ctx context.Context) ([]*model.Message, error) {
	var messages []*model.Message
	err := m.DB.SelectContext(ctx, &messages, "SELECT * FROM messages")
	return messages, err
}
