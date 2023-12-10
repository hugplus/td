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

// PremiumGetBoostsListRequest represents TL type `premium.getBoostsList#60f67660`.
//
// See https://core.telegram.org/method/premium.getBoostsList for reference.
type PremiumGetBoostsListRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Gifts field of PremiumGetBoostsListRequest.
	Gifts bool
	// Peer field of PremiumGetBoostsListRequest.
	Peer InputPeerClass
	// Offset field of PremiumGetBoostsListRequest.
	Offset string
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// PremiumGetBoostsListRequestTypeID is TL type id of PremiumGetBoostsListRequest.
const PremiumGetBoostsListRequestTypeID = 0x60f67660

// Ensuring interfaces in compile-time for PremiumGetBoostsListRequest.
var (
	_ bin.Encoder     = &PremiumGetBoostsListRequest{}
	_ bin.Decoder     = &PremiumGetBoostsListRequest{}
	_ bin.BareEncoder = &PremiumGetBoostsListRequest{}
	_ bin.BareDecoder = &PremiumGetBoostsListRequest{}
)

func (g *PremiumGetBoostsListRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Gifts == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PremiumGetBoostsListRequest) String() string {
	if g == nil {
		return "PremiumGetBoostsListRequest(nil)"
	}
	type Alias PremiumGetBoostsListRequest
	return fmt.Sprintf("PremiumGetBoostsListRequest%+v", Alias(*g))
}

// FillFrom fills PremiumGetBoostsListRequest from given interface.
func (g *PremiumGetBoostsListRequest) FillFrom(from interface {
	GetGifts() (value bool)
	GetPeer() (value InputPeerClass)
	GetOffset() (value string)
	GetLimit() (value int)
}) {
	g.Gifts = from.GetGifts()
	g.Peer = from.GetPeer()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumGetBoostsListRequest) TypeID() uint32 {
	return PremiumGetBoostsListRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumGetBoostsListRequest) TypeName() string {
	return "premium.getBoostsList"
}

// TypeInfo returns info about TL type.
func (g *PremiumGetBoostsListRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premium.getBoostsList",
		ID:   PremiumGetBoostsListRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Gifts",
			SchemaName: "gifts",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *PremiumGetBoostsListRequest) SetFlags() {
	if !(g.Gifts == false) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *PremiumGetBoostsListRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode premium.getBoostsList#60f67660 as nil")
	}
	b.PutID(PremiumGetBoostsListRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PremiumGetBoostsListRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode premium.getBoostsList#60f67660 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premium.getBoostsList#60f67660: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode premium.getBoostsList#60f67660: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premium.getBoostsList#60f67660: field peer: %w", err)
	}
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *PremiumGetBoostsListRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode premium.getBoostsList#60f67660 to nil")
	}
	if err := b.ConsumeID(PremiumGetBoostsListRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode premium.getBoostsList#60f67660: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PremiumGetBoostsListRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode premium.getBoostsList#60f67660 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode premium.getBoostsList#60f67660: field flags: %w", err)
		}
	}
	g.Gifts = g.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode premium.getBoostsList#60f67660: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode premium.getBoostsList#60f67660: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode premium.getBoostsList#60f67660: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// SetGifts sets value of Gifts conditional field.
func (g *PremiumGetBoostsListRequest) SetGifts(value bool) {
	if value {
		g.Flags.Set(0)
		g.Gifts = true
	} else {
		g.Flags.Unset(0)
		g.Gifts = false
	}
}

// GetGifts returns value of Gifts conditional field.
func (g *PremiumGetBoostsListRequest) GetGifts() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (g *PremiumGetBoostsListRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetOffset returns value of Offset field.
func (g *PremiumGetBoostsListRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *PremiumGetBoostsListRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// PremiumGetBoostsList invokes method premium.getBoostsList#60f67660 returning error if any.
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/premium.getBoostsList for reference.
func (c *Client) PremiumGetBoostsList(ctx context.Context, request *PremiumGetBoostsListRequest) (*PremiumBoostsList, error) {
	var result PremiumBoostsList

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
