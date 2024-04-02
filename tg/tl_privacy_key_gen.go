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

// PrivacyKeyStatusTimestamp represents TL type `privacyKeyStatusTimestamp#bc2eab30`.
// Whether we can see the last online timestamp of this user
//
// See https://core.telegram.org/constructor/privacyKeyStatusTimestamp for reference.
type PrivacyKeyStatusTimestamp struct {
}

// PrivacyKeyStatusTimestampTypeID is TL type id of PrivacyKeyStatusTimestamp.
const PrivacyKeyStatusTimestampTypeID = 0xbc2eab30

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyStatusTimestamp) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyStatusTimestamp.
var (
	_ bin.Encoder     = &PrivacyKeyStatusTimestamp{}
	_ bin.Decoder     = &PrivacyKeyStatusTimestamp{}
	_ bin.BareEncoder = &PrivacyKeyStatusTimestamp{}
	_ bin.BareDecoder = &PrivacyKeyStatusTimestamp{}

	_ PrivacyKeyClass = &PrivacyKeyStatusTimestamp{}
)

func (p *PrivacyKeyStatusTimestamp) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyStatusTimestamp) String() string {
	if p == nil {
		return "PrivacyKeyStatusTimestamp(nil)"
	}
	type Alias PrivacyKeyStatusTimestamp
	return fmt.Sprintf("PrivacyKeyStatusTimestamp%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyStatusTimestamp) TypeID() uint32 {
	return PrivacyKeyStatusTimestampTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyStatusTimestamp) TypeName() string {
	return "privacyKeyStatusTimestamp"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyStatusTimestamp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyStatusTimestamp",
		ID:   PrivacyKeyStatusTimestampTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyStatusTimestamp) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyStatusTimestamp#bc2eab30 as nil")
	}
	b.PutID(PrivacyKeyStatusTimestampTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyStatusTimestamp) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyStatusTimestamp#bc2eab30 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyStatusTimestamp) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyStatusTimestamp#bc2eab30 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyStatusTimestampTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyStatusTimestamp#bc2eab30: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyStatusTimestamp) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyStatusTimestamp#bc2eab30 to nil")
	}
	return nil
}

// PrivacyKeyChatInvite represents TL type `privacyKeyChatInvite#500e6dfa`.
// Whether the user can be invited to chats
//
// See https://core.telegram.org/constructor/privacyKeyChatInvite for reference.
type PrivacyKeyChatInvite struct {
}

// PrivacyKeyChatInviteTypeID is TL type id of PrivacyKeyChatInvite.
const PrivacyKeyChatInviteTypeID = 0x500e6dfa

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyChatInvite) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyChatInvite.
var (
	_ bin.Encoder     = &PrivacyKeyChatInvite{}
	_ bin.Decoder     = &PrivacyKeyChatInvite{}
	_ bin.BareEncoder = &PrivacyKeyChatInvite{}
	_ bin.BareDecoder = &PrivacyKeyChatInvite{}

	_ PrivacyKeyClass = &PrivacyKeyChatInvite{}
)

func (p *PrivacyKeyChatInvite) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyChatInvite) String() string {
	if p == nil {
		return "PrivacyKeyChatInvite(nil)"
	}
	type Alias PrivacyKeyChatInvite
	return fmt.Sprintf("PrivacyKeyChatInvite%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyChatInvite) TypeID() uint32 {
	return PrivacyKeyChatInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyChatInvite) TypeName() string {
	return "privacyKeyChatInvite"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyChatInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyChatInvite",
		ID:   PrivacyKeyChatInviteTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyChatInvite) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyChatInvite#500e6dfa as nil")
	}
	b.PutID(PrivacyKeyChatInviteTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyChatInvite) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyChatInvite#500e6dfa as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyChatInvite) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyChatInvite#500e6dfa to nil")
	}
	if err := b.ConsumeID(PrivacyKeyChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyChatInvite#500e6dfa: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyChatInvite) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyChatInvite#500e6dfa to nil")
	}
	return nil
}

// PrivacyKeyPhoneCall represents TL type `privacyKeyPhoneCall#3d662b7b`.
// Whether the user accepts phone calls
//
// See https://core.telegram.org/constructor/privacyKeyPhoneCall for reference.
type PrivacyKeyPhoneCall struct {
}

