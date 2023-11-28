package subgraph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	"github.com/wundergraph/cosmo-federation-demos/demos/go/pkg/subgraphs/family/subgraph/generated"
	"github.com/wundergraph/cosmo-federation-demos/demos/go/pkg/subgraphs/family/subgraph/model"
)

// FindEmployees is the resolver for the findEmployees field.
func (r *queryResolver) FindEmployees(ctx context.Context, criteria *model.SearchInput) ([]*model.Employee, error) {
	output := make([]*model.Employee, 0, len(employees))
	for _, employee := range employees {
		if criteria.HasPets != nil && *criteria.HasPets != (len(employee.Details.Pets) > 0) {
			continue
		}
		if criteria.Nationality != nil && *criteria.Nationality != employee.Details.Nationality {
			continue
		}
		nested := criteria.Nested
		if nested != nil {
			if nested.MaritalStatus != nil {
				if employee.Details.MaritalStatus == nil || *nested.MaritalStatus != *employee.Details.MaritalStatus {
					continue
				}
			}
			if nested.HasChildren != nil && *nested.HasChildren != employee.Details.HasChildren {
				continue
			}
		}
		output = append(output, employee)
	}
	return output, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
