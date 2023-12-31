// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"simple-bank/ent/account"
	"simple-bank/ent/entry"
	"simple-bank/ent/predicate"
	"simple-bank/ent/transfer"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAge sets the "age" field.
func (au *AccountUpdate) SetAge(i int) *AccountUpdate {
	au.mutation.ResetAge()
	au.mutation.SetAge(i)
	return au
}

// AddAge adds i to the "age" field.
func (au *AccountUpdate) AddAge(i int) *AccountUpdate {
	au.mutation.AddAge(i)
	return au
}

// SetOwner sets the "owner" field.
func (au *AccountUpdate) SetOwner(s string) *AccountUpdate {
	au.mutation.SetOwner(s)
	return au
}

// SetBalance sets the "balance" field.
func (au *AccountUpdate) SetBalance(i int64) *AccountUpdate {
	au.mutation.ResetBalance()
	au.mutation.SetBalance(i)
	return au
}

// AddBalance adds i to the "balance" field.
func (au *AccountUpdate) AddBalance(i int64) *AccountUpdate {
	au.mutation.AddBalance(i)
	return au
}

// SetCurrency sets the "currency" field.
func (au *AccountUpdate) SetCurrency(s string) *AccountUpdate {
	au.mutation.SetCurrency(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AccountUpdate) SetCreatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetCountryCode sets the "country_code" field.
func (au *AccountUpdate) SetCountryCode(i int) *AccountUpdate {
	au.mutation.ResetCountryCode()
	au.mutation.SetCountryCode(i)
	return au
}

// AddCountryCode adds i to the "country_code" field.
func (au *AccountUpdate) AddCountryCode(i int) *AccountUpdate {
	au.mutation.AddCountryCode(i)
	return au
}

// AddEntryIDs adds the "entry" edge to the Entry entity by IDs.
func (au *AccountUpdate) AddEntryIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.AddEntryIDs(ids...)
	return au
}

// AddEntry adds the "entry" edges to the Entry entity.
func (au *AccountUpdate) AddEntry(e ...*Entry) *AccountUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.AddEntryIDs(ids...)
}

// AddTransferIDs adds the "transfer" edge to the Transfer entity by IDs.
func (au *AccountUpdate) AddTransferIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.AddTransferIDs(ids...)
	return au
}

// AddTransfer adds the "transfer" edges to the Transfer entity.
func (au *AccountUpdate) AddTransfer(t ...*Transfer) *AccountUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTransferIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearEntry clears all "entry" edges to the Entry entity.
func (au *AccountUpdate) ClearEntry() *AccountUpdate {
	au.mutation.ClearEntry()
	return au
}

// RemoveEntryIDs removes the "entry" edge to Entry entities by IDs.
func (au *AccountUpdate) RemoveEntryIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.RemoveEntryIDs(ids...)
	return au
}

// RemoveEntry removes "entry" edges to Entry entities.
func (au *AccountUpdate) RemoveEntry(e ...*Entry) *AccountUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.RemoveEntryIDs(ids...)
}

// ClearTransfer clears all "transfer" edges to the Transfer entity.
func (au *AccountUpdate) ClearTransfer() *AccountUpdate {
	au.mutation.ClearTransfer()
	return au
}

// RemoveTransferIDs removes the "transfer" edge to Transfer entities by IDs.
func (au *AccountUpdate) RemoveTransferIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.RemoveTransferIDs(ids...)
	return au
}

