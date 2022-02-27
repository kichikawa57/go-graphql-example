// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kichikawa/ent/follow"
)

// Follow is the model entity for the Follow schema.
type Follow struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FollowerID holds the value of the "follower_id" field.
	FollowerID uuid.UUID `json:"follower_id,omitempty"`
	// FollowedID holds the value of the "followed_id" field.
	FollowedID  uuid.UUID `json:"followed_id,omitempty"`
	follower_id *uuid.UUID
	followed_id *uuid.UUID
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Follow) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case follow.FieldCreatedAt, follow.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case follow.FieldID, follow.FieldFollowerID, follow.FieldFollowedID:
			values[i] = new(uuid.UUID)
		case follow.ForeignKeys[0]: // follower_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case follow.ForeignKeys[1]: // followed_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Follow", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Follow fields.
func (f *Follow) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case follow.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case follow.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case follow.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case follow.FieldFollowerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field follower_id", values[i])
			} else if value != nil {
				f.FollowerID = *value
			}
		case follow.FieldFollowedID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field followed_id", values[i])
			} else if value != nil {
				f.FollowedID = *value
			}
		case follow.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field follower_id", values[i])
			} else if value.Valid {
				f.follower_id = new(uuid.UUID)
				*f.follower_id = *value.S.(*uuid.UUID)
			}
		case follow.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field followed_id", values[i])
			} else if value.Valid {
				f.followed_id = new(uuid.UUID)
				*f.followed_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Follow.
// Note that you need to call Follow.Unwrap() before calling this method if this Follow
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Follow) Update() *FollowUpdateOne {
	return (&FollowClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Follow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Follow) Unwrap() *Follow {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Follow is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Follow) String() string {
	var builder strings.Builder
	builder.WriteString("Follow(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", follower_id=")
	builder.WriteString(fmt.Sprintf("%v", f.FollowerID))
	builder.WriteString(", followed_id=")
	builder.WriteString(fmt.Sprintf("%v", f.FollowedID))
	builder.WriteByte(')')
	return builder.String()
}

// Follows is a parsable slice of Follow.
type Follows []*Follow

func (f Follows) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}