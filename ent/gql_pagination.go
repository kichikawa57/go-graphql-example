// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/google/uuid"
	"github.com/kichikawa/ent/comment"
	"github.com/kichikawa/ent/follow"
	"github.com/kichikawa/ent/good"
	"github.com/kichikawa/ent/refreshtoken"
	"github.com/kichikawa/ent/tweet"
	"github.com/kichikawa/ent/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    uuid.UUID `msgpack:"i"`
	Value Value     `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// CommentEdge is the edge representation of Comment.
type CommentEdge struct {
	Node   *Comment `json:"node"`
	Cursor Cursor   `json:"cursor"`
}

// CommentConnection is the connection containing edges to Comment.
type CommentConnection struct {
	Edges      []*CommentEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

// CommentPaginateOption enables pagination customization.
type CommentPaginateOption func(*commentPager) error

// WithCommentOrder configures pagination ordering.
func WithCommentOrder(order *CommentOrder) CommentPaginateOption {
	if order == nil {
		order = DefaultCommentOrder
	}
	o := *order
	return func(pager *commentPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCommentOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCommentFilter configures pagination filter.
func WithCommentFilter(filter func(*CommentQuery) (*CommentQuery, error)) CommentPaginateOption {
	return func(pager *commentPager) error {
		if filter == nil {
			return errors.New("CommentQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type commentPager struct {
	order  *CommentOrder
	filter func(*CommentQuery) (*CommentQuery, error)
}

func newCommentPager(opts []CommentPaginateOption) (*commentPager, error) {
	pager := &commentPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCommentOrder
	}
	return pager, nil
}

func (p *commentPager) applyFilter(query *CommentQuery) (*CommentQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *commentPager) toCursor(c *Comment) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *commentPager) applyCursors(query *CommentQuery, after, before *Cursor) *CommentQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultCommentOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *commentPager) applyOrder(query *CommentQuery, reverse bool) *CommentQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultCommentOrder.Field {
		query = query.Order(direction.orderFunc(DefaultCommentOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Comment.
func (c *CommentQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CommentPaginateOption,
) (*CommentConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCommentPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}

	conn := &CommentConnection{Edges: []*CommentEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := c.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := c.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	c = pager.applyCursors(c, after, before)
	c = pager.applyOrder(c, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		c = c.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := c.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Comment
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Comment {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Comment {
			return nodes[i]
		}
	}

	conn.Edges = make([]*CommentEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &CommentEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// CommentOrderField defines the ordering field of Comment.
type CommentOrderField struct {
	field    string
	toCursor func(*Comment) Cursor
}

// CommentOrder defines the ordering of Comment.
type CommentOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *CommentOrderField `json:"field"`
}

// DefaultCommentOrder is the default ordering of Comment.
var DefaultCommentOrder = &CommentOrder{
	Direction: OrderDirectionAsc,
	Field: &CommentOrderField{
		field: comment.FieldID,
		toCursor: func(c *Comment) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Comment into CommentEdge.
func (c *Comment) ToEdge(order *CommentOrder) *CommentEdge {
	if order == nil {
		order = DefaultCommentOrder
	}
	return &CommentEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// FollowEdge is the edge representation of Follow.
type FollowEdge struct {
	Node   *Follow `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// FollowConnection is the connection containing edges to Follow.
