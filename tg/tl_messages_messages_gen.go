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
)

// MessagesMessages represents TL type `messages.messages#8c718e87`.
// Full list of messages with auxilary data.
//
// See https://core.telegram.org/constructor/messages.messages for reference.
type MessagesMessages struct {
	// List of messages
	Messages []MessageClass
	// List of chats mentioned in dialogs
	Chats []ChatClass
	// List of users mentioned in messages and chats
	Users []UserClass
}

// MessagesMessagesTypeID is TL type id of MessagesMessages.
const MessagesMessagesTypeID = 0x8c718e87

// construct implements constructor of MessagesMessagesClass.
func (m MessagesMessages) construct() MessagesMessagesClass { return &m }

// Ensuring interfaces in compile-time for MessagesMessages.
var (
	_ bin.Encoder     = &MessagesMessages{}
	_ bin.Decoder     = &MessagesMessages{}
	_ bin.BareEncoder = &MessagesMessages{}
	_ bin.BareDecoder = &MessagesMessages{}

	_ MessagesMessagesClass = &MessagesMessages{}
)

func (m *MessagesMessages) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Messages == nil) {
		return false
	}
	if !(m.Chats == nil) {
		return false
	}
	if !(m.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMessages) String() string {
	if m == nil {
		return "MessagesMessages(nil)"
	}
	type Alias MessagesMessages
	return fmt.Sprintf("MessagesMessages%+v", Alias(*m))
}

// FillFrom fills MessagesMessages from given interface.
func (m *MessagesMessages) FillFrom(from interface {
	GetMessages() (value []MessageClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	m.Messages = from.GetMessages()
	m.Chats = from.GetChats()
	m.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMessages) TypeID() uint32 {
	return MessagesMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMessages) TypeName() string {
	return "messages.messages"
}

// TypeInfo returns info about TL type.
func (m *MessagesMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.messages",
		ID:   MessagesMessagesTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagesMessages) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messages#8c718e87 as nil")
	}
	b.PutID(MessagesMessagesTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagesMessages) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messages#8c718e87 as nil")
	}
	b.PutVectorHeader(len(m.Messages))
	for idx, v := range m.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messages#8c718e87: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messages#8c718e87: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Chats))
	for idx, v := range m.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messages#8c718e87: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messages#8c718e87: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messages#8c718e87: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messages#8c718e87: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMessages) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messages#8c718e87 to nil")
	}
	if err := b.ConsumeID(MessagesMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.messages#8c718e87: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagesMessages) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messages#8c718e87 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messages#8c718e87: field messages: %w", err)
		}

		if headerLen > 0 {
			m.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messages#8c718e87: field messages: %w", err)
			}
			m.Messages = append(m.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messages#8c718e87: field chats: %w", err)
		}

		if headerLen > 0 {
			m.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messages#8c718e87: field chats: %w", err)
			}
			m.Chats = append(m.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messages#8c718e87: field users: %w", err)
		}

		if headerLen > 0 {
			m.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messages#8c718e87: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	return nil
}

// GetMessages returns value of Messages field.
func (m *MessagesMessages) GetMessages() (value []MessageClass) {
	return m.Messages
}

// GetChats returns value of Chats field.
func (m *MessagesMessages) GetChats() (value []ChatClass) {
	return m.Chats
}

// GetUsers returns value of Users field.
func (m *MessagesMessages) GetUsers() (value []UserClass) {
	return m.Users
}

