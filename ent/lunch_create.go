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

// LunchCreate is the builder for creating a Lunch entity.
type LunchCreate struct {
	config
	mutation *LunchMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lc *LunchCreate) SetCreatedAt(t time.Time) *LunchCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LunchCreate) SetNillableCreatedAt(t *time.Time) *LunchCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LunchCreate) SetUpdatedAt(t time.Time) *LunchCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LunchCreate) SetNillableUpdatedAt(t *time.Time) *LunchCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetPayer sets the "payer" field.
func (lc *LunchCreate) SetPayer(s string) *LunchCreate {
	lc.mutation.SetPayer(s)
	return lc
}

// SetRestaurantName sets the "restaurantName" field.
func (lc *LunchCreate) SetRestaurantName(s string) *LunchCreate {
	lc.mutation.SetRestaurantName(s)
	return lc
}

// SetCafeName sets the "cafeName" field.
func (lc *LunchCreate) SetCafeName(s string) *LunchCreate {
	lc.mutation.SetCafeName(s)
	return lc
}

// SetPaymentTime sets the "paymentTime" field.
func (lc *LunchCreate) SetPaymentTime(t time.Time) *LunchCreate {
	lc.mutation.SetPaymentTime(t)
	return lc
}

// AddParticipantIDs adds the "participant" edge to the LunchParticipant entity by IDs.
func (lc *LunchCreate) AddParticipantIDs(ids ...int) *LunchCreate {
	lc.mutation.AddParticipantIDs(ids...)
	return lc
}

// AddParticipant adds the "participant" edges to the LunchParticipant entity.
func (lc *LunchCreate) AddParticipant(l ...*LunchParticipant) *LunchCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lc.AddParticipantIDs(ids...)
}

// Mutation returns the LunchMutation object of the builder.
func (lc *LunchCreate) Mutation() *LunchMutation {
	return lc.mutation
}

// Save creates the Lunch in the database.
func (lc *LunchCreate) Save(ctx context.Context) (*Lunch, error) {
	lc.defaults()
	return withHooks[*Lunch, LunchMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LunchCreate) SaveX(ctx context.Context) *Lunch {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LunchCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LunchCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LunchCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := lunch.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := lunch.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LunchCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Lunch.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Lunch.updated_at"`)}
	}
	if _, ok := lc.mutation.Payer(); !ok {
		return &ValidationError{Name: "payer", err: errors.New(`ent: missing required field "Lunch.payer"`)}
	}
	if _, ok := lc.mutation.RestaurantName(); !ok {
		return &ValidationError{Name: "restaurantName", err: errors.New(`ent: missing required field "Lunch.restaurantName"`)}
	}
	if _, ok := lc.mutation.CafeName(); !ok {
		return &ValidationError{Name: "cafeName", err: errors.New(`ent: missing required field "Lunch.cafeName"`)}
	}
	if _, ok := lc.mutation.PaymentTime(); !ok {
		return &ValidationError{Name: "paymentTime", err: errors.New(`ent: missing required field "Lunch.paymentTime"`)}
	}
	return nil
}

func (lc *LunchCreate) sqlSave(ctx context.Context) (*Lunch, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LunchCreate) createSpec() (*Lunch, *sqlgraph.CreateSpec) {
	var (
		_node = &Lunch{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lunch.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunch.FieldID,
			},
		}
	)
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(lunch.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(lunch.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.Payer(); ok {
		_spec.SetField(lunch.FieldPayer, field.TypeString, value)
		_node.Payer = value
	}
	if value, ok := lc.mutation.RestaurantName(); ok {
		_spec.SetField(lunch.FieldRestaurantName, field.TypeString, value)
		_node.RestaurantName = value
	}
	if value, ok := lc.mutation.CafeName(); ok {
		_spec.SetField(lunch.FieldCafeName, field.TypeString, value)
		_node.CafeName = value
	}
	if value, ok := lc.mutation.PaymentTime(); ok {
		_spec.SetField(lunch.FieldPaymentTime, field.TypeTime, value)
		_node.PaymentTime = value
	}
	if nodes := lc.mutation.ParticipantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lunch.ParticipantTable,
			Columns: []string{lunch.ParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchparticipant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LunchCreateBulk is the builder for creating many Lunch entities in bulk.
type LunchCreateBulk struct {
	config
	builders []*LunchCreate
}

// Save creates the Lunch entities in the database.
func (lcb *LunchCreateBulk) Save(ctx context.Context) ([]*Lunch, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lunch, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LunchMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LunchCreateBulk) SaveX(ctx context.Context) []*Lunch {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LunchCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LunchCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
