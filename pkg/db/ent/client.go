// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/miningdetail"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/mininggeneral"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/miningunsold"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/profit"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/withdraw"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Detail is the client for interacting with the Detail builders.
	Detail *DetailClient
	// General is the client for interacting with the General builders.
	General *GeneralClient
	// MiningDetail is the client for interacting with the MiningDetail builders.
	MiningDetail *MiningDetailClient
	// MiningGeneral is the client for interacting with the MiningGeneral builders.
	MiningGeneral *MiningGeneralClient
	// MiningUnsold is the client for interacting with the MiningUnsold builders.
	MiningUnsold *MiningUnsoldClient
	// Profit is the client for interacting with the Profit builders.
	Profit *ProfitClient
	// Withdraw is the client for interacting with the Withdraw builders.
	Withdraw *WithdrawClient
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
	c.Detail = NewDetailClient(c.config)
	c.General = NewGeneralClient(c.config)
	c.MiningDetail = NewMiningDetailClient(c.config)
	c.MiningGeneral = NewMiningGeneralClient(c.config)
	c.MiningUnsold = NewMiningUnsoldClient(c.config)
	c.Profit = NewProfitClient(c.config)
	c.Withdraw = NewWithdrawClient(c.config)
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
		ctx:           ctx,
		config:        cfg,
		Detail:        NewDetailClient(cfg),
		General:       NewGeneralClient(cfg),
		MiningDetail:  NewMiningDetailClient(cfg),
		MiningGeneral: NewMiningGeneralClient(cfg),
		MiningUnsold:  NewMiningUnsoldClient(cfg),
		Profit:        NewProfitClient(cfg),
		Withdraw:      NewWithdrawClient(cfg),
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
		ctx:           ctx,
		config:        cfg,
		Detail:        NewDetailClient(cfg),
		General:       NewGeneralClient(cfg),
		MiningDetail:  NewMiningDetailClient(cfg),
		MiningGeneral: NewMiningGeneralClient(cfg),
		MiningUnsold:  NewMiningUnsoldClient(cfg),
		Profit:        NewProfitClient(cfg),
		Withdraw:      NewWithdrawClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Detail.
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
	c.Detail.Use(hooks...)
	c.General.Use(hooks...)
	c.MiningDetail.Use(hooks...)
	c.MiningGeneral.Use(hooks...)
	c.MiningUnsold.Use(hooks...)
	c.Profit.Use(hooks...)
	c.Withdraw.Use(hooks...)
}

// DetailClient is a client for the Detail schema.
type DetailClient struct {
	config
}