// MapMessages returns field Messages wrapped in MessageClassArray helper.
func (m *MessagesMessages) MapMessages() (value MessageClassArray) {
	return MessageClassArray(m.Messages)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (m *MessagesMessages) MapChats() (value ChatClassArray) {
	return ChatClassArray(m.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (m *MessagesMessages) MapUsers() (value UserClassArray) {
	return UserClassArray(m.Users)
}

// MessagesMessagesSlice represents TL type `messages.messagesSlice#3a54685e`.
// Incomplete list of messages and auxiliary data.
//
// See https://core.telegram.org/constructor/messages.messagesSlice for reference.
type MessagesMessagesSlice struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, indicates that the results may be inexact
	Inexact bool
	// Total number of messages in the list
	Count int
	// Rate to use in the offset_rate parameter in the next call to messages.searchGlobal¹
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.searchGlobal
	//
	// Use SetNextRate and GetNextRate helpers.
	NextRate int
	// Indicates the absolute position of messages[0] within the total result set with count
	// count. This is useful, for example, if the result was fetched using offset_id, and we
	// need to display a progress/total counter (like photo 134 of 200, for all media in a
	// chat, we could simply use photo ${offset_id_offset} of ${count}.
	//
	// Use SetOffsetIDOffset and GetOffsetIDOffset helpers.
	OffsetIDOffset int
	// List of messages
	Messages []MessageClass
	// List of chats mentioned in messages
	Chats []ChatClass
	// List of users mentioned in messages and chats
	Users []UserClass
}

// MessagesMessagesSliceTypeID is TL type id of MessagesMessagesSlice.
const MessagesMessagesSliceTypeID = 0x3a54685e

// construct implements constructor of MessagesMessagesClass.
func (m MessagesMessagesSlice) construct() MessagesMessagesClass { return &m }

// Ensuring interfaces in compile-time for MessagesMessagesSlice.
var (
	_ bin.Encoder     = &MessagesMessagesSlice{}
	_ bin.Decoder     = &MessagesMessagesSlice{}
	_ bin.BareEncoder = &MessagesMessagesSlice{}
	_ bin.BareDecoder = &MessagesMessagesSlice{}

	_ MessagesMessagesClass = &MessagesMessagesSlice{}
)

func (m *MessagesMessagesSlice) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Inexact == false) {
		return false
	}
	if !(m.Count == 0) {
		return false
	}
	if !(m.NextRate == 0) {
		return false
	}
	if !(m.OffsetIDOffset == 0) {
		return false
	}
	if !(m.Messages == nil) {
		return false
	}
	if !(m.Chats == nil) {
		return false
	}
	if !(m.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMessagesSlice) String() string {
	if m == nil {
		return "MessagesMessagesSlice(nil)"
	}
	type Alias MessagesMessagesSlice
	return fmt.Sprintf("MessagesMessagesSlice%+v", Alias(*m))
}

// FillFrom fills MessagesMessagesSlice from given interface.
func (m *MessagesMessagesSlice) FillFrom(from interface {
	GetInexact() (value bool)
	GetCount() (value int)
	GetNextRate() (value int, ok bool)
	GetOffsetIDOffset() (value int, ok bool)
	GetMessages() (value []MessageClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	m.Inexact = from.GetInexact()
	m.Count = from.GetCount()
	if val, ok := from.GetNextRate(); ok {
		m.NextRate = val
	}

	if val, ok := from.GetOffsetIDOffset(); ok {
		m.OffsetIDOffset = val
	}

	m.Messages = from.GetMessages()
	m.Chats = from.GetChats()
	m.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMessagesSlice) TypeID() uint32 {
	return MessagesMessagesSliceTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMessagesSlice) TypeName() string {
	return "messages.messagesSlice"
}

// TypeInfo returns info about TL type.
func (m *MessagesMessagesSlice) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.messagesSlice",
		ID:   MessagesMessagesSliceTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Inexact",
			SchemaName: "inexact",
			Null:       !m.Flags.Has(1),
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "NextRate",
			SchemaName: "next_rate",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "OffsetIDOffset",
			SchemaName: "offset_id_offset",
			Null:       !m.Flags.Has(2),
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagesMessagesSlice) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messagesSlice#3a54685e as nil")
	}
	b.PutID(MessagesMessagesSliceTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagesMessagesSlice) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messagesSlice#3a54685e as nil")
	}
	if !(m.Inexact == false) {
		m.Flags.Set(1)
	}
	if !(m.NextRate == 0) {
		m.Flags.Set(0)
	}
	if !(m.OffsetIDOffset == 0) {
		m.Flags.Set(2)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.messagesSlice#3a54685e: field flags: %w", err)
	}
	b.PutInt(m.Count)
	if m.Flags.Has(0) {
		b.PutInt(m.NextRate)
	}
	if m.Flags.Has(2) {
		b.PutInt(m.OffsetIDOffset)
	}
	b.PutVectorHeader(len(m.Messages))
	for idx, v := range m.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messagesSlice#3a54685e: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messagesSlice#3a54685e: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Chats))
	for idx, v := range m.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messagesSlice#3a54685e: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messagesSlice#3a54685e: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messagesSlice#3a54685e: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messagesSlice#3a54685e: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMessagesSlice) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messagesSlice#3a54685e to nil")
	}
	if err := b.ConsumeID(MessagesMessagesSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagesMessagesSlice) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messagesSlice#3a54685e to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field flags: %w", err)
		}
	}
	m.Inexact = m.Flags.Has(1)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field count: %w", err)
		}
		m.Count = value
	}
	if m.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field next_rate: %w", err)
		}
		m.NextRate = value
	}
	if m.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field offset_id_offset: %w", err)
		}
		m.OffsetIDOffset = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field messages: %w", err)
		}

		if headerLen > 0 {
			m.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field messages: %w", err)
			}
			m.Messages = append(m.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field chats: %w", err)
		}

		if headerLen > 0 {
			m.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field chats: %w", err)
			}
			m.Chats = append(m.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field users: %w", err)
		}

		if headerLen > 0 {
			m.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messagesSlice#3a54685e: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	return nil
}

