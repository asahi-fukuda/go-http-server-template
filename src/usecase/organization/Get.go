package organization

import (
	"context"

	"github.com/myuoncorp/go-http-server-template/domain/model"
	"github.com/myuoncorp/go-http-server-template/domain/repository"
)

// Get
type Get struct {
	OrganizationRepository repository.OrganizationRepository
}

type (
	// GetInput
	GetInput struct{}
	// GetOutput
	GetOutput struct {
		Organization *model.Organization
	}
)

// Execute
func (u *Get) Execute(ctx context.Context, input *GetInput) (*GetOutput, error) {
	organization, err := u.OrganizationRepository.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &GetOutput{Organization: organization}, nil
}
