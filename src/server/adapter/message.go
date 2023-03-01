package adapter

import (
	"github.com/myuoncorp/go-http-server-template/domain/model"
	"github.com/myuoncorp/go-http-server-template/oapistub"
)

// *model.Messageを*oapistub.Messageに変換
func Message(m *model.Message) *oapistub.Message {
	return &oapistub.Message{
		Id:        m.ID,
		Name:      m.Name,
		Message:   m.Message,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// []*model.Messageを[]*oapistub.Messageに変換
func MessageList(ms []*model.Message) []*oapistub.Message {
	ds := make([]*oapistub.Message, 0, len(ms))
	for _, m := range ms {
		ds = append(ds, Message(m))
	}
	return ds
}

// []*oapistub.Messageを[]oapistub.Messageに変換
func B(ms []*oapistub.Message) []oapistub.Message {
	ds := make([]oapistub.Message, len(ms))
	for _, m := range ms {
		ds = append(ds, oapistub.Message{
			Id:        m.Id,
			Name:      m.Name,
			Message:   m.Message,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}
	return ds
}
