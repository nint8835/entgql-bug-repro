// Code generated by entc, DO NOT EDIT.

package othertest

import (
	"entgo.io/ent/dialect/sql"
	"github.com/nint8835/entgql-bug-repro/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TestEQ applies the EQ predicate on the "test" field.
func TestEQ(v Test) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTest), v))
	})
}

// TestNEQ applies the NEQ predicate on the "test" field.
func TestNEQ(v Test) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTest), v))
	})
}

// TestIn applies the In predicate on the "test" field.
func TestIn(vs ...Test) predicate.OtherTest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherTest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTest), v...))
	})
}

// TestNotIn applies the NotIn predicate on the "test" field.
func TestNotIn(vs ...Test) predicate.OtherTest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OtherTest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTest), v...))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OtherTest) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OtherTest) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
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
func Not(p predicate.OtherTest) predicate.OtherTest {
	return predicate.OtherTest(func(s *sql.Selector) {
		p(s.Not())
	})
}
