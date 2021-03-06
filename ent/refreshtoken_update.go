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
	"github.com/kichikawa/ent/predicate"
	"github.com/kichikawa/ent/refreshtoken"
)

// RefreshTokenUpdate is the builder for updating RefreshToken entities.
type RefreshTokenUpdate struct {
	config
	hooks    []Hook
	mutation *RefreshTokenMutation
}

// Where appends a list predicates to the RefreshTokenUpdate builder.
func (rtu *RefreshTokenUpdate) Where(ps ...predicate.RefreshToken) *RefreshTokenUpdate {
	rtu.mutation.Where(ps...)
	return rtu
}

// SetUpdatedAt sets the "updated_at" field.
func (rtu *RefreshTokenUpdate) SetUpdatedAt(t time.Time) *RefreshTokenUpdate {
	rtu.mutation.SetUpdatedAt(t)
	return rtu
}

// SetToken sets the "token" field.
func (rtu *RefreshTokenUpdate) SetToken(s string) *RefreshTokenUpdate {
	rtu.mutation.SetToken(s)
	return rtu
}

// SetExpiresAt sets the "expires_at" field.
func (rtu *RefreshTokenUpdate) SetExpiresAt(t time.Time) *RefreshTokenUpdate {
	rtu.mutation.SetExpiresAt(t)
	return rtu
}

// Mutation returns the RefreshTokenMutation object of the builder.
func (rtu *RefreshTokenUpdate) Mutation() *RefreshTokenMutation {
	return rtu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtu *RefreshTokenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	rtu.defaults()
	if len(rtu.hooks) == 0 {
		affected, err = rtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefreshTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtu.mutation = mutation
			affected, err = rtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtu.hooks) - 1; i >= 0; i-- {
			if rtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *RefreshTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *RefreshTokenUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *RefreshTokenUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtu *RefreshTokenUpdate) defaults() {
	if _, ok := rtu.mutation.UpdatedAt(); !ok {
		v := refreshtoken.UpdateDefaultUpdatedAt()
		rtu.mutation.SetUpdatedAt(v)
	}
}

func (rtu *RefreshTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   refreshtoken.Table,
			Columns: refreshtoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: refreshtoken.FieldID,
			},
		},
	}
	if ps := rtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: refreshtoken.FieldUpdatedAt,
		})
	}
	if value, ok := rtu.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refreshtoken.FieldToken,
		})
	}
	if value, ok := rtu.mutation.ExpiresAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: refreshtoken.FieldExpiresAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refreshtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RefreshTokenUpdateOne is the builder for updating a single RefreshToken entity.
type RefreshTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RefreshTokenMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (rtuo *RefreshTokenUpdateOne) SetUpdatedAt(t time.Time) *RefreshTokenUpdateOne {
	rtuo.mutation.SetUpdatedAt(t)
	return rtuo
}

// SetToken sets the "token" field.
func (rtuo *RefreshTokenUpdateOne) SetToken(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetToken(s)
	return rtuo
}

// SetExpiresAt sets the "expires_at" field.
func (rtuo *RefreshTokenUpdateOne) SetExpiresAt(t time.Time) *RefreshTokenUpdateOne {
	rtuo.mutation.SetExpiresAt(t)
	return rtuo
}

// Mutation returns the RefreshTokenMutation object of the builder.
func (rtuo *RefreshTokenUpdateOne) Mutation() *RefreshTokenMutation {
	return rtuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtuo *RefreshTokenUpdateOne) Select(field string, fields ...string) *RefreshTokenUpdateOne {
	rtuo.fields = append([]string{field}, fields...)
	return rtuo
}

// Save executes the query and returns the updated RefreshToken entity.
func (rtuo *RefreshTokenUpdateOne) Save(ctx context.Context) (*RefreshToken, error) {
	var (
		err  error
		node *RefreshToken
	)
	rtuo.defaults()
	if len(rtuo.hooks) == 0 {
		node, err = rtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefreshTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtuo.mutation = mutation
			node, err = rtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtuo.hooks) - 1; i >= 0; i-- {
			if rtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *RefreshTokenUpdateOne) SaveX(ctx context.Context) *RefreshToken {
	node, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtuo *RefreshTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *RefreshTokenUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtuo *RefreshTokenUpdateOne) defaults() {
	if _, ok := rtuo.mutation.UpdatedAt(); !ok {
		v := refreshtoken.UpdateDefaultUpdatedAt()
		rtuo.mutation.SetUpdatedAt(v)
	}
}

func (rtuo *RefreshTokenUpdateOne) sqlSave(ctx context.Context) (_node *RefreshToken, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   refreshtoken.Table,
			Columns: refreshtoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: refreshtoken.FieldID,
			},
		},
	}
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RefreshToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, refreshtoken.FieldID)
		for _, f := range fields {
			if !refreshtoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != refreshtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: refreshtoken.FieldUpdatedAt,
		})
	}
	if value, ok := rtuo.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: refreshtoken.FieldToken,
		})
	}
	if value, ok := rtuo.mutation.ExpiresAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: refreshtoken.FieldExpiresAt,
		})
	}
	_node = &RefreshToken{config: rtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refreshtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
