// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (gr *Group) Users(ctx context.Context) ([]*User, error) {
	result, err := gr.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryUsers().All(ctx)
	}
	return result, err
}

func (u *User) Pets(ctx context.Context) ([]*Pet, error) {
	result, err := u.Edges.PetsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryPets().All(ctx)
	}
	return result, err
}

func (u *User) Groups(ctx context.Context) ([]*Group, error) {
	result, err := u.Edges.GroupsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryGroups().All(ctx)
	}
	return result, err
}
