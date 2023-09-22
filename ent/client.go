// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"simple-bank/ent/migrate"

	"simple-bank/ent/account"
	"simple-bank/ent/entry"
	"simple-bank/ent/transfer"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Entry is the client for interacting with the Entry builders.
	Entry *EntryClient
	// Transfer is the client for interacting with the Transfer builders.
	Transfer *TransferClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Entry = NewEntryClient(c.config)
	c.Transfer = NewTransferClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Account:  NewAccountClient(cfg),
		Entry:    NewEntryClient(cfg),
		Transfer: NewTransferClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Account:  NewAccountClient(cfg),
		Entry:    NewEntryClient(cfg),
		Transfer: NewTransferClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Entry.Use(hooks...)
	c.Transfer.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Account.Intercept(interceptors...)
	c.Entry.Intercept(interceptors...)
	c.Transfer.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AccountMutation:
		return c.Account.mutate(ctx, m)
	case *EntryMutation:
		return c.Entry.mutate(ctx, m)
	case *TransferMutation:
		return c.Transfer.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `account.Intercept(f(g(h())))`.
func (c *AccountClient) Intercept(interceptors ...Interceptor) {
	c.inters.Account = append(c.inters.Account, interceptors...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uuid.UUID) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id uuid.UUID) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAccount},
		inters: c.Interceptors(),
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uuid.UUID) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uuid.UUID) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEntry queries the entry edge of a Account.
func (c *AccountClient) QueryEntry(a *Account) *EntryQuery {
	query := (&EntryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(entry.Table, entry.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.EntryTable, account.EntryColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTransfer queries the transfer edge of a Account.
func (c *AccountClient) QueryTransfer(a *Account) *TransferQuery {
	query := (&TransferClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(transfer.Table, transfer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.TransferTable, account.TransferColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// Interceptors returns the client interceptors.
func (c *AccountClient) Interceptors() []Interceptor {
	return c.inters.Account
}

func (c *AccountClient) mutate(ctx context.Context, m *AccountMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AccountCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AccountDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Account mutation op: %q", m.Op())
	}
}

// EntryClient is a client for the Entry schema.
type EntryClient struct {
	config
}

// NewEntryClient returns a client for the Entry from the given config.
func NewEntryClient(c config) *EntryClient {
	return &EntryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `entry.Hooks(f(g(h())))`.
func (c *EntryClient) Use(hooks ...Hook) {
	c.hooks.Entry = append(c.hooks.Entry, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `entry.Intercept(f(g(h())))`.
func (c *EntryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Entry = append(c.inters.Entry, interceptors...)
}

// Create returns a builder for creating a Entry entity.
func (c *EntryClient) Create() *EntryCreate {
	mutation := newEntryMutation(c.config, OpCreate)
	return &EntryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Entry entities.
func (c *EntryClient) CreateBulk(builders ...*EntryCreate) *EntryCreateBulk {
	return &EntryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Entry.
func (c *EntryClient) Update() *EntryUpdate {
	mutation := newEntryMutation(c.config, OpUpdate)
	return &EntryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EntryClient) UpdateOne(e *Entry) *EntryUpdateOne {
	mutation := newEntryMutation(c.config, OpUpdateOne, withEntry(e))
	return &EntryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EntryClient) UpdateOneID(id uuid.UUID) *EntryUpdateOne {
	mutation := newEntryMutation(c.config, OpUpdateOne, withEntryID(id))
	return &EntryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Entry.
func (c *EntryClient) Delete() *EntryDelete {
	mutation := newEntryMutation(c.config, OpDelete)
	return &EntryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EntryClient) DeleteOne(e *Entry) *EntryDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EntryClient) DeleteOneID(id uuid.UUID) *EntryDeleteOne {
	builder := c.Delete().Where(entry.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EntryDeleteOne{builder}
}

// Query returns a query builder for Entry.
func (c *EntryClient) Query() *EntryQuery {
	return &EntryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEntry},
		inters: c.Interceptors(),
	}
}

// Get returns a Entry entity by its id.
func (c *EntryClient) Get(ctx context.Context, id uuid.UUID) (*Entry, error) {
	return c.Query().Where(entry.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EntryClient) GetX(ctx context.Context, id uuid.UUID) *Entry {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Entry.
func (c *EntryClient) QueryOwner(e *Entry) *AccountQuery {
	query := (&AccountClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(entry.Table, entry.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entry.OwnerTable, entry.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EntryClient) Hooks() []Hook {
	return c.hooks.Entry
}

// Interceptors returns the client interceptors.
func (c *EntryClient) Interceptors() []Interceptor {
	return c.inters.Entry
}

func (c *EntryClient) mutate(ctx context.Context, m *EntryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EntryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EntryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EntryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EntryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Entry mutation op: %q", m.Op())
	}
}

// TransferClient is a client for the Transfer schema.
type TransferClient struct {
	config
}

// NewTransferClient returns a client for the Transfer from the given config.
func NewTransferClient(c config) *TransferClient {
	return &TransferClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transfer.Hooks(f(g(h())))`.
func (c *TransferClient) Use(hooks ...Hook) {
	c.hooks.Transfer = append(c.hooks.Transfer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `transfer.Intercept(f(g(h())))`.
func (c *TransferClient) Intercept(interceptors ...Interceptor) {
	c.inters.Transfer = append(c.inters.Transfer, interceptors...)
}

// Create returns a builder for creating a Transfer entity.
func (c *TransferClient) Create() *TransferCreate {
	mutation := newTransferMutation(c.config, OpCreate)
	return &TransferCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transfer entities.
func (c *TransferClient) CreateBulk(builders ...*TransferCreate) *TransferCreateBulk {
	return &TransferCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transfer.
func (c *TransferClient) Update() *TransferUpdate {
	mutation := newTransferMutation(c.config, OpUpdate)
	return &TransferUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransferClient) UpdateOne(t *Transfer) *TransferUpdateOne {
	mutation := newTransferMutation(c.config, OpUpdateOne, withTransfer(t))
	return &TransferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransferClient) UpdateOneID(id uuid.UUID) *TransferUpdateOne {
	mutation := newTransferMutation(c.config, OpUpdateOne, withTransferID(id))
	return &TransferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transfer.
func (c *TransferClient) Delete() *TransferDelete {
	mutation := newTransferMutation(c.config, OpDelete)
	return &TransferDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransferClient) DeleteOne(t *Transfer) *TransferDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TransferClient) DeleteOneID(id uuid.UUID) *TransferDeleteOne {
	builder := c.Delete().Where(transfer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransferDeleteOne{builder}
}

// Query returns a query builder for Transfer.
func (c *TransferClient) Query() *TransferQuery {
	return &TransferQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTransfer},
		inters: c.Interceptors(),
	}
}

// Get returns a Transfer entity by its id.
func (c *TransferClient) Get(ctx context.Context, id uuid.UUID) (*Transfer, error) {
	return c.Query().Where(transfer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransferClient) GetX(ctx context.Context, id uuid.UUID) *Transfer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TransferClient) Hooks() []Hook {
	return c.hooks.Transfer
}

// Interceptors returns the client interceptors.
func (c *TransferClient) Interceptors() []Interceptor {
	return c.inters.Transfer
}

func (c *TransferClient) mutate(ctx context.Context, m *TransferMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TransferCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TransferUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TransferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TransferDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Transfer mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Account, Entry, Transfer []ent.Hook
	}
	inters struct {
		Account, Entry, Transfer []ent.Interceptor
	}
)