// PrivacyKeyPhoneCallTypeID is TL type id of PrivacyKeyPhoneCall.
const PrivacyKeyPhoneCallTypeID = 0x3d662b7b

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyPhoneCall) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyPhoneCall.
var (
	_ bin.Encoder     = &PrivacyKeyPhoneCall{}
	_ bin.Decoder     = &PrivacyKeyPhoneCall{}
	_ bin.BareEncoder = &PrivacyKeyPhoneCall{}
	_ bin.BareDecoder = &PrivacyKeyPhoneCall{}

	_ PrivacyKeyClass = &PrivacyKeyPhoneCall{}
)

func (p *PrivacyKeyPhoneCall) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyPhoneCall) String() string {
	if p == nil {
		return "PrivacyKeyPhoneCall(nil)"
	}
	type Alias PrivacyKeyPhoneCall
	return fmt.Sprintf("PrivacyKeyPhoneCall%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyPhoneCall) TypeID() uint32 {
	return PrivacyKeyPhoneCallTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyPhoneCall) TypeName() string {
	return "privacyKeyPhoneCall"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyPhoneCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyPhoneCall",
		ID:   PrivacyKeyPhoneCallTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyPhoneCall) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneCall#3d662b7b as nil")
	}
	b.PutID(PrivacyKeyPhoneCallTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyPhoneCall) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneCall#3d662b7b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyPhoneCall) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneCall#3d662b7b to nil")
	}
	if err := b.ConsumeID(PrivacyKeyPhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyPhoneCall#3d662b7b: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyPhoneCall) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneCall#3d662b7b to nil")
	}
	return nil
}

// PrivacyKeyPhoneP2P represents TL type `privacyKeyPhoneP2P#39491cc8`.
// Whether P2P connections in phone calls with this user are allowed
//
// See https://core.telegram.org/constructor/privacyKeyPhoneP2P for reference.
type PrivacyKeyPhoneP2P struct {
}

// PrivacyKeyPhoneP2PTypeID is TL type id of PrivacyKeyPhoneP2P.
const PrivacyKeyPhoneP2PTypeID = 0x39491cc8

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyPhoneP2P) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyPhoneP2P.
var (
	_ bin.Encoder     = &PrivacyKeyPhoneP2P{}
	_ bin.Decoder     = &PrivacyKeyPhoneP2P{}
	_ bin.BareEncoder = &PrivacyKeyPhoneP2P{}
	_ bin.BareDecoder = &PrivacyKeyPhoneP2P{}

	_ PrivacyKeyClass = &PrivacyKeyPhoneP2P{}
)

func (p *PrivacyKeyPhoneP2P) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyPhoneP2P) String() string {
	if p == nil {
		return "PrivacyKeyPhoneP2P(nil)"
	}
	type Alias PrivacyKeyPhoneP2P
	return fmt.Sprintf("PrivacyKeyPhoneP2P%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyPhoneP2P) TypeID() uint32 {
	return PrivacyKeyPhoneP2PTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyPhoneP2P) TypeName() string {
	return "privacyKeyPhoneP2P"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyPhoneP2P) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyPhoneP2P",
		ID:   PrivacyKeyPhoneP2PTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyPhoneP2P) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneP2P#39491cc8 as nil")
	}
	b.PutID(PrivacyKeyPhoneP2PTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyPhoneP2P) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneP2P#39491cc8 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyPhoneP2P) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneP2P#39491cc8 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyPhoneP2PTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyPhoneP2P#39491cc8: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyPhoneP2P) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneP2P#39491cc8 to nil")
	}
	return nil
}

// PrivacyKeyForwards represents TL type `privacyKeyForwards#69ec56a3`.
// Whether messages forwarded from the user will be anonymously forwarded¹
//
// Links:
//  1. https://telegram.org/blog/unsend-privacy-emoji#anonymous-forwarding
//
// See https://core.telegram.org/constructor/privacyKeyForwards for reference.
type PrivacyKeyForwards struct {
}

