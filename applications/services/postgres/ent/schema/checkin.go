package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CheckIn holds the schema definition for the CheckIn entity.
type CheckIn struct {
	ent.Schema
}

// Fields of the CheckIn.
func (CheckIn) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("attendanceCode").
			Positive().Min(1000000).Max(9999999),
		field.String("studentId"),
		field.Enum("status").
			Values("success", "outOfTime", "notFound", "outOfRange", "error"),
		field.Int64("checkinTime"),
	}
}

// Edges of the CheckIn.
func (CheckIn) Edges() []ent.Edge {
	return nil
}
