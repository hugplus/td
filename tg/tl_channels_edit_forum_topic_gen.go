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

// ChannelsEditForumTopicRequest represents TL type `channels.editForumTopic#f4dfa185`.
// Edit forum topic¹; requires manage_topics rights².
//
// Links:
//  1. https://core.telegram.org/api/forum
//  2. https://core.telegram.org/api/rights
//
// See https://core.telegram.org/method/channels.editForumTopic for reference.
type ChannelsEditForumTopicRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Supergroup
	Channel InputChannelClass
	// Topic ID
	TopicID int
	// If present, will update the topic title (maximum UTF-8 length: 128).
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// If present, updates the custom emoji¹ used as topic icon. Telegram Premium² users
	// can use any custom emoji, other users can only use the custom emojis contained in the
	// inputStickerSetEmojiDefaultTopicIcons³ emoji pack. Pass 0 to switch to the fallback
	// topic icon.
	//
	// Links:
	//  1) https://core.telegram.org/api/custom-emoji
	//  2) https://core.telegram.org/api/premium
	//  3) https://core.telegram.org/constructor/inputStickerSetEmojiDefaultTopicIcons
	//
	// Use SetIconEmojiID and GetIconEmojiID helpers.
	IconEmojiID int64
	// If present, will update the open/closed status of the topic.
	//
	// Use SetClosed and GetClosed helpers.
	Closed bool
	// If present, will hide/unhide the topic (only valid for the "General" topic, id=1).
	//
	// Use SetHidden and GetHidden helpers.
	Hidden bool
}

// ChannelsEditForumTopicRequestTypeID is TL type id of ChannelsEditForumTopicRequest.
const ChannelsEditForumTopicRequestTypeID = 0xf4dfa185

// Ensuring interfaces in compile-time for ChannelsEditForumTopicRequest.
var (
	_ bin.Encoder     = &ChannelsEditForumTopicRequest{}
	_ bin.Decoder     = &ChannelsEditForumTopicRequest{}
	_ bin.BareEncoder = &ChannelsEditForumTopicRequest{}
	_ bin.BareDecoder = &ChannelsEditForumTopicRequest{}
)

func (e *ChannelsEditForumTopicRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.TopicID == 0) {
		return false
	}
	if !(e.Title == "") {
		return false
	}
	if !(e.IconEmojiID == 0) {
		return false
	}
	if !(e.Closed == false) {
		return false
	}
	if !(e.Hidden == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsEditForumTopicRequest) String() string {
	if e == nil {
		return "ChannelsEditForumTopicRequest(nil)"
	}
	type Alias ChannelsEditForumTopicRequest
	return fmt.Sprintf("ChannelsEditForumTopicRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsEditForumTopicRequest from given interface.
func (e *ChannelsEditForumTopicRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetTopicID() (value int)
	GetTitle() (value string, ok bool)
	GetIconEmojiID() (value int64, ok bool)
	GetClosed() (value bool, ok bool)
	GetHidden() (value bool, ok bool)
}) {
	e.Channel = from.GetChannel()
	e.TopicID = from.GetTopicID()
	if val, ok := from.GetTitle(); ok {
		e.Title = val
	}

	if val, ok := from.GetIconEmojiID(); ok {
		e.IconEmojiID = val
	}

	if val, ok := from.GetClosed(); ok {
		e.Closed = val
	}

	if val, ok := from.GetHidden(); ok {
		e.Hidden = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsEditForumTopicRequest) TypeID() uint32 {
	return ChannelsEditForumTopicRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsEditForumTopicRequest) TypeName() string {
	return "channels.editForumTopic"
}

// TypeInfo returns info about TL type.
func (e *ChannelsEditForumTopicRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.editForumTopic",
		ID:   ChannelsEditForumTopicRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "TopicID",
			SchemaName: "topic_id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "IconEmojiID",
			SchemaName: "icon_emoji_id",
			Null:       !e.Flags.Has(1),
		},
		{
			Name:       "Closed",
			SchemaName: "closed",
			Null:       !e.Flags.Has(2),
		},
		{
			Name:       "Hidden",
			SchemaName: "hidden",
			Null:       !e.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *ChannelsEditForumTopicRequest) SetFlags() {
	if !(e.Title == "") {
		e.Flags.Set(0)
	}
	if !(e.IconEmojiID == 0) {
		e.Flags.Set(1)
	}
	if !(e.Closed == false) {
		e.Flags.Set(2)
	}
	if !(e.Hidden == false) {
		e.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (e *ChannelsEditForumTopicRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editForumTopic#f4dfa185 as nil")
	}
	b.PutID(ChannelsEditForumTopicRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ChannelsEditForumTopicRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editForumTopic#f4dfa185 as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editForumTopic#f4dfa185: field flags: %w", err)
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editForumTopic#f4dfa185: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editForumTopic#f4dfa185: field channel: %w", err)
	}
	b.PutInt(e.TopicID)
	if e.Flags.Has(0) {
		b.PutString(e.Title)
	}
	if e.Flags.Has(1) {
		b.PutLong(e.IconEmojiID)
	}
	if e.Flags.Has(2) {
		b.PutBool(e.Closed)
	}
	if e.Flags.Has(3) {
		b.PutBool(e.Hidden)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditForumTopicRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editForumTopic#f4dfa185 to nil")
	}
	if err := b.ConsumeID(ChannelsEditForumTopicRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editForumTopic#f4dfa185: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ChannelsEditForumTopicRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editForumTopic#f4dfa185 to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.editForumTopic#f4dfa185: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editForumTopic#f4dfa185: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editForumTopic#f4dfa185: field topic_id: %w", err)
		}
		e.TopicID = value
	}
	if e.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editForumTopic#f4dfa185: field title: %w", err)
		}
		e.Title = value
	}
	if e.Flags.Has(1) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editForumTopic#f4dfa185: field icon_emoji_id: %w", err)
		}
		e.IconEmojiID = value
	}
	if e.Flags.Has(2) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editForumTopic#f4dfa185: field closed: %w", err)
		}
		e.Closed = value
	}
	if e.Flags.Has(3) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editForumTopic#f4dfa185: field hidden: %w", err)
		}
		e.Hidden = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (e *ChannelsEditForumTopicRequest) GetChannel() (value InputChannelClass) {
	if e == nil {
		return
	}
	return e.Channel
}