// PrivacyKeyForwardsTypeID is TL type id of PrivacyKeyForwards.
const PrivacyKeyForwardsTypeID = 0x69ec56a3

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyForwards) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyForwards.
var (
	_ bin.Encoder     = &PrivacyKeyForwards{}
	_ bin.Decoder     = &PrivacyKeyForwards{}
	_ bin.BareEncoder = &PrivacyKeyForwards{}
	_ bin.BareDecoder = &PrivacyKeyForwards{}

	_ PrivacyKeyClass = &PrivacyKeyForwards{}
)

func (p *PrivacyKeyForwards) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyForwards) String() string {
	if p == nil {
		return "PrivacyKeyForwards(nil)"
	}
	type Alias PrivacyKeyForwards
	return fmt.Sprintf("PrivacyKeyForwards%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyForwards) TypeID() uint32 {
	return PrivacyKeyForwardsTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyForwards) TypeName() string {
	return "privacyKeyForwards"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyForwards) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyForwards",
		ID:   PrivacyKeyForwardsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyForwards) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyForwards#69ec56a3 as nil")
	}
	b.PutID(PrivacyKeyForwardsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyForwards) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyForwards#69ec56a3 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyForwards) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyForwards#69ec56a3 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyForwardsTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyForwards#69ec56a3: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyForwards) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyForwards#69ec56a3 to nil")
	}
	return nil
}

// PrivacyKeyProfilePhoto represents TL type `privacyKeyProfilePhoto#96151fed`.
// Whether the profile picture of the user is visible
//
// See https://core.telegram.org/constructor/privacyKeyProfilePhoto for reference.
type PrivacyKeyProfilePhoto struct {
}

// PrivacyKeyProfilePhotoTypeID is TL type id of PrivacyKeyProfilePhoto.
const PrivacyKeyProfilePhotoTypeID = 0x96151fed

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyProfilePhoto) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyProfilePhoto.
var (
	_ bin.Encoder     = &PrivacyKeyProfilePhoto{}
	_ bin.Decoder     = &PrivacyKeyProfilePhoto{}
	_ bin.BareEncoder = &PrivacyKeyProfilePhoto{}
	_ bin.BareDecoder = &PrivacyKeyProfilePhoto{}

	_ PrivacyKeyClass = &PrivacyKeyProfilePhoto{}
)

func (p *PrivacyKeyProfilePhoto) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyProfilePhoto) String() string {
	if p == nil {
		return "PrivacyKeyProfilePhoto(nil)"
	}
	type Alias PrivacyKeyProfilePhoto
	return fmt.Sprintf("PrivacyKeyProfilePhoto%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyProfilePhoto) TypeID() uint32 {
	return PrivacyKeyProfilePhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyProfilePhoto) TypeName() string {
	return "privacyKeyProfilePhoto"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyProfilePhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyProfilePhoto",
		ID:   PrivacyKeyProfilePhotoTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyProfilePhoto) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyProfilePhoto#96151fed as nil")
	}
	b.PutID(PrivacyKeyProfilePhotoTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyProfilePhoto) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyProfilePhoto#96151fed as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyProfilePhoto) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyProfilePhoto#96151fed to nil")
	}
	if err := b.ConsumeID(PrivacyKeyProfilePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyProfilePhoto#96151fed: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyProfilePhoto) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyProfilePhoto#96151fed to nil")
	}
	return nil
}

// PrivacyKeyPhoneNumber represents TL type `privacyKeyPhoneNumber#d19ae46d`.
// Whether the user allows us to see his phone number
//
// See https://core.telegram.org/constructor/privacyKeyPhoneNumber for reference.
type PrivacyKeyPhoneNumber struct {
}

// PrivacyKeyPhoneNumberTypeID is TL type id of PrivacyKeyPhoneNumber.
const PrivacyKeyPhoneNumberTypeID = 0xd19ae46d

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyPhoneNumber) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyPhoneNumber.
var (
	_ bin.Encoder     = &PrivacyKeyPhoneNumber{}
	_ bin.Decoder     = &PrivacyKeyPhoneNumber{}
	_ bin.BareEncoder = &PrivacyKeyPhoneNumber{}
	_ bin.BareDecoder = &PrivacyKeyPhoneNumber{}

	_ PrivacyKeyClass = &PrivacyKeyPhoneNumber{}
)