// SetInexact sets value of Inexact conditional field.
func (m *MessagesMessagesSlice) SetInexact(value bool) {
	if value {
		m.Flags.Set(1)
		m.Inexact = true
	} else {
		m.Flags.Unset(1)
		m.Inexact = false
	}
}

// GetInexact returns value of Inexact conditional field.
func (m *MessagesMessagesSlice) GetInexact() (value bool) {
	return m.Flags.Has(1)
}

// GetCount returns value of Count field.
func (m *MessagesMessagesSlice) GetCount() (value int) {
	return m.Count
}

// SetNextRate sets value of NextRate conditional field.
func (m *MessagesMessagesSlice) SetNextRate(value int) {
	m.Flags.Set(0)
	m.NextRate = value
}

// GetNextRate returns value of NextRate conditional field and
// boolean which is true if field was set.
func (m *MessagesMessagesSlice) GetNextRate() (value int, ok bool) {
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.NextRate, true
}

// SetOffsetIDOffset sets value of OffsetIDOffset conditional field.
func (m *MessagesMessagesSlice) SetOffsetIDOffset(value int) {
	m.Flags.Set(2)
	m.OffsetIDOffset = value
}

// GetOffsetIDOffset returns value of OffsetIDOffset conditional field and
// boolean which is true if field was set.
func (m *MessagesMessagesSlice) GetOffsetIDOffset() (value int, ok bool) {
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.OffsetIDOffset, true
}

// GetMessages returns value of Messages field.
func (m *MessagesMessagesSlice) GetMessages() (value []MessageClass) {
	return m.Messages
}

// GetChats returns value of Chats field.
func (m *MessagesMessagesSlice) GetChats() (value []ChatClass) {
	return m.Chats
}

// GetUsers returns value of Users field.
func (m *MessagesMessagesSlice) GetUsers() (value []UserClass) {
	return m.Users
}

