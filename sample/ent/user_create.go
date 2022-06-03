// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cybozu-go/scim/sample/ent/email"
	"github.com/cybozu-go/scim/sample/ent/group"
	"github.com/cybozu-go/scim/sample/ent/name"
	"github.com/cybozu-go/scim/sample/ent/user"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetActive sets the "active" field.
func (uc *UserCreate) SetActive(b bool) *UserCreate {
	uc.mutation.SetActive(b)
	return uc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uc *UserCreate) SetNillableActive(b *bool) *UserCreate {
	if b != nil {
		uc.SetActive(*b)
	}
	return uc
}

// SetDisplayName sets the "displayName" field.
func (uc *UserCreate) SetDisplayName(s string) *UserCreate {
	uc.mutation.SetDisplayName(s)
	return uc
}

// SetNillableDisplayName sets the "displayName" field if the given value is not nil.
func (uc *UserCreate) SetNillableDisplayName(s *string) *UserCreate {
	if s != nil {
		uc.SetDisplayName(*s)
	}
	return uc
}

// SetExternalID sets the "externalID" field.
func (uc *UserCreate) SetExternalID(s string) *UserCreate {
	uc.mutation.SetExternalID(s)
	return uc
}

// SetNillableExternalID sets the "externalID" field if the given value is not nil.
func (uc *UserCreate) SetNillableExternalID(s *string) *UserCreate {
	if s != nil {
		uc.SetExternalID(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetPreferredLanguage sets the "preferredLanguage" field.
func (uc *UserCreate) SetPreferredLanguage(s string) *UserCreate {
	uc.mutation.SetPreferredLanguage(s)
	return uc
}

// SetNillablePreferredLanguage sets the "preferredLanguage" field if the given value is not nil.
func (uc *UserCreate) SetNillablePreferredLanguage(s *string) *UserCreate {
	if s != nil {
		uc.SetPreferredLanguage(*s)
	}
	return uc
}

// SetLocale sets the "locale" field.
func (uc *UserCreate) SetLocale(s string) *UserCreate {
	uc.mutation.SetLocale(s)
	return uc
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (uc *UserCreate) SetNillableLocale(s *string) *UserCreate {
	if s != nil {
		uc.SetLocale(*s)
	}
	return uc
}

// SetTimezone sets the "timezone" field.
func (uc *UserCreate) SetTimezone(s string) *UserCreate {
	uc.mutation.SetTimezone(s)
	return uc
}

// SetNillableTimezone sets the "timezone" field if the given value is not nil.
func (uc *UserCreate) SetNillableTimezone(s *string) *UserCreate {
	if s != nil {
		uc.SetTimezone(*s)
	}
	return uc
}

// SetUserType sets the "userType" field.
func (uc *UserCreate) SetUserType(s string) *UserCreate {
	uc.mutation.SetUserType(s)
	return uc
}

// SetNillableUserType sets the "userType" field if the given value is not nil.
func (uc *UserCreate) SetNillableUserType(s *string) *UserCreate {
	if s != nil {
		uc.SetUserType(*s)
	}
	return uc
}

// SetUserName sets the "userName" field.
func (uc *UserCreate) SetUserName(s string) *UserCreate {
	uc.mutation.SetUserName(s)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uc *UserCreate) AddGroupIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddGroupIDs(ids...)
	return uc
}

// AddGroups adds the "groups" edges to the Group entity.
func (uc *UserCreate) AddGroups(g ...*Group) *UserCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroupIDs(ids...)
}

// AddEmailIDs adds the "emails" edge to the Email entity by IDs.
func (uc *UserCreate) AddEmailIDs(ids ...int) *UserCreate {
	uc.mutation.AddEmailIDs(ids...)
	return uc
}

// AddEmails adds the "emails" edges to the Email entity.
func (uc *UserCreate) AddEmails(e ...*Email) *UserCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddEmailIDs(ids...)
}

// AddNameIDs adds the "names" edge to the Name entity by IDs.
func (uc *UserCreate) AddNameIDs(ids ...int) *UserCreate {
	uc.mutation.AddNameIDs(ids...)
	return uc
}

// AddNames adds the "names" edges to the Name entity.
func (uc *UserCreate) AddNames(n ...*Name) *UserCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uc.AddNameIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Active(); !ok {
		v := user.DefaultActive
		uc.mutation.SetActive(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "User.active"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if _, ok := uc.mutation.UserName(); !ok {
		return &ValidationError{Name: "userName", err: errors.New(`ent: missing required field "User.userName"`)}
	}
	if v, ok := uc.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "userName", err: fmt.Errorf(`ent: validator failed for field "User.userName": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := uc.mutation.DisplayName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDisplayName,
		})
		_node.DisplayName = value
	}
	if value, ok := uc.mutation.ExternalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldExternalID,
		})
		_node.ExternalID = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.PreferredLanguage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPreferredLanguage,
		})
		_node.PreferredLanguage = value
	}
	if value, ok := uc.mutation.Locale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLocale,
		})
		_node.Locale = value
	}
	if value, ok := uc.mutation.Timezone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTimezone,
		})
		_node.Timezone = value
	}
	if value, ok := uc.mutation.UserType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserType,
		})
		_node.UserType = value
	}
	if value, ok := uc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserName,
		})
		_node.UserName = value
	}
	if nodes := uc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EmailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EmailsTable,
			Columns: []string{user.EmailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: email.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.NamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NamesTable,
			Columns: []string{user.NamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: name.FieldID,
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

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}