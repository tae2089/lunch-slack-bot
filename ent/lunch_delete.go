// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bc-labs-lunch-bot/ent/lunch"
	"bc-labs-lunch-bot/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LunchDelete is the builder for deleting a Lunch entity.
type LunchDelete struct {
	config
	hooks    []Hook
	mutation *LunchMutation
}

// Where appends a list predicates to the LunchDelete builder.
func (ld *LunchDelete) Where(ps ...predicate.Lunch) *LunchDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LunchDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LunchMutation](ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LunchDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LunchDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: lunch.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunch.FieldID,
			},
		},
	}
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LunchDeleteOne is the builder for deleting a single Lunch entity.
type LunchDeleteOne struct {
	ld *LunchDelete
}

// Exec executes the deletion query.
func (ldo *LunchDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lunch.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LunchDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}