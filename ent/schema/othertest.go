package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OtherTest holds the schema definition for the OtherTest entity.
type OtherTest struct {
	ent.Schema
}

// Fields of the OtherTest.
func (OtherTest) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("test").
			Values("c", "d"),
	}
}

// Edges of the OtherTest.
func (OtherTest) Edges() []ent.Edge {
	return nil
}
