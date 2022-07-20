// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/service-template/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/service-template/pkg/db/ent/template"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Template is the client for interacting with the Template builders.
	Template *TemplateClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Template = NewTemplateClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		Template: NewTemplateClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		Template: NewTemplateClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Template.
//		Query().
//		Count(ctx)
//
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
	c.Template.Use(hooks...)
}

// TemplateClient is a client for the Template schema.
type TemplateClient struct {
	config
}

// NewTemplateClient returns a client for the Template from the given config.
func NewTemplateClient(c config) *TemplateClient {
	return &TemplateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `template.Hooks(f(g(h())))`.
func (c *TemplateClient) Use(hooks ...Hook) {
	c.hooks.Template = append(c.hooks.Template, hooks...)
}

// Create returns a create builder for Template.
func (c *TemplateClient) Create() *TemplateCreate {
	mutation := newTemplateMutation(c.config, OpCreate)
	return &TemplateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Template entities.
func (c *TemplateClient) CreateBulk(builders ...*TemplateCreate) *TemplateCreateBulk {
	return &TemplateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Template.
func (c *TemplateClient) Update() *TemplateUpdate {
	mutation := newTemplateMutation(c.config, OpUpdate)
	return &TemplateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TemplateClient) UpdateOne(t *Template) *TemplateUpdateOne {
	mutation := newTemplateMutation(c.config, OpUpdateOne, withTemplate(t))
	return &TemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TemplateClient) UpdateOneID(id uuid.UUID) *TemplateUpdateOne {
	mutation := newTemplateMutation(c.config, OpUpdateOne, withTemplateID(id))
	return &TemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Template.
func (c *TemplateClient) Delete() *TemplateDelete {
	mutation := newTemplateMutation(c.config, OpDelete)
	return &TemplateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TemplateClient) DeleteOne(t *Template) *TemplateDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TemplateClient) DeleteOneID(id uuid.UUID) *TemplateDeleteOne {
	builder := c.Delete().Where(template.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TemplateDeleteOne{builder}
}

// Query returns a query builder for Template.
func (c *TemplateClient) Query() *TemplateQuery {
	return &TemplateQuery{
		config: c.config,
	}
}

// Get returns a Template entity by its id.
func (c *TemplateClient) Get(ctx context.Context, id uuid.UUID) (*Template, error) {
	return c.Query().Where(template.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TemplateClient) GetX(ctx context.Context, id uuid.UUID) *Template {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TemplateClient) Hooks() []Hook {
	hooks := c.hooks.Template
	return append(hooks[:len(hooks):len(hooks)], template.Hooks[:]...)
}