func (p *PrivacyKeyPhoneNumber) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyPhoneNumber) String() string {
	if p == nil {
		return "PrivacyKeyPhoneNumber(nil)"
	}
	type Alias PrivacyKeyPhoneNumber
	return fmt.Sprintf("PrivacyKeyPhoneNumber%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyPhoneNumber) TypeID() uint32 {
	return PrivacyKeyPhoneNumberTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyPhoneNumber) TypeName() string {
	return "privacyKeyPhoneNumber"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyPhoneNumber) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyPhoneNumber",
		ID:   PrivacyKeyPhoneNumberTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyPhoneNumber) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneNumber#d19ae46d as nil")
	}
	b.PutID(PrivacyKeyPhoneNumberTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyPhoneNumber) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneNumber#d19ae46d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyPhoneNumber) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneNumber#d19ae46d to nil")
	}
	if err := b.ConsumeID(PrivacyKeyPhoneNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyPhoneNumber#d19ae46d: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyPhoneNumber) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneNumber#d19ae46d to nil")
	}
	return nil
}

// PrivacyKeyAddedByPhone represents TL type `privacyKeyAddedByPhone#42ffd42b`.
// Whether this user can be added to our contact list by their phone number
//
// See https://core.telegram.org/constructor/privacyKeyAddedByPhone for reference.
type PrivacyKeyAddedByPhone struct {
}

// PrivacyKeyAddedByPhoneTypeID is TL type id of PrivacyKeyAddedByPhone.
const PrivacyKeyAddedByPhoneTypeID = 0x42ffd42b

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyAddedByPhone) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyAddedByPhone.
var (
	_ bin.Encoder     = &PrivacyKeyAddedByPhone{}
	_ bin.Decoder     = &PrivacyKeyAddedByPhone{}
	_ bin.BareEncoder = &PrivacyKeyAddedByPhone{}
	_ bin.BareDecoder = &PrivacyKeyAddedByPhone{}

	_ PrivacyKeyClass = &PrivacyKeyAddedByPhone{}
)

func (p *PrivacyKeyAddedByPhone) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyAddedByPhone) String() string {
	if p == nil {
		return "PrivacyKeyAddedByPhone(nil)"
	}
	type Alias PrivacyKeyAddedByPhone
	return fmt.Sprintf("PrivacyKeyAddedByPhone%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyAddedByPhone) TypeID() uint32 {
	return PrivacyKeyAddedByPhoneTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyAddedByPhone) TypeName() string {
	return "privacyKeyAddedByPhone"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyAddedByPhone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyAddedByPhone",
		ID:   PrivacyKeyAddedByPhoneTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyAddedByPhone) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyAddedByPhone#42ffd42b as nil")
	}
	b.PutID(PrivacyKeyAddedByPhoneTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyAddedByPhone) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyAddedByPhone#42ffd42b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyAddedByPhone) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyAddedByPhone#42ffd42b to nil")
	}
	if err := b.ConsumeID(PrivacyKeyAddedByPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyAddedByPhone#42ffd42b: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyAddedByPhone) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyAddedByPhone#42ffd42b to nil")
	}
	return nil
}

// PrivacyKeyVoiceMessages represents TL type `privacyKeyVoiceMessages#697f414`.
// Whether the user accepts voice messages
//
// See https://core.telegram.org/constructor/privacyKeyVoiceMessages for reference.
type PrivacyKeyVoiceMessages struct {
}

// PrivacyKeyVoiceMessagesTypeID is TL type id of PrivacyKeyVoiceMessages.
const PrivacyKeyVoiceMessagesTypeID = 0x697f414

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyVoiceMessages) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyVoiceMessages.
var (
	_ bin.Encoder     = &PrivacyKeyVoiceMessages{}
	_ bin.Decoder     = &PrivacyKeyVoiceMessages{}
	_ bin.BareEncoder = &PrivacyKeyVoiceMessages{}
	_ bin.BareDecoder = &PrivacyKeyVoiceMessages{}

	_ PrivacyKeyClass = &PrivacyKeyVoiceMessages{}
)

