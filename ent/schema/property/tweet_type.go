package property

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
)

type TweetType string

const (
	TweetTypeUnPublic TweetType = "unpublic"
	TweetTypePublic   TweetType = "public"
)

func (TweetType) Values() (kinds []string) {
	for _, s := range []TweetType{TweetTypeUnPublic, TweetTypePublic} {
		kinds = append(kinds, string(s))
	}
	return
}

func (u *TweetType) UnmarshalGQL(v interface{}) error {
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
func (u TweetType) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(u)))
}
