package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Pnu struct {
	ent.Schema
}

func (Pnu) Fields() []ent.Field {
	return []ent.Field{
		// ID
		field.String("id").Unique(),

		// Location Information in string
		field.String("jibun"),
		field.String("bchk"),
		field.String("pnu"),
		field.String("col_adm_se"),

		// Geometry
		field.Other("geometry", &Geometry{}).
			SchemaType(Geometry{}.SchemaType()),
	}
}

// Edges of the Pnu
func (Pnu) Edges() []ent.Edge {
	return nil
}

func (Pnu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "pnu"},
	}
}
