// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// StatsGetStoryStatsRequest represents TL type `stats.getStoryStats#374fef40`.
//
// See https://core.telegram.org/method/stats.getStoryStats for reference.
type StatsGetStoryStatsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Dark field of StatsGetStoryStatsRequest.
	Dark bool
	// Peer field of StatsGetStoryStatsRequest.
	Peer InputPeerClass
	// ID field of StatsGetStoryStatsRequest.
	ID int
}

// StatsGetStoryStatsRequestTypeID is TL type id of StatsGetStoryStatsRequest.
const StatsGetStoryStatsRequestTypeID = 0x374fef40

// Ensuring interfaces in compile-time for StatsGetStoryStatsRequest.
var (
	_ bin.Encoder     = &StatsGetStoryStatsRequest{}
	_ bin.Decoder     = &StatsGetStoryStatsRequest{}
	_ bin.BareEncoder = &StatsGetStoryStatsRequest{}
	_ bin.BareDecoder = &StatsGetStoryStatsRequest{}
)

func (g *StatsGetStoryStatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Dark == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetStoryStatsRequest) String() string {
	if g == nil {
		return "StatsGetStoryStatsRequest(nil)"
	}
	type Alias StatsGetStoryStatsRequest
	return fmt.Sprintf("StatsGetStoryStatsRequest%+v", Alias(*g))
}

// FillFrom fills StatsGetStoryStatsRequest from given interface.
func (g *StatsGetStoryStatsRequest) FillFrom(from interface {
	GetDark() (value bool)
	GetPeer() (value InputPeerClass)
	GetID() (value int)
}) {
	g.Dark = from.GetDark()
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGetStoryStatsRequest) TypeID() uint32 {
	return StatsGetStoryStatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGetStoryStatsRequest) TypeName() string {
	return "stats.getStoryStats"
}

// TypeInfo returns info about TL type.
func (g *StatsGetStoryStatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.getStoryStats",
		ID:   StatsGetStoryStatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dark",
			SchemaName: "dark",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *StatsGetStoryStatsRequest) SetFlags() {
	if !(g.Dark == false) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *StatsGetStoryStatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getStoryStats#374fef40 as nil")
	}
	b.PutID(StatsGetStoryStatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StatsGetStoryStatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getStoryStats#374fef40 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getStoryStats#374fef40: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stats.getStoryStats#374fef40: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getStoryStats#374fef40: field peer: %w", err)
	}
	b.PutInt(g.ID)
	return nil
}

// Decode implements bin.Decoder.
func (g *StatsGetStoryStatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getStoryStats#374fef40 to nil")
	}
	if err := b.ConsumeID(StatsGetStoryStatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getStoryStats#374fef40: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StatsGetStoryStatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getStoryStats#374fef40 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.getStoryStats#374fef40: field flags: %w", err)
		}
	}
	g.Dark = g.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getStoryStats#374fef40: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getStoryStats#374fef40: field id: %w", err)
		}
		g.ID = value
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (g *StatsGetStoryStatsRequest) SetDark(value bool) {
	if value {
		g.Flags.Set(0)
		g.Dark = true
	} else {
		g.Flags.Unset(0)
		g.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (g *StatsGetStoryStatsRequest) GetDark() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (g *StatsGetStoryStatsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *StatsGetStoryStatsRequest) GetID() (value int) {
	if g == nil {
		return
	}
	return g.ID
}

// StatsGetStoryStats invokes method stats.getStoryStats#374fef40 returning error if any.
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/stats.getStoryStats for reference.
func (c *Client) StatsGetStoryStats(ctx context.Context, request *StatsGetStoryStatsRequest) (*StatsStoryStats, error) {
	var result StatsStoryStats

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
