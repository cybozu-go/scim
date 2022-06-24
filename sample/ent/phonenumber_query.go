// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cybozu-go/scim/sample/ent/phonenumber"
	"github.com/cybozu-go/scim/sample/ent/predicate"
	"github.com/cybozu-go/scim/sample/ent/user"
	"github.com/google/uuid"
)

// PhoneNumberQuery is the builder for querying PhoneNumber entities.
type PhoneNumberQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PhoneNumber
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PhoneNumberQuery builder.
func (pnq *PhoneNumberQuery) Where(ps ...predicate.PhoneNumber) *PhoneNumberQuery {
	pnq.predicates = append(pnq.predicates, ps...)
	return pnq
}

// Limit adds a limit step to the query.
func (pnq *PhoneNumberQuery) Limit(limit int) *PhoneNumberQuery {
	pnq.limit = &limit
	return pnq
}

// Offset adds an offset step to the query.
func (pnq *PhoneNumberQuery) Offset(offset int) *PhoneNumberQuery {
	pnq.offset = &offset
	return pnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pnq *PhoneNumberQuery) Unique(unique bool) *PhoneNumberQuery {
	pnq.unique = &unique
	return pnq
}

// Order adds an order step to the query.
func (pnq *PhoneNumberQuery) Order(o ...OrderFunc) *PhoneNumberQuery {
	pnq.order = append(pnq.order, o...)
	return pnq
}

