package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/myuoncorp/go-http-server-template/domain/model"
	"github.com/myuoncorp/go-http-server-template/domain/repository"
)

type OrganizationRepository struct {
	db *sqlx.DB
}

func NewOrganizationRepository(db *sqlx.DB) *OrganizationRepository {
	return &OrganizationRepository{db}
}

var _ repository.OrganizationRepository = new(OrganizationRepository)

func (*OrganizationRepository) Get(ctx context.Context) (*model.Organization, error) {
	panic("unimplemented")
}

func (*OrganizationRepository) List(ctx context.Context) ([]*model.Organization, error) {
	panic("unimplemented")
}
