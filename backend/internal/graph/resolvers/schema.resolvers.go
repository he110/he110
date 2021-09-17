package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"He110/PersonalWebSite/internal/graph"
	"context"
	"fmt"
)

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
