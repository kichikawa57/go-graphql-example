package auth_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kichikawa/auth"
)

func TestGenerate(t *testing.T) {
	token, tokenErr := auth.Generate(uuid.New(), time.Now())

	if tokenErr != nil {
		t.Fatalf("failed test %s", tokenErr)
	}

	if token == "" {
		t.Fatalf("failed test no token exists.")
	}

	fmt.Println("token", token)
}

func TestValidate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		token, tokenErr := auth.Generate(uuid.New(), time.Now())

		if tokenErr != nil {
			t.Fatalf("failed test %s", tokenErr)
		}

		if token == "" {
			t.Fatalf("failed test no token exists.")
		}

		_, validateErr := auth.Validate(token)

		if validateErr != nil {
			t.Fatalf("failed test %s", validateErr)
		}

		fmt.Println("token", token)
	})

	t.Run("FailToExpirationOfTerm", func(t *testing.T) {
		token, tokenErr := auth.Generate(uuid.New(), time.Now().Add(-31*time.Minute))

		if tokenErr != nil {
			t.Fatalf("failed test %s", tokenErr)
		}

		if token == "" {
			t.Fatalf("failed test no token exists.")
		}

		_, validateErr := auth.Validate(token)

		if validateErr == nil {
			t.Fatalf("failed test %s", validateErr)
		}
	})
}

func TestParse(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		token, tokenErr := auth.Generate(uuid.New(), time.Now())

		if tokenErr != nil {
			t.Fatalf("failed test %s", tokenErr)
		}

		if token == "" {
			t.Fatalf("failed test no token exists.")
		}

		id, validateErr := auth.Parse(token)

		if validateErr != nil {
			t.Fatalf("failed test %s", validateErr)
		}

		fmt.Println("id", id)
	})

	t.Run("FailToExpirationOfTerm", func(t *testing.T) {
		token, tokenErr := auth.Generate(uuid.New(), time.Now().Add(-31*time.Minute))

		if tokenErr != nil {
			t.Fatalf("failed test %s", tokenErr)
		}

		if token == "" {
			t.Fatalf("failed test no token exists.")
		}

		_, validateErr := auth.Parse(token)

		if validateErr == nil {
			t.Fatalf("failed test %s", validateErr)
		}
	})
}
