// Code generated by entc, DO NOT EDIT.

package name

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cybozu-go/scim/sample/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FamilyName applies equality check predicate on the "familyName" field. It's identical to FamilyNameEQ.
func FamilyName(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFamilyName), v))
	})
}

// Formatted applies equality check predicate on the "formatted" field. It's identical to FormattedEQ.
func Formatted(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFormatted), v))
	})
}

// GivenName applies equality check predicate on the "givenName" field. It's identical to GivenNameEQ.
func GivenName(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGivenName), v))
	})
}

// HonorificPrefix applies equality check predicate on the "honorificPrefix" field. It's identical to HonorificPrefixEQ.
func HonorificPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificSuffix applies equality check predicate on the "honorificSuffix" field. It's identical to HonorificSuffixEQ.
func HonorificSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHonorificSuffix), v))
	})
}

// MiddleName applies equality check predicate on the "middleName" field. It's identical to MiddleNameEQ.
func MiddleName(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiddleName), v))
	})
}

// FamilyNameEQ applies the EQ predicate on the "familyName" field.
func FamilyNameEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFamilyName), v))
	})
}

// FamilyNameNEQ applies the NEQ predicate on the "familyName" field.
func FamilyNameNEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFamilyName), v))
	})
}

// FamilyNameIn applies the In predicate on the "familyName" field.
func FamilyNameIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFamilyName), v...))
	})
}

// FamilyNameNotIn applies the NotIn predicate on the "familyName" field.
func FamilyNameNotIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFamilyName), v...))
	})
}

// FamilyNameGT applies the GT predicate on the "familyName" field.
func FamilyNameGT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFamilyName), v))
	})
}

// FamilyNameGTE applies the GTE predicate on the "familyName" field.
func FamilyNameGTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFamilyName), v))
	})
}

// FamilyNameLT applies the LT predicate on the "familyName" field.
func FamilyNameLT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFamilyName), v))
	})
}

// FamilyNameLTE applies the LTE predicate on the "familyName" field.
func FamilyNameLTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFamilyName), v))
	})
}

// FamilyNameContains applies the Contains predicate on the "familyName" field.
func FamilyNameContains(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFamilyName), v))
	})
}

// FamilyNameHasPrefix applies the HasPrefix predicate on the "familyName" field.
func FamilyNameHasPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFamilyName), v))
	})
}

// FamilyNameHasSuffix applies the HasSuffix predicate on the "familyName" field.
func FamilyNameHasSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFamilyName), v))
	})
}

// FamilyNameIsNil applies the IsNil predicate on the "familyName" field.
func FamilyNameIsNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFamilyName)))
	})
}

// FamilyNameNotNil applies the NotNil predicate on the "familyName" field.
func FamilyNameNotNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFamilyName)))
	})
}

// FamilyNameEqualFold applies the EqualFold predicate on the "familyName" field.
func FamilyNameEqualFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFamilyName), v))
	})
}

// FamilyNameContainsFold applies the ContainsFold predicate on the "familyName" field.
func FamilyNameContainsFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFamilyName), v))
	})
}

// FormattedEQ applies the EQ predicate on the "formatted" field.
func FormattedEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFormatted), v))
	})
}

// FormattedNEQ applies the NEQ predicate on the "formatted" field.
func FormattedNEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFormatted), v))
	})
}

// FormattedIn applies the In predicate on the "formatted" field.
func FormattedIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFormatted), v...))
	})
}

// FormattedNotIn applies the NotIn predicate on the "formatted" field.
func FormattedNotIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFormatted), v...))
	})
}

// FormattedGT applies the GT predicate on the "formatted" field.
func FormattedGT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFormatted), v))
	})
}

// FormattedGTE applies the GTE predicate on the "formatted" field.
func FormattedGTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFormatted), v))
	})
}

// FormattedLT applies the LT predicate on the "formatted" field.
func FormattedLT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFormatted), v))
	})
}

// FormattedLTE applies the LTE predicate on the "formatted" field.
func FormattedLTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFormatted), v))
	})
}

// FormattedContains applies the Contains predicate on the "formatted" field.
func FormattedContains(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFormatted), v))
	})
}

// FormattedHasPrefix applies the HasPrefix predicate on the "formatted" field.
func FormattedHasPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFormatted), v))
	})
}

