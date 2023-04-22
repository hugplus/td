// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// SetBotInfoShortDescriptionRequest represents TL type `setBotInfoShortDescription#3a96bae3`.
type SetBotInfoShortDescriptionRequest struct {
	// Identifier of the target bot
	BotUserID int64
	// A two-letter ISO 639-1 language code. If empty, the short description will be shown to
	// all users, for which language there are no dedicated description
	LanguageCode string
	// New bot's short description on the specified language
	ShortDescription string
}

// SetBotInfoShortDescriptionRequestTypeID is TL type id of SetBotInfoShortDescriptionRequest.
const SetBotInfoShortDescriptionRequestTypeID = 0x3a96bae3

// Ensuring interfaces in compile-time for SetBotInfoShortDescriptionRequest.
var (
	_ bin.Encoder     = &SetBotInfoShortDescriptionRequest{}
	_ bin.Decoder     = &SetBotInfoShortDescriptionRequest{}
	_ bin.BareEncoder = &SetBotInfoShortDescriptionRequest{}
	_ bin.BareDecoder = &SetBotInfoShortDescriptionRequest{}
)

func (s *SetBotInfoShortDescriptionRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.BotUserID == 0) {
		return false
	}
	if !(s.LanguageCode == "") {
		return false
	}
	if !(s.ShortDescription == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBotInfoShortDescriptionRequest) String() string {
	if s == nil {
		return "SetBotInfoShortDescriptionRequest(nil)"
	}
	type Alias SetBotInfoShortDescriptionRequest
	return fmt.Sprintf("SetBotInfoShortDescriptionRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBotInfoShortDescriptionRequest) TypeID() uint32 {
	return SetBotInfoShortDescriptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBotInfoShortDescriptionRequest) TypeName() string {
	return "setBotInfoShortDescription"
}

// TypeInfo returns info about TL type.
func (s *SetBotInfoShortDescriptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBotInfoShortDescription",
		ID:   SetBotInfoShortDescriptionRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
		{
			Name:       "ShortDescription",
			SchemaName: "short_description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBotInfoShortDescriptionRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotInfoShortDescription#3a96bae3 as nil")
	}
	b.PutID(SetBotInfoShortDescriptionRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBotInfoShortDescriptionRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotInfoShortDescription#3a96bae3 as nil")
	}
	b.PutInt53(s.BotUserID)
	b.PutString(s.LanguageCode)
	b.PutString(s.ShortDescription)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBotInfoShortDescriptionRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotInfoShortDescription#3a96bae3 to nil")
	}
	if err := b.ConsumeID(SetBotInfoShortDescriptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBotInfoShortDescription#3a96bae3: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBotInfoShortDescriptionRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotInfoShortDescription#3a96bae3 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setBotInfoShortDescription#3a96bae3: field bot_user_id: %w", err)
		}
		s.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setBotInfoShortDescription#3a96bae3: field language_code: %w", err)
		}
		s.LanguageCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setBotInfoShortDescription#3a96bae3: field short_description: %w", err)
		}
		s.ShortDescription = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBotInfoShortDescriptionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotInfoShortDescription#3a96bae3 as nil")
	}
	b.ObjStart()
	b.PutID("setBotInfoShortDescription")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(s.BotUserID)
	b.Comma()
	b.FieldStart("language_code")
	b.PutString(s.LanguageCode)
	b.Comma()
	b.FieldStart("short_description")
	b.PutString(s.ShortDescription)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBotInfoShortDescriptionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotInfoShortDescription#3a96bae3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBotInfoShortDescription"); err != nil {
				return fmt.Errorf("unable to decode setBotInfoShortDescription#3a96bae3: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setBotInfoShortDescription#3a96bae3: field bot_user_id: %w", err)
			}
			s.BotUserID = value
		case "language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setBotInfoShortDescription#3a96bae3: field language_code: %w", err)
			}
			s.LanguageCode = value
		case "short_description":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setBotInfoShortDescription#3a96bae3: field short_description: %w", err)
			}
			s.ShortDescription = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (s *SetBotInfoShortDescriptionRequest) GetBotUserID() (value int64) {
	if s == nil {
		return
	}
	return s.BotUserID
}

// GetLanguageCode returns value of LanguageCode field.
func (s *SetBotInfoShortDescriptionRequest) GetLanguageCode() (value string) {
	if s == nil {
		return
	}
	return s.LanguageCode
}

// GetShortDescription returns value of ShortDescription field.
func (s *SetBotInfoShortDescriptionRequest) GetShortDescription() (value string) {
	if s == nil {
		return
	}
	return s.ShortDescription
}

// SetBotInfoShortDescription invokes method setBotInfoShortDescription#3a96bae3 returning error if any.
func (c *Client) SetBotInfoShortDescription(ctx context.Context, request *SetBotInfoShortDescriptionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
