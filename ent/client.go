// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/kichikawa/ent/migrate"

	"github.com/kichikawa/ent/comment"
	"github.com/kichikawa/ent/follow"
	"github.com/kichikawa/ent/good"
	"github.com/kichikawa/ent/refreshtoken"
	"github.com/kichikawa/ent/tweet"
	"github.com/kichikawa/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// Follow is the client for interacting with the Follow builders.
	Follow *FollowClient
	// Good is the client for interacting with the Good builders.
	Good *GoodClient
	// RefreshToken is the client for interacting with the RefreshToken builders.
	RefreshToken *RefreshTokenClient
	// Tweet is the client for interacting with the Tweet builders.
	Tweet *TweetClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Comment = NewCommentClient(c.config)
	c.Follow = NewFollowClient(c.config)
	c.Good = NewGoodClient(c.config)
	c.RefreshToken = NewRefreshTokenClient(c.config)
	c.Tweet = NewTweetClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Comment:      NewCommentClient(cfg),
		Follow:       NewFollowClient(cfg),
		Good:         NewGoodClient(cfg),
		RefreshToken: NewRefreshTokenClient(cfg),
		Tweet:        NewTweetClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Comment:      NewCommentClient(cfg),
		Follow:       NewFollowClient(cfg),
		Good:         NewGoodClient(cfg),
		RefreshToken: NewRefreshTokenClient(cfg),
		Tweet:        NewTweetClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Comment.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Comment.Use(hooks...)
	c.Follow.Use(hooks...)
	c.Good.Use(hooks...)
	c.RefreshToken.Use(hooks...)
	c.Tweet.Use(hooks...)
	c.User.Use(hooks...)
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Create returns a create builder for Comment.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id uuid.UUID) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CommentClient) DeleteOneID(id uuid.UUID) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id uuid.UUID) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id uuid.UUID) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// FollowClient is a client for the Follow schema.
type FollowClient struct {
	config
}

