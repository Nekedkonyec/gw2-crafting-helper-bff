package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nekedkonyec/gw2-crafting-helper-bff/graph/generated"
	"github.com/nekedkonyec/gw2-crafting-helper-bff/graph/model"
)

func (r *queryResolver) ListItems(ctx context.Context, id string) ([]*model.Item, error) {
	var items []*model.Item
	dummyItem := model.Item{
		ID:   id,
		Name: "Dummy Item",
		Icon: "https://asdasd.com",
	}
	items = append(items, &dummyItem)
	return items, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
