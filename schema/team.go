package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Team struct {
	ent.Schema
}

func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name").Optional(),
		field.String("user_id").Optional(),
		field.String("slug").Optional(),
	}
}

func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("teams").
			Field("user_id").
			Unique(),
	}
}
