package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/droslean/99designs-gqlgen-2360/graph/generated"
	"github.com/droslean/99designs-gqlgen-2360/graph/model"
)

// QueryPerson is the resolver for the queryPerson field.
func (r *queryResolver) QueryPerson(ctx context.Context, filter *model.PersonFilter, first *int, offset *int) ([]model.Person, error) {
	var ret []model.Person

	j := []byte(`{"id":"1234","name":"Bob","address":"Test address 123"}`)
	fmt.Println(string(j))

	// This will generate error:
	// json: cannot unmarshal object into Go value of type model.Person
	// var p model.Person
	// if err := json.Unmarshal(j, &p); err != nil {
	// 	....
	// }
	// So I need to marshal it to Bob, Sara or George struct.
	// But how I can determine this from the json bytes?
	// It can be either of them.

	return ret, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
