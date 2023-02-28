package repository

import (
	"context"

	"github.com/myuoncorp/go-http-server-template/domain/model"
)

//go:generate $HOME/go/bin/mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

type OrganizationRepository interface {
	Get(ctx context.Context) (*model.Organization, error)
	List(ctx context.Context) ([]*model.Organization, error)
}
