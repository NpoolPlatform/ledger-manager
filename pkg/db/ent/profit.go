// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/profit"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Profit is the model entity for the Profit schema.
type Profit struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Incoming holds the value of the "incoming" field.
	Incoming *decimal.Decimal `json:"incoming,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case profit.FieldIncoming:
			values[i] = &sql.NullScanner{S: new(decimal.Decimal)}
		case profit.FieldCreatedAt, profit.FieldUpdatedAt, profit.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case profit.FieldID, profit.FieldAppID, profit.FieldUserID, profit.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Profit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profit fields.
func (pr *Profit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profit.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case profit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = uint32(value.Int64)
			}
		case profit.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = uint32(value.Int64)
			}
		case profit.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = uint32(value.Int64)
			}
		case profit.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				pr.AppID = *value
			}
		case profit.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				pr.UserID = *value
			}
		case profit.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				pr.CoinTypeID = *value
			}
		case profit.FieldIncoming:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field incoming", values[i])
			} else if value.Valid {
				pr.Incoming = new(decimal.Decimal)
				*pr.Incoming = *value.S.(*decimal.Decimal)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Profit.
// Note that you need to call Profit.Unwrap() before calling this method if this Profit
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profit) Update() *ProfitUpdateOne {
	return (&ProfitClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Profit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profit) Unwrap() *Profit {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profit is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profit) String() string {
	var builder strings.Builder
	builder.WriteString("Profit(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.DeletedAt))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.UserID))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.CoinTypeID))
	if v := pr.Incoming; v != nil {
		builder.WriteString(", incoming=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Profits is a parsable slice of Profit.
type Profits []*Profit

func (pr Profits) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
