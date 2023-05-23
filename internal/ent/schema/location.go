package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/vektah/gqlparser/v2/ast"
	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Location holds the schema definition for the Instance entity.
type Location struct {
	ent.Schema
}

func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Comment("ID for the location.").
			GoType(gidx.PrefixedID("")).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(LocationPrefix) }).
			Unique().
			Immutable(),
		field.Text("name").
			NotEmpty().
			Comment("The name for the location.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("description").
			Comment("An optional description for the location.").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("owner_id").
			GoType(gidx.PrefixedID("")).
			Comment("The ID of the resource owner for the location.").
			Immutable().
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
				entx.EventsHookAdditionalSubject(),
			),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entx.GraphKeyDirective("id"),
		prefixIDDirective(LocationPrefix),
		entx.EventsHookSubjectName("location"),
		entgql.RelayConnection(),
		schema.Comment("Representation of a location."),
		entgql.Implements("MetadataNode"),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a location."),
			entgql.MutationUpdate().Description("Input information to update a location."),
		),
	}
}

func (Location) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}

func prefixIDDirective(prefix string) entgql.Annotation {
	var args []*ast.Argument
	if prefix != "" {
		args = append(args, &ast.Argument{
			Name: "prefix",
			Value: &ast.Value{
				Raw:  prefix,
				Kind: ast.StringValue,
			},
		})
	}

	return entgql.Directives(entgql.NewDirective("prefixedID", args...))
}
