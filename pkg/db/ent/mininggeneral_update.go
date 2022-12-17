// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/mininggeneral"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// MiningGeneralUpdate is the builder for updating MiningGeneral entities.
type MiningGeneralUpdate struct {
	config
	hooks     []Hook
	mutation  *MiningGeneralMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MiningGeneralUpdate builder.
func (mgu *MiningGeneralUpdate) Where(ps ...predicate.MiningGeneral) *MiningGeneralUpdate {
	mgu.mutation.Where(ps...)
	return mgu
}

// SetCreatedAt sets the "created_at" field.
func (mgu *MiningGeneralUpdate) SetCreatedAt(u uint32) *MiningGeneralUpdate {
	mgu.mutation.ResetCreatedAt()
	mgu.mutation.SetCreatedAt(u)
	return mgu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mgu *MiningGeneralUpdate) SetNillableCreatedAt(u *uint32) *MiningGeneralUpdate {
	if u != nil {
		mgu.SetCreatedAt(*u)
	}
	return mgu
}

// AddCreatedAt adds u to the "created_at" field.
func (mgu *MiningGeneralUpdate) AddCreatedAt(u int32) *MiningGeneralUpdate {
	mgu.mutation.AddCreatedAt(u)
	return mgu
}

// SetUpdatedAt sets the "updated_at" field.
func (mgu *MiningGeneralUpdate) SetUpdatedAt(u uint32) *MiningGeneralUpdate {
	mgu.mutation.ResetUpdatedAt()
	mgu.mutation.SetUpdatedAt(u)
	return mgu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (mgu *MiningGeneralUpdate) AddUpdatedAt(u int32) *MiningGeneralUpdate {
	mgu.mutation.AddUpdatedAt(u)
	return mgu
}

// SetDeletedAt sets the "deleted_at" field.
func (mgu *MiningGeneralUpdate) SetDeletedAt(u uint32) *MiningGeneralUpdate {
	mgu.mutation.ResetDeletedAt()
	mgu.mutation.SetDeletedAt(u)
	return mgu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mgu *MiningGeneralUpdate) SetNillableDeletedAt(u *uint32) *MiningGeneralUpdate {
	if u != nil {
		mgu.SetDeletedAt(*u)
	}
	return mgu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (mgu *MiningGeneralUpdate) AddDeletedAt(u int32) *MiningGeneralUpdate {
	mgu.mutation.AddDeletedAt(u)
	return mgu
}

// SetGoodID sets the "good_id" field.
func (mgu *MiningGeneralUpdate) SetGoodID(u uuid.UUID) *MiningGeneralUpdate {
	mgu.mutation.SetGoodID(u)
	return mgu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (mgu *MiningGeneralUpdate) SetNillableGoodID(u *uuid.UUID) *MiningGeneralUpdate {
	if u != nil {
		mgu.SetGoodID(*u)
	}
	return mgu
}

// ClearGoodID clears the value of the "good_id" field.
func (mgu *MiningGeneralUpdate) ClearGoodID() *MiningGeneralUpdate {
	mgu.mutation.ClearGoodID()
	return mgu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (mgu *MiningGeneralUpdate) SetCoinTypeID(u uuid.UUID) *MiningGeneralUpdate {
	mgu.mutation.SetCoinTypeID(u)
	return mgu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (mgu *MiningGeneralUpdate) SetNillableCoinTypeID(u *uuid.UUID) *MiningGeneralUpdate {
	if u != nil {
		mgu.SetCoinTypeID(*u)
	}
	return mgu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (mgu *MiningGeneralUpdate) ClearCoinTypeID() *MiningGeneralUpdate {
	mgu.mutation.ClearCoinTypeID()
	return mgu
}

// SetAmount sets the "amount" field.
func (mgu *MiningGeneralUpdate) SetAmount(d decimal.Decimal) *MiningGeneralUpdate {
	mgu.mutation.ResetAmount()
	mgu.mutation.SetAmount(d)
	return mgu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (mgu *MiningGeneralUpdate) SetNillableAmount(d *decimal.Decimal) *MiningGeneralUpdate {
	if d != nil {
		mgu.SetAmount(*d)
	}
	return mgu
}

// AddAmount adds d to the "amount" field.
func (mgu *MiningGeneralUpdate) AddAmount(d decimal.Decimal) *MiningGeneralUpdate {
	mgu.mutation.AddAmount(d)
	return mgu
}

// ClearAmount clears the value of the "amount" field.
func (mgu *MiningGeneralUpdate) ClearAmount() *MiningGeneralUpdate {
	mgu.mutation.ClearAmount()
	return mgu
}

// SetToPlatform sets the "to_platform" field.
func (mgu *MiningGeneralUpdate) SetToPlatform(d decimal.Decimal) *MiningGeneralUpdate {
	mgu.mutation.ResetToPlatform()
	mgu.mutation.SetToPlatform(d)
	return mgu
}

// SetNillableToPlatform sets the "to_platform" field if the given value is not nil.
func (mgu *MiningGeneralUpdate) SetNillableToPlatform(d *decimal.Decimal) *MiningGeneralUpdate {
	if d != nil {
		mgu.SetToPlatform(*d)
	}
	return mgu
}

// AddToPlatform adds d to the "to_platform" field.
func (mgu *MiningGeneralUpdate) AddToPlatform(d decimal.Decimal) *MiningGeneralUpdate {
	mgu.mutation.AddToPlatform(d)
	return mgu
}

// ClearToPlatform clears the value of the "to_platform" field.
func (mgu *MiningGeneralUpdate) ClearToPlatform() *MiningGeneralUpdate {
	mgu.mutation.ClearToPlatform()
	return mgu
}

// SetToUser sets the "to_user" field.
func (mgu *MiningGeneralUpdate) SetToUser(d decimal.Decimal) *MiningGeneralUpdate {
	mgu.mutation.ResetToUser()
	mgu.mutation.SetToUser(d)
	return mgu
}

// SetNillableToUser sets the "to_user" field if the given value is not nil.
func (mgu *MiningGeneralUpdate) SetNillableToUser(d *decimal.Decimal) *MiningGeneralUpdate {
	if d != nil {
		mgu.SetToUser(*d)
	}
	return mgu
}

// AddToUser adds d to the "to_user" field.
func (mgu *MiningGeneralUpdate) AddToUser(d decimal.Decimal) *MiningGeneralUpdate {
	mgu.mutation.AddToUser(d)
	return mgu
}

// ClearToUser clears the value of the "to_user" field.
func (mgu *MiningGeneralUpdate) ClearToUser() *MiningGeneralUpdate {
	mgu.mutation.ClearToUser()
	return mgu
}

// Mutation returns the MiningGeneralMutation object of the builder.
func (mgu *MiningGeneralUpdate) Mutation() *MiningGeneralMutation {
	return mgu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mgu *MiningGeneralUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := mgu.defaults(); err != nil {
		return 0, err
	}
	if len(mgu.hooks) == 0 {
		affected, err = mgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiningGeneralMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mgu.mutation = mutation
			affected, err = mgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mgu.hooks) - 1; i >= 0; i-- {
			if mgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mgu *MiningGeneralUpdate) SaveX(ctx context.Context) int {
	affected, err := mgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mgu *MiningGeneralUpdate) Exec(ctx context.Context) error {
	_, err := mgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mgu *MiningGeneralUpdate) ExecX(ctx context.Context) {
	if err := mgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mgu *MiningGeneralUpdate) defaults() error {
	if _, ok := mgu.mutation.UpdatedAt(); !ok {
		if mininggeneral.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized mininggeneral.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := mininggeneral.UpdateDefaultUpdatedAt()
		mgu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mgu *MiningGeneralUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MiningGeneralUpdate {
	mgu.modifiers = append(mgu.modifiers, modifiers...)
	return mgu
}

func (mgu *MiningGeneralUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mininggeneral.Table,
			Columns: mininggeneral.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mininggeneral.FieldID,
			},
		},
	}
	if ps := mgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mgu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldCreatedAt,
		})
	}
	if value, ok := mgu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldCreatedAt,
		})
	}
	if value, ok := mgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldUpdatedAt,
		})
	}
	if value, ok := mgu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldUpdatedAt,
		})
	}
	if value, ok := mgu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldDeletedAt,
		})
	}
	if value, ok := mgu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldDeletedAt,
		})
	}
	if value, ok := mgu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mininggeneral.FieldGoodID,
		})
	}
	if mgu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: mininggeneral.FieldGoodID,
		})
	}
	if value, ok := mgu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mininggeneral.FieldCoinTypeID,
		})
	}
	if mgu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: mininggeneral.FieldCoinTypeID,
		})
	}
	if value, ok := mgu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldAmount,
		})
	}
	if value, ok := mgu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldAmount,
		})
	}
	if mgu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: mininggeneral.FieldAmount,
		})
	}
	if value, ok := mgu.mutation.ToPlatform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToPlatform,
		})
	}
	if value, ok := mgu.mutation.AddedToPlatform(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToPlatform,
		})
	}
	if mgu.mutation.ToPlatformCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: mininggeneral.FieldToPlatform,
		})
	}
	if value, ok := mgu.mutation.ToUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToUser,
		})
	}
	if value, ok := mgu.mutation.AddedToUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToUser,
		})
	}
	if mgu.mutation.ToUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: mininggeneral.FieldToUser,
		})
	}
	_spec.Modifiers = mgu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, mgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mininggeneral.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MiningGeneralUpdateOne is the builder for updating a single MiningGeneral entity.
type MiningGeneralUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MiningGeneralMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (mguo *MiningGeneralUpdateOne) SetCreatedAt(u uint32) *MiningGeneralUpdateOne {
	mguo.mutation.ResetCreatedAt()
	mguo.mutation.SetCreatedAt(u)
	return mguo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mguo *MiningGeneralUpdateOne) SetNillableCreatedAt(u *uint32) *MiningGeneralUpdateOne {
	if u != nil {
		mguo.SetCreatedAt(*u)
	}
	return mguo
}

// AddCreatedAt adds u to the "created_at" field.
func (mguo *MiningGeneralUpdateOne) AddCreatedAt(u int32) *MiningGeneralUpdateOne {
	mguo.mutation.AddCreatedAt(u)
	return mguo
}

// SetUpdatedAt sets the "updated_at" field.
func (mguo *MiningGeneralUpdateOne) SetUpdatedAt(u uint32) *MiningGeneralUpdateOne {
	mguo.mutation.ResetUpdatedAt()
	mguo.mutation.SetUpdatedAt(u)
	return mguo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (mguo *MiningGeneralUpdateOne) AddUpdatedAt(u int32) *MiningGeneralUpdateOne {
	mguo.mutation.AddUpdatedAt(u)
	return mguo
}

// SetDeletedAt sets the "deleted_at" field.
func (mguo *MiningGeneralUpdateOne) SetDeletedAt(u uint32) *MiningGeneralUpdateOne {
	mguo.mutation.ResetDeletedAt()
	mguo.mutation.SetDeletedAt(u)
	return mguo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mguo *MiningGeneralUpdateOne) SetNillableDeletedAt(u *uint32) *MiningGeneralUpdateOne {
	if u != nil {
		mguo.SetDeletedAt(*u)
	}
	return mguo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (mguo *MiningGeneralUpdateOne) AddDeletedAt(u int32) *MiningGeneralUpdateOne {
	mguo.mutation.AddDeletedAt(u)
	return mguo
}

// SetGoodID sets the "good_id" field.
func (mguo *MiningGeneralUpdateOne) SetGoodID(u uuid.UUID) *MiningGeneralUpdateOne {
	mguo.mutation.SetGoodID(u)
	return mguo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (mguo *MiningGeneralUpdateOne) SetNillableGoodID(u *uuid.UUID) *MiningGeneralUpdateOne {
	if u != nil {
		mguo.SetGoodID(*u)
	}
	return mguo
}

// ClearGoodID clears the value of the "good_id" field.
func (mguo *MiningGeneralUpdateOne) ClearGoodID() *MiningGeneralUpdateOne {
	mguo.mutation.ClearGoodID()
	return mguo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (mguo *MiningGeneralUpdateOne) SetCoinTypeID(u uuid.UUID) *MiningGeneralUpdateOne {
	mguo.mutation.SetCoinTypeID(u)
	return mguo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (mguo *MiningGeneralUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *MiningGeneralUpdateOne {
	if u != nil {
		mguo.SetCoinTypeID(*u)
	}
	return mguo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (mguo *MiningGeneralUpdateOne) ClearCoinTypeID() *MiningGeneralUpdateOne {
	mguo.mutation.ClearCoinTypeID()
	return mguo
}

// SetAmount sets the "amount" field.
func (mguo *MiningGeneralUpdateOne) SetAmount(d decimal.Decimal) *MiningGeneralUpdateOne {
	mguo.mutation.ResetAmount()
	mguo.mutation.SetAmount(d)
	return mguo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (mguo *MiningGeneralUpdateOne) SetNillableAmount(d *decimal.Decimal) *MiningGeneralUpdateOne {
	if d != nil {
		mguo.SetAmount(*d)
	}
	return mguo
}

// AddAmount adds d to the "amount" field.
func (mguo *MiningGeneralUpdateOne) AddAmount(d decimal.Decimal) *MiningGeneralUpdateOne {
	mguo.mutation.AddAmount(d)
	return mguo
}

// ClearAmount clears the value of the "amount" field.
func (mguo *MiningGeneralUpdateOne) ClearAmount() *MiningGeneralUpdateOne {
	mguo.mutation.ClearAmount()
	return mguo
}

// SetToPlatform sets the "to_platform" field.
func (mguo *MiningGeneralUpdateOne) SetToPlatform(d decimal.Decimal) *MiningGeneralUpdateOne {
	mguo.mutation.ResetToPlatform()
	mguo.mutation.SetToPlatform(d)
	return mguo
}

// SetNillableToPlatform sets the "to_platform" field if the given value is not nil.
func (mguo *MiningGeneralUpdateOne) SetNillableToPlatform(d *decimal.Decimal) *MiningGeneralUpdateOne {
	if d != nil {
		mguo.SetToPlatform(*d)
	}
	return mguo
}

// AddToPlatform adds d to the "to_platform" field.
func (mguo *MiningGeneralUpdateOne) AddToPlatform(d decimal.Decimal) *MiningGeneralUpdateOne {
	mguo.mutation.AddToPlatform(d)
	return mguo
}

// ClearToPlatform clears the value of the "to_platform" field.
func (mguo *MiningGeneralUpdateOne) ClearToPlatform() *MiningGeneralUpdateOne {
	mguo.mutation.ClearToPlatform()
	return mguo
}

// SetToUser sets the "to_user" field.
func (mguo *MiningGeneralUpdateOne) SetToUser(d decimal.Decimal) *MiningGeneralUpdateOne {
	mguo.mutation.ResetToUser()
	mguo.mutation.SetToUser(d)
	return mguo
}

// SetNillableToUser sets the "to_user" field if the given value is not nil.
func (mguo *MiningGeneralUpdateOne) SetNillableToUser(d *decimal.Decimal) *MiningGeneralUpdateOne {
	if d != nil {
		mguo.SetToUser(*d)
	}
	return mguo
}

// AddToUser adds d to the "to_user" field.
func (mguo *MiningGeneralUpdateOne) AddToUser(d decimal.Decimal) *MiningGeneralUpdateOne {
	mguo.mutation.AddToUser(d)
	return mguo
}

// ClearToUser clears the value of the "to_user" field.
func (mguo *MiningGeneralUpdateOne) ClearToUser() *MiningGeneralUpdateOne {
	mguo.mutation.ClearToUser()
	return mguo
}

// Mutation returns the MiningGeneralMutation object of the builder.
func (mguo *MiningGeneralUpdateOne) Mutation() *MiningGeneralMutation {
	return mguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mguo *MiningGeneralUpdateOne) Select(field string, fields ...string) *MiningGeneralUpdateOne {
	mguo.fields = append([]string{field}, fields...)
	return mguo
}

// Save executes the query and returns the updated MiningGeneral entity.
func (mguo *MiningGeneralUpdateOne) Save(ctx context.Context) (*MiningGeneral, error) {
	var (
		err  error
		node *MiningGeneral
	)
	if err := mguo.defaults(); err != nil {
		return nil, err
	}
	if len(mguo.hooks) == 0 {
		node, err = mguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiningGeneralMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mguo.mutation = mutation
			node, err = mguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mguo.hooks) - 1; i >= 0; i-- {
			if mguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MiningGeneral)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MiningGeneralMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mguo *MiningGeneralUpdateOne) SaveX(ctx context.Context) *MiningGeneral {
	node, err := mguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mguo *MiningGeneralUpdateOne) Exec(ctx context.Context) error {
	_, err := mguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mguo *MiningGeneralUpdateOne) ExecX(ctx context.Context) {
	if err := mguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mguo *MiningGeneralUpdateOne) defaults() error {
	if _, ok := mguo.mutation.UpdatedAt(); !ok {
		if mininggeneral.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized mininggeneral.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := mininggeneral.UpdateDefaultUpdatedAt()
		mguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mguo *MiningGeneralUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MiningGeneralUpdateOne {
	mguo.modifiers = append(mguo.modifiers, modifiers...)
	return mguo
}

func (mguo *MiningGeneralUpdateOne) sqlSave(ctx context.Context) (_node *MiningGeneral, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mininggeneral.Table,
			Columns: mininggeneral.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mininggeneral.FieldID,
			},
		},
	}
	id, ok := mguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MiningGeneral.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mininggeneral.FieldID)
		for _, f := range fields {
			if !mininggeneral.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mininggeneral.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldCreatedAt,
		})
	}
	if value, ok := mguo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldCreatedAt,
		})
	}
	if value, ok := mguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldUpdatedAt,
		})
	}
	if value, ok := mguo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldUpdatedAt,
		})
	}
	if value, ok := mguo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldDeletedAt,
		})
	}
	if value, ok := mguo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldDeletedAt,
		})
	}
	if value, ok := mguo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mininggeneral.FieldGoodID,
		})
	}
	if mguo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: mininggeneral.FieldGoodID,
		})
	}
	if value, ok := mguo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mininggeneral.FieldCoinTypeID,
		})
	}
	if mguo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: mininggeneral.FieldCoinTypeID,
		})
	}
	if value, ok := mguo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldAmount,
		})
	}
	if value, ok := mguo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldAmount,
		})
	}
	if mguo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: mininggeneral.FieldAmount,
		})
	}
	if value, ok := mguo.mutation.ToPlatform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToPlatform,
		})
	}
	if value, ok := mguo.mutation.AddedToPlatform(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToPlatform,
		})
	}
	if mguo.mutation.ToPlatformCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: mininggeneral.FieldToPlatform,
		})
	}
	if value, ok := mguo.mutation.ToUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToUser,
		})
	}
	if value, ok := mguo.mutation.AddedToUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToUser,
		})
	}
	if mguo.mutation.ToUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: mininggeneral.FieldToUser,
		})
	}
	_spec.Modifiers = mguo.modifiers
	_node = &MiningGeneral{config: mguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mininggeneral.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
