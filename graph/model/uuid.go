package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
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
	uuid, _ := uuid.FromBytes(Bytes(u))
	_, _ = io.WriteString(w, strconv.Quote(uuid.String()))
}

func Bytes(bytes [16]byte) []byte {
	strs := []byte{}
	for _, b := range bytes {
		strs = append(strs, b)
	}
	return strs
}
