// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/detail"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Detail is the model entity for the Detail schema.
type Detail struct {
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
	// IoType holds the value of the "io_type" field.
	IoType string `json:"io_type,omitempty"`
	// IoSubType holds the value of the "io_sub_type" field.
	IoSubType string `json:"io_sub_type,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// FromCoinTypeID holds the value of the "from_coin_type_id" field.
	FromCoinTypeID uuid.UUID `json:"from_coin_type_id,omitempty"`
	// CoinUsdCurrency holds the value of the "coin_usd_currency" field.
	CoinUsdCurrency *decimal.Decimal `json:"coin_usd_currency,omitempty"`
	// IoExtra holds the value of the "io_extra" field.
	IoExtra string `json:"io_extra,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Detail) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case detail.FieldCoinUsdCurrency:
			values[i] = &sql.NullScanner{S: new(decimal.Decimal)}
		case detail.FieldAmount:
			values[i] = new(decimal.Decimal)
		case detail.FieldCreatedAt, detail.FieldUpdatedAt, detail.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case detail.FieldIoType, detail.FieldIoSubType, detail.FieldIoExtra:
			values[i] = new(sql.NullString)
		case detail.FieldID, detail.FieldAppID, detail.FieldUserID, detail.FieldCoinTypeID, detail.FieldFromCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Detail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Detail fields.
func (d *Detail) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case detail.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case detail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = uint32(value.Int64)
			}
		case detail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = uint32(value.Int64)
			}
		case detail.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				d.DeletedAt = uint32(value.Int64)
			}
		case detail.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				d.AppID = *value
			}
		case detail.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				d.UserID = *value
			}
		case detail.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				d.CoinTypeID = *value
			}
		case detail.FieldIoType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field io_type", values[i])
			} else if value.Valid {
				d.IoType = value.String
			}
		case detail.FieldIoSubType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field io_sub_type", values[i])
			} else if value.Valid {
				d.IoSubType = value.String
			}
		case detail.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				d.Amount = *value
			}
		case detail.FieldFromCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field from_coin_type_id", values[i])
			} else if value != nil {
				d.FromCoinTypeID = *value
			}
		case detail.FieldCoinUsdCurrency:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field coin_usd_currency", values[i])
			} else if value.Valid {
				d.CoinUsdCurrency = new(decimal.Decimal)
				*d.CoinUsdCurrency = *value.S.(*decimal.Decimal)
			}
		case detail.FieldIoExtra:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field io_extra", values[i])
			} else if value.Valid {
				d.IoExtra = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Detail.
// Note that you need to call Detail.Unwrap() before calling this method if this Detail
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Detail) Update() *DetailUpdateOne {
	return (&DetailClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Detail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Detail) Unwrap() *Detail {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Detail is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Detail) String() string {
	var builder strings.Builder
	builder.WriteString("Detail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", d.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", d.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", d.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", d.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", d.UserID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", d.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("io_type=")
	builder.WriteString(d.IoType)
	builder.WriteString(", ")
	builder.WriteString("io_sub_type=")
	builder.WriteString(d.IoSubType)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", d.Amount))
	builder.WriteString(", ")
	builder.WriteString("from_coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", d.FromCoinTypeID))
	builder.WriteString(", ")
	if v := d.CoinUsdCurrency; v != nil {
		builder.WriteString("coin_usd_currency=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("io_extra=")
	builder.WriteString(d.IoExtra)
	builder.WriteByte(')')
	return builder.String()
}

// Details is a parsable slice of Detail.
type Details []*Detail

func (d Details) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
