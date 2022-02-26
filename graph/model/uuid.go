package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
	"github.com/kichikawa/shared"
)

type UUID uuid.UUID

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *UUID) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("uuid must be string")
	}

	if _, err := uuid.Parse(str); err != nil {
		return fmt.Errorf("not in uuid format: %w", err)
	}

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (u UUID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(shared.ConvertUUIDToString(uuid.UUID(u))))
}
