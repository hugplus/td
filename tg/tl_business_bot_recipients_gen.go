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

// BusinessBotRecipients represents TL type `businessBotRecipients#b88cf373`.
//
// See https://core.telegram.org/constructor/businessBotRecipients for reference.
type BusinessBotRecipients struct {
	// Flags field of BusinessBotRecipients.
	Flags bin.Fields
	// ExistingChats field of BusinessBotRecipients.
	ExistingChats bool
	// NewChats field of BusinessBotRecipients.
	NewChats bool
	// Contacts field of BusinessBotRecipients.
	Contacts bool
	// NonContacts field of BusinessBotRecipients.
	NonContacts bool
	// ExcludeSelected field of BusinessBotRecipients.
	ExcludeSelected bool
	// Users field of BusinessBotRecipients.
	//
	// Use SetUsers and GetUsers helpers.
	Users []int64
	// ExcludeUsers field of BusinessBotRecipients.
	//
	// Use SetExcludeUsers and GetExcludeUsers helpers.
	ExcludeUsers []int64
}

// BusinessBotRecipientsTypeID is TL type id of BusinessBotRecipients.
const BusinessBotRecipientsTypeID = 0xb88cf373

// Ensuring interfaces in compile-time for BusinessBotRecipients.
var (
	_ bin.Encoder     = &BusinessBotRecipients{}
	_ bin.Decoder     = &BusinessBotRecipients{}
	_ bin.BareEncoder = &BusinessBotRecipients{}
	_ bin.BareDecoder = &BusinessBotRecipients{}
)

func (b *BusinessBotRecipients) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.ExistingChats == false) {
		return false
	}
	if !(b.NewChats == false) {
		return false
	}
	if !(b.Contacts == false) {
		return false
	}
	if !(b.NonContacts == false) {
		return false
	}
	if !(b.ExcludeSelected == false) {
		return false
	}
	if !(b.Users == nil) {
		return false
	}
	if !(b.ExcludeUsers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessBotRecipients) String() string {
	if b == nil {
		return "BusinessBotRecipients(nil)"
	}
	type Alias BusinessBotRecipients
	return fmt.Sprintf("BusinessBotRecipients%+v", Alias(*b))
}

// FillFrom fills BusinessBotRecipients from given interface.
func (b *BusinessBotRecipients) FillFrom(from interface {
	GetExistingChats() (value bool)
	GetNewChats() (value bool)
	GetContacts() (value bool)
	GetNonContacts() (value bool)
	GetExcludeSelected() (value bool)
	GetUsers() (value []int64, ok bool)
	GetExcludeUsers() (value []int64, ok bool)
}) {
	b.ExistingChats = from.GetExistingChats()
	b.NewChats = from.GetNewChats()
	b.Contacts = from.GetContacts()
	b.NonContacts = from.GetNonContacts()
	b.ExcludeSelected = from.GetExcludeSelected()
	if val, ok := from.GetUsers(); ok {
		b.Users = val
	}

	if val, ok := from.GetExcludeUsers(); ok {
		b.ExcludeUsers = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessBotRecipients) TypeID() uint32 {
	return BusinessBotRecipientsTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessBotRecipients) TypeName() string {
	return "businessBotRecipients"
}

// TypeInfo returns info about TL type.
func (b *BusinessBotRecipients) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessBotRecipients",
		ID:   BusinessBotRecipientsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ExistingChats",
			SchemaName: "existing_chats",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "NewChats",
			SchemaName: "new_chats",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "Contacts",
			SchemaName: "contacts",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "NonContacts",
			SchemaName: "non_contacts",
			Null:       !b.Flags.Has(3),
		},
		{
			Name:       "ExcludeSelected",
			SchemaName: "exclude_selected",
			Null:       !b.Flags.Has(5),
		},
		{
			Name:       "Users",
			SchemaName: "users",
			Null:       !b.Flags.Has(4),
		},
		{
			Name:       "ExcludeUsers",
			SchemaName: "exclude_users",
			Null:       !b.Flags.Has(6),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (b *BusinessBotRecipients) SetFlags() {
	if !(b.ExistingChats == false) {
		b.Flags.Set(0)
	}
	if !(b.NewChats == false) {
		b.Flags.Set(1)
	}
	if !(b.Contacts == false) {
		b.Flags.Set(2)
	}
	if !(b.NonContacts == false) {
		b.Flags.Set(3)
	}
	if !(b.ExcludeSelected == false) {
		b.Flags.Set(5)
	}
	if !(b.Users == nil) {
		b.Flags.Set(4)
	}
	if !(b.ExcludeUsers == nil) {
		b.Flags.Set(6)
	}
}

// Encode implements bin.Encoder.
func (b *BusinessBotRecipients) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessBotRecipients#b88cf373 as nil")
	}
	buf.PutID(BusinessBotRecipientsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessBotRecipients) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessBotRecipients#b88cf373 as nil")
	}
	b.SetFlags()
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessBotRecipients#b88cf373: field flags: %w", err)
	}
	if b.Flags.Has(4) {
		buf.PutVectorHeader(len(b.Users))
		for _, v := range b.Users {
			buf.PutLong(v)
		}
	}
	if b.Flags.Has(6) {
		buf.PutVectorHeader(len(b.ExcludeUsers))
		for _, v := range b.ExcludeUsers {
			buf.PutLong(v)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessBotRecipients) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessBotRecipients#b88cf373 to nil")
	}
	if err := buf.ConsumeID(BusinessBotRecipientsTypeID); err != nil {
		return fmt.Errorf("unable to decode businessBotRecipients#b88cf373: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessBotRecipients) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessBotRecipients#b88cf373 to nil")
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessBotRecipients#b88cf373: field flags: %w", err)
		}
	}
	b.ExistingChats = b.Flags.Has(0)
	b.NewChats = b.Flags.Has(1)
	b.Contacts = b.Flags.Has(2)
	b.NonContacts = b.Flags.Has(3)
	b.ExcludeSelected = b.Flags.Has(5)
	if b.Flags.Has(4) {
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode businessBotRecipients#b88cf373: field users: %w", err)
		}

		if headerLen > 0 {
			b.Users = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := buf.Long()
			if err != nil {
				return fmt.Errorf("unable to decode businessBotRecipients#b88cf373: field users: %w", err)
			}
			b.Users = append(b.Users, value)
		}
	}
	if b.Flags.Has(6) {
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode businessBotRecipients#b88cf373: field exclude_users: %w", err)
		}

		if headerLen > 0 {
			b.ExcludeUsers = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := buf.Long()
			if err != nil {
				return fmt.Errorf("unable to decode businessBotRecipients#b88cf373: field exclude_users: %w", err)
			}
			b.ExcludeUsers = append(b.ExcludeUsers, value)
		}
	}
	return nil
}

