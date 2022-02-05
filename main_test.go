package main_test

import (
	"testing"

	"github.com/kichikawa/test/helper"
)

func TestMain(t *testing.T) {
	helper.Request(t, "test.json")
}