// NewDetailClient returns a client for the Detail from the given config.
func NewDetailClient(c config) *DetailClient {
	return &DetailClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `detail.Hooks(f(g(h())))`.
func (c *DetailClient) Use(hooks ...Hook) {
	c.hooks.Detail = append(c.hooks.Detail, hooks...)
}

// Create returns a builder for creating a Detail entity.
func (c *DetailClient) Create() *DetailCreate {
	mutation := newDetailMutation(c.config, OpCreate)
	return &DetailCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Detail entities.
func (c *DetailClient) CreateBulk(builders ...*DetailCreate) *DetailCreateBulk {
	return &DetailCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Detail.
func (c *DetailClient) Update() *DetailUpdate {
	mutation := newDetailMutation(c.config, OpUpdate)
	return &DetailUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DetailClient) UpdateOne(d *Detail) *DetailUpdateOne {
	mutation := newDetailMutation(c.config, OpUpdateOne, withDetail(d))
	return &DetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DetailClient) UpdateOneID(id uuid.UUID) *DetailUpdateOne {
	mutation := newDetailMutation(c.config, OpUpdateOne, withDetailID(id))
	return &DetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Detail.
func (c *DetailClient) Delete() *DetailDelete {
	mutation := newDetailMutation(c.config, OpDelete)
	return &DetailDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DetailClient) DeleteOne(d *Detail) *DetailDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DetailClient) DeleteOneID(id uuid.UUID) *DetailDeleteOne {
	builder := c.Delete().Where(detail.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DetailDeleteOne{builder}
}

// Query returns a query builder for Detail.
func (c *DetailClient) Query() *DetailQuery {
	return &DetailQuery{
		config: c.config,
	}
}

// Get returns a Detail entity by its id.
func (c *DetailClient) Get(ctx context.Context, id uuid.UUID) (*Detail, error) {
	return c.Query().Where(detail.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DetailClient) GetX(ctx context.Context, id uuid.UUID) *Detail {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DetailClient) Hooks() []Hook {
	hooks := c.hooks.Detail
	return append(hooks[:len(hooks):len(hooks)], detail.Hooks[:]...)
}

// GeneralClient is a client for the General schema.
type GeneralClient struct {
	config
}

// NewGeneralClient returns a client for the General from the given config.
func NewGeneralClient(c config) *GeneralClient {
	return &GeneralClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `general.Hooks(f(g(h())))`.
func (c *GeneralClient) Use(hooks ...Hook) {
	c.hooks.General = append(c.hooks.General, hooks...)
}

// Create returns a builder for creating a General entity.
func (c *GeneralClient) Create() *GeneralCreate {
	mutation := newGeneralMutation(c.config, OpCreate)
	return &GeneralCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of General entities.
func (c *GeneralClient) CreateBulk(builders ...*GeneralCreate) *GeneralCreateBulk {
	return &GeneralCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for General.
func (c *GeneralClient) Update() *GeneralUpdate {
	mutation := newGeneralMutation(c.config, OpUpdate)
	return &GeneralUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GeneralClient) UpdateOne(ge *General) *GeneralUpdateOne {
	mutation := newGeneralMutation(c.config, OpUpdateOne, withGeneral(ge))
	return &GeneralUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GeneralClient) UpdateOneID(id uuid.UUID) *GeneralUpdateOne {
	mutation := newGeneralMutation(c.config, OpUpdateOne, withGeneralID(id))
	return &GeneralUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for General.
func (c *GeneralClient) Delete() *GeneralDelete {
	mutation := newGeneralMutation(c.config, OpDelete)
	return &GeneralDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GeneralClient) DeleteOne(ge *General) *GeneralDeleteOne {
	return c.DeleteOneID(ge.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GeneralClient) DeleteOneID(id uuid.UUID) *GeneralDeleteOne {
	builder := c.Delete().Where(general.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GeneralDeleteOne{builder}
}

// Query returns a query builder for General.
func (c *GeneralClient) Query() *GeneralQuery {
	return &GeneralQuery{
		config: c.config,
	}
}

// Get returns a General entity by its id.
func (c *GeneralClient) Get(ctx context.Context, id uuid.UUID) (*General, error) {
	return c.Query().Where(general.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GeneralClient) GetX(ctx context.Context, id uuid.UUID) *General {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GeneralClient) Hooks() []Hook {
	hooks := c.hooks.General
	return append(hooks[:len(hooks):len(hooks)], general.Hooks[:]...)
}

// MiningDetailClient is a client for the MiningDetail schema.
type MiningDetailClient struct {
	config
}

// NewMiningDetailClient returns a client for the MiningDetail from the given config.
func NewMiningDetailClient(c config) *MiningDetailClient {
	return &MiningDetailClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `miningdetail.Hooks(f(g(h())))`.
func (c *MiningDetailClient) Use(hooks ...Hook) {
	c.hooks.MiningDetail = append(c.hooks.MiningDetail, hooks...)
}

// Create returns a builder for creating a MiningDetail entity.
func (c *MiningDetailClient) Create() *MiningDetailCreate {
	mutation := newMiningDetailMutation(c.config, OpCreate)
	return &MiningDetailCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MiningDetail entities.
func (c *MiningDetailClient) CreateBulk(builders ...*MiningDetailCreate) *MiningDetailCreateBulk {
	return &MiningDetailCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MiningDetail.
func (c *MiningDetailClient) Update() *MiningDetailUpdate {
	mutation := newMiningDetailMutation(c.config, OpUpdate)
	return &MiningDetailUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MiningDetailClient) UpdateOne(md *MiningDetail) *MiningDetailUpdateOne {
	mutation := newMiningDetailMutation(c.config, OpUpdateOne, withMiningDetail(md))
	return &MiningDetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MiningDetailClient) UpdateOneID(id uuid.UUID) *MiningDetailUpdateOne {
	mutation := newMiningDetailMutation(c.config, OpUpdateOne, withMiningDetailID(id))
	return &MiningDetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MiningDetail.
func (c *MiningDetailClient) Delete() *MiningDetailDelete {
	mutation := newMiningDetailMutation(c.config, OpDelete)
	return &MiningDetailDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MiningDetailClient) DeleteOne(md *MiningDetail) *MiningDetailDeleteOne {
	return c.DeleteOneID(md.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MiningDetailClient) DeleteOneID(id uuid.UUID) *MiningDetailDeleteOne {
	builder := c.Delete().Where(miningdetail.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MiningDetailDeleteOne{builder}
}

// Query returns a query builder for MiningDetail.
func (c *MiningDetailClient) Query() *MiningDetailQuery {
	return &MiningDetailQuery{
		config: c.config,
	}
}

// Get returns a MiningDetail entity by its id.
func (c *MiningDetailClient) Get(ctx context.Context, id uuid.UUID) (*MiningDetail, error) {
	return c.Query().Where(miningdetail.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MiningDetailClient) GetX(ctx context.Context, id uuid.UUID) *MiningDetail {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MiningDetailClient) Hooks() []Hook {
	hooks := c.hooks.MiningDetail
	return append(hooks[:len(hooks):len(hooks)], miningdetail.Hooks[:]...)
}

// MiningGeneralClient is a client for the MiningGeneral schema.
type MiningGeneralClient struct {
	config
}

// NewMiningGeneralClient returns a client for the MiningGeneral from the given config.
func NewMiningGeneralClient(c config) *MiningGeneralClient {
	return &MiningGeneralClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mininggeneral.Hooks(f(g(h())))`.
func (c *MiningGeneralClient) Use(hooks ...Hook) {
	c.hooks.MiningGeneral = append(c.hooks.MiningGeneral, hooks...)
}

// Create returns a builder for creating a MiningGeneral entity.
func (c *MiningGeneralClient) Create() *MiningGeneralCreate {
	mutation := newMiningGeneralMutation(c.config, OpCreate)
	return &MiningGeneralCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MiningGeneral entities.
func (c *MiningGeneralClient) CreateBulk(builders ...*MiningGeneralCreate) *MiningGeneralCreateBulk {
	return &MiningGeneralCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MiningGeneral.
func (c *MiningGeneralClient) Update() *MiningGeneralUpdate {
	mutation := newMiningGeneralMutation(c.config, OpUpdate)
	return &MiningGeneralUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MiningGeneralClient) UpdateOne(mg *MiningGeneral) *MiningGeneralUpdateOne {
	mutation := newMiningGeneralMutation(c.config, OpUpdateOne, withMiningGeneral(mg))
	return &MiningGeneralUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MiningGeneralClient) UpdateOneID(id uuid.UUID) *MiningGeneralUpdateOne {
	mutation := newMiningGeneralMutation(c.config, OpUpdateOne, withMiningGeneralID(id))
	return &MiningGeneralUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MiningGeneral.
func (c *MiningGeneralClient) Delete() *MiningGeneralDelete {
	mutation := newMiningGeneralMutation(c.config, OpDelete)
	return &MiningGeneralDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MiningGeneralClient) DeleteOne(mg *MiningGeneral) *MiningGeneralDeleteOne {
	return c.DeleteOneID(mg.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MiningGeneralClient) DeleteOneID(id uuid.UUID) *MiningGeneralDeleteOne {
	builder := c.Delete().Where(mininggeneral.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MiningGeneralDeleteOne{builder}
}

// Query returns a query builder for MiningGeneral.
func (c *MiningGeneralClient) Query() *MiningGeneralQuery {
	return &MiningGeneralQuery{
		config: c.config,
	}
}

// Get returns a MiningGeneral entity by its id.
func (c *MiningGeneralClient) Get(ctx context.Context, id uuid.UUID) (*MiningGeneral, error) {
	return c.Query().Where(mininggeneral.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MiningGeneralClient) GetX(ctx context.Context, id uuid.UUID) *MiningGeneral {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MiningGeneralClient) Hooks() []Hook {
	hooks := c.hooks.MiningGeneral
	return append(hooks[:len(hooks):len(hooks)], mininggeneral.Hooks[:]...)
}

// MiningUnsoldClient is a client for the MiningUnsold schema.
type MiningUnsoldClient struct {
	config
}

// NewMiningUnsoldClient returns a client for the MiningUnsold from the given config.
func NewMiningUnsoldClient(c config) *MiningUnsoldClient {
	return &MiningUnsoldClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `miningunsold.Hooks(f(g(h())))`.
func (c *MiningUnsoldClient) Use(hooks ...Hook) {
	c.hooks.MiningUnsold = append(c.hooks.MiningUnsold, hooks...)
}

// Create returns a builder for creating a MiningUnsold entity.
func (c *MiningUnsoldClient) Create() *MiningUnsoldCreate {
	mutation := newMiningUnsoldMutation(c.config, OpCreate)
	return &MiningUnsoldCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MiningUnsold entities.
func (c *MiningUnsoldClient) CreateBulk(builders ...*MiningUnsoldCreate) *MiningUnsoldCreateBulk {
	return &MiningUnsoldCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MiningUnsold.
func (c *MiningUnsoldClient) Update() *MiningUnsoldUpdate {
	mutation := newMiningUnsoldMutation(c.config, OpUpdate)
	return &MiningUnsoldUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MiningUnsoldClient) UpdateOne(mu *MiningUnsold) *MiningUnsoldUpdateOne {
	mutation := newMiningUnsoldMutation(c.config, OpUpdateOne, withMiningUnsold(mu))
	return &MiningUnsoldUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MiningUnsoldClient) UpdateOneID(id uuid.UUID) *MiningUnsoldUpdateOne {
	mutation := newMiningUnsoldMutation(c.config, OpUpdateOne, withMiningUnsoldID(id))
	return &MiningUnsoldUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MiningUnsold.
func (c *MiningUnsoldClient) Delete() *MiningUnsoldDelete {
	mutation := newMiningUnsoldMutation(c.config, OpDelete)
	return &MiningUnsoldDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MiningUnsoldClient) DeleteOne(mu *MiningUnsold) *MiningUnsoldDeleteOne {
	return c.DeleteOneID(mu.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MiningUnsoldClient) DeleteOneID(id uuid.UUID) *MiningUnsoldDeleteOne {
	builder := c.Delete().Where(miningunsold.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MiningUnsoldDeleteOne{builder}
}

// Query returns a query builder for MiningUnsold.
func (c *MiningUnsoldClient) Query() *MiningUnsoldQuery {
	return &MiningUnsoldQuery{
		config: c.config,
	}
}

// Get returns a MiningUnsold entity by its id.
func (c *MiningUnsoldClient) Get(ctx context.Context, id uuid.UUID) (*MiningUnsold, error) {
	return c.Query().Where(miningunsold.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MiningUnsoldClient) GetX(ctx context.Context, id uuid.UUID) *MiningUnsold {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MiningUnsoldClient) Hooks() []Hook {
	hooks := c.hooks.MiningUnsold
	return append(hooks[:len(hooks):len(hooks)], miningunsold.Hooks[:]...)
}

// ProfitClient is a client for the Profit schema.
type ProfitClient struct {
	config
}

// NewProfitClient returns a client for the Profit from the given config.
func NewProfitClient(c config) *ProfitClient {
	return &ProfitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profit.Hooks(f(g(h())))`.
func (c *ProfitClient) Use(hooks ...Hook) {
	c.hooks.Profit = append(c.hooks.Profit, hooks...)
}

// Create returns a builder for creating a Profit entity.
func (c *ProfitClient) Create() *ProfitCreate {
	mutation := newProfitMutation(c.config, OpCreate)
	return &ProfitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Profit entities.
func (c *ProfitClient) CreateBulk(builders ...*ProfitCreate) *ProfitCreateBulk {
	return &ProfitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Profit.
func (c *ProfitClient) Update() *ProfitUpdate {
	mutation := newProfitMutation(c.config, OpUpdate)
	return &ProfitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfitClient) UpdateOne(pr *Profit) *ProfitUpdateOne {
	mutation := newProfitMutation(c.config, OpUpdateOne, withProfit(pr))
	return &ProfitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfitClient) UpdateOneID(id uuid.UUID) *ProfitUpdateOne {
	mutation := newProfitMutation(c.config, OpUpdateOne, withProfitID(id))
	return &ProfitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Profit.
func (c *ProfitClient) Delete() *ProfitDelete {
	mutation := newProfitMutation(c.config, OpDelete)
	return &ProfitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProfitClient) DeleteOne(pr *Profit) *ProfitDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ProfitClient) DeleteOneID(id uuid.UUID) *ProfitDeleteOne {
	builder := c.Delete().Where(profit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfitDeleteOne{builder}
}

// Query returns a query builder for Profit.
func (c *ProfitClient) Query() *ProfitQuery {
	return &ProfitQuery{
		config: c.config,
	}
}

// Get returns a Profit entity by its id.
func (c *ProfitClient) Get(ctx context.Context, id uuid.UUID) (*Profit, error) {
	return c.Query().Where(profit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfitClient) GetX(ctx context.Context, id uuid.UUID) *Profit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ProfitClient) Hooks() []Hook {
	hooks := c.hooks.Profit
	return append(hooks[:len(hooks):len(hooks)], profit.Hooks[:]...)
}

// WithdrawClient is a client for the Withdraw schema.
type WithdrawClient struct {
	config
}

// NewWithdrawClient returns a client for the Withdraw from the given config.
func NewWithdrawClient(c config) *WithdrawClient {
	return &WithdrawClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `withdraw.Hooks(f(g(h())))`.
func (c *WithdrawClient) Use(hooks ...Hook) {
	c.hooks.Withdraw = append(c.hooks.Withdraw, hooks...)
}

// Create returns a builder for creating a Withdraw entity.
func (c *WithdrawClient) Create() *WithdrawCreate {
	mutation := newWithdrawMutation(c.config, OpCreate)
	return &WithdrawCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Withdraw entities.
func (c *WithdrawClient) CreateBulk(builders ...*WithdrawCreate) *WithdrawCreateBulk {
	return &WithdrawCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Withdraw.
func (c *WithdrawClient) Update() *WithdrawUpdate {
	mutation := newWithdrawMutation(c.config, OpUpdate)
	return &WithdrawUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WithdrawClient) UpdateOne(w *Withdraw) *WithdrawUpdateOne {
	mutation := newWithdrawMutation(c.config, OpUpdateOne, withWithdraw(w))
	return &WithdrawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WithdrawClient) UpdateOneID(id uuid.UUID) *WithdrawUpdateOne {
	mutation := newWithdrawMutation(c.config, OpUpdateOne, withWithdrawID(id))
	return &WithdrawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Withdraw.
func (c *WithdrawClient) Delete() *WithdrawDelete {
	mutation := newWithdrawMutation(c.config, OpDelete)
	return &WithdrawDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WithdrawClient) DeleteOne(w *Withdraw) *WithdrawDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *WithdrawClient) DeleteOneID(id uuid.UUID) *WithdrawDeleteOne {
	builder := c.Delete().Where(withdraw.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WithdrawDeleteOne{builder}
}

// Query returns a query builder for Withdraw.
func (c *WithdrawClient) Query() *WithdrawQuery {
	return &WithdrawQuery{
		config: c.config,
	}
}

// Get returns a Withdraw entity by its id.
func (c *WithdrawClient) Get(ctx context.Context, id uuid.UUID) (*Withdraw, error) {
	return c.Query().Where(withdraw.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WithdrawClient) GetX(ctx context.Context, id uuid.UUID) *Withdraw {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WithdrawClient) Hooks() []Hook {
	hooks := c.hooks.Withdraw
	return append(hooks[:len(hooks):len(hooks)], withdraw.Hooks[:]...)
}
