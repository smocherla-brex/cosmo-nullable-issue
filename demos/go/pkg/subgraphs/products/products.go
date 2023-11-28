package products

import (
	"github.com/99designs/gqlgen/graphql"

	"github.com/wundergraph/cosmo-federation-demos/demos/go/pkg/subgraphs/products/subgraph"
	"github.com/wundergraph/cosmo-federation-demos/demos/go/pkg/subgraphs/products/subgraph/generated"
)

func NewSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{Resolvers: &subgraph.Resolver{}})
}