// MapMessages returns field Messages wrapped in MessageClassArray helper.
func (m *MessagesMessagesSlice) MapMessages() (value MessageClassArray) {
	return MessageClassArray(m.Messages)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (m *MessagesMessagesSlice) MapChats() (value ChatClassArray) {
	return ChatClassArray(m.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (m *MessagesMessagesSlice) MapUsers() (value UserClassArray) {
	return UserClassArray(m.Users)
}

// MessagesChannelMessages represents TL type `messages.channelMessages#64479808`.
// Channel messages
//
// See https://core.telegram.org/constructor/messages.channelMessages for reference.
type MessagesChannelMessages struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, returned results may be inexact
	Inexact bool
	// Event count after generation¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int
	// Total number of results were found server-side (may not be all included here)
	Count int
	// Indicates the absolute position of messages[0] within the total result set with count
	// count. This is useful, for example, if the result was fetched using offset_id, and we
	// need to display a progress/total counter (like photo 134 of 200, for all media in a
	// chat, we could simply use photo ${offset_id_offset} of ${count}.
	//
	// Use SetOffsetIDOffset and GetOffsetIDOffset helpers.
	OffsetIDOffset int
	// Found messages
	Messages []MessageClass
	// Chats
	Chats []ChatClass
	// Users
	Users []UserClass
}

// MessagesChannelMessagesTypeID is TL type id of MessagesChannelMessages.
const MessagesChannelMessagesTypeID = 0x64479808

// construct implements constructor of MessagesMessagesClass.
func (c MessagesChannelMessages) construct() MessagesMessagesClass { return &c }

// Ensuring interfaces in compile-time for MessagesChannelMessages.
var (
	_ bin.Encoder     = &MessagesChannelMessages{}
	_ bin.Decoder     = &MessagesChannelMessages{}
	_ bin.BareEncoder = &MessagesChannelMessages{}
	_ bin.BareDecoder = &MessagesChannelMessages{}

	_ MessagesMessagesClass = &MessagesChannelMessages{}
)

func (c *MessagesChannelMessages) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Inexact == false) {
		return false
	}
	if !(c.Pts == 0) {
		return false
	}
	if !(c.Count == 0) {
		return false
	}
	if !(c.OffsetIDOffset == 0) {
		return false
	}
	if !(c.Messages == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesChannelMessages) String() string {
	if c == nil {
		return "MessagesChannelMessages(nil)"
	}
	type Alias MessagesChannelMessages
	return fmt.Sprintf("MessagesChannelMessages%+v", Alias(*c))
}

// FillFrom fills MessagesChannelMessages from given interface.
func (c *MessagesChannelMessages) FillFrom(from interface {
	GetInexact() (value bool)
	GetPts() (value int)
	GetCount() (value int)
	GetOffsetIDOffset() (value int, ok bool)
	GetMessages() (value []MessageClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	c.Inexact = from.GetInexact()
	c.Pts = from.GetPts()
	c.Count = from.GetCount()
	if val, ok := from.GetOffsetIDOffset(); ok {
		c.OffsetIDOffset = val
	}

	c.Messages = from.GetMessages()
	c.Chats = from.GetChats()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesChannelMessages) TypeID() uint32 {
	return MessagesChannelMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesChannelMessages) TypeName() string {
	return "messages.channelMessages"
}

// TypeInfo returns info about TL type.
func (c *MessagesChannelMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.channelMessages",
		ID:   MessagesChannelMessagesTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Inexact",
			SchemaName: "inexact",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "Pts",
			SchemaName: "pts",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "OffsetIDOffset",
			SchemaName: "offset_id_offset",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesChannelMessages) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.channelMessages#64479808 as nil")
	}
	b.PutID(MessagesChannelMessagesTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesChannelMessages) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.channelMessages#64479808 as nil")
	}
	if !(c.Inexact == false) {
		c.Flags.Set(1)
	}
	if !(c.OffsetIDOffset == 0) {
		c.Flags.Set(2)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.channelMessages#64479808: field flags: %w", err)
	}
	b.PutInt(c.Pts)
	b.PutInt(c.Count)
	if c.Flags.Has(2) {
		b.PutInt(c.OffsetIDOffset)
	}
	b.PutVectorHeader(len(c.Messages))
	for idx, v := range c.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.channelMessages#64479808: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.channelMessages#64479808: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.channelMessages#64479808: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.channelMessages#64479808: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.channelMessages#64479808: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.channelMessages#64479808: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesChannelMessages) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.channelMessages#64479808 to nil")
	}
	if err := b.ConsumeID(MessagesChannelMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.channelMessages#64479808: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesChannelMessages) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.channelMessages#64479808 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.channelMessages#64479808: field flags: %w", err)
		}
	}
	c.Inexact = c.Flags.Has(1)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.channelMessages#64479808: field pts: %w", err)
		}
		c.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.channelMessages#64479808: field count: %w", err)
		}
		c.Count = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.channelMessages#64479808: field offset_id_offset: %w", err)
		}
		c.OffsetIDOffset = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.channelMessages#64479808: field messages: %w", err)
		}

		if headerLen > 0 {
			c.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.channelMessages#64479808: field messages: %w", err)
			}
			c.Messages = append(c.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.channelMessages#64479808: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.channelMessages#64479808: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.channelMessages#64479808: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.channelMessages#64479808: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// SetInexact sets value of Inexact conditional field.
func (c *MessagesChannelMessages) SetInexact(value bool) {
	if value {
		c.Flags.Set(1)
		c.Inexact = true
	} else {
		c.Flags.Unset(1)
		c.Inexact = false
	}
}

// GetInexact returns value of Inexact conditional field.
func (c *MessagesChannelMessages) GetInexact() (value bool) {
	return c.Flags.Has(1)
}

// GetPts returns value of Pts field.
func (c *MessagesChannelMessages) GetPts() (value int) {
	return c.Pts
}

// GetCount returns value of Count field.
func (c *MessagesChannelMessages) GetCount() (value int) {
	return c.Count
}

// SetOffsetIDOffset sets value of OffsetIDOffset conditional field.
func (c *MessagesChannelMessages) SetOffsetIDOffset(value int) {
	c.Flags.Set(2)
	c.OffsetIDOffset = value
}

// GetOffsetIDOffset returns value of OffsetIDOffset conditional field and
// boolean which is true if field was set.
func (c *MessagesChannelMessages) GetOffsetIDOffset() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.OffsetIDOffset, true
}

