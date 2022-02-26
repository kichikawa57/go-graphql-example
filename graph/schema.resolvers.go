package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kichikawa/ent"
	"github.com/kichikawa/graph/generated"
	"github.com/kichikawa/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) (*ent.Pet, error) {
	users, usersErr := r.Client.User.Query().All(ctx)

	if usersErr != nil {
		fmt.Println("usersErr =====================>", usersErr)

		return nil, usersErr
	}

	fmt.Println("users =====================>", len(users))

	return &ent.Pet{Name: "test"}, nil
}

func (r *queryResolver) GetUserList(ctx context.Context) ([]*ent.User, error) {
	users, usersErr := r.Client.User.Query().All(ctx)

	if usersErr != nil {
		return nil, usersErr
	}

	return users, nil
}

func (r *userResolver) ID(ctx context.Context, obj *ent.User) (model.UserId, error) {
	return model.UserId(obj.ID), nil
}

func (r *userResolver) Status(ctx context.Context, obj *ent.User) (model.UserStatus, error) {
	return model.UserStatus(string(obj.Status)), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