func (p *PrivacyKeyVoiceMessages) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyVoiceMessages) String() string {
	if p == nil {
		return "PrivacyKeyVoiceMessages(nil)"
	}
	type Alias PrivacyKeyVoiceMessages
	return fmt.Sprintf("PrivacyKeyVoiceMessages%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyVoiceMessages) TypeID() uint32 {
	return PrivacyKeyVoiceMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyVoiceMessages) TypeName() string {
	return "privacyKeyVoiceMessages"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyVoiceMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyVoiceMessages",
		ID:   PrivacyKeyVoiceMessagesTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyVoiceMessages) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyVoiceMessages#697f414 as nil")
	}
	b.PutID(PrivacyKeyVoiceMessagesTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyVoiceMessages) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyVoiceMessages#697f414 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyVoiceMessages) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyVoiceMessages#697f414 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyVoiceMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyVoiceMessages#697f414: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyVoiceMessages) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyVoiceMessages#697f414 to nil")
	}
	return nil
}

// PrivacyKeyAbout represents TL type `privacyKeyAbout#a486b761`.
// Whether people can see your bio
//
// See https://core.telegram.org/constructor/privacyKeyAbout for reference.
type PrivacyKeyAbout struct {
}

// PrivacyKeyAboutTypeID is TL type id of PrivacyKeyAbout.
const PrivacyKeyAboutTypeID = 0xa486b761

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyAbout) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyAbout.
var (
	_ bin.Encoder     = &PrivacyKeyAbout{}
	_ bin.Decoder     = &PrivacyKeyAbout{}
	_ bin.BareEncoder = &PrivacyKeyAbout{}
	_ bin.BareDecoder = &PrivacyKeyAbout{}

	_ PrivacyKeyClass = &PrivacyKeyAbout{}
)

func (p *PrivacyKeyAbout) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyAbout) String() string {
	if p == nil {
		return "PrivacyKeyAbout(nil)"
	}
	type Alias PrivacyKeyAbout
	return fmt.Sprintf("PrivacyKeyAbout%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyAbout) TypeID() uint32 {
	return PrivacyKeyAboutTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyAbout) TypeName() string {
	return "privacyKeyAbout"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyAbout) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyAbout",
		ID:   PrivacyKeyAboutTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyAbout) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyAbout#a486b761 as nil")
	}
	b.PutID(PrivacyKeyAboutTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyAbout) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyAbout#a486b761 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyAbout) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyAbout#a486b761 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyAboutTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyAbout#a486b761: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyAbout) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyAbout#a486b761 to nil")
	}
	return nil
}

// PrivacyKeyBirthday represents TL type `privacyKeyBirthday#2000a518`.
//
// See https://core.telegram.org/constructor/privacyKeyBirthday for reference.
type PrivacyKeyBirthday struct {
}

// PrivacyKeyBirthdayTypeID is TL type id of PrivacyKeyBirthday.
const PrivacyKeyBirthdayTypeID = 0x2000a518

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyBirthday) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyBirthday.
var (
	_ bin.Encoder     = &PrivacyKeyBirthday{}
	_ bin.Decoder     = &PrivacyKeyBirthday{}
	_ bin.BareEncoder = &PrivacyKeyBirthday{}
	_ bin.BareDecoder = &PrivacyKeyBirthday{}

	_ PrivacyKeyClass = &PrivacyKeyBirthday{}
)

func (p *PrivacyKeyBirthday) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyKeyBirthday) String() string {
	if p == nil {
		return "PrivacyKeyBirthday(nil)"
	}
	type Alias PrivacyKeyBirthday
	return fmt.Sprintf("PrivacyKeyBirthday%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrivacyKeyBirthday) TypeID() uint32 {
	return PrivacyKeyBirthdayTypeID
}

// TypeName returns name of type in TL schema.
func (*PrivacyKeyBirthday) TypeName() string {
	return "privacyKeyBirthday"
}

// TypeInfo returns info about TL type.
func (p *PrivacyKeyBirthday) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "privacyKeyBirthday",
		ID:   PrivacyKeyBirthdayTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrivacyKeyBirthday) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyBirthday#2000a518 as nil")
	}
	b.PutID(PrivacyKeyBirthdayTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrivacyKeyBirthday) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyBirthday#2000a518 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyBirthday) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyBirthday#2000a518 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyBirthdayTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyBirthday#2000a518: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrivacyKeyBirthday) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyBirthday#2000a518 to nil")
	}
	return nil
}

