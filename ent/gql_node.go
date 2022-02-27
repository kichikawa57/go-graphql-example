// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"github.com/kichikawa/ent/comment"
	"github.com/kichikawa/ent/follow"
	"github.com/kichikawa/ent/good"
	"github.com/kichikawa/ent/refreshtoken"
	"github.com/kichikawa/ent/tweet"
	"github.com/kichikawa/ent/user"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     uuid.UUID `json:"id,omitempty"`     // node id.
	Type   string    `json:"type,omitempty"`   // node type.
	Fields []*Field  `json:"fields,omitempty"` // node fields.
	Edges  []*Edge   `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string      `json:"type,omitempty"` // edge type.
	Name string      `json:"name,omitempty"` // edge name.
	IDs  []uuid.UUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (c *Comment) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Comment",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(c.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Text); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "text",
		Value: string(buf),
	}
	return node, nil
}

func (f *Follow) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     f.ID,
		Type:   "Follow",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(f.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.FollowerID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "uuid.UUID",
		Name:  "follower_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.FollowedID); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "followed_id",
		Value: string(buf),
	}
	return node, nil
}

func (_go *Good) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     _go.ID,
		Type:   "Good",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(_go.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(_go.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(_go.UserID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "uuid.UUID",
		Name:  "user_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(_go.TweetID); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "tweet_id",
		Value: string(buf),
	}
	return node, nil
}

func (rt *RefreshToken) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     rt.ID,
		Type:   "RefreshToken",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(rt.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(rt.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(rt.Token); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "token",
		Value: string(buf),
	}
	if buf, err = json.Marshal(rt.ExpiresAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "expires_at",
		Value: string(buf),
	}
	return node, nil
}

func (t *Tweet) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     t.ID,
		Type:   "Tweet",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(t.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.UserID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "uuid.UUID",
		Name:  "user_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Text); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "text",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Type); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "property.TweetType",
		Name:  "type",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Good",
		Name: "goods",
	}
	err = t.QueryGoods().
		Select(good.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Comment",
		Name: "comments",
	}
	err = t.QueryComments().
		Select(comment.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 5),
	}
	var buf []byte
	if buf, err = json.Marshal(u.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.AccountName); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "property.UserAccountName",
		Name:  "account_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Email); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "property.UserEmail",
		Name:  "email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Password); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "password",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Age); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "age",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Tweet",
		Name: "tweet",
	}
	err = u.QueryTweet().
		Select(tweet.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Good",
		Name: "good",
	}
	err = u.QueryGood().
		Select(good.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Comment",
		Name: "comment",
	}
	err = u.QueryComment().
		Select(comment.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Follow",
		Name: "follower",
	}
	err = u.QueryFollower().
		Select(follow.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "Follow",
		Name: "followed",
	}
	err = u.QueryFollowed().
		Select(follow.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id uuid.UUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, uuid.UUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, uuid.UUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, uuid.UUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id uuid.UUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id uuid.UUID) (Noder, error) {
	switch table {
	case comment.Table:
		n, err := c.Comment.Query().
			Where(comment.ID(id)).
			CollectFields(ctx, "Comment").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case follow.Table:
		n, err := c.Follow.Query().
			Where(follow.ID(id)).
			CollectFields(ctx, "Follow").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case good.Table:
		n, err := c.Good.Query().
			Where(good.ID(id)).
			CollectFields(ctx, "Good").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case refreshtoken.Table:
		n, err := c.RefreshToken.Query().
			Where(refreshtoken.ID(id)).
			CollectFields(ctx, "RefreshToken").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case tweet.Table:
		n, err := c.Tweet.Query().
			Where(tweet.ID(id)).
			CollectFields(ctx, "Tweet").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		n, err := c.User.Query().
			Where(user.ID(id)).
			CollectFields(ctx, "User").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]uuid.UUID)
	id2idx := make(map[uuid.UUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []uuid.UUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[uuid.UUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case comment.Table:
		nodes, err := c.Comment.Query().
			Where(comment.IDIn(ids...)).
			CollectFields(ctx, "Comment").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case follow.Table:
		nodes, err := c.Follow.Query().
			Where(follow.IDIn(ids...)).
			CollectFields(ctx, "Follow").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case good.Table:
		nodes, err := c.Good.Query().
			Where(good.IDIn(ids...)).
			CollectFields(ctx, "Good").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case refreshtoken.Table:
		nodes, err := c.RefreshToken.Query().
			Where(refreshtoken.IDIn(ids...)).
			CollectFields(ctx, "RefreshToken").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case tweet.Table:
		nodes, err := c.Tweet.Query().
			Where(tweet.IDIn(ids...)).
			CollectFields(ctx, "Tweet").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		nodes, err := c.User.Query().
			Where(user.IDIn(ids...)).
			CollectFields(ctx, "User").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