// SetExistingChats sets value of ExistingChats conditional field.
func (b *BusinessBotRecipients) SetExistingChats(value bool) {
	if value {
		b.Flags.Set(0)
		b.ExistingChats = true
	} else {
		b.Flags.Unset(0)
		b.ExistingChats = false
	}
}

// GetExistingChats returns value of ExistingChats conditional field.
func (b *BusinessBotRecipients) GetExistingChats() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(0)
}

// SetNewChats sets value of NewChats conditional field.
func (b *BusinessBotRecipients) SetNewChats(value bool) {
	if value {
		b.Flags.Set(1)
		b.NewChats = true
	} else {
		b.Flags.Unset(1)
		b.NewChats = false
	}
}

// GetNewChats returns value of NewChats conditional field.
func (b *BusinessBotRecipients) GetNewChats() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(1)
}

// SetContacts sets value of Contacts conditional field.
func (b *BusinessBotRecipients) SetContacts(value bool) {
	if value {
		b.Flags.Set(2)
		b.Contacts = true
	} else {
		b.Flags.Unset(2)
		b.Contacts = false
	}
}

// GetContacts returns value of Contacts conditional field.
func (b *BusinessBotRecipients) GetContacts() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(2)
}

// SetNonContacts sets value of NonContacts conditional field.
func (b *BusinessBotRecipients) SetNonContacts(value bool) {
	if value {
		b.Flags.Set(3)
		b.NonContacts = true
	} else {
		b.Flags.Unset(3)
		b.NonContacts = false
	}
}

// GetNonContacts returns value of NonContacts conditional field.
func (b *BusinessBotRecipients) GetNonContacts() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(3)
}

// SetExcludeSelected sets value of ExcludeSelected conditional field.
func (b *BusinessBotRecipients) SetExcludeSelected(value bool) {
	if value {
		b.Flags.Set(5)
		b.ExcludeSelected = true
	} else {
		b.Flags.Unset(5)
		b.ExcludeSelected = false
	}
}

// GetExcludeSelected returns value of ExcludeSelected conditional field.
func (b *BusinessBotRecipients) GetExcludeSelected() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(5)
}

// SetUsers sets value of Users conditional field.
func (b *BusinessBotRecipients) SetUsers(value []int64) {
	b.Flags.Set(4)
	b.Users = value
}

// GetUsers returns value of Users conditional field and
// boolean which is true if field was set.
func (b *BusinessBotRecipients) GetUsers() (value []int64, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(4) {
		return value, false
	}
	return b.Users, true
}

// SetExcludeUsers sets value of ExcludeUsers conditional field.
func (b *BusinessBotRecipients) SetExcludeUsers(value []int64) {
	b.Flags.Set(6)
	b.ExcludeUsers = value
}

// GetExcludeUsers returns value of ExcludeUsers conditional field and
// boolean which is true if field was set.
func (b *BusinessBotRecipients) GetExcludeUsers() (value []int64, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(6) {
		return value, false
	}
	return b.ExcludeUsers, true
}
