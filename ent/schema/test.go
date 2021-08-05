package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Test holds the schema definition for the Test entity.
type Test struct {
	ent.Schema
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("test").
			Values("a", "b"),
	}
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return nil
}