// QueryUser chains the current query on the "user" edge.
func (pnq *PhoneNumberQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: pnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(phonenumber.Table, phonenumber.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, phonenumber.UserTable, phonenumber.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(pnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PhoneNumber entity from the query.
// Returns a *NotFoundError when no PhoneNumber was found.
func (pnq *PhoneNumberQuery) First(ctx context.Context) (*PhoneNumber, error) {
	nodes, err := pnq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{phonenumber.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pnq *PhoneNumberQuery) FirstX(ctx context.Context) *PhoneNumber {
	node, err := pnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PhoneNumber ID from the query.
// Returns a *NotFoundError when no PhoneNumber ID was found.
func (pnq *PhoneNumberQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pnq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{phonenumber.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pnq *PhoneNumberQuery) FirstIDX(ctx context.Context) int {
	id, err := pnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PhoneNumber entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PhoneNumber entity is found.
// Returns a *NotFoundError when no PhoneNumber entities are found.
func (pnq *PhoneNumberQuery) Only(ctx context.Context) (*PhoneNumber, error) {
	nodes, err := pnq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{phonenumber.Label}
	default:
		return nil, &NotSingularError{phonenumber.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pnq *PhoneNumberQuery) OnlyX(ctx context.Context) *PhoneNumber {
	node, err := pnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PhoneNumber ID in the query.
// Returns a *NotSingularError when more than one PhoneNumber ID is found.
// Returns a *NotFoundError when no entities are found.
func (pnq *PhoneNumberQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pnq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = &NotSingularError{phonenumber.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pnq *PhoneNumberQuery) OnlyIDX(ctx context.Context) int {
	id, err := pnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PhoneNumbers.
func (pnq *PhoneNumberQuery) All(ctx context.Context) ([]*PhoneNumber, error) {
	if err := pnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pnq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pnq *PhoneNumberQuery) AllX(ctx context.Context) []*PhoneNumber {
	nodes, err := pnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PhoneNumber IDs.
func (pnq *PhoneNumberQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pnq.Select(phonenumber.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pnq *PhoneNumberQuery) IDsX(ctx context.Context) []int {
	ids, err := pnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pnq *PhoneNumberQuery) Count(ctx context.Context) (int, error) {
	if err := pnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pnq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pnq *PhoneNumberQuery) CountX(ctx context.Context) int {
	count, err := pnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pnq *PhoneNumberQuery) Exist(ctx context.Context) (bool, error) {
	if err := pnq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pnq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pnq *PhoneNumberQuery) ExistX(ctx context.Context) bool {
	exist, err := pnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PhoneNumberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pnq *PhoneNumberQuery) Clone() *PhoneNumberQuery {
	if pnq == nil {
		return nil
	}
	return &PhoneNumberQuery{
		config:     pnq.config,
		limit:      pnq.limit,
		offset:     pnq.offset,
		order:      append([]OrderFunc{}, pnq.order...),
		predicates: append([]predicate.PhoneNumber{}, pnq.predicates...),
		withUser:   pnq.withUser.Clone(),
		// clone intermediate query.
		sql:    pnq.sql.Clone(),
		path:   pnq.path,
		unique: pnq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (pnq *PhoneNumberQuery) WithUser(opts ...func(*UserQuery)) *PhoneNumberQuery {
	query := &UserQuery{config: pnq.config}
	for _, opt := range opts {
		opt(query)
	}
	pnq.withUser = query
	return pnq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PhoneNumber.Query().
//		GroupBy(phonenumber.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pnq *PhoneNumberQuery) GroupBy(field string, fields ...string) *PhoneNumberGroupBy {
	group := &PhoneNumberGroupBy{config: pnq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pnq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//	}
//
//	client.PhoneNumber.Query().
//		Select(phonenumber.FieldValue).
//		Scan(ctx, &v)
//
func (pnq *PhoneNumberQuery) Select(fields ...string) *PhoneNumberSelect {
	pnq.fields = append(pnq.fields, fields...)
	return &PhoneNumberSelect{PhoneNumberQuery: pnq}
}

func (pnq *PhoneNumberQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pnq.fields {
		if !phonenumber.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pnq.path != nil {
		prev, err := pnq.path(ctx)
		if err != nil {
			return err
		}
		pnq.sql = prev
	}
	return nil
}

func (pnq *PhoneNumberQuery) sqlAll(ctx context.Context) ([]*PhoneNumber, error) {
	var (
		nodes       = []*PhoneNumber{}
		withFKs     = pnq.withFKs
		_spec       = pnq.querySpec()
		loadedTypes = [1]bool{
			pnq.withUser != nil,
		}
	)
	if pnq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, phonenumber.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PhoneNumber{config: pnq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pnq.withUser; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*PhoneNumber)
		for i := range nodes {
			if nodes[i].phone_number_user == nil {
				continue
			}
			fk := *nodes[i].phone_number_user
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "phone_number_user" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (pnq *PhoneNumberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pnq.querySpec()
	_spec.Node.Columns = pnq.fields
	if len(pnq.fields) > 0 {
		_spec.Unique = pnq.unique != nil && *pnq.unique
	}
	return sqlgraph.CountNodes(ctx, pnq.driver, _spec)
}

func (pnq *PhoneNumberQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pnq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pnq *PhoneNumberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   phonenumber.Table,
			Columns: phonenumber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: phonenumber.FieldID,
			},
		},
		From:   pnq.sql,
		Unique: true,
	}
	if unique := pnq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pnq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phonenumber.FieldID)
		for i := range fields {
			if fields[i] != phonenumber.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pnq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pnq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pnq *PhoneNumberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pnq.driver.Dialect())
	t1 := builder.Table(phonenumber.Table)
	columns := pnq.fields
	if len(columns) == 0 {
		columns = phonenumber.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pnq.sql != nil {
		selector = pnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pnq.unique != nil && *pnq.unique {
		selector.Distinct()
	}
	for _, p := range pnq.predicates {
		p(selector)
	}
	for _, p := range pnq.order {
		p(selector)
	}
	if offset := pnq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pnq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PhoneNumberGroupBy is the group-by builder for PhoneNumber entities.
type PhoneNumberGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pngb *PhoneNumberGroupBy) Aggregate(fns ...AggregateFunc) *PhoneNumberGroupBy {
	pngb.fns = append(pngb.fns, fns...)
	return pngb
}

// Scan applies the group-by query and scans the result into the given value.
func (pngb *PhoneNumberGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pngb.path(ctx)
	if err != nil {
		return err
	}
	pngb.sql = query
	return pngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pngb *PhoneNumberGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pngb.fields) > 1 {
		return nil, errors.New("ent: PhoneNumberGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) StringsX(ctx context.Context) []string {
	v, err := pngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pngb *PhoneNumberGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = fmt.Errorf("ent: PhoneNumberGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) StringX(ctx context.Context) string {
	v, err := pngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pngb *PhoneNumberGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pngb.fields) > 1 {
		return nil, errors.New("ent: PhoneNumberGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) IntsX(ctx context.Context) []int {
	v, err := pngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pngb *PhoneNumberGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = fmt.Errorf("ent: PhoneNumberGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) IntX(ctx context.Context) int {
	v, err := pngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pngb *PhoneNumberGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pngb.fields) > 1 {
		return nil, errors.New("ent: PhoneNumberGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pngb *PhoneNumberGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = fmt.Errorf("ent: PhoneNumberGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pngb *PhoneNumberGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pngb.fields) > 1 {
		return nil, errors.New("ent: PhoneNumberGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pngb *PhoneNumberGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = fmt.Errorf("ent: PhoneNumberGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pngb *PhoneNumberGroupBy) BoolX(ctx context.Context) bool {
	v, err := pngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pngb *PhoneNumberGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pngb.fields {
		if !phonenumber.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pngb *PhoneNumberGroupBy) sqlQuery() *sql.Selector {
	selector := pngb.sql.Select()
	aggregation := make([]string, 0, len(pngb.fns))
	for _, fn := range pngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pngb.fields)+len(pngb.fns))
		for _, f := range pngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pngb.fields...)...)
}

// PhoneNumberSelect is the builder for selecting fields of PhoneNumber entities.
type PhoneNumberSelect struct {
	*PhoneNumberQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pns *PhoneNumberSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pns.prepareQuery(ctx); err != nil {
		return err
	}
	pns.sql = pns.PhoneNumberQuery.sqlQuery(ctx)
	return pns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pns *PhoneNumberSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pns *PhoneNumberSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pns.fields) > 1 {
		return nil, errors.New("ent: PhoneNumberSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pns *PhoneNumberSelect) StringsX(ctx context.Context) []string {
	v, err := pns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pns *PhoneNumberSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = fmt.Errorf("ent: PhoneNumberSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pns *PhoneNumberSelect) StringX(ctx context.Context) string {
	v, err := pns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pns *PhoneNumberSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pns.fields) > 1 {
		return nil, errors.New("ent: PhoneNumberSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pns *PhoneNumberSelect) IntsX(ctx context.Context) []int {
	v, err := pns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pns *PhoneNumberSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = fmt.Errorf("ent: PhoneNumberSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pns *PhoneNumberSelect) IntX(ctx context.Context) int {
	v, err := pns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pns *PhoneNumberSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pns.fields) > 1 {
		return nil, errors.New("ent: PhoneNumberSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pns *PhoneNumberSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pns *PhoneNumberSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = fmt.Errorf("ent: PhoneNumberSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pns *PhoneNumberSelect) Float64X(ctx context.Context) float64 {
	v, err := pns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pns *PhoneNumberSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pns.fields) > 1 {
		return nil, errors.New("ent: PhoneNumberSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pns *PhoneNumberSelect) BoolsX(ctx context.Context) []bool {
	v, err := pns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pns *PhoneNumberSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{phonenumber.Label}
	default:
		err = fmt.Errorf("ent: PhoneNumberSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pns *PhoneNumberSelect) BoolX(ctx context.Context) bool {
	v, err := pns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pns *PhoneNumberSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pns.sql.Query()
	if err := pns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}