// Code generated by entc, DO NOT EDIT.

package email

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cybozu-go/scim/sample/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Display applies equality check predicate on the "display" field. It's identical to DisplayEQ.
func Display(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplay), v))
	})
}

// Primary applies equality check predicate on the "primary" field. It's identical to PrimaryEQ.
func Primary(v bool) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimary), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// DisplayEQ applies the EQ predicate on the "display" field.
func DisplayEQ(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplay), v))
	})
}

// DisplayNEQ applies the NEQ predicate on the "display" field.
func DisplayNEQ(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisplay), v))
	})
}

// DisplayIn applies the In predicate on the "display" field.
func DisplayIn(vs ...string) predicate.Email {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Email(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDisplay), v...))
	})
}

// DisplayNotIn applies the NotIn predicate on the "display" field.
func DisplayNotIn(vs ...string) predicate.Email {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Email(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDisplay), v...))
	})
}

// DisplayGT applies the GT predicate on the "display" field.
func DisplayGT(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDisplay), v))
	})
}

// DisplayGTE applies the GTE predicate on the "display" field.
func DisplayGTE(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDisplay), v))
	})
}

// DisplayLT applies the LT predicate on the "display" field.
func DisplayLT(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDisplay), v))
	})
}

// DisplayLTE applies the LTE predicate on the "display" field.
func DisplayLTE(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDisplay), v))
	})
}

// DisplayContains applies the Contains predicate on the "display" field.
func DisplayContains(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDisplay), v))
	})
}

// DisplayHasPrefix applies the HasPrefix predicate on the "display" field.
func DisplayHasPrefix(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDisplay), v))
	})
}

// DisplayHasSuffix applies the HasSuffix predicate on the "display" field.
func DisplayHasSuffix(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDisplay), v))
	})
}

// DisplayIsNil applies the IsNil predicate on the "display" field.
func DisplayIsNil() predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDisplay)))
	})
}

// DisplayNotNil applies the NotNil predicate on the "display" field.
func DisplayNotNil() predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDisplay)))
	})
}

// DisplayEqualFold applies the EqualFold predicate on the "display" field.
func DisplayEqualFold(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDisplay), v))
	})
}

// DisplayContainsFold applies the ContainsFold predicate on the "display" field.
func DisplayContainsFold(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDisplay), v))
	})
}

// PrimaryEQ applies the EQ predicate on the "primary" field.
func PrimaryEQ(v bool) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimary), v))
	})
}

// PrimaryNEQ applies the NEQ predicate on the "primary" field.
func PrimaryNEQ(v bool) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrimary), v))
	})
}

// PrimaryIsNil applies the IsNil predicate on the "primary" field.
func PrimaryIsNil() predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPrimary)))
	})
}

// PrimaryNotNil applies the NotNil predicate on the "primary" field.
func PrimaryNotNil() predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPrimary)))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Email {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Email(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Email {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Email(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldType)))
	})
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldType)))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Email {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Email(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Email {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Email(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldValue), v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldValue), v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldValue), v))
	})
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldValue), v))
	})
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldValue), v))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Email) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Email) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Email) predicate.Email {
	return predicate.Email(func(s *sql.Selector) {
		p(s.Not())
	})
}