// NewFollowClient returns a client for the Follow from the given config.
func NewFollowClient(c config) *FollowClient {
	return &FollowClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `follow.Hooks(f(g(h())))`.
func (c *FollowClient) Use(hooks ...Hook) {
	c.hooks.Follow = append(c.hooks.Follow, hooks...)
}

// Create returns a create builder for Follow.
func (c *FollowClient) Create() *FollowCreate {
	mutation := newFollowMutation(c.config, OpCreate)
	return &FollowCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Follow entities.
func (c *FollowClient) CreateBulk(builders ...*FollowCreate) *FollowCreateBulk {
	return &FollowCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Follow.
func (c *FollowClient) Update() *FollowUpdate {
	mutation := newFollowMutation(c.config, OpUpdate)
	return &FollowUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FollowClient) UpdateOne(f *Follow) *FollowUpdateOne {
	mutation := newFollowMutation(c.config, OpUpdateOne, withFollow(f))
	return &FollowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FollowClient) UpdateOneID(id uuid.UUID) *FollowUpdateOne {
	mutation := newFollowMutation(c.config, OpUpdateOne, withFollowID(id))
	return &FollowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Follow.
func (c *FollowClient) Delete() *FollowDelete {
	mutation := newFollowMutation(c.config, OpDelete)
	return &FollowDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FollowClient) DeleteOne(f *Follow) *FollowDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FollowClient) DeleteOneID(id uuid.UUID) *FollowDeleteOne {
	builder := c.Delete().Where(follow.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FollowDeleteOne{builder}
}

// Query returns a query builder for Follow.
func (c *FollowClient) Query() *FollowQuery {
	return &FollowQuery{
		config: c.config,
	}
}

// Get returns a Follow entity by its id.
func (c *FollowClient) Get(ctx context.Context, id uuid.UUID) (*Follow, error) {
	return c.Query().Where(follow.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FollowClient) GetX(ctx context.Context, id uuid.UUID) *Follow {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FollowClient) Hooks() []Hook {
	return c.hooks.Follow
}

// GoodClient is a client for the Good schema.
type GoodClient struct {
	config
}

// NewGoodClient returns a client for the Good from the given config.
func NewGoodClient(c config) *GoodClient {
	return &GoodClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `good.Hooks(f(g(h())))`.
func (c *GoodClient) Use(hooks ...Hook) {
	c.hooks.Good = append(c.hooks.Good, hooks...)
}

// Create returns a create builder for Good.
func (c *GoodClient) Create() *GoodCreate {
	mutation := newGoodMutation(c.config, OpCreate)
	return &GoodCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Good entities.
func (c *GoodClient) CreateBulk(builders ...*GoodCreate) *GoodCreateBulk {
	return &GoodCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Good.
func (c *GoodClient) Update() *GoodUpdate {
	mutation := newGoodMutation(c.config, OpUpdate)
	return &GoodUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GoodClient) UpdateOne(_go *Good) *GoodUpdateOne {
	mutation := newGoodMutation(c.config, OpUpdateOne, withGood(_go))
	return &GoodUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GoodClient) UpdateOneID(id uuid.UUID) *GoodUpdateOne {
	mutation := newGoodMutation(c.config, OpUpdateOne, withGoodID(id))
	return &GoodUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Good.
func (c *GoodClient) Delete() *GoodDelete {
	mutation := newGoodMutation(c.config, OpDelete)
	return &GoodDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GoodClient) DeleteOne(_go *Good) *GoodDeleteOne {
	return c.DeleteOneID(_go.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GoodClient) DeleteOneID(id uuid.UUID) *GoodDeleteOne {
	builder := c.Delete().Where(good.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GoodDeleteOne{builder}
}

// Query returns a query builder for Good.
func (c *GoodClient) Query() *GoodQuery {
	return &GoodQuery{
		config: c.config,
	}
}

// Get returns a Good entity by its id.
func (c *GoodClient) Get(ctx context.Context, id uuid.UUID) (*Good, error) {
	return c.Query().Where(good.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GoodClient) GetX(ctx context.Context, id uuid.UUID) *Good {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GoodClient) Hooks() []Hook {
	return c.hooks.Good
}

// RefreshTokenClient is a client for the RefreshToken schema.
type RefreshTokenClient struct {
	config
}

// NewRefreshTokenClient returns a client for the RefreshToken from the given config.
func NewRefreshTokenClient(c config) *RefreshTokenClient {
	return &RefreshTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `refreshtoken.Hooks(f(g(h())))`.
func (c *RefreshTokenClient) Use(hooks ...Hook) {
	c.hooks.RefreshToken = append(c.hooks.RefreshToken, hooks...)
}

// Create returns a create builder for RefreshToken.
func (c *RefreshTokenClient) Create() *RefreshTokenCreate {
	mutation := newRefreshTokenMutation(c.config, OpCreate)
	return &RefreshTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RefreshToken entities.
func (c *RefreshTokenClient) CreateBulk(builders ...*RefreshTokenCreate) *RefreshTokenCreateBulk {
	return &RefreshTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RefreshToken.
func (c *RefreshTokenClient) Update() *RefreshTokenUpdate {
	mutation := newRefreshTokenMutation(c.config, OpUpdate)
	return &RefreshTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RefreshTokenClient) UpdateOne(rt *RefreshToken) *RefreshTokenUpdateOne {
	mutation := newRefreshTokenMutation(c.config, OpUpdateOne, withRefreshToken(rt))
	return &RefreshTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RefreshTokenClient) UpdateOneID(id uuid.UUID) *RefreshTokenUpdateOne {
	mutation := newRefreshTokenMutation(c.config, OpUpdateOne, withRefreshTokenID(id))
	return &RefreshTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RefreshToken.
func (c *RefreshTokenClient) Delete() *RefreshTokenDelete {
	mutation := newRefreshTokenMutation(c.config, OpDelete)
	return &RefreshTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RefreshTokenClient) DeleteOne(rt *RefreshToken) *RefreshTokenDeleteOne {
	return c.DeleteOneID(rt.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RefreshTokenClient) DeleteOneID(id uuid.UUID) *RefreshTokenDeleteOne {
	builder := c.Delete().Where(refreshtoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RefreshTokenDeleteOne{builder}
}

// Query returns a query builder for RefreshToken.
func (c *RefreshTokenClient) Query() *RefreshTokenQuery {
	return &RefreshTokenQuery{
		config: c.config,
	}
}

// Get returns a RefreshToken entity by its id.
func (c *RefreshTokenClient) Get(ctx context.Context, id uuid.UUID) (*RefreshToken, error) {
	return c.Query().Where(refreshtoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RefreshTokenClient) GetX(ctx context.Context, id uuid.UUID) *RefreshToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RefreshTokenClient) Hooks() []Hook {
	return c.hooks.RefreshToken
}

// TweetClient is a client for the Tweet schema.
type TweetClient struct {
	config
}

// NewTweetClient returns a client for the Tweet from the given config.
func NewTweetClient(c config) *TweetClient {
	return &TweetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tweet.Hooks(f(g(h())))`.
func (c *TweetClient) Use(hooks ...Hook) {
	c.hooks.Tweet = append(c.hooks.Tweet, hooks...)
}

// Create returns a create builder for Tweet.
func (c *TweetClient) Create() *TweetCreate {
	mutation := newTweetMutation(c.config, OpCreate)
	return &TweetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tweet entities.
func (c *TweetClient) CreateBulk(builders ...*TweetCreate) *TweetCreateBulk {
	return &TweetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tweet.
func (c *TweetClient) Update() *TweetUpdate {
	mutation := newTweetMutation(c.config, OpUpdate)
	return &TweetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TweetClient) UpdateOne(t *Tweet) *TweetUpdateOne {
	mutation := newTweetMutation(c.config, OpUpdateOne, withTweet(t))
	return &TweetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TweetClient) UpdateOneID(id uuid.UUID) *TweetUpdateOne {
	mutation := newTweetMutation(c.config, OpUpdateOne, withTweetID(id))
	return &TweetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tweet.
func (c *TweetClient) Delete() *TweetDelete {
	mutation := newTweetMutation(c.config, OpDelete)
	return &TweetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TweetClient) DeleteOne(t *Tweet) *TweetDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TweetClient) DeleteOneID(id uuid.UUID) *TweetDeleteOne {
	builder := c.Delete().Where(tweet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TweetDeleteOne{builder}
}

// Query returns a query builder for Tweet.
func (c *TweetClient) Query() *TweetQuery {
	return &TweetQuery{
		config: c.config,
	}
}

// Get returns a Tweet entity by its id.
func (c *TweetClient) Get(ctx context.Context, id uuid.UUID) (*Tweet, error) {
	return c.Query().Where(tweet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TweetClient) GetX(ctx context.Context, id uuid.UUID) *Tweet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGoods queries the goods edge of a Tweet.
func (c *TweetClient) QueryGoods(t *Tweet) *GoodQuery {
	query := &GoodQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, id),
			sqlgraph.To(good.Table, good.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tweet.GoodsTable, tweet.GoodsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a Tweet.
func (c *TweetClient) QueryComments(t *Tweet) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tweet.CommentsTable, tweet.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TweetClient) Hooks() []Hook {
	return c.hooks.Tweet
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTweet queries the tweet edge of a User.
func (c *UserClient) QueryTweet(u *User) *TweetQuery {
	query := &TweetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TweetTable, user.TweetColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGood queries the good edge of a User.
func (c *UserClient) QueryGood(u *User) *GoodQuery {
	query := &GoodQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(good.Table, good.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.GoodTable, user.GoodColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComment queries the comment edge of a User.
func (c *UserClient) QueryComment(u *User) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CommentTable, user.CommentColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollower queries the follower edge of a User.
func (c *UserClient) QueryFollower(u *User) *FollowQuery {
	query := &FollowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(follow.Table, follow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FollowerTable, user.FollowerColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollowed queries the followed edge of a User.
func (c *UserClient) QueryFollowed(u *User) *FollowQuery {
	query := &FollowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(follow.Table, follow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FollowedTable, user.FollowedColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