// GetMessages returns value of Messages field.
func (c *MessagesChannelMessages) GetMessages() (value []MessageClass) {
	return c.Messages
}

// GetChats returns value of Chats field.
func (c *MessagesChannelMessages) GetChats() (value []ChatClass) {
	return c.Chats
}

// GetUsers returns value of Users field.
func (c *MessagesChannelMessages) GetUsers() (value []UserClass) {
	return c.Users
}

// MapMessages returns field Messages wrapped in MessageClassArray helper.
func (c *MessagesChannelMessages) MapMessages() (value MessageClassArray) {
	return MessageClassArray(c.Messages)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (c *MessagesChannelMessages) MapChats() (value ChatClassArray) {
	return ChatClassArray(c.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *MessagesChannelMessages) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}

// MessagesMessagesNotModified represents TL type `messages.messagesNotModified#74535f21`.
// No new messages matching the query were found
//
// See https://core.telegram.org/constructor/messages.messagesNotModified for reference.
type MessagesMessagesNotModified struct {
	// Number of results found server-side by the given query
	Count int
}

// MessagesMessagesNotModifiedTypeID is TL type id of MessagesMessagesNotModified.
const MessagesMessagesNotModifiedTypeID = 0x74535f21

// construct implements constructor of MessagesMessagesClass.
func (m MessagesMessagesNotModified) construct() MessagesMessagesClass { return &m }

// Ensuring interfaces in compile-time for MessagesMessagesNotModified.
var (
	_ bin.Encoder     = &MessagesMessagesNotModified{}
	_ bin.Decoder     = &MessagesMessagesNotModified{}
	_ bin.BareEncoder = &MessagesMessagesNotModified{}
	_ bin.BareDecoder = &MessagesMessagesNotModified{}

	_ MessagesMessagesClass = &MessagesMessagesNotModified{}
)

func (m *MessagesMessagesNotModified) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMessagesNotModified) String() string {
	if m == nil {
		return "MessagesMessagesNotModified(nil)"
	}
	type Alias MessagesMessagesNotModified
	return fmt.Sprintf("MessagesMessagesNotModified%+v", Alias(*m))
}

// FillFrom fills MessagesMessagesNotModified from given interface.
func (m *MessagesMessagesNotModified) FillFrom(from interface {
	GetCount() (value int)
}) {
	m.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMessagesNotModified) TypeID() uint32 {
	return MessagesMessagesNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMessagesNotModified) TypeName() string {
	return "messages.messagesNotModified"
}

// TypeInfo returns info about TL type.
func (m *MessagesMessagesNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.messagesNotModified",
		ID:   MessagesMessagesNotModifiedTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagesMessagesNotModified) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messagesNotModified#74535f21 as nil")
	}
	b.PutID(MessagesMessagesNotModifiedTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagesMessagesNotModified) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messagesNotModified#74535f21 as nil")
	}
	b.PutInt(m.Count)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMessagesNotModified) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messagesNotModified#74535f21 to nil")
	}
	if err := b.ConsumeID(MessagesMessagesNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.messagesNotModified#74535f21: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagesMessagesNotModified) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messagesNotModified#74535f21 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messagesNotModified#74535f21: field count: %w", err)
		}
		m.Count = value
	}
	return nil
}

