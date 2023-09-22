// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"simple-bank/ent/account"
	"simple-bank/ent/entry"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntryCreate is the builder for creating a Entry entity.
type EntryCreate struct {
	config
	mutation *EntryMutation
	hooks    []Hook
}

// SetAccountID sets the "account_id" field.
func (ec *EntryCreate) SetAccountID(u uuid.UUID) *EntryCreate {
	ec.mutation.SetAccountID(u)
	return ec
}

// SetAmount sets the "amount" field.
func (ec *EntryCreate) SetAmount(i int32) *EntryCreate {
	ec.mutation.SetAmount(i)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EntryCreate) SetCreatedAt(t time.Time) *EntryCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EntryCreate) SetNillableCreatedAt(t *time.Time) *EntryCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EntryCreate) SetID(u uuid.UUID) *EntryCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EntryCreate) SetNillableID(u *uuid.UUID) *EntryCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// SetOwnerID sets the "owner" edge to the Account entity by ID.
func (ec *EntryCreate) SetOwnerID(id uuid.UUID) *EntryCreate {
	ec.mutation.SetOwnerID(id)
	return ec
}

// SetOwner sets the "owner" edge to the Account entity.
func (ec *EntryCreate) SetOwner(a *Account) *EntryCreate {
	return ec.SetOwnerID(a.ID)
}

// Mutation returns the EntryMutation object of the builder.
func (ec *EntryCreate) Mutation() *EntryMutation {
	return ec.mutation
}

// Save creates the Entry in the database.
func (ec *EntryCreate) Save(ctx context.Context) (*Entry, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EntryCreate) SaveX(ctx context.Context) *Entry {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EntryCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EntryCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EntryCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := entry.DefaultCreatedAt
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := entry.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EntryCreate) check() error {
	if _, ok := ec.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "Entry.account_id"`)}
	}
	if _, ok := ec.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Entry.amount"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Entry.created_at"`)}
	}
	if _, ok := ec.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Entry.owner"`)}
	}
	return nil
}

func (ec *EntryCreate) sqlSave(ctx context.Context) (*Entry, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EntryCreate) createSpec() (*Entry, *sqlgraph.CreateSpec) {
	var (
		_node = &Entry{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(entry.Table, sqlgraph.NewFieldSpec(entry.FieldID, field.TypeUUID))
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.Amount(); ok {
		_spec.SetField(entry.FieldAmount, field.TypeInt32, value)
		_node.Amount = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(entry.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ec.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.OwnerTable,
			Columns: []string{entry.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EntryCreateBulk is the builder for creating many Entry entities in bulk.
type EntryCreateBulk struct {
	config
	builders []*EntryCreate
}

// Save creates the Entry entities in the database.
func (ecb *EntryCreateBulk) Save(ctx context.Context) ([]*Entry, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Entry, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EntryCreateBulk) SaveX(ctx context.Context) []*Entry {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EntryCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EntryCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
