// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/mininggeneral"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// MiningGeneral is the model entity for the MiningGeneral schema.
type MiningGeneral struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// ToPlatform holds the value of the "to_platform" field.
	ToPlatform decimal.Decimal `json:"to_platform,omitempty"`
	// ToUser holds the value of the "to_user" field.
	ToUser decimal.Decimal `json:"to_user,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MiningGeneral) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mininggeneral.FieldAmount, mininggeneral.FieldToPlatform, mininggeneral.FieldToUser:
			values[i] = new(decimal.Decimal)
		case mininggeneral.FieldCreatedAt, mininggeneral.FieldUpdatedAt, mininggeneral.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case mininggeneral.FieldID, mininggeneral.FieldGoodID, mininggeneral.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MiningGeneral", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MiningGeneral fields.
func (mg *MiningGeneral) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mininggeneral.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				mg.ID = *value
			}
		case mininggeneral.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mg.CreatedAt = uint32(value.Int64)
			}
		case mininggeneral.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mg.UpdatedAt = uint32(value.Int64)
			}
		case mininggeneral.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mg.DeletedAt = uint32(value.Int64)
			}
		case mininggeneral.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				mg.GoodID = *value
			}
		case mininggeneral.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				mg.CoinTypeID = *value
			}
		case mininggeneral.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				mg.Amount = *value
			}
		case mininggeneral.FieldToPlatform:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field to_platform", values[i])
			} else if value != nil {
				mg.ToPlatform = *value
			}
		case mininggeneral.FieldToUser:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field to_user", values[i])
			} else if value != nil {
				mg.ToUser = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MiningGeneral.
// Note that you need to call MiningGeneral.Unwrap() before calling this method if this MiningGeneral
// was returned from a transaction, and the transaction was committed or rolled back.
func (mg *MiningGeneral) Update() *MiningGeneralUpdateOne {
	return (&MiningGeneralClient{config: mg.config}).UpdateOne(mg)
}

// Unwrap unwraps the MiningGeneral entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mg *MiningGeneral) Unwrap() *MiningGeneral {
	_tx, ok := mg.config.driver.(*txDriver)
	if !ok {
		panic("ent: MiningGeneral is not a transactional entity")
	}
	mg.config.driver = _tx.drv
	return mg
}

// String implements the fmt.Stringer.
func (mg *MiningGeneral) String() string {
	var builder strings.Builder
	builder.WriteString("MiningGeneral(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mg.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", mg.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", mg.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", mg.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", mg.GoodID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", mg.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", mg.Amount))
	builder.WriteString(", ")
	builder.WriteString("to_platform=")
	builder.WriteString(fmt.Sprintf("%v", mg.ToPlatform))
	builder.WriteString(", ")
	builder.WriteString("to_user=")
	builder.WriteString(fmt.Sprintf("%v", mg.ToUser))
	builder.WriteByte(')')
	return builder.String()
}

// MiningGenerals is a parsable slice of MiningGeneral.
type MiningGenerals []*MiningGeneral

func (mg MiningGenerals) config(cfg config) {
	for _i := range mg {
		mg[_i].config = cfg
	}
}
