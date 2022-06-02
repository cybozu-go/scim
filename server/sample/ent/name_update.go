// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cybozu-go/scim/server/sample/ent/name"
	"github.com/cybozu-go/scim/server/sample/ent/predicate"
	"github.com/cybozu-go/scim/server/sample/ent/user"
	"github.com/google/uuid"
)

// NameUpdate is the builder for updating Name entities.
type NameUpdate struct {
	config
	hooks    []Hook
	mutation *NameMutation
}

// Where appends a list predicates to the NameUpdate builder.
func (nu *NameUpdate) Where(ps ...predicate.Name) *NameUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetFamilyName sets the "familyName" field.
func (nu *NameUpdate) SetFamilyName(s string) *NameUpdate {
	nu.mutation.SetFamilyName(s)
	return nu
}

// SetNillableFamilyName sets the "familyName" field if the given value is not nil.
func (nu *NameUpdate) SetNillableFamilyName(s *string) *NameUpdate {
	if s != nil {
		nu.SetFamilyName(*s)
	}
	return nu
}

// ClearFamilyName clears the value of the "familyName" field.
func (nu *NameUpdate) ClearFamilyName() *NameUpdate {
	nu.mutation.ClearFamilyName()
	return nu
}

// SetFormatted sets the "formatted" field.
func (nu *NameUpdate) SetFormatted(s string) *NameUpdate {
	nu.mutation.SetFormatted(s)
	return nu
}

// SetNillableFormatted sets the "formatted" field if the given value is not nil.
func (nu *NameUpdate) SetNillableFormatted(s *string) *NameUpdate {
	if s != nil {
		nu.SetFormatted(*s)
	}
	return nu
}

// ClearFormatted clears the value of the "formatted" field.
func (nu *NameUpdate) ClearFormatted() *NameUpdate {
	nu.mutation.ClearFormatted()
	return nu
}

// SetGivenName sets the "givenName" field.
func (nu *NameUpdate) SetGivenName(s string) *NameUpdate {
	nu.mutation.SetGivenName(s)
	return nu
}

// SetNillableGivenName sets the "givenName" field if the given value is not nil.
func (nu *NameUpdate) SetNillableGivenName(s *string) *NameUpdate {
	if s != nil {
		nu.SetGivenName(*s)
	}
	return nu
}

// ClearGivenName clears the value of the "givenName" field.
func (nu *NameUpdate) ClearGivenName() *NameUpdate {
	nu.mutation.ClearGivenName()
	return nu
}

// SetHonorificPrefix sets the "honorificPrefix" field.
func (nu *NameUpdate) SetHonorificPrefix(s string) *NameUpdate {
	nu.mutation.SetHonorificPrefix(s)
	return nu
}

// SetNillableHonorificPrefix sets the "honorificPrefix" field if the given value is not nil.
func (nu *NameUpdate) SetNillableHonorificPrefix(s *string) *NameUpdate {
	if s != nil {
		nu.SetHonorificPrefix(*s)
	}
	return nu
}

// ClearHonorificPrefix clears the value of the "honorificPrefix" field.
func (nu *NameUpdate) ClearHonorificPrefix() *NameUpdate {
	nu.mutation.ClearHonorificPrefix()
	return nu
}

// SetHonorificSuffix sets the "honorificSuffix" field.
func (nu *NameUpdate) SetHonorificSuffix(s string) *NameUpdate {
	nu.mutation.SetHonorificSuffix(s)
	return nu
}

// SetNillableHonorificSuffix sets the "honorificSuffix" field if the given value is not nil.
func (nu *NameUpdate) SetNillableHonorificSuffix(s *string) *NameUpdate {
	if s != nil {
		nu.SetHonorificSuffix(*s)
	}
	return nu
}

// ClearHonorificSuffix clears the value of the "honorificSuffix" field.
func (nu *NameUpdate) ClearHonorificSuffix() *NameUpdate {
	nu.mutation.ClearHonorificSuffix()
	return nu
}

// SetMiddleName sets the "middleName" field.
func (nu *NameUpdate) SetMiddleName(s string) *NameUpdate {
	nu.mutation.SetMiddleName(s)
	return nu
}

// SetNillableMiddleName sets the "middleName" field if the given value is not nil.
func (nu *NameUpdate) SetNillableMiddleName(s *string) *NameUpdate {
	if s != nil {
		nu.SetMiddleName(*s)
	}
	return nu
}

// ClearMiddleName clears the value of the "middleName" field.
func (nu *NameUpdate) ClearMiddleName() *NameUpdate {
	nu.mutation.ClearMiddleName()
	return nu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (nu *NameUpdate) SetUsersID(id uuid.UUID) *NameUpdate {
	nu.mutation.SetUsersID(id)
	return nu
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (nu *NameUpdate) SetNillableUsersID(id *uuid.UUID) *NameUpdate {
	if id != nil {
		nu = nu.SetUsersID(*id)
	}
	return nu
}

// SetUsers sets the "users" edge to the User entity.
func (nu *NameUpdate) SetUsers(u *User) *NameUpdate {
	return nu.SetUsersID(u.ID)
}

// Mutation returns the NameMutation object of the builder.
func (nu *NameUpdate) Mutation() *NameMutation {
	return nu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (nu *NameUpdate) ClearUsers() *NameUpdate {
	nu.mutation.ClearUsers()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NameUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			if nu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NameUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NameUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NameUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   name.Table,
			Columns: name.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: name.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.FamilyName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldFamilyName,
		})
	}
	if nu.mutation.FamilyNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldFamilyName,
		})
	}
	if value, ok := nu.mutation.Formatted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldFormatted,
		})
	}
	if nu.mutation.FormattedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldFormatted,
		})
	}
	if value, ok := nu.mutation.GivenName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldGivenName,
		})
	}
	if nu.mutation.GivenNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldGivenName,
		})
	}
	if value, ok := nu.mutation.HonorificPrefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldHonorificPrefix,
		})
	}
	if nu.mutation.HonorificPrefixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldHonorificPrefix,
		})
	}
	if value, ok := nu.mutation.HonorificSuffix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldHonorificSuffix,
		})
	}
	if nu.mutation.HonorificSuffixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldHonorificSuffix,
		})
	}
	if value, ok := nu.mutation.MiddleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldMiddleName,
		})
	}
	if nu.mutation.MiddleNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldMiddleName,
		})
	}
	if nu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   name.UsersTable,
			Columns: []string{name.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   name.UsersTable,
			Columns: []string{name.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{name.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NameUpdateOne is the builder for updating a single Name entity.
type NameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NameMutation
}

// SetFamilyName sets the "familyName" field.
func (nuo *NameUpdateOne) SetFamilyName(s string) *NameUpdateOne {
	nuo.mutation.SetFamilyName(s)
	return nuo
}

// SetNillableFamilyName sets the "familyName" field if the given value is not nil.
func (nuo *NameUpdateOne) SetNillableFamilyName(s *string) *NameUpdateOne {
	if s != nil {
		nuo.SetFamilyName(*s)
	}
	return nuo
}

// ClearFamilyName clears the value of the "familyName" field.
func (nuo *NameUpdateOne) ClearFamilyName() *NameUpdateOne {
	nuo.mutation.ClearFamilyName()
	return nuo
}

// SetFormatted sets the "formatted" field.
func (nuo *NameUpdateOne) SetFormatted(s string) *NameUpdateOne {
	nuo.mutation.SetFormatted(s)
	return nuo
}

// SetNillableFormatted sets the "formatted" field if the given value is not nil.
func (nuo *NameUpdateOne) SetNillableFormatted(s *string) *NameUpdateOne {
	if s != nil {
		nuo.SetFormatted(*s)
	}
	return nuo
}

// ClearFormatted clears the value of the "formatted" field.
func (nuo *NameUpdateOne) ClearFormatted() *NameUpdateOne {
	nuo.mutation.ClearFormatted()
	return nuo
}

// SetGivenName sets the "givenName" field.
func (nuo *NameUpdateOne) SetGivenName(s string) *NameUpdateOne {
	nuo.mutation.SetGivenName(s)
	return nuo
}

// SetNillableGivenName sets the "givenName" field if the given value is not nil.
func (nuo *NameUpdateOne) SetNillableGivenName(s *string) *NameUpdateOne {
	if s != nil {
		nuo.SetGivenName(*s)
	}
	return nuo
}

// ClearGivenName clears the value of the "givenName" field.
func (nuo *NameUpdateOne) ClearGivenName() *NameUpdateOne {
	nuo.mutation.ClearGivenName()
	return nuo
}

// SetHonorificPrefix sets the "honorificPrefix" field.
func (nuo *NameUpdateOne) SetHonorificPrefix(s string) *NameUpdateOne {
	nuo.mutation.SetHonorificPrefix(s)
	return nuo
}

// SetNillableHonorificPrefix sets the "honorificPrefix" field if the given value is not nil.
func (nuo *NameUpdateOne) SetNillableHonorificPrefix(s *string) *NameUpdateOne {
	if s != nil {
		nuo.SetHonorificPrefix(*s)
	}
	return nuo
}

// ClearHonorificPrefix clears the value of the "honorificPrefix" field.
func (nuo *NameUpdateOne) ClearHonorificPrefix() *NameUpdateOne {
	nuo.mutation.ClearHonorificPrefix()
	return nuo
}

// SetHonorificSuffix sets the "honorificSuffix" field.
func (nuo *NameUpdateOne) SetHonorificSuffix(s string) *NameUpdateOne {
	nuo.mutation.SetHonorificSuffix(s)
	return nuo
}

// SetNillableHonorificSuffix sets the "honorificSuffix" field if the given value is not nil.
func (nuo *NameUpdateOne) SetNillableHonorificSuffix(s *string) *NameUpdateOne {
	if s != nil {
		nuo.SetHonorificSuffix(*s)
	}
	return nuo
}

// ClearHonorificSuffix clears the value of the "honorificSuffix" field.
func (nuo *NameUpdateOne) ClearHonorificSuffix() *NameUpdateOne {
	nuo.mutation.ClearHonorificSuffix()
	return nuo
}

// SetMiddleName sets the "middleName" field.
func (nuo *NameUpdateOne) SetMiddleName(s string) *NameUpdateOne {
	nuo.mutation.SetMiddleName(s)
	return nuo
}

// SetNillableMiddleName sets the "middleName" field if the given value is not nil.
func (nuo *NameUpdateOne) SetNillableMiddleName(s *string) *NameUpdateOne {
	if s != nil {
		nuo.SetMiddleName(*s)
	}
	return nuo
}

// ClearMiddleName clears the value of the "middleName" field.
func (nuo *NameUpdateOne) ClearMiddleName() *NameUpdateOne {
	nuo.mutation.ClearMiddleName()
	return nuo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (nuo *NameUpdateOne) SetUsersID(id uuid.UUID) *NameUpdateOne {
	nuo.mutation.SetUsersID(id)
	return nuo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (nuo *NameUpdateOne) SetNillableUsersID(id *uuid.UUID) *NameUpdateOne {
	if id != nil {
		nuo = nuo.SetUsersID(*id)
	}
	return nuo
}

// SetUsers sets the "users" edge to the User entity.
func (nuo *NameUpdateOne) SetUsers(u *User) *NameUpdateOne {
	return nuo.SetUsersID(u.ID)
}

// Mutation returns the NameMutation object of the builder.
func (nuo *NameUpdateOne) Mutation() *NameMutation {
	return nuo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (nuo *NameUpdateOne) ClearUsers() *NameUpdateOne {
	nuo.mutation.ClearUsers()
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NameUpdateOne) Select(field string, fields ...string) *NameUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Name entity.
func (nuo *NameUpdateOne) Save(ctx context.Context) (*Name, error) {
	var (
		err  error
		node *Name
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			if nuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NameUpdateOne) SaveX(ctx context.Context) *Name {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NameUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NameUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NameUpdateOne) sqlSave(ctx context.Context) (_node *Name, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   name.Table,
			Columns: name.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: name.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Name.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, name.FieldID)
		for _, f := range fields {
			if !name.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != name.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.FamilyName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldFamilyName,
		})
	}
	if nuo.mutation.FamilyNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldFamilyName,
		})
	}
	if value, ok := nuo.mutation.Formatted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldFormatted,
		})
	}
	if nuo.mutation.FormattedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldFormatted,
		})
	}
	if value, ok := nuo.mutation.GivenName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldGivenName,
		})
	}
	if nuo.mutation.GivenNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldGivenName,
		})
	}
	if value, ok := nuo.mutation.HonorificPrefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldHonorificPrefix,
		})
	}
	if nuo.mutation.HonorificPrefixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldHonorificPrefix,
		})
	}
	if value, ok := nuo.mutation.HonorificSuffix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldHonorificSuffix,
		})
	}
	if nuo.mutation.HonorificSuffixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldHonorificSuffix,
		})
	}
	if value, ok := nuo.mutation.MiddleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldMiddleName,
		})
	}
	if nuo.mutation.MiddleNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: name.FieldMiddleName,
		})
	}
	if nuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   name.UsersTable,
			Columns: []string{name.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   name.UsersTable,
			Columns: []string{name.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Name{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{name.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