// GetTopicID returns value of TopicID field.
func (e *ChannelsEditForumTopicRequest) GetTopicID() (value int) {
	if e == nil {
		return
	}
	return e.TopicID
}

// SetTitle sets value of Title conditional field.
func (e *ChannelsEditForumTopicRequest) SetTitle(value string) {
	e.Flags.Set(0)
	e.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (e *ChannelsEditForumTopicRequest) GetTitle() (value string, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.Title, true
}

// SetIconEmojiID sets value of IconEmojiID conditional field.
func (e *ChannelsEditForumTopicRequest) SetIconEmojiID(value int64) {
	e.Flags.Set(1)
	e.IconEmojiID = value
}

// GetIconEmojiID returns value of IconEmojiID conditional field and
// boolean which is true if field was set.
func (e *ChannelsEditForumTopicRequest) GetIconEmojiID() (value int64, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(1) {
		return value, false
	}
	return e.IconEmojiID, true
}

// SetClosed sets value of Closed conditional field.
func (e *ChannelsEditForumTopicRequest) SetClosed(value bool) {
	e.Flags.Set(2)
	e.Closed = value
}

// GetClosed returns value of Closed conditional field and
// boolean which is true if field was set.
func (e *ChannelsEditForumTopicRequest) GetClosed() (value bool, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(2) {
		return value, false
	}
	return e.Closed, true
}

// SetHidden sets value of Hidden conditional field.
func (e *ChannelsEditForumTopicRequest) SetHidden(value bool) {
	e.Flags.Set(3)
	e.Hidden = value
}

// GetHidden returns value of Hidden conditional field and
// boolean which is true if field was set.
func (e *ChannelsEditForumTopicRequest) GetHidden() (value bool, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(3) {
		return value, false
	}
	return e.Hidden, true
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsEditForumTopicRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// ChannelsEditForumTopic invokes method channels.editForumTopic#f4dfa185 returning error if any.
// Edit forum topic¹; requires manage_topics rights².
//
// Links:
//  1. https://core.telegram.org/api/forum
//  2. https://core.telegram.org/api/rights
//
// Possible errors:
//
//	400 CHANNEL_FORUM_MISSING: This supergroup is not a forum.
//	403 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 DOCUMENT_INVALID: The specified document is invalid.
//	400 GENERAL_MODIFY_ICON_FORBIDDEN: You can't modify the icon of the "General" topic.
//	400 TOPIC_CLOSE_SEPARATELY: The close flag cannot be provided together with any of the other flags.
//	400 TOPIC_HIDE_SEPARATELY: The hide flag cannot be provided together with any of the other flags.
//	400 TOPIC_ID_INVALID: The specified topic ID is invalid.
//	400 TOPIC_NOT_MODIFIED: The updated topic info is equal to the current topic info, nothing was changed.
//
// See https://core.telegram.org/method/channels.editForumTopic for reference.
// Can be used by bots.
func (c *Client) ChannelsEditForumTopic(ctx context.Context, request *ChannelsEditForumTopicRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