// PrivacyKeyClassName is schema name of PrivacyKeyClass.
const PrivacyKeyClassName = "PrivacyKey"

// PrivacyKeyClass represents PrivacyKey generic type.
//
// See https://core.telegram.org/type/PrivacyKey for reference.
//
// Example:
//
//	g, err := tg.DecodePrivacyKey(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.PrivacyKeyStatusTimestamp: // privacyKeyStatusTimestamp#bc2eab30
//	case *tg.PrivacyKeyChatInvite: // privacyKeyChatInvite#500e6dfa
//	case *tg.PrivacyKeyPhoneCall: // privacyKeyPhoneCall#3d662b7b
//	case *tg.PrivacyKeyPhoneP2P: // privacyKeyPhoneP2P#39491cc8
//	case *tg.PrivacyKeyForwards: // privacyKeyForwards#69ec56a3
//	case *tg.PrivacyKeyProfilePhoto: // privacyKeyProfilePhoto#96151fed
//	case *tg.PrivacyKeyPhoneNumber: // privacyKeyPhoneNumber#d19ae46d
//	case *tg.PrivacyKeyAddedByPhone: // privacyKeyAddedByPhone#42ffd42b
//	case *tg.PrivacyKeyVoiceMessages: // privacyKeyVoiceMessages#697f414
//	case *tg.PrivacyKeyAbout: // privacyKeyAbout#a486b761
//	case *tg.PrivacyKeyBirthday: // privacyKeyBirthday#2000a518
//	default: panic(v)
//	}
type PrivacyKeyClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PrivacyKeyClass

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

// DecodePrivacyKey implements binary de-serialization for PrivacyKeyClass.
func DecodePrivacyKey(buf *bin.Buffer) (PrivacyKeyClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PrivacyKeyStatusTimestampTypeID:
		// Decoding privacyKeyStatusTimestamp#bc2eab30.
		v := PrivacyKeyStatusTimestamp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyChatInviteTypeID:
		// Decoding privacyKeyChatInvite#500e6dfa.
		v := PrivacyKeyChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyPhoneCallTypeID:
		// Decoding privacyKeyPhoneCall#3d662b7b.
		v := PrivacyKeyPhoneCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyPhoneP2PTypeID:
		// Decoding privacyKeyPhoneP2P#39491cc8.
		v := PrivacyKeyPhoneP2P{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyForwardsTypeID:
		// Decoding privacyKeyForwards#69ec56a3.
		v := PrivacyKeyForwards{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyProfilePhotoTypeID:
		// Decoding privacyKeyProfilePhoto#96151fed.
		v := PrivacyKeyProfilePhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyPhoneNumberTypeID:
		// Decoding privacyKeyPhoneNumber#d19ae46d.
		v := PrivacyKeyPhoneNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyAddedByPhoneTypeID:
		// Decoding privacyKeyAddedByPhone#42ffd42b.
		v := PrivacyKeyAddedByPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyVoiceMessagesTypeID:
		// Decoding privacyKeyVoiceMessages#697f414.
		v := PrivacyKeyVoiceMessages{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyAboutTypeID:
		// Decoding privacyKeyAbout#a486b761.
		v := PrivacyKeyAbout{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyBirthdayTypeID:
		// Decoding privacyKeyBirthday#2000a518.
		v := PrivacyKeyBirthday{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", bin.NewUnexpectedID(id))
	}
}

// PrivacyKey boxes the PrivacyKeyClass providing a helper.
type PrivacyKeyBox struct {
	PrivacyKey PrivacyKeyClass
}

// Decode implements bin.Decoder for PrivacyKeyBox.
func (b *PrivacyKeyBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PrivacyKeyBox to nil")
	}
	v, err := DecodePrivacyKey(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PrivacyKey = v
	return nil
}

// Encode implements bin.Encode for PrivacyKeyBox.
func (b *PrivacyKeyBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PrivacyKey == nil {
		return fmt.Errorf("unable to encode PrivacyKeyClass as nil")
	}
	return b.PrivacyKey.Encode(buf)
}
