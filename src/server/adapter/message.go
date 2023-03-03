package adapter

import (
	"github.com/myuoncorp/go-http-server-template/domain/model"
	"github.com/myuoncorp/go-http-server-template/oapistub"
)

// *model.Messageを*oapistub.Messageに変換
func ToMessage(m *model.Message) *oapistub.Message {
	return &oapistub.Message{
		Id:        m.ID,
		Name:      m.Name,
		Message:   m.Message,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// []*model.Messageを[]*oapistub.Messageに変換
func ToOapiMessage(ms []*model.Message) []oapistub.Message {
	ds := make([]oapistub.Message, len(ms))
	for _, m := range ms {
		ds = append(ds, *ToMessage(m))
	}
	return ds
}