// RemoveTransfer removes "transfer" edges to Transfer entities.
func (au *AccountUpdate) RemoveTransfer(t ...*Transfer) *AccountUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTransferIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Age(); ok {
		_spec.SetField(account.FieldAge, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedAge(); ok {
		_spec.AddField(account.FieldAge, field.TypeInt, value)
	}
	if value, ok := au.mutation.Owner(); ok {
		_spec.SetField(account.FieldOwner, field.TypeString, value)
	}
	if value, ok := au.mutation.Balance(); ok {
		_spec.SetField(account.FieldBalance, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedBalance(); ok {
		_spec.AddField(account.FieldBalance, field.TypeInt64, value)
	}
	if value, ok := au.mutation.Currency(); ok {
		_spec.SetField(account.FieldCurrency, field.TypeString, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.CountryCode(); ok {
		_spec.SetField(account.FieldCountryCode, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedCountryCode(); ok {
		_spec.AddField(account.FieldCountryCode, field.TypeInt, value)
	}
	if au.mutation.EntryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntryTable,
			Columns: []string{account.EntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedEntryIDs(); len(nodes) > 0 && !au.mutation.EntryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntryTable,
			Columns: []string{account.EntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EntryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntryTable,
			Columns: []string{account.EntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TransferCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransferTable,
			Columns: []string{account.TransferColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transfer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTransferIDs(); len(nodes) > 0 && !au.mutation.TransferCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransferTable,
			Columns: []string{account.TransferColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transfer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TransferIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransferTable,
			Columns: []string{account.TransferColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transfer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetAge sets the "age" field.
func (auo *AccountUpdateOne) SetAge(i int) *AccountUpdateOne {
	auo.mutation.ResetAge()
	auo.mutation.SetAge(i)
	return auo
}

// AddAge adds i to the "age" field.
func (auo *AccountUpdateOne) AddAge(i int) *AccountUpdateOne {
	auo.mutation.AddAge(i)
	return auo
}

// SetOwner sets the "owner" field.
func (auo *AccountUpdateOne) SetOwner(s string) *AccountUpdateOne {
	auo.mutation.SetOwner(s)
	return auo
}

// SetBalance sets the "balance" field.
func (auo *AccountUpdateOne) SetBalance(i int64) *AccountUpdateOne {
	auo.mutation.ResetBalance()
	auo.mutation.SetBalance(i)
	return auo
}

// AddBalance adds i to the "balance" field.
func (auo *AccountUpdateOne) AddBalance(i int64) *AccountUpdateOne {
	auo.mutation.AddBalance(i)
	return auo
}

// SetCurrency sets the "currency" field.
func (auo *AccountUpdateOne) SetCurrency(s string) *AccountUpdateOne {
	auo.mutation.SetCurrency(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AccountUpdateOne) SetCreatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetCountryCode sets the "country_code" field.
func (auo *AccountUpdateOne) SetCountryCode(i int) *AccountUpdateOne {
	auo.mutation.ResetCountryCode()
	auo.mutation.SetCountryCode(i)
	return auo
}

// AddCountryCode adds i to the "country_code" field.
func (auo *AccountUpdateOne) AddCountryCode(i int) *AccountUpdateOne {
	auo.mutation.AddCountryCode(i)
	return auo
}

// AddEntryIDs adds the "entry" edge to the Entry entity by IDs.
func (auo *AccountUpdateOne) AddEntryIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.AddEntryIDs(ids...)
	return auo
}

// AddEntry adds the "entry" edges to the Entry entity.
func (auo *AccountUpdateOne) AddEntry(e ...*Entry) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.AddEntryIDs(ids...)
}

// AddTransferIDs adds the "transfer" edge to the Transfer entity by IDs.
func (auo *AccountUpdateOne) AddTransferIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.AddTransferIDs(ids...)
	return auo
}

// AddTransfer adds the "transfer" edges to the Transfer entity.
func (auo *AccountUpdateOne) AddTransfer(t ...*Transfer) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTransferIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearEntry clears all "entry" edges to the Entry entity.
func (auo *AccountUpdateOne) ClearEntry() *AccountUpdateOne {
	auo.mutation.ClearEntry()
	return auo
}

// RemoveEntryIDs removes the "entry" edge to Entry entities by IDs.
func (auo *AccountUpdateOne) RemoveEntryIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.RemoveEntryIDs(ids...)
	return auo
}

// RemoveEntry removes "entry" edges to Entry entities.
func (auo *AccountUpdateOne) RemoveEntry(e ...*Entry) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.RemoveEntryIDs(ids...)
}

// ClearTransfer clears all "transfer" edges to the Transfer entity.
func (auo *AccountUpdateOne) ClearTransfer() *AccountUpdateOne {
	auo.mutation.ClearTransfer()
	return auo
}

// RemoveTransferIDs removes the "transfer" edge to Transfer entities by IDs.
func (auo *AccountUpdateOne) RemoveTransferIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.RemoveTransferIDs(ids...)
	return auo
}

// RemoveTransfer removes "transfer" edges to Transfer entities.
func (auo *AccountUpdateOne) RemoveTransfer(t ...*Transfer) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTransferIDs(ids...)
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Age(); ok {
		_spec.SetField(account.FieldAge, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedAge(); ok {
		_spec.AddField(account.FieldAge, field.TypeInt, value)
	}
	if value, ok := auo.mutation.Owner(); ok {
		_spec.SetField(account.FieldOwner, field.TypeString, value)
	}
	if value, ok := auo.mutation.Balance(); ok {
		_spec.SetField(account.FieldBalance, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedBalance(); ok {
		_spec.AddField(account.FieldBalance, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.Currency(); ok {
		_spec.SetField(account.FieldCurrency, field.TypeString, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.CountryCode(); ok {
		_spec.SetField(account.FieldCountryCode, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedCountryCode(); ok {
		_spec.AddField(account.FieldCountryCode, field.TypeInt, value)
	}
	if auo.mutation.EntryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntryTable,
			Columns: []string{account.EntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedEntryIDs(); len(nodes) > 0 && !auo.mutation.EntryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntryTable,
			Columns: []string{account.EntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EntryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntryTable,
			Columns: []string{account.EntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TransferCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransferTable,
			Columns: []string{account.TransferColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transfer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTransferIDs(); len(nodes) > 0 && !auo.mutation.TransferCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransferTable,
			Columns: []string{account.TransferColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transfer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TransferIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransferTable,
			Columns: []string{account.TransferColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transfer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
