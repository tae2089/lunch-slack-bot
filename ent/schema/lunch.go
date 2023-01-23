package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Lunch holds the schema definition for the Lunch entity.
type Lunch struct {
	ent.Schema
}

// Fields of the Lunch.
func (Lunch) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.String("payer"),
		field.String("restaurantName"),
		field.String("cafeName"),
		field.Time("paymentTime"),
	}
}

// Edges of the Lunch.
func (Lunch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("participant", LunchParticipant.Type),
	}
}
