package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// LunchParticipant holds the schema definition for the LunchParticipant entity.
type LunchParticipant struct {
	ent.Schema
}

// Fields of the LunchParticipant.
func (LunchParticipant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.Int("payment_id").
			Optional(),
	}
}

// Edges of the LunchParticipant.
func (LunchParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment", Lunch.Type).
			Ref("participant").
			Unique().
			Field("payment_id"),
	}
}
