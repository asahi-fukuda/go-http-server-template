package organization

import (
	"context"

	"github.com/myuoncorp/go-http-server-template/domain/model"
	"github.com/myuoncorp/go-http-server-template/domain/repository"
)

// List
type List struct {
	OrganizationRepository repository.OrganizationRepository
}

type (
	// ListInput
	ListInput struct{}
	// ListOutput
	ListOutput struct {
		Organizations []*model.Organization
	}
)

// Execute
func (u *List) Execute(ctx context.Context, input *ListInput) (*ListOutput, error) {
	organizations, err := u.OrganizationRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return &ListOutput{Organizations: organizations}, nil
}