// GetCount returns value of Count field.
func (m *MessagesMessagesNotModified) GetCount() (value int) {
	return m.Count
}

// MessagesMessagesClass represents messages.Messages generic type.
//
// See https://core.telegram.org/type/messages.Messages for reference.
//
// Example:
//  g, err := tg.DecodeMessagesMessages(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesMessages: // messages.messages#8c718e87
//  case *tg.MessagesMessagesSlice: // messages.messagesSlice#3a54685e
//  case *tg.MessagesChannelMessages: // messages.channelMessages#64479808
//  case *tg.MessagesMessagesNotModified: // messages.messagesNotModified#74535f21
//  default: panic(v)
//  }
type MessagesMessagesClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesMessagesClass

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

	// AsModified tries to map MessagesMessagesClass to ModifiedMessagesMessages.
	AsModified() (ModifiedMessagesMessages, bool)
}

// ModifiedMessagesMessages represents Modified subset of MessagesMessagesClass.
type ModifiedMessagesMessages interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesMessagesClass

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

	// List of messages
	GetMessages() (value []MessageClass)

	// List of chats mentioned in dialogs
	GetChats() (value []ChatClass)

	// List of users mentioned in messages and chats
	GetUsers() (value []UserClass)
}

// AsModified tries to map MessagesMessages to ModifiedMessagesMessages.
func (m *MessagesMessages) AsModified() (ModifiedMessagesMessages, bool) {
	value, ok := (MessagesMessagesClass(m)).(ModifiedMessagesMessages)
	return value, ok
}

// AsModified tries to map MessagesMessagesSlice to ModifiedMessagesMessages.
func (m *MessagesMessagesSlice) AsModified() (ModifiedMessagesMessages, bool) {
	value, ok := (MessagesMessagesClass(m)).(ModifiedMessagesMessages)
	return value, ok
}

// AsModified tries to map MessagesChannelMessages to ModifiedMessagesMessages.
func (c *MessagesChannelMessages) AsModified() (ModifiedMessagesMessages, bool) {
	value, ok := (MessagesMessagesClass(c)).(ModifiedMessagesMessages)
	return value, ok
}

// AsModified tries to map MessagesMessagesNotModified to ModifiedMessagesMessages.
func (m *MessagesMessagesNotModified) AsModified() (ModifiedMessagesMessages, bool) {
	value, ok := (MessagesMessagesClass(m)).(ModifiedMessagesMessages)
	return value, ok
}

// DecodeMessagesMessages implements binary de-serialization for MessagesMessagesClass.
func DecodeMessagesMessages(buf *bin.Buffer) (MessagesMessagesClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesMessagesTypeID:
		// Decoding messages.messages#8c718e87.
		v := MessagesMessages{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesMessagesClass: %w", err)
		}
		return &v, nil
	case MessagesMessagesSliceTypeID:
		// Decoding messages.messagesSlice#3a54685e.
		v := MessagesMessagesSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesMessagesClass: %w", err)
		}
		return &v, nil
	case MessagesChannelMessagesTypeID:
		// Decoding messages.channelMessages#64479808.
		v := MessagesChannelMessages{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesMessagesClass: %w", err)
		}
		return &v, nil
	case MessagesMessagesNotModifiedTypeID:
		// Decoding messages.messagesNotModified#74535f21.
		v := MessagesMessagesNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesMessagesClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesMessagesClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesMessages boxes the MessagesMessagesClass providing a helper.
type MessagesMessagesBox struct {
	Messages MessagesMessagesClass
}

// Decode implements bin.Decoder for MessagesMessagesBox.
func (b *MessagesMessagesBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesMessagesBox to nil")
	}
	v, err := DecodeMessagesMessages(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Messages = v
	return nil
}

// Encode implements bin.Encode for MessagesMessagesBox.
func (b *MessagesMessagesBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Messages == nil {
		return fmt.Errorf("unable to encode MessagesMessagesClass as nil")
	}
	return b.Messages.Encode(buf)
}
