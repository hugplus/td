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

// InputReplyToMessage represents TL type `inputReplyToMessage#22c0f6d5`.
// Reply to a message.
//
// See https://core.telegram.org/constructor/inputReplyToMessage for reference.
type InputReplyToMessage struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The message ID to reply to.
	ReplyToMsgID int
	// This field must contain the topic ID only when replying to messages in forum topics
	// different from the "General" topic (i.e. reply_to_msg_id is set and reply_to_msg_id !=
	// topicID and topicID != 1).  If the replied-to message is deleted before the method
	// finishes execution, the value in this field will be used to send the message to the
	// correct topic, instead of the "General" topic.
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
	// ReplyToPeerID field of InputReplyToMessage.
	//
	// Use SetReplyToPeerID and GetReplyToPeerID helpers.
	ReplyToPeerID InputPeerClass
	// QuoteText field of InputReplyToMessage.
	//
	// Use SetQuoteText and GetQuoteText helpers.
	QuoteText string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetQuoteEntities and GetQuoteEntities helpers.
	QuoteEntities []MessageEntityClass
	// QuoteOffset field of InputReplyToMessage.
	//
	// Use SetQuoteOffset and GetQuoteOffset helpers.
	QuoteOffset int
}

// InputReplyToMessageTypeID is TL type id of InputReplyToMessage.
const InputReplyToMessageTypeID = 0x22c0f6d5

// construct implements constructor of InputReplyToClass.
func (i InputReplyToMessage) construct() InputReplyToClass { return &i }

// Ensuring interfaces in compile-time for InputReplyToMessage.
var (
	_ bin.Encoder     = &InputReplyToMessage{}
	_ bin.Decoder     = &InputReplyToMessage{}
	_ bin.BareEncoder = &InputReplyToMessage{}
	_ bin.BareDecoder = &InputReplyToMessage{}

	_ InputReplyToClass = &InputReplyToMessage{}
)

func (i *InputReplyToMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.ReplyToMsgID == 0) {
		return false
	}
	if !(i.TopMsgID == 0) {
		return false
	}
	if !(i.ReplyToPeerID == nil) {
		return false
	}
	if !(i.QuoteText == "") {
		return false
	}
	if !(i.QuoteEntities == nil) {
		return false
	}
	if !(i.QuoteOffset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReplyToMessage) String() string {
	if i == nil {
		return "InputReplyToMessage(nil)"
	}
	type Alias InputReplyToMessage
	return fmt.Sprintf("InputReplyToMessage%+v", Alias(*i))
}