type FollowConnection struct {
	Edges      []*FollowEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

// FollowPaginateOption enables pagination customization.
type FollowPaginateOption func(*followPager) error

// WithFollowOrder configures pagination ordering.
func WithFollowOrder(order *FollowOrder) FollowPaginateOption {
	if order == nil {
		order = DefaultFollowOrder
	}
	o := *order
	return func(pager *followPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultFollowOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithFollowFilter configures pagination filter.
func WithFollowFilter(filter func(*FollowQuery) (*FollowQuery, error)) FollowPaginateOption {
	return func(pager *followPager) error {
		if filter == nil {
			return errors.New("FollowQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type followPager struct {
	order  *FollowOrder
	filter func(*FollowQuery) (*FollowQuery, error)
}

func newFollowPager(opts []FollowPaginateOption) (*followPager, error) {
	pager := &followPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultFollowOrder
	}
	return pager, nil
}

func (p *followPager) applyFilter(query *FollowQuery) (*FollowQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *followPager) toCursor(f *Follow) Cursor {
	return p.order.Field.toCursor(f)
}

func (p *followPager) applyCursors(query *FollowQuery, after, before *Cursor) *FollowQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultFollowOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *followPager) applyOrder(query *FollowQuery, reverse bool) *FollowQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultFollowOrder.Field {
		query = query.Order(direction.orderFunc(DefaultFollowOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Follow.
func (f *FollowQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...FollowPaginateOption,
) (*FollowConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newFollowPager(opts)
	if err != nil {
		return nil, err
	}

	if f, err = pager.applyFilter(f); err != nil {
		return nil, err
	}

	conn := &FollowConnection{Edges: []*FollowEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := f.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := f.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	f = pager.applyCursors(f, after, before)
	f = pager.applyOrder(f, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		f = f.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		f = f.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := f.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Follow
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Follow {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Follow {
			return nodes[i]
		}
	}

	conn.Edges = make([]*FollowEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &FollowEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// FollowOrderField defines the ordering field of Follow.
type FollowOrderField struct {
	field    string
	toCursor func(*Follow) Cursor
}

// FollowOrder defines the ordering of Follow.
type FollowOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *FollowOrderField `json:"field"`
}

// DefaultFollowOrder is the default ordering of Follow.
var DefaultFollowOrder = &FollowOrder{
	Direction: OrderDirectionAsc,
	Field: &FollowOrderField{
		field: follow.FieldID,
		toCursor: func(f *Follow) Cursor {
			return Cursor{ID: f.ID}
		},
	},
}

// ToEdge converts Follow into FollowEdge.
func (f *Follow) ToEdge(order *FollowOrder) *FollowEdge {
	if order == nil {
		order = DefaultFollowOrder
	}
	return &FollowEdge{
		Node:   f,
		Cursor: order.Field.toCursor(f),
	}
}

// GoodEdge is the edge representation of Good.
type GoodEdge struct {
	Node   *Good  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// GoodConnection is the connection containing edges to Good.
type GoodConnection struct {
	Edges      []*GoodEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// GoodPaginateOption enables pagination customization.
type GoodPaginateOption func(*goodPager) error

// WithGoodOrder configures pagination ordering.
func WithGoodOrder(order *GoodOrder) GoodPaginateOption {
	if order == nil {
		order = DefaultGoodOrder
	}
	o := *order
	return func(pager *goodPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultGoodOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithGoodFilter configures pagination filter.
func WithGoodFilter(filter func(*GoodQuery) (*GoodQuery, error)) GoodPaginateOption {
	return func(pager *goodPager) error {
		if filter == nil {
			return errors.New("GoodQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type goodPager struct {
	order  *GoodOrder
	filter func(*GoodQuery) (*GoodQuery, error)
}

func newGoodPager(opts []GoodPaginateOption) (*goodPager, error) {
	pager := &goodPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultGoodOrder
	}
	return pager, nil
}

func (p *goodPager) applyFilter(query *GoodQuery) (*GoodQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *goodPager) toCursor(_go *Good) Cursor {
	return p.order.Field.toCursor(_go)
}

func (p *goodPager) applyCursors(query *GoodQuery, after, before *Cursor) *GoodQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultGoodOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *goodPager) applyOrder(query *GoodQuery, reverse bool) *GoodQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultGoodOrder.Field {
		query = query.Order(direction.orderFunc(DefaultGoodOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Good.
func (_go *GoodQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...GoodPaginateOption,
) (*GoodConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newGoodPager(opts)
	if err != nil {
		return nil, err
	}

	if _go, err = pager.applyFilter(_go); err != nil {
		return nil, err
	}

	conn := &GoodConnection{Edges: []*GoodEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := _go.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := _go.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	_go = pager.applyCursors(_go, after, before)
	_go = pager.applyOrder(_go, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		_go = _go.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		_go = _go.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := _go.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Good
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Good {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Good {
			return nodes[i]
		}
	}

	conn.Edges = make([]*GoodEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &GoodEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// GoodOrderField defines the ordering field of Good.
type GoodOrderField struct {
	field    string
	toCursor func(*Good) Cursor
}

// GoodOrder defines the ordering of Good.
type GoodOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *GoodOrderField `json:"field"`
}

// DefaultGoodOrder is the default ordering of Good.
var DefaultGoodOrder = &GoodOrder{
	Direction: OrderDirectionAsc,
	Field: &GoodOrderField{
		field: good.FieldID,
		toCursor: func(_go *Good) Cursor {
			return Cursor{ID: _go.ID}
		},
	},
}

// ToEdge converts Good into GoodEdge.
func (_go *Good) ToEdge(order *GoodOrder) *GoodEdge {
	if order == nil {
		order = DefaultGoodOrder
	}
	return &GoodEdge{
		Node:   _go,
		Cursor: order.Field.toCursor(_go),
	}
}

// RefreshTokenEdge is the edge representation of RefreshToken.
type RefreshTokenEdge struct {
	Node   *RefreshToken `json:"node"`
	Cursor Cursor        `json:"cursor"`
}

// RefreshTokenConnection is the connection containing edges to RefreshToken.
type RefreshTokenConnection struct {
	Edges      []*RefreshTokenEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

// RefreshTokenPaginateOption enables pagination customization.
type RefreshTokenPaginateOption func(*refreshTokenPager) error

// WithRefreshTokenOrder configures pagination ordering.
func WithRefreshTokenOrder(order *RefreshTokenOrder) RefreshTokenPaginateOption {
	if order == nil {
		order = DefaultRefreshTokenOrder
	}
	o := *order
	return func(pager *refreshTokenPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultRefreshTokenOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithRefreshTokenFilter configures pagination filter.
func WithRefreshTokenFilter(filter func(*RefreshTokenQuery) (*RefreshTokenQuery, error)) RefreshTokenPaginateOption {
	return func(pager *refreshTokenPager) error {
		if filter == nil {
			return errors.New("RefreshTokenQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type refreshTokenPager struct {
	order  *RefreshTokenOrder
	filter func(*RefreshTokenQuery) (*RefreshTokenQuery, error)
}

func newRefreshTokenPager(opts []RefreshTokenPaginateOption) (*refreshTokenPager, error) {
	pager := &refreshTokenPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultRefreshTokenOrder
	}
	return pager, nil
}

func (p *refreshTokenPager) applyFilter(query *RefreshTokenQuery) (*RefreshTokenQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *refreshTokenPager) toCursor(rt *RefreshToken) Cursor {
	return p.order.Field.toCursor(rt)
}

func (p *refreshTokenPager) applyCursors(query *RefreshTokenQuery, after, before *Cursor) *RefreshTokenQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultRefreshTokenOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *refreshTokenPager) applyOrder(query *RefreshTokenQuery, reverse bool) *RefreshTokenQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultRefreshTokenOrder.Field {
		query = query.Order(direction.orderFunc(DefaultRefreshTokenOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to RefreshToken.
func (rt *RefreshTokenQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...RefreshTokenPaginateOption,
) (*RefreshTokenConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newRefreshTokenPager(opts)
	if err != nil {
		return nil, err
	}

	if rt, err = pager.applyFilter(rt); err != nil {
		return nil, err
	}

	conn := &RefreshTokenConnection{Edges: []*RefreshTokenEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := rt.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := rt.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	rt = pager.applyCursors(rt, after, before)
	rt = pager.applyOrder(rt, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		rt = rt.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		rt = rt.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := rt.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *RefreshToken
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *RefreshToken {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *RefreshToken {
			return nodes[i]
		}
	}

	conn.Edges = make([]*RefreshTokenEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &RefreshTokenEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// RefreshTokenOrderField defines the ordering field of RefreshToken.
type RefreshTokenOrderField struct {
	field    string
	toCursor func(*RefreshToken) Cursor
}

// RefreshTokenOrder defines the ordering of RefreshToken.
type RefreshTokenOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *RefreshTokenOrderField `json:"field"`
}

// DefaultRefreshTokenOrder is the default ordering of RefreshToken.
var DefaultRefreshTokenOrder = &RefreshTokenOrder{
	Direction: OrderDirectionAsc,
	Field: &RefreshTokenOrderField{
		field: refreshtoken.FieldID,
		toCursor: func(rt *RefreshToken) Cursor {
			return Cursor{ID: rt.ID}
		},
	},
}

// ToEdge converts RefreshToken into RefreshTokenEdge.
func (rt *RefreshToken) ToEdge(order *RefreshTokenOrder) *RefreshTokenEdge {
	if order == nil {
		order = DefaultRefreshTokenOrder
	}
	return &RefreshTokenEdge{
		Node:   rt,
		Cursor: order.Field.toCursor(rt),
	}
}

// TweetEdge is the edge representation of Tweet.
type TweetEdge struct {
	Node   *Tweet `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// TweetConnection is the connection containing edges to Tweet.
type TweetConnection struct {
	Edges      []*TweetEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

// TweetPaginateOption enables pagination customization.
type TweetPaginateOption func(*tweetPager) error

// WithTweetOrder configures pagination ordering.
func WithTweetOrder(order *TweetOrder) TweetPaginateOption {
	if order == nil {
		order = DefaultTweetOrder
	}
	o := *order
	return func(pager *tweetPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultTweetOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithTweetFilter configures pagination filter.
func WithTweetFilter(filter func(*TweetQuery) (*TweetQuery, error)) TweetPaginateOption {
	return func(pager *tweetPager) error {
		if filter == nil {
			return errors.New("TweetQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type tweetPager struct {
	order  *TweetOrder
	filter func(*TweetQuery) (*TweetQuery, error)
}

func newTweetPager(opts []TweetPaginateOption) (*tweetPager, error) {
	pager := &tweetPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultTweetOrder
	}
	return pager, nil
}

func (p *tweetPager) applyFilter(query *TweetQuery) (*TweetQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *tweetPager) toCursor(t *Tweet) Cursor {
	return p.order.Field.toCursor(t)
}

func (p *tweetPager) applyCursors(query *TweetQuery, after, before *Cursor) *TweetQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultTweetOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *tweetPager) applyOrder(query *TweetQuery, reverse bool) *TweetQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultTweetOrder.Field {
		query = query.Order(direction.orderFunc(DefaultTweetOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Tweet.
func (t *TweetQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TweetPaginateOption,
) (*TweetConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTweetPager(opts)
	if err != nil {
		return nil, err
	}

	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}

	conn := &TweetConnection{Edges: []*TweetEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := t.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := t.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	t = pager.applyCursors(t, after, before)
	t = pager.applyOrder(t, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		t = t.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := t.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Tweet
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Tweet {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Tweet {
			return nodes[i]
		}
	}

	conn.Edges = make([]*TweetEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &TweetEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// TweetOrderField defines the ordering field of Tweet.
type TweetOrderField struct {
	field    string
	toCursor func(*Tweet) Cursor
}

// TweetOrder defines the ordering of Tweet.
type TweetOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *TweetOrderField `json:"field"`
}

// DefaultTweetOrder is the default ordering of Tweet.
var DefaultTweetOrder = &TweetOrder{
	Direction: OrderDirectionAsc,
	Field: &TweetOrderField{
		field: tweet.FieldID,
		toCursor: func(t *Tweet) Cursor {
			return Cursor{ID: t.ID}
		},
	},
}

// ToEdge converts Tweet into TweetEdge.
func (t *Tweet) ToEdge(order *TweetOrder) *TweetEdge {
	if order == nil {
		order = DefaultTweetOrder
	}
	return &TweetEdge{
		Node:   t,
		Cursor: order.Field.toCursor(t),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}

	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}

	conn := &UserConnection{Edges: []*UserEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := u.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := u.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		u = u.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := u.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}

	conn.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
