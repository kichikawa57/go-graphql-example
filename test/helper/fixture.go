package helper

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kichikawa/ent"
	"github.com/kichikawa/ent/schema/property"
)

func UserFixture(db DB) {
	bulks := make([]*ent.UserCreate, 5)

	for i := 0; i < 5; i++ {
		bulks[i] = db.Client.User.
			Create().
			SetName(property.UserName(fmt.Sprintf("name %d", i))).
			SetEmail(property.UserEmail(fmt.Sprintf("hoge%d@email", i))).
			SetAge(10).
			SetPassword("password").
			SetAccountName(property.UserAccountName(fmt.Sprintf("hoge%d", i)))
	}

	db.Client.User.CreateBulk(bulks...).Save(db.Ctx)
}

func TweetFixture(db DB) {
	bulks := make([]*ent.TweetCreate, 5)

	for i := 0; i < 5; i++ {
		bulks[i] = db.Client.Tweet.
			Create().
			SetText(fmt.Sprintf("text%d", i)).
			SetType(property.TweetTypePublic).
			SetUserID(uuid.New())
	}

	db.Client.Tweet.CreateBulk(bulks...).Save(db.Ctx)
}

func GoodFixture(db DB) {
	bulks := make([]*ent.GoodCreate, 5)

	for i := 0; i < 5; i++ {
		bulks[i] = db.Client.Good.
			Create().
			SetTweetID(uuid.New()).
			SetUserID(uuid.New())
	}

	db.Client.Good.CreateBulk(bulks...).Save(db.Ctx)
}

func CommentFixture(db DB) {
	bulks := make([]*ent.CommentCreate, 5)

	for i := 0; i < 5; i++ {
		bulks[i] = db.Client.Comment.
			Create().
			SetTweetID(uuid.New()).
			SetUserID(uuid.New()).
			SetText(fmt.Sprintf("text%d", i))
	}

	db.Client.Comment.CreateBulk(bulks...).Save(db.Ctx)
}

func FollowFixture(db DB) {
	bulks := make([]*ent.FollowCreate, 5)

	for i := 0; i < 5; i++ {
		bulks[i] = db.Client.Follow.
			Create().
			SetFollowedID(uuid.New()).
			SetFollowerID(uuid.New())
	}

	db.Client.Follow.CreateBulk(bulks...).Save(db.Ctx)
}

func Fixture(db DB) {
	UserFixture(db)
	TweetFixture(db)
	GoodFixture(db)
	CommentFixture(db)
	FollowFixture(db)
}