// FillFrom fills InputReplyToMessage from given interface.
func (i *InputReplyToMessage) FillFrom(from interface {
	GetReplyToMsgID() (value int)
	GetTopMsgID() (value int, ok bool)
	GetReplyToPeerID() (value InputPeerClass, ok bool)
	GetQuoteText() (value string, ok bool)
	GetQuoteEntities() (value []MessageEntityClass, ok bool)
	GetQuoteOffset() (value int, ok bool)
}) {
	i.ReplyToMsgID = from.GetReplyToMsgID()
	if val, ok := from.GetTopMsgID(); ok {
		i.TopMsgID = val
	}

	if val, ok := from.GetReplyToPeerID(); ok {
		i.ReplyToPeerID = val
	}

	if val, ok := from.GetQuoteText(); ok {
		i.QuoteText = val
	}

	if val, ok := from.GetQuoteEntities(); ok {
		i.QuoteEntities = val
	}

	if val, ok := from.GetQuoteOffset(); ok {
		i.QuoteOffset = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReplyToMessage) TypeID() uint32 {
	return InputReplyToMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReplyToMessage) TypeName() string {
	return "inputReplyToMessage"
}

// TypeInfo returns info about TL type.
func (i *InputReplyToMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReplyToMessage",
		ID:   InputReplyToMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReplyToMsgID",
			SchemaName: "reply_to_msg_id",
		},
		{
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "ReplyToPeerID",
			SchemaName: "reply_to_peer_id",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "QuoteText",
			SchemaName: "quote_text",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "QuoteEntities",
			SchemaName: "quote_entities",
			Null:       !i.Flags.Has(3),
		},
		{
			Name:       "QuoteOffset",
			SchemaName: "quote_offset",
			Null:       !i.Flags.Has(4),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputReplyToMessage) SetFlags() {
	if !(i.TopMsgID == 0) {
		i.Flags.Set(0)
	}
	if !(i.ReplyToPeerID == nil) {
		i.Flags.Set(1)
	}
	if !(i.QuoteText == "") {
		i.Flags.Set(2)
	}
	if !(i.QuoteEntities == nil) {
		i.Flags.Set(3)
	}
	if !(i.QuoteOffset == 0) {
		i.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (i *InputReplyToMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReplyToMessage#22c0f6d5 as nil")
	}
	b.PutID(InputReplyToMessageTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReplyToMessage) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReplyToMessage#22c0f6d5 as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputReplyToMessage#22c0f6d5: field flags: %w", err)
	}
	b.PutInt(i.ReplyToMsgID)
	if i.Flags.Has(0) {
		b.PutInt(i.TopMsgID)
	}
	if i.Flags.Has(1) {
		if i.ReplyToPeerID == nil {
			return fmt.Errorf("unable to encode inputReplyToMessage#22c0f6d5: field reply_to_peer_id is nil")
		}
		if err := i.ReplyToPeerID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputReplyToMessage#22c0f6d5: field reply_to_peer_id: %w", err)
		}
	}
	if i.Flags.Has(2) {
		b.PutString(i.QuoteText)
	}
	if i.Flags.Has(3) {
		b.PutVectorHeader(len(i.QuoteEntities))
		for idx, v := range i.QuoteEntities {
			if v == nil {
				return fmt.Errorf("unable to encode inputReplyToMessage#22c0f6d5: field quote_entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputReplyToMessage#22c0f6d5: field quote_entities element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(4) {
		b.PutInt(i.QuoteOffset)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReplyToMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReplyToMessage#22c0f6d5 to nil")
	}
	if err := b.ConsumeID(InputReplyToMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReplyToMessage) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReplyToMessage#22c0f6d5 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: field reply_to_msg_id: %w", err)
		}
		i.ReplyToMsgID = value
	}
	if i.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: field top_msg_id: %w", err)
		}
		i.TopMsgID = value
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: field reply_to_peer_id: %w", err)
		}
		i.ReplyToPeerID = value
	}
	if i.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: field quote_text: %w", err)
		}
		i.QuoteText = value
	}
	if i.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: field quote_entities: %w", err)
		}

		if headerLen > 0 {
			i.QuoteEntities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: field quote_entities: %w", err)
			}
			i.QuoteEntities = append(i.QuoteEntities, value)
		}
	}
	if i.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputReplyToMessage#22c0f6d5: field quote_offset: %w", err)
		}
		i.QuoteOffset = value
	}
	return nil
}

// GetReplyToMsgID returns value of ReplyToMsgID field.
func (i *InputReplyToMessage) GetReplyToMsgID() (value int) {
	if i == nil {
		return
	}
	return i.ReplyToMsgID
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (i *InputReplyToMessage) SetTopMsgID(value int) {
	i.Flags.Set(0)
	i.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (i *InputReplyToMessage) GetTopMsgID() (value int, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.TopMsgID, true
}

// SetReplyToPeerID sets value of ReplyToPeerID conditional field.
func (i *InputReplyToMessage) SetReplyToPeerID(value InputPeerClass) {
	i.Flags.Set(1)
	i.ReplyToPeerID = value
}

// GetReplyToPeerID returns value of ReplyToPeerID conditional field and
// boolean which is true if field was set.
func (i *InputReplyToMessage) GetReplyToPeerID() (value InputPeerClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.ReplyToPeerID, true
}

// SetQuoteText sets value of QuoteText conditional field.
func (i *InputReplyToMessage) SetQuoteText(value string) {
	i.Flags.Set(2)
	i.QuoteText = value
}

// GetQuoteText returns value of QuoteText conditional field and
// boolean which is true if field was set.
func (i *InputReplyToMessage) GetQuoteText() (value string, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.QuoteText, true
}

// SetQuoteEntities sets value of QuoteEntities conditional field.
func (i *InputReplyToMessage) SetQuoteEntities(value []MessageEntityClass) {
	i.Flags.Set(3)
	i.QuoteEntities = value
}

// GetQuoteEntities returns value of QuoteEntities conditional field and
// boolean which is true if field was set.
func (i *InputReplyToMessage) GetQuoteEntities() (value []MessageEntityClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.QuoteEntities, true
}

// SetQuoteOffset sets value of QuoteOffset conditional field.
func (i *InputReplyToMessage) SetQuoteOffset(value int) {
	i.Flags.Set(4)
	i.QuoteOffset = value
}

// GetQuoteOffset returns value of QuoteOffset conditional field and
// boolean which is true if field was set.
func (i *InputReplyToMessage) GetQuoteOffset() (value int, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(4) {
		return value, false
	}
	return i.QuoteOffset, true
}

// MapQuoteEntities returns field QuoteEntities wrapped in MessageEntityClassArray helper.
func (i *InputReplyToMessage) MapQuoteEntities() (value MessageEntityClassArray, ok bool) {
	if !i.Flags.Has(3) {
		return value, false
	}
	return MessageEntityClassArray(i.QuoteEntities), true
}

// InputReplyToStory represents TL type `inputReplyToStory#15b0f283`.
// Reply to a story.
//
// See https://core.telegram.org/constructor/inputReplyToStory for reference.
type InputReplyToStory struct {
	// ID of the user that posted the story.
	UserID InputUserClass
	// ID of the story to reply to.
	StoryID int
}

// InputReplyToStoryTypeID is TL type id of InputReplyToStory.
const InputReplyToStoryTypeID = 0x15b0f283

// construct implements constructor of InputReplyToClass.
func (i InputReplyToStory) construct() InputReplyToClass { return &i }

// Ensuring interfaces in compile-time for InputReplyToStory.
var (
	_ bin.Encoder     = &InputReplyToStory{}
	_ bin.Decoder     = &InputReplyToStory{}
	_ bin.BareEncoder = &InputReplyToStory{}
	_ bin.BareDecoder = &InputReplyToStory{}

	_ InputReplyToClass = &InputReplyToStory{}
)

func (i *InputReplyToStory) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.UserID == nil) {
		return false
	}
	if !(i.StoryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReplyToStory) String() string {
	if i == nil {
		return "InputReplyToStory(nil)"
	}
	type Alias InputReplyToStory
	return fmt.Sprintf("InputReplyToStory%+v", Alias(*i))
}

// FillFrom fills InputReplyToStory from given interface.
func (i *InputReplyToStory) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetStoryID() (value int)
}) {
	i.UserID = from.GetUserID()
	i.StoryID = from.GetStoryID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReplyToStory) TypeID() uint32 {
	return InputReplyToStoryTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReplyToStory) TypeName() string {
	return "inputReplyToStory"
}

