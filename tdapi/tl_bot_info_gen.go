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

// BotInfo represents TL type `botInfo#bfdb89a`.
type BotInfo struct {
	// The text that is shown on the bot's profile page and is sent together with the link
	// when users share the bot
	ShortDescription string
	// Contains information about a bot
	Description string
	// Photo shown in the chat with the bot if the chat is empty; may be null
	Photo Photo
	// Animation shown in the chat with the bot if the chat is empty; may be null
	Animation Animation
	// Information about a button to show instead of the bot commands menu button; may be
	// null if ordinary bot commands menu must be shown
	MenuButton BotMenuButton
	// List of the bot commands
	Commands []BotCommand
	// Default administrator rights for adding the bot to basic group and supergroup chats;
	// may be null
	DefaultGroupAdministratorRights ChatAdministratorRights
	// Default administrator rights for adding the bot to channels; may be null
	DefaultChannelAdministratorRights ChatAdministratorRights
	// The internal link, which can be used to edit bot commands; may be null
	EditCommandsLink InternalLinkTypeClass
	// The internal link, which can be used to edit bot description; may be null
	EditDescriptionLink InternalLinkTypeClass
	// The internal link, which can be used to edit the photo or animation shown in the chat
	// with the bot if the chat is empty; may be null
	EditDescriptionMediaLink InternalLinkTypeClass
	// The internal link, which can be used to edit bot settings; may be null
	EditSettingsLink InternalLinkTypeClass
}

// BotInfoTypeID is TL type id of BotInfo.
const BotInfoTypeID = 0xbfdb89a

// Ensuring interfaces in compile-time for BotInfo.
var (
	_ bin.Encoder     = &BotInfo{}
	_ bin.Decoder     = &BotInfo{}
	_ bin.BareEncoder = &BotInfo{}
	_ bin.BareDecoder = &BotInfo{}
)