// FormattedHasSuffix applies the HasSuffix predicate on the "formatted" field.
func FormattedHasSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFormatted), v))
	})
}

// FormattedIsNil applies the IsNil predicate on the "formatted" field.
func FormattedIsNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFormatted)))
	})
}

// FormattedNotNil applies the NotNil predicate on the "formatted" field.
func FormattedNotNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFormatted)))
	})
}

// FormattedEqualFold applies the EqualFold predicate on the "formatted" field.
func FormattedEqualFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFormatted), v))
	})
}

// FormattedContainsFold applies the ContainsFold predicate on the "formatted" field.
func FormattedContainsFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFormatted), v))
	})
}

// GivenNameEQ applies the EQ predicate on the "givenName" field.
func GivenNameEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGivenName), v))
	})
}

// GivenNameNEQ applies the NEQ predicate on the "givenName" field.
func GivenNameNEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGivenName), v))
	})
}

// GivenNameIn applies the In predicate on the "givenName" field.
func GivenNameIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGivenName), v...))
	})
}

// GivenNameNotIn applies the NotIn predicate on the "givenName" field.
func GivenNameNotIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGivenName), v...))
	})
}

// GivenNameGT applies the GT predicate on the "givenName" field.
func GivenNameGT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGivenName), v))
	})
}

// GivenNameGTE applies the GTE predicate on the "givenName" field.
func GivenNameGTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGivenName), v))
	})
}

// GivenNameLT applies the LT predicate on the "givenName" field.
func GivenNameLT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGivenName), v))
	})
}

// GivenNameLTE applies the LTE predicate on the "givenName" field.
func GivenNameLTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGivenName), v))
	})
}

// GivenNameContains applies the Contains predicate on the "givenName" field.
func GivenNameContains(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGivenName), v))
	})
}

// GivenNameHasPrefix applies the HasPrefix predicate on the "givenName" field.
func GivenNameHasPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGivenName), v))
	})
}

// GivenNameHasSuffix applies the HasSuffix predicate on the "givenName" field.
func GivenNameHasSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGivenName), v))
	})
}

// GivenNameIsNil applies the IsNil predicate on the "givenName" field.
func GivenNameIsNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGivenName)))
	})
}

// GivenNameNotNil applies the NotNil predicate on the "givenName" field.
func GivenNameNotNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGivenName)))
	})
}

// GivenNameEqualFold applies the EqualFold predicate on the "givenName" field.
func GivenNameEqualFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGivenName), v))
	})
}

// GivenNameContainsFold applies the ContainsFold predicate on the "givenName" field.
func GivenNameContainsFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGivenName), v))
	})
}

// HonorificPrefixEQ applies the EQ predicate on the "honorificPrefix" field.
func HonorificPrefixEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixNEQ applies the NEQ predicate on the "honorificPrefix" field.
func HonorificPrefixNEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixIn applies the In predicate on the "honorificPrefix" field.
func HonorificPrefixIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHonorificPrefix), v...))
	})
}

// HonorificPrefixNotIn applies the NotIn predicate on the "honorificPrefix" field.
func HonorificPrefixNotIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHonorificPrefix), v...))
	})
}

// HonorificPrefixGT applies the GT predicate on the "honorificPrefix" field.
func HonorificPrefixGT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixGTE applies the GTE predicate on the "honorificPrefix" field.
func HonorificPrefixGTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixLT applies the LT predicate on the "honorificPrefix" field.
func HonorificPrefixLT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixLTE applies the LTE predicate on the "honorificPrefix" field.
func HonorificPrefixLTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixContains applies the Contains predicate on the "honorificPrefix" field.
func HonorificPrefixContains(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixHasPrefix applies the HasPrefix predicate on the "honorificPrefix" field.
func HonorificPrefixHasPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixHasSuffix applies the HasSuffix predicate on the "honorificPrefix" field.
func HonorificPrefixHasSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixIsNil applies the IsNil predicate on the "honorificPrefix" field.
func HonorificPrefixIsNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHonorificPrefix)))
	})
}

// HonorificPrefixNotNil applies the NotNil predicate on the "honorificPrefix" field.
func HonorificPrefixNotNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHonorificPrefix)))
	})
}