// TypeInfo returns info about TL type.
func (i *InputReplyToStory) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReplyToStory",
		ID:   InputReplyToStoryTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReplyToStory) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReplyToStory#15b0f283 as nil")
	}
	b.PutID(InputReplyToStoryTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReplyToStory) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReplyToStory#15b0f283 as nil")
	}
	if i.UserID == nil {
		return fmt.Errorf("unable to encode inputReplyToStory#15b0f283: field user_id is nil")
	}
	if err := i.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputReplyToStory#15b0f283: field user_id: %w", err)
	}
	b.PutInt(i.StoryID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReplyToStory) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReplyToStory#15b0f283 to nil")
	}
	if err := b.ConsumeID(InputReplyToStoryTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReplyToStory#15b0f283: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReplyToStory) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReplyToStory#15b0f283 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputReplyToStory#15b0f283: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputReplyToStory#15b0f283: field story_id: %w", err)
		}
		i.StoryID = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (i *InputReplyToStory) GetUserID() (value InputUserClass) {
	if i == nil {
		return
	}
	return i.UserID
}

// GetStoryID returns value of StoryID field.
func (i *InputReplyToStory) GetStoryID() (value int) {
	if i == nil {
		return
	}
	return i.StoryID
}

// InputReplyToClassName is schema name of InputReplyToClass.
const InputReplyToClassName = "InputReplyTo"

// InputReplyToClass represents InputReplyTo generic type.
//
// See https://core.telegram.org/type/InputReplyTo for reference.
//
// Example:
//
//	g, err := tg.DecodeInputReplyTo(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputReplyToMessage: // inputReplyToMessage#22c0f6d5
//	case *tg.InputReplyToStory: // inputReplyToStory#15b0f283
//	default: panic(v)
//	}
type InputReplyToClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputReplyToClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputReplyTo implements binary de-serialization for InputReplyToClass.
func DecodeInputReplyTo(buf *bin.Buffer) (InputReplyToClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputReplyToMessageTypeID:
		// Decoding inputReplyToMessage#22c0f6d5.
		v := InputReplyToMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputReplyToClass: %w", err)
		}
		return &v, nil
	case InputReplyToStoryTypeID:
		// Decoding inputReplyToStory#15b0f283.
		v := InputReplyToStory{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputReplyToClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputReplyToClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputReplyTo boxes the InputReplyToClass providing a helper.
type InputReplyToBox struct {
	InputReplyTo InputReplyToClass
}

// Decode implements bin.Decoder for InputReplyToBox.
func (b *InputReplyToBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputReplyToBox to nil")
	}
	v, err := DecodeInputReplyTo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputReplyTo = v
	return nil
}

// Encode implements bin.Encode for InputReplyToBox.
func (b *InputReplyToBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputReplyTo == nil {
		return fmt.Errorf("unable to encode InputReplyToClass as nil")
	}
	return b.InputReplyTo.Encode(buf)
}
