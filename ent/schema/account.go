package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Accounts entity.
type Account struct {
	ent.Schema
}

// Fields of the Accounts.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().StorageKey("oid"),
		field.Int("age"),
		field.String("owner"),
		field.Int64("balance"),
		field.String("currency"),
		field.Time("created_at").Default(time.Now()),
		field.Int("country_code"),
	}
}

// Edges of the Accounts.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("entry", Entry.Type),
		edge.To("transfer", Transfer.Type),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner").Unique(),
	}
}