func (b *BotInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ShortDescription == "") {
		return false
	}
	if !(b.Description == "") {
		return false
	}
	if !(b.Photo.Zero()) {
		return false
	}
	if !(b.Animation.Zero()) {
		return false
	}
	if !(b.MenuButton.Zero()) {
		return false
	}
	if !(b.Commands == nil) {
		return false
	}
	if !(b.DefaultGroupAdministratorRights.Zero()) {
		return false
	}
	if !(b.DefaultChannelAdministratorRights.Zero()) {
		return false
	}
	if !(b.EditCommandsLink == nil) {
		return false
	}
	if !(b.EditDescriptionLink == nil) {
		return false
	}
	if !(b.EditDescriptionMediaLink == nil) {
		return false
	}
	if !(b.EditSettingsLink == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotInfo) String() string {
	if b == nil {
		return "BotInfo(nil)"
	}
	type Alias BotInfo
	return fmt.Sprintf("BotInfo%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotInfo) TypeID() uint32 {
	return BotInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*BotInfo) TypeName() string {
	return "botInfo"
}

// TypeInfo returns info about TL type.
func (b *BotInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botInfo",
		ID:   BotInfoTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortDescription",
			SchemaName: "short_description",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Animation",
			SchemaName: "animation",
		},
		{
			Name:       "MenuButton",
			SchemaName: "menu_button",
		},
		{
			Name:       "Commands",
			SchemaName: "commands",
		},
		{
			Name:       "DefaultGroupAdministratorRights",
			SchemaName: "default_group_administrator_rights",
		},
		{
			Name:       "DefaultChannelAdministratorRights",
			SchemaName: "default_channel_administrator_rights",
		},
		{
			Name:       "EditCommandsLink",
			SchemaName: "edit_commands_link",
		},
		{
			Name:       "EditDescriptionLink",
			SchemaName: "edit_description_link",
		},
		{
			Name:       "EditDescriptionMediaLink",
			SchemaName: "edit_description_media_link",
		},
		{
			Name:       "EditSettingsLink",
			SchemaName: "edit_settings_link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInfo#bfdb89a as nil")
	}
	buf.PutID(BotInfoTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotInfo) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInfo#bfdb89a as nil")
	}
	buf.PutString(b.ShortDescription)
	buf.PutString(b.Description)
	if err := b.Photo.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field photo: %w", err)
	}
	if err := b.Animation.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field animation: %w", err)
	}
	if err := b.MenuButton.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field menu_button: %w", err)
	}
	buf.PutInt(len(b.Commands))
	for idx, v := range b.Commands {
		if err := v.EncodeBare(buf); err != nil {
			return fmt.Errorf("unable to encode bare botInfo#bfdb89a: field commands element with index %d: %w", idx, err)
		}
	}
	if err := b.DefaultGroupAdministratorRights.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field default_group_administrator_rights: %w", err)
	}
	if err := b.DefaultChannelAdministratorRights.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field default_channel_administrator_rights: %w", err)
	}
	if b.EditCommandsLink == nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_commands_link is nil")
	}
	if err := b.EditCommandsLink.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_commands_link: %w", err)
	}
	if b.EditDescriptionLink == nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_description_link is nil")
	}
	if err := b.EditDescriptionLink.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_description_link: %w", err)
	}
	if b.EditDescriptionMediaLink == nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_description_media_link is nil")
	}
	if err := b.EditDescriptionMediaLink.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_description_media_link: %w", err)
	}
	if b.EditSettingsLink == nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_settings_link is nil")
	}
	if err := b.EditSettingsLink.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_settings_link: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInfo#bfdb89a to nil")
	}
	if err := buf.ConsumeID(BotInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode botInfo#bfdb89a: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotInfo) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInfo#bfdb89a to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field short_description: %w", err)
		}
		b.ShortDescription = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field description: %w", err)
		}
		b.Description = value
	}
	{
		if err := b.Photo.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field photo: %w", err)
		}
	}
	{
		if err := b.Animation.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field animation: %w", err)
		}
	}
	{
		if err := b.MenuButton.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field menu_button: %w", err)
		}
	}
	{
		headerLen, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field commands: %w", err)
		}

		if headerLen > 0 {
			b.Commands = make([]BotCommand, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotCommand
			if err := value.DecodeBare(buf); err != nil {
				return fmt.Errorf("unable to decode bare botInfo#bfdb89a: field commands: %w", err)
			}
			b.Commands = append(b.Commands, value)
		}
	}
	{
		if err := b.DefaultGroupAdministratorRights.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field default_group_administrator_rights: %w", err)
		}
	}
	{
		if err := b.DefaultChannelAdministratorRights.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field default_channel_administrator_rights: %w", err)
		}
	}
	{
		value, err := DecodeInternalLinkType(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field edit_commands_link: %w", err)
		}
		b.EditCommandsLink = value
	}
	{
		value, err := DecodeInternalLinkType(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field edit_description_link: %w", err)
		}
		b.EditDescriptionLink = value
	}
	{
		value, err := DecodeInternalLinkType(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field edit_description_media_link: %w", err)
		}
		b.EditDescriptionMediaLink = value
	}
	{
		value, err := DecodeInternalLinkType(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#bfdb89a: field edit_settings_link: %w", err)
		}
		b.EditSettingsLink = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotInfo) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botInfo#bfdb89a as nil")
	}
	buf.ObjStart()
	buf.PutID("botInfo")
	buf.Comma()
	buf.FieldStart("short_description")
	buf.PutString(b.ShortDescription)
	buf.Comma()
	buf.FieldStart("description")
	buf.PutString(b.Description)
	buf.Comma()
	buf.FieldStart("photo")
	if err := b.Photo.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field photo: %w", err)
	}
	buf.Comma()
	buf.FieldStart("animation")
	if err := b.Animation.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field animation: %w", err)
	}
	buf.Comma()
	buf.FieldStart("menu_button")
	if err := b.MenuButton.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field menu_button: %w", err)
	}
	buf.Comma()
	buf.FieldStart("commands")
	buf.ArrStart()
	for idx, v := range b.Commands {
		if err := v.EncodeTDLibJSON(buf); err != nil {
			return fmt.Errorf("unable to encode botInfo#bfdb89a: field commands element with index %d: %w", idx, err)
		}
		buf.Comma()
	}
	buf.StripComma()
	buf.ArrEnd()
	buf.Comma()
	buf.FieldStart("default_group_administrator_rights")
	if err := b.DefaultGroupAdministratorRights.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field default_group_administrator_rights: %w", err)
	}
	buf.Comma()
	buf.FieldStart("default_channel_administrator_rights")
	if err := b.DefaultChannelAdministratorRights.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field default_channel_administrator_rights: %w", err)
	}
	buf.Comma()
	buf.FieldStart("edit_commands_link")
	if b.EditCommandsLink == nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_commands_link is nil")
	}
	if err := b.EditCommandsLink.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_commands_link: %w", err)
	}
	buf.Comma()
	buf.FieldStart("edit_description_link")
	if b.EditDescriptionLink == nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_description_link is nil")
	}
	if err := b.EditDescriptionLink.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_description_link: %w", err)
	}
	buf.Comma()
	buf.FieldStart("edit_description_media_link")
	if b.EditDescriptionMediaLink == nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_description_media_link is nil")
	}
	if err := b.EditDescriptionMediaLink.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_description_media_link: %w", err)
	}
	buf.Comma()
	buf.FieldStart("edit_settings_link")
	if b.EditSettingsLink == nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_settings_link is nil")
	}
	if err := b.EditSettingsLink.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#bfdb89a: field edit_settings_link: %w", err)
	}
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotInfo) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botInfo#bfdb89a to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botInfo"); err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: %w", err)
			}
		case "short_description":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field short_description: %w", err)
			}
			b.ShortDescription = value
		case "description":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field description: %w", err)
			}
			b.Description = value
		case "photo":
			if err := b.Photo.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field photo: %w", err)
			}
		case "animation":
			if err := b.Animation.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field animation: %w", err)
			}
		case "menu_button":
			if err := b.MenuButton.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field menu_button: %w", err)
			}
		case "commands":
			if err := buf.Arr(func(buf tdjson.Decoder) error {
				var value BotCommand
				if err := value.DecodeTDLibJSON(buf); err != nil {
					return fmt.Errorf("unable to decode botInfo#bfdb89a: field commands: %w", err)
				}
				b.Commands = append(b.Commands, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field commands: %w", err)
			}
		case "default_group_administrator_rights":
			if err := b.DefaultGroupAdministratorRights.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field default_group_administrator_rights: %w", err)
			}
		case "default_channel_administrator_rights":
			if err := b.DefaultChannelAdministratorRights.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field default_channel_administrator_rights: %w", err)
			}
		case "edit_commands_link":
			value, err := DecodeTDLibJSONInternalLinkType(buf)
			if err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field edit_commands_link: %w", err)
			}
			b.EditCommandsLink = value
		case "edit_description_link":
			value, err := DecodeTDLibJSONInternalLinkType(buf)
			if err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field edit_description_link: %w", err)
			}
			b.EditDescriptionLink = value
		case "edit_description_media_link":
			value, err := DecodeTDLibJSONInternalLinkType(buf)
			if err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field edit_description_media_link: %w", err)
			}
			b.EditDescriptionMediaLink = value
		case "edit_settings_link":
			value, err := DecodeTDLibJSONInternalLinkType(buf)
			if err != nil {
				return fmt.Errorf("unable to decode botInfo#bfdb89a: field edit_settings_link: %w", err)
			}
			b.EditSettingsLink = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetShortDescription returns value of ShortDescription field.