// HonorificPrefixEqualFold applies the EqualFold predicate on the "honorificPrefix" field.
func HonorificPrefixEqualFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificPrefixContainsFold applies the ContainsFold predicate on the "honorificPrefix" field.
func HonorificPrefixContainsFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHonorificPrefix), v))
	})
}

// HonorificSuffixEQ applies the EQ predicate on the "honorificSuffix" field.
func HonorificSuffixEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixNEQ applies the NEQ predicate on the "honorificSuffix" field.
func HonorificSuffixNEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixIn applies the In predicate on the "honorificSuffix" field.
func HonorificSuffixIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHonorificSuffix), v...))
	})
}

// HonorificSuffixNotIn applies the NotIn predicate on the "honorificSuffix" field.
func HonorificSuffixNotIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHonorificSuffix), v...))
	})
}

// HonorificSuffixGT applies the GT predicate on the "honorificSuffix" field.
func HonorificSuffixGT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixGTE applies the GTE predicate on the "honorificSuffix" field.
func HonorificSuffixGTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixLT applies the LT predicate on the "honorificSuffix" field.
func HonorificSuffixLT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixLTE applies the LTE predicate on the "honorificSuffix" field.
func HonorificSuffixLTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixContains applies the Contains predicate on the "honorificSuffix" field.
func HonorificSuffixContains(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixHasPrefix applies the HasPrefix predicate on the "honorificSuffix" field.
func HonorificSuffixHasPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixHasSuffix applies the HasSuffix predicate on the "honorificSuffix" field.
func HonorificSuffixHasSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixIsNil applies the IsNil predicate on the "honorificSuffix" field.
func HonorificSuffixIsNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHonorificSuffix)))
	})
}

// HonorificSuffixNotNil applies the NotNil predicate on the "honorificSuffix" field.
func HonorificSuffixNotNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHonorificSuffix)))
	})
}

// HonorificSuffixEqualFold applies the EqualFold predicate on the "honorificSuffix" field.
func HonorificSuffixEqualFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHonorificSuffix), v))
	})
}

// HonorificSuffixContainsFold applies the ContainsFold predicate on the "honorificSuffix" field.
func HonorificSuffixContainsFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHonorificSuffix), v))
	})
}

// MiddleNameEQ applies the EQ predicate on the "middleName" field.
func MiddleNameEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiddleName), v))
	})
}

// MiddleNameNEQ applies the NEQ predicate on the "middleName" field.
func MiddleNameNEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMiddleName), v))
	})
}

// MiddleNameIn applies the In predicate on the "middleName" field.
func MiddleNameIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMiddleName), v...))
	})
}

// MiddleNameNotIn applies the NotIn predicate on the "middleName" field.
func MiddleNameNotIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMiddleName), v...))
	})
}

// MiddleNameGT applies the GT predicate on the "middleName" field.
func MiddleNameGT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMiddleName), v))
	})
}

// MiddleNameGTE applies the GTE predicate on the "middleName" field.
func MiddleNameGTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMiddleName), v))
	})
}

// MiddleNameLT applies the LT predicate on the "middleName" field.
func MiddleNameLT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMiddleName), v))
	})
}

// MiddleNameLTE applies the LTE predicate on the "middleName" field.
func MiddleNameLTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMiddleName), v))
	})
}

// MiddleNameContains applies the Contains predicate on the "middleName" field.
func MiddleNameContains(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMiddleName), v))
	})
}

// MiddleNameHasPrefix applies the HasPrefix predicate on the "middleName" field.
func MiddleNameHasPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMiddleName), v))
	})
}

// MiddleNameHasSuffix applies the HasSuffix predicate on the "middleName" field.
func MiddleNameHasSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMiddleName), v))
	})
}

// MiddleNameIsNil applies the IsNil predicate on the "middleName" field.
func MiddleNameIsNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMiddleName)))
	})
}

// MiddleNameNotNil applies the NotNil predicate on the "middleName" field.
func MiddleNameNotNil() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMiddleName)))
	})
}

// MiddleNameEqualFold applies the EqualFold predicate on the "middleName" field.
func MiddleNameEqualFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMiddleName), v))
	})
}

// MiddleNameContainsFold applies the ContainsFold predicate on the "middleName" field.
func MiddleNameContainsFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMiddleName), v))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
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
func And(predicates ...predicate.Name) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Name) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
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
func Not(p predicate.Name) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		p(s.Not())
	})
}