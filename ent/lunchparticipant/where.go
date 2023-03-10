// Code generated by ent, DO NOT EDIT.

package lunchparticipant

import (
	"bc-labs-lunch-bot/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldUpdatedAt, v))
}

// PaymentID applies equality check predicate on the "payment_id" field. It's identical to PaymentIDEQ.
func PaymentID(v int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldPaymentID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldContainsFold(FieldName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldLTE(FieldUpdatedAt, v))
}

// PaymentIDEQ applies the EQ predicate on the "payment_id" field.
func PaymentIDEQ(v int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldEQ(FieldPaymentID, v))
}

// PaymentIDNEQ applies the NEQ predicate on the "payment_id" field.
func PaymentIDNEQ(v int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNEQ(FieldPaymentID, v))
}

// PaymentIDIn applies the In predicate on the "payment_id" field.
func PaymentIDIn(vs ...int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldIn(FieldPaymentID, vs...))
}

// PaymentIDNotIn applies the NotIn predicate on the "payment_id" field.
func PaymentIDNotIn(vs ...int) predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNotIn(FieldPaymentID, vs...))
}

// PaymentIDIsNil applies the IsNil predicate on the "payment_id" field.
func PaymentIDIsNil() predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldIsNull(FieldPaymentID))
}

// PaymentIDNotNil applies the NotNil predicate on the "payment_id" field.
func PaymentIDNotNil() predicate.LunchParticipant {
	return predicate.LunchParticipant(sql.FieldNotNull(FieldPaymentID))
}

// HasPayment applies the HasEdge predicate on the "payment" edge.
func HasPayment() predicate.LunchParticipant {
	return predicate.LunchParticipant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentWith applies the HasEdge predicate on the "payment" edge with a given conditions (other predicates).
func HasPaymentWith(preds ...predicate.Lunch) predicate.LunchParticipant {
	return predicate.LunchParticipant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LunchParticipant) predicate.LunchParticipant {
	return predicate.LunchParticipant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LunchParticipant) predicate.LunchParticipant {
	return predicate.LunchParticipant(func(s *sql.Selector) {
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
func Not(p predicate.LunchParticipant) predicate.LunchParticipant {
	return predicate.LunchParticipant(func(s *sql.Selector) {
		p(s.Not())
	})
}