func (b *BotInfo) GetShortDescription() (value string) {
	if b == nil {
		return
	}
	return b.ShortDescription
}

// GetDescription returns value of Description field.
func (b *BotInfo) GetDescription() (value string) {
	if b == nil {
		return
	}
	return b.Description
}

// GetPhoto returns value of Photo field.
func (b *BotInfo) GetPhoto() (value Photo) {
	if b == nil {
		return
	}
	return b.Photo
}

// GetAnimation returns value of Animation field.
func (b *BotInfo) GetAnimation() (value Animation) {
	if b == nil {
		return
	}
	return b.Animation
}

// GetMenuButton returns value of MenuButton field.
func (b *BotInfo) GetMenuButton() (value BotMenuButton) {
	if b == nil {
		return
	}
	return b.MenuButton
}

// GetCommands returns value of Commands field.
func (b *BotInfo) GetCommands() (value []BotCommand) {
	if b == nil {
		return
	}
	return b.Commands
}

// GetDefaultGroupAdministratorRights returns value of DefaultGroupAdministratorRights field.
func (b *BotInfo) GetDefaultGroupAdministratorRights() (value ChatAdministratorRights) {
	if b == nil {
		return
	}
	return b.DefaultGroupAdministratorRights
}

// GetDefaultChannelAdministratorRights returns value of DefaultChannelAdministratorRights field.
func (b *BotInfo) GetDefaultChannelAdministratorRights() (value ChatAdministratorRights) {
	if b == nil {
		return
	}
	return b.DefaultChannelAdministratorRights
}

// GetEditCommandsLink returns value of EditCommandsLink field.
func (b *BotInfo) GetEditCommandsLink() (value InternalLinkTypeClass) {
	if b == nil {
		return
	}
	return b.EditCommandsLink
}

// GetEditDescriptionLink returns value of EditDescriptionLink field.
func (b *BotInfo) GetEditDescriptionLink() (value InternalLinkTypeClass) {
	if b == nil {
		return
	}
	return b.EditDescriptionLink
}

// GetEditDescriptionMediaLink returns value of EditDescriptionMediaLink field.
func (b *BotInfo) GetEditDescriptionMediaLink() (value InternalLinkTypeClass) {
	if b == nil {
		return
	}
	return b.EditDescriptionMediaLink
}

// GetEditSettingsLink returns value of EditSettingsLink field.
func (b *BotInfo) GetEditSettingsLink() (value InternalLinkTypeClass) {
	if b == nil {
		return
	}
	return b.EditSettingsLink
}
