// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kichikawa/ent/good"
	"github.com/kichikawa/ent/predicate"
)

// GoodUpdate is the builder for updating Good entities.
type GoodUpdate struct {
	config
	hooks    []Hook
	mutation *GoodMutation
}

// Where appends a list predicates to the GoodUpdate builder.
func (gu *GoodUpdate) Where(ps ...predicate.Good) *GoodUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GoodUpdate) SetUpdatedAt(t time.Time) *GoodUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetUserID sets the "user_id" field.
func (gu *GoodUpdate) SetUserID(u uuid.UUID) *GoodUpdate {
	gu.mutation.SetUserID(u)
	return gu
}

// SetTweetID sets the "tweet_id" field.
func (gu *GoodUpdate) SetTweetID(u uuid.UUID) *GoodUpdate {
	gu.mutation.SetTweetID(u)
	return gu
}

// Mutation returns the GoodMutation object of the builder.
func (gu *GoodUpdate) Mutation() *GoodMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GoodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gu.defaults()
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GoodUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GoodUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GoodUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GoodUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := good.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

func (gu *GoodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   good.Table,
			Columns: good.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: good.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: good.FieldUpdatedAt,
		})
	}
	if value, ok := gu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: good.FieldUserID,
		})
	}
	if value, ok := gu.mutation.TweetID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: good.FieldTweetID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{good.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GoodUpdateOne is the builder for updating a single Good entity.
type GoodUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoodMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GoodUpdateOne) SetUpdatedAt(t time.Time) *GoodUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetUserID sets the "user_id" field.
func (guo *GoodUpdateOne) SetUserID(u uuid.UUID) *GoodUpdateOne {
	guo.mutation.SetUserID(u)
	return guo
}

// SetTweetID sets the "tweet_id" field.
func (guo *GoodUpdateOne) SetTweetID(u uuid.UUID) *GoodUpdateOne {
	guo.mutation.SetTweetID(u)
	return guo
}

// Mutation returns the GoodMutation object of the builder.
func (guo *GoodUpdateOne) Mutation() *GoodMutation {
	return guo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GoodUpdateOne) Select(field string, fields ...string) *GoodUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Good entity.
func (guo *GoodUpdateOne) Save(ctx context.Context) (*Good, error) {
	var (
		err  error
		node *Good
	)
	guo.defaults()
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GoodUpdateOne) SaveX(ctx context.Context) *Good {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GoodUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GoodUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GoodUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := good.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

func (guo *GoodUpdateOne) sqlSave(ctx context.Context) (_node *Good, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   good.Table,
			Columns: good.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: good.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Good.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, good.FieldID)
		for _, f := range fields {
			if !good.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != good.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: good.FieldUpdatedAt,
		})
	}
	if value, ok := guo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: good.FieldUserID,
		})
	}
	if value, ok := guo.mutation.TweetID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: good.FieldTweetID,
		})
	}
	_node = &Good{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{good.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
