// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kichikawa/ent/comment"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// TweetID holds the value of the "tweet_id" field.
	TweetID uuid.UUID `json:"tweet_id,omitempty"`
	// Text holds the value of the "text" field.
	Text     string `json:"text,omitempty"`
	tweet_id *uuid.UUID
	user_id  *uuid.UUID
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldText:
			values[i] = new(sql.NullString)
		case comment.FieldCreatedAt, comment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case comment.FieldID, comment.FieldUserID, comment.FieldTweetID:
			values[i] = new(uuid.UUID)
		case comment.ForeignKeys[0]: // tweet_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case comment.ForeignKeys[1]: // user_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case comment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case comment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case comment.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				c.UserID = *value
			}
		case comment.FieldTweetID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field tweet_id", values[i])
			} else if value != nil {
				c.TweetID = *value
			}
		case comment.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				c.Text = value.String
			}
		case comment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field tweet_id", values[i])
			} else if value.Valid {
				c.tweet_id = new(uuid.UUID)
				*c.tweet_id = *value.S.(*uuid.UUID)
			}
		case comment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.user_id = new(uuid.UUID)
				*c.user_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteString(", tweet_id=")
	builder.WriteString(fmt.Sprintf("%v", c.TweetID))
	builder.WriteString(", text=")
	builder.WriteString(c.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
