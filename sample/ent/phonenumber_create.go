// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cybozu-go/scim/sample/ent/phonenumber"
	"github.com/cybozu-go/scim/sample/ent/user"
	"github.com/google/uuid"
)

// PhoneNumberCreate is the builder for creating a PhoneNumber entity.
type PhoneNumberCreate struct {
	config
	mutation *PhoneNumberMutation
	hooks    []Hook
}

// SetValue sets the "value" field.
func (pnc *PhoneNumberCreate) SetValue(s string) *PhoneNumberCreate {
	pnc.mutation.SetValue(s)
	return pnc
}

// SetDisplay sets the "display" field.
func (pnc *PhoneNumberCreate) SetDisplay(s string) *PhoneNumberCreate {
	pnc.mutation.SetDisplay(s)
	return pnc
}

// SetType sets the "type" field.
func (pnc *PhoneNumberCreate) SetType(s string) *PhoneNumberCreate {
	pnc.mutation.SetType(s)
	return pnc
}

// SetPrimary sets the "primary" field.
func (pnc *PhoneNumberCreate) SetPrimary(b bool) *PhoneNumberCreate {
	pnc.mutation.SetPrimary(b)
	return pnc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pnc *PhoneNumberCreate) SetUserID(id uuid.UUID) *PhoneNumberCreate {
	pnc.mutation.SetUserID(id)
	return pnc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pnc *PhoneNumberCreate) SetNillableUserID(id *uuid.UUID) *PhoneNumberCreate {
	if id != nil {
		pnc = pnc.SetUserID(*id)
	}
	return pnc
}

// SetUser sets the "user" edge to the User entity.
func (pnc *PhoneNumberCreate) SetUser(u *User) *PhoneNumberCreate {
	return pnc.SetUserID(u.ID)
}

// Mutation returns the PhoneNumberMutation object of the builder.
func (pnc *PhoneNumberCreate) Mutation() *PhoneNumberMutation {
	return pnc.mutation
}

// Save creates the PhoneNumber in the database.
func (pnc *PhoneNumberCreate) Save(ctx context.Context) (*PhoneNumber, error) {
	var (
		err  error
		node *PhoneNumber
	)
	if len(pnc.hooks) == 0 {
		if err = pnc.check(); err != nil {
			return nil, err
		}
		node, err = pnc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhoneNumberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pnc.check(); err != nil {
				return nil, err
			}
			pnc.mutation = mutation
			if node, err = pnc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pnc.hooks) - 1; i >= 0; i-- {
			if pnc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pnc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pnc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pnc *PhoneNumberCreate) SaveX(ctx context.Context) *PhoneNumber {
	v, err := pnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pnc *PhoneNumberCreate) Exec(ctx context.Context) error {
	_, err := pnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pnc *PhoneNumberCreate) ExecX(ctx context.Context) {
	if err := pnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pnc *PhoneNumberCreate) check() error {
	if _, ok := pnc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "PhoneNumber.value"`)}
	}
	if _, ok := pnc.mutation.Display(); !ok {
		return &ValidationError{Name: "display", err: errors.New(`ent: missing required field "PhoneNumber.display"`)}
	}
	if _, ok := pnc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "PhoneNumber.type"`)}
	}
	if _, ok := pnc.mutation.Primary(); !ok {
		return &ValidationError{Name: "primary", err: errors.New(`ent: missing required field "PhoneNumber.primary"`)}
	}
	return nil
}

func (pnc *PhoneNumberCreate) sqlSave(ctx context.Context) (*PhoneNumber, error) {
	_node, _spec := pnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pnc *PhoneNumberCreate) createSpec() (*PhoneNumber, *sqlgraph.CreateSpec) {
	var (
		_node = &PhoneNumber{config: pnc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: phonenumber.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: phonenumber.FieldID,
			},
		}
	)
	if value, ok := pnc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: phonenumber.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := pnc.mutation.Display(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: phonenumber.FieldDisplay,
		})
		_node.Display = value
	}
	if value, ok := pnc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: phonenumber.FieldType,
		})
		_node.Type = value
	}
	if value, ok := pnc.mutation.Primary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: phonenumber.FieldPrimary,
		})
		_node.Primary = value
	}
	if nodes := pnc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phonenumber.UserTable,
			Columns: []string{phonenumber.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.phone_number_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PhoneNumberCreateBulk is the builder for creating many PhoneNumber entities in bulk.
type PhoneNumberCreateBulk struct {
	config
	builders []*PhoneNumberCreate
}

// Save creates the PhoneNumber entities in the database.
func (pncb *PhoneNumberCreateBulk) Save(ctx context.Context) ([]*PhoneNumber, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pncb.builders))
	nodes := make([]*PhoneNumber, len(pncb.builders))
	mutators := make([]Mutator, len(pncb.builders))
	for i := range pncb.builders {
		func(i int, root context.Context) {
			builder := pncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PhoneNumberMutation)
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
					_, err = mutators[i+1].Mutate(root, pncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pncb *PhoneNumberCreateBulk) SaveX(ctx context.Context) []*PhoneNumber {
	v, err := pncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pncb *PhoneNumberCreateBulk) Exec(ctx context.Context) error {
	_, err := pncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pncb *PhoneNumberCreateBulk) ExecX(ctx context.Context) {
	if err := pncb.Exec(ctx); err != nil {
		panic(err)
	}
}