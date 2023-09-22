package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Transfer holds the schema definition for the Transfer entity.
type Transfer struct {
	ent.Schema
}

// Fields of the Transfer.
func (Transfer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.Int32("from_account_id").Unique(),
		field.Int32("to_account_id").Unique(),
		field.Int32("amount"),
		field.Time("created_at"),
	}
}

// Edge of the Transfer.
func (Transfer) Edge() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Account.Type).Ref("transfer").Unique().Field("from_account_id"),
		edge.From("owner", Account.Type).Ref("transfer").Unique().Field("to_account_id"),
	}
}

func (Transfer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("from_account_id").Unique(),
		index.Fields("to_account_id").Unique(),
	}
}
