// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bc-labs-lunch-bot/ent/lunch"
	"bc-labs-lunch-bot/ent/lunchparticipant"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LunchParticipantCreate is the builder for creating a LunchParticipant entity.
type LunchParticipantCreate struct {
	config
	mutation *LunchParticipantMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (lpc *LunchParticipantCreate) SetName(s string) *LunchParticipantCreate {
	lpc.mutation.SetName(s)
	return lpc
}

// SetCreatedAt sets the "created_at" field.
func (lpc *LunchParticipantCreate) SetCreatedAt(t time.Time) *LunchParticipantCreate {
	lpc.mutation.SetCreatedAt(t)
	return lpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lpc *LunchParticipantCreate) SetNillableCreatedAt(t *time.Time) *LunchParticipantCreate {
	if t != nil {
		lpc.SetCreatedAt(*t)
	}
	return lpc
}

// SetUpdatedAt sets the "updated_at" field.
func (lpc *LunchParticipantCreate) SetUpdatedAt(t time.Time) *LunchParticipantCreate {
	lpc.mutation.SetUpdatedAt(t)
	return lpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lpc *LunchParticipantCreate) SetNillableUpdatedAt(t *time.Time) *LunchParticipantCreate {
	if t != nil {
		lpc.SetUpdatedAt(*t)
	}
	return lpc
}

// SetPaymentID sets the "payment_id" field.
func (lpc *LunchParticipantCreate) SetPaymentID(i int) *LunchParticipantCreate {
	lpc.mutation.SetPaymentID(i)
	return lpc
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (lpc *LunchParticipantCreate) SetNillablePaymentID(i *int) *LunchParticipantCreate {
	if i != nil {
		lpc.SetPaymentID(*i)
	}
	return lpc
}

// SetPayment sets the "payment" edge to the Lunch entity.
func (lpc *LunchParticipantCreate) SetPayment(l *Lunch) *LunchParticipantCreate {
	return lpc.SetPaymentID(l.ID)
}

// Mutation returns the LunchParticipantMutation object of the builder.
func (lpc *LunchParticipantCreate) Mutation() *LunchParticipantMutation {
	return lpc.mutation
}

// Save creates the LunchParticipant in the database.
func (lpc *LunchParticipantCreate) Save(ctx context.Context) (*LunchParticipant, error) {
	lpc.defaults()
	return withHooks[*LunchParticipant, LunchParticipantMutation](ctx, lpc.sqlSave, lpc.mutation, lpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lpc *LunchParticipantCreate) SaveX(ctx context.Context) *LunchParticipant {
	v, err := lpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lpc *LunchParticipantCreate) Exec(ctx context.Context) error {
	_, err := lpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpc *LunchParticipantCreate) ExecX(ctx context.Context) {
	if err := lpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lpc *LunchParticipantCreate) defaults() {
	if _, ok := lpc.mutation.CreatedAt(); !ok {
		v := lunchparticipant.DefaultCreatedAt()
		lpc.mutation.SetCreatedAt(v)
	}
	if _, ok := lpc.mutation.UpdatedAt(); !ok {
		v := lunchparticipant.DefaultUpdatedAt()
		lpc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lpc *LunchParticipantCreate) check() error {
	if _, ok := lpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "LunchParticipant.name"`)}
	}
	if _, ok := lpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LunchParticipant.created_at"`)}
	}
	if _, ok := lpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LunchParticipant.updated_at"`)}
	}
	return nil
}

func (lpc *LunchParticipantCreate) sqlSave(ctx context.Context) (*LunchParticipant, error) {
	if err := lpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lpc.mutation.id = &_node.ID
	lpc.mutation.done = true
	return _node, nil
}

func (lpc *LunchParticipantCreate) createSpec() (*LunchParticipant, *sqlgraph.CreateSpec) {
	var (
		_node = &LunchParticipant{config: lpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lunchparticipant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchparticipant.FieldID,
			},
		}
	)
	if value, ok := lpc.mutation.Name(); ok {
		_spec.SetField(lunchparticipant.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := lpc.mutation.CreatedAt(); ok {
		_spec.SetField(lunchparticipant.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lpc.mutation.UpdatedAt(); ok {
		_spec.SetField(lunchparticipant.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := lpc.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchparticipant.PaymentTable,
			Columns: []string{lunchparticipant.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunch.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PaymentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LunchParticipantCreateBulk is the builder for creating many LunchParticipant entities in bulk.
type LunchParticipantCreateBulk struct {
	config
	builders []*LunchParticipantCreate
}

// Save creates the LunchParticipant entities in the database.
func (lpcb *LunchParticipantCreateBulk) Save(ctx context.Context) ([]*LunchParticipant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lpcb.builders))
	nodes := make([]*LunchParticipant, len(lpcb.builders))
	mutators := make([]Mutator, len(lpcb.builders))
	for i := range lpcb.builders {
		func(i int, root context.Context) {
			builder := lpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LunchParticipantMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lpcb *LunchParticipantCreateBulk) SaveX(ctx context.Context) []*LunchParticipant {
	v, err := lpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lpcb *LunchParticipantCreateBulk) Exec(ctx context.Context) error {
	_, err := lpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpcb *LunchParticipantCreateBulk) ExecX(ctx context.Context) {
	if err := lpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
