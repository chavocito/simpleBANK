package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Entry holds the schema definition for the Entry entity.
type Entry struct {
	ent.Schema
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.UUID("account_id", uuid.UUID{}),
		field.Int32("amount"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Entry.
func (Entry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Account.Type).Ref("entry").Unique().Field("account_id").Required(),
	}
}

func (Entry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id"),
	}
}
