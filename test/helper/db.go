package helper

import (
	"context"

	"github.com/kichikawa/ent"
)

type DB struct {
	Ctx    context.Context
	Client *ent.Client
}

type DBRepository interface {
	DeleteAll()
	TeadownFixture()
}

func (d DB) TeadownFixture() {
	d.DeleteAll()
	Fixture(d)
}

func (d DB) DeleteAll() {
	d.Client.User.Delete().Exec(d.Ctx)
	d.Client.Comment.Delete().Exec(d.Ctx)
	d.Client.Tweet.Delete().Exec(d.Ctx)
	d.Client.RefreshToken.Delete().Exec(d.Ctx)
	d.Client.Good.Delete().Exec(d.Ctx)
	d.Client.Follow.Delete().Exec(d.Ctx)
}
