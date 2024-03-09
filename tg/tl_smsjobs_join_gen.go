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

// SMSJobsJoinRequest represents TL type `smsjobs.join#a74ece2d`.
//
// See https://core.telegram.org/method/smsjobs.join for reference.
type SMSJobsJoinRequest struct {
}

// SMSJobsJoinRequestTypeID is TL type id of SMSJobsJoinRequest.
const SMSJobsJoinRequestTypeID = 0xa74ece2d

// Ensuring interfaces in compile-time for SMSJobsJoinRequest.
var (
	_ bin.Encoder     = &SMSJobsJoinRequest{}
	_ bin.Decoder     = &SMSJobsJoinRequest{}
	_ bin.BareEncoder = &SMSJobsJoinRequest{}
	_ bin.BareDecoder = &SMSJobsJoinRequest{}
)

func (j *SMSJobsJoinRequest) Zero() bool {
	if j == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (j *SMSJobsJoinRequest) String() string {
	if j == nil {
		return "SMSJobsJoinRequest(nil)"
	}
	type Alias SMSJobsJoinRequest
	return fmt.Sprintf("SMSJobsJoinRequest%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SMSJobsJoinRequest) TypeID() uint32 {
	return SMSJobsJoinRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SMSJobsJoinRequest) TypeName() string {
	return "smsjobs.join"
}

// TypeInfo returns info about TL type.
func (j *SMSJobsJoinRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "smsjobs.join",
		ID:   SMSJobsJoinRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (j *SMSJobsJoinRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode smsjobs.join#a74ece2d as nil")
	}
	b.PutID(SMSJobsJoinRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *SMSJobsJoinRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode smsjobs.join#a74ece2d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *SMSJobsJoinRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode smsjobs.join#a74ece2d to nil")
	}
	if err := b.ConsumeID(SMSJobsJoinRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode smsjobs.join#a74ece2d: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *SMSJobsJoinRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode smsjobs.join#a74ece2d to nil")
	}
	return nil
}

// SMSJobsJoin invokes method smsjobs.join#a74ece2d returning error if any.
//
// See https://core.telegram.org/method/smsjobs.join for reference.
func (c *Client) SMSJobsJoin(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &SMSJobsJoinRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
