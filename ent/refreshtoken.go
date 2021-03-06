// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kichikawa/ent/refreshtoken"
)

// RefreshToken is the model entity for the RefreshToken schema.
type RefreshToken struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RefreshToken) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case refreshtoken.FieldToken:
			values[i] = new(sql.NullString)
		case refreshtoken.FieldCreatedAt, refreshtoken.FieldUpdatedAt, refreshtoken.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		case refreshtoken.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RefreshToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RefreshToken fields.
func (rt *RefreshToken) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case refreshtoken.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rt.ID = *value
			}
		case refreshtoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rt.CreatedAt = value.Time
			}
		case refreshtoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rt.UpdatedAt = value.Time
			}
		case refreshtoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				rt.Token = value.String
			}
		case refreshtoken.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				rt.ExpiresAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RefreshToken.
// Note that you need to call RefreshToken.Unwrap() before calling this method if this RefreshToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *RefreshToken) Update() *RefreshTokenUpdateOne {
	return (&RefreshTokenClient{config: rt.config}).UpdateOne(rt)
}

// Unwrap unwraps the RefreshToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rt *RefreshToken) Unwrap() *RefreshToken {
	tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("ent: RefreshToken is not a transactional entity")
	}
	rt.config.driver = tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *RefreshToken) String() string {
	var builder strings.Builder
	builder.WriteString("RefreshToken(")
	builder.WriteString(fmt.Sprintf("id=%v", rt.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(rt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(rt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", token=")
	builder.WriteString(rt.Token)
	builder.WriteString(", expires_at=")
	builder.WriteString(rt.ExpiresAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RefreshTokens is a parsable slice of RefreshToken.
type RefreshTokens []*RefreshToken

func (rt RefreshTokens) config(cfg config) {
	for _i := range rt {
		rt[_i].config = cfg
	}
}
