// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"simple-bank/ent/transfer"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Transfer is the model entity for the Transfer schema.
type Transfer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// FromAccountID holds the value of the "from_account_id" field.
	FromAccountID int32 `json:"from_account_id,omitempty"`
	// ToAccountID holds the value of the "to_account_id" field.
	ToAccountID int32 `json:"to_account_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int32 `json:"amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt        time.Time `json:"created_at,omitempty"`
	account_transfer *uuid.UUID
	selectValues     sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transfer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transfer.FieldFromAccountID, transfer.FieldToAccountID, transfer.FieldAmount:
			values[i] = new(sql.NullInt64)
		case transfer.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case transfer.FieldID:
			values[i] = new(uuid.UUID)
		case transfer.ForeignKeys[0]: // account_transfer
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transfer fields.
func (t *Transfer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transfer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case transfer.FieldFromAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from_account_id", values[i])
			} else if value.Valid {
				t.FromAccountID = int32(value.Int64)
			}
		case transfer.FieldToAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field to_account_id", values[i])
			} else if value.Valid {
				t.ToAccountID = int32(value.Int64)
			}
		case transfer.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				t.Amount = int32(value.Int64)
			}
		case transfer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case transfer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field account_transfer", values[i])
			} else if value.Valid {
				t.account_transfer = new(uuid.UUID)
				*t.account_transfer = *value.S.(*uuid.UUID)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Transfer.
// This includes values selected through modifiers, order, etc.
func (t *Transfer) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Transfer.
// Note that you need to call Transfer.Unwrap() before calling this method if this Transfer
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transfer) Update() *TransferUpdateOne {
	return NewTransferClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Transfer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transfer) Unwrap() *Transfer {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transfer is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transfer) String() string {
	var builder strings.Builder
	builder.WriteString("Transfer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("from_account_id=")
	builder.WriteString(fmt.Sprintf("%v", t.FromAccountID))
	builder.WriteString(", ")
	builder.WriteString("to_account_id=")
	builder.WriteString(fmt.Sprintf("%v", t.ToAccountID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Transfers is a parsable slice of Transfer.
type Transfers []*Transfer
