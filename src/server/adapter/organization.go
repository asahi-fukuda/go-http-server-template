package adapter

import (
	"github.com/myuoncorp/go-http-server-template/domain/model"
	"github.com/myuoncorp/go-http-server-template/oapistub"
)

// Organization adapter
func Organization(m *model.Organization) oapistub.Organization {
	return oapistub.Organization{
		Id:   m.ID,
		Name: m.Name,
	}
}

// OrganizationList adapter
func OrganizationList(ms []*model.Organization) []oapistub.Organization {
	ds := make([]oapistub.Organization, 0, len(ms))
	for _, m := range ms {
		ds = append(ds, Organization(m))
	}
	return ds
}
