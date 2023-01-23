// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bc-labs-lunch-bot/ent/lunch"
	"bc-labs-lunch-bot/ent/lunchparticipant"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// LunchParticipant is the model entity for the LunchParticipant schema.
type LunchParticipant struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// PaymentID holds the value of the "payment_id" field.
	PaymentID int `json:"payment_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LunchParticipantQuery when eager-loading is set.
	Edges LunchParticipantEdges `json:"edges"`
}

// LunchParticipantEdges holds the relations/edges for other nodes in the graph.
type LunchParticipantEdges struct {
	// Payment holds the value of the payment edge.
	Payment *Lunch `json:"payment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PaymentOrErr returns the Payment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LunchParticipantEdges) PaymentOrErr() (*Lunch, error) {
	if e.loadedTypes[0] {
		if e.Payment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: lunch.Label}
		}
		return e.Payment, nil
	}
	return nil, &NotLoadedError{edge: "payment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LunchParticipant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lunchparticipant.FieldID, lunchparticipant.FieldPaymentID:
			values[i] = new(sql.NullInt64)
		case lunchparticipant.FieldName:
			values[i] = new(sql.NullString)
		case lunchparticipant.FieldCreatedAt, lunchparticipant.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LunchParticipant", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LunchParticipant fields.
func (lp *LunchParticipant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lunchparticipant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lp.ID = int(value.Int64)
		case lunchparticipant.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				lp.Name = value.String
			}
		case lunchparticipant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lp.CreatedAt = value.Time
			}
		case lunchparticipant.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				lp.UpdatedAt = value.Time
			}
		case lunchparticipant.FieldPaymentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payment_id", values[i])
			} else if value.Valid {
				lp.PaymentID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPayment queries the "payment" edge of the LunchParticipant entity.
func (lp *LunchParticipant) QueryPayment() *LunchQuery {
	return NewLunchParticipantClient(lp.config).QueryPayment(lp)
}

// Update returns a builder for updating this LunchParticipant.
// Note that you need to call LunchParticipant.Unwrap() before calling this method if this LunchParticipant
// was returned from a transaction, and the transaction was committed or rolled back.
func (lp *LunchParticipant) Update() *LunchParticipantUpdateOne {
	return NewLunchParticipantClient(lp.config).UpdateOne(lp)
}

// Unwrap unwraps the LunchParticipant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lp *LunchParticipant) Unwrap() *LunchParticipant {
	_tx, ok := lp.config.driver.(*txDriver)
	if !ok {
		panic("ent: LunchParticipant is not a transactional entity")
	}
	lp.config.driver = _tx.drv
	return lp
}

// String implements the fmt.Stringer.
func (lp *LunchParticipant) String() string {
	var builder strings.Builder
	builder.WriteString("LunchParticipant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lp.ID))
	builder.WriteString("name=")
	builder.WriteString(lp.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(lp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(lp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("payment_id=")
	builder.WriteString(fmt.Sprintf("%v", lp.PaymentID))
	builder.WriteByte(')')
	return builder.String()
}

// LunchParticipants is a parsable slice of LunchParticipant.
type LunchParticipants []*LunchParticipant

func (lp LunchParticipants) config(cfg config) {
	for _i := range lp {
		lp[_i].config = cfg
	}
}
