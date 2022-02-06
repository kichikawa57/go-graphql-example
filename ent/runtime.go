// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/kichikawa/ent/pet"
	"github.com/kichikawa/ent/schema"
	"github.com/kichikawa/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petMixin := schema.Pet{}.Mixin()
	petMixinFields0 := petMixin[0].Fields()
	_ = petMixinFields0
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescCreatedAt is the schema descriptor for created_at field.
	petDescCreatedAt := petMixinFields0[0].Descriptor()
	// pet.DefaultCreatedAt holds the default value on creation for the created_at field.
	pet.DefaultCreatedAt = petDescCreatedAt.Default.(func() time.Time)
	// petDescUpdatedAt is the schema descriptor for updated_at field.
	petDescUpdatedAt := petMixinFields0[1].Descriptor()
	// pet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pet.DefaultUpdatedAt = petDescUpdatedAt.Default.(func() time.Time)
	// pet.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pet.UpdateDefaultUpdatedAt = petDescUpdatedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}