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

// InputChatPhotoEmpty represents TL type `inputChatPhotoEmpty#1ca48f57`.
// Empty constructor, remove group photo.
//
// See https://core.telegram.org/constructor/inputChatPhotoEmpty for reference.
type InputChatPhotoEmpty struct {
}

// InputChatPhotoEmptyTypeID is TL type id of InputChatPhotoEmpty.
const InputChatPhotoEmptyTypeID = 0x1ca48f57

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhotoEmpty) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhotoEmpty.
var (
	_ bin.Encoder     = &InputChatPhotoEmpty{}
	_ bin.Decoder     = &InputChatPhotoEmpty{}
	_ bin.BareEncoder = &InputChatPhotoEmpty{}
	_ bin.BareDecoder = &InputChatPhotoEmpty{}

	_ InputChatPhotoClass = &InputChatPhotoEmpty{}
)

func (i *InputChatPhotoEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatPhotoEmpty) String() string {
	if i == nil {
		return "InputChatPhotoEmpty(nil)"
	}
	type Alias InputChatPhotoEmpty
	return fmt.Sprintf("InputChatPhotoEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatPhotoEmpty) TypeID() uint32 {
	return InputChatPhotoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatPhotoEmpty) TypeName() string {
	return "inputChatPhotoEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputChatPhotoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatPhotoEmpty",
		ID:   InputChatPhotoEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChatPhotoEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoEmpty#1ca48f57 as nil")
	}
	b.PutID(InputChatPhotoEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatPhotoEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoEmpty#1ca48f57 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatPhotoEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoEmpty#1ca48f57 to nil")
	}
	if err := b.ConsumeID(InputChatPhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatPhotoEmpty#1ca48f57: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatPhotoEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoEmpty#1ca48f57 to nil")
	}
	return nil
}

// InputChatUploadedPhoto represents TL type `inputChatUploadedPhoto#c642724e`.
// New photo to be set as group profile photo.
//
// See https://core.telegram.org/constructor/inputChatUploadedPhoto for reference.
type InputChatUploadedPhoto struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// File saved in parts using the method upload.saveFilePart¹
	//
	// Links:
	//  1) https://core.telegram.org/method/upload.saveFilePart
	//
	// Use SetFile and GetFile helpers.
	File InputFileClass
	// Square video for animated profile picture
	//
	// Use SetVideo and GetVideo helpers.
	Video InputFileClass
	// Timestamp that should be shown as static preview to the user (seconds)
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
}

// InputChatUploadedPhotoTypeID is TL type id of InputChatUploadedPhoto.
const InputChatUploadedPhotoTypeID = 0xc642724e

// construct implements constructor of InputChatPhotoClass.
func (i InputChatUploadedPhoto) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatUploadedPhoto.
var (
	_ bin.Encoder     = &InputChatUploadedPhoto{}
	_ bin.Decoder     = &InputChatUploadedPhoto{}
	_ bin.BareEncoder = &InputChatUploadedPhoto{}
	_ bin.BareDecoder = &InputChatUploadedPhoto{}

	_ InputChatPhotoClass = &InputChatUploadedPhoto{}
)

func (i *InputChatUploadedPhoto) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.File == nil) {
		return false
	}
	if !(i.Video == nil) {
		return false
	}
	if !(i.VideoStartTs == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatUploadedPhoto) String() string {
	if i == nil {
		return "InputChatUploadedPhoto(nil)"
	}
	type Alias InputChatUploadedPhoto
	return fmt.Sprintf("InputChatUploadedPhoto%+v", Alias(*i))
}

// FillFrom fills InputChatUploadedPhoto from given interface.
func (i *InputChatUploadedPhoto) FillFrom(from interface {
	GetFile() (value InputFileClass, ok bool)
	GetVideo() (value InputFileClass, ok bool)
	GetVideoStartTs() (value float64, ok bool)
}) {
	if val, ok := from.GetFile(); ok {
		i.File = val
	}

	if val, ok := from.GetVideo(); ok {
		i.Video = val
	}

	if val, ok := from.GetVideoStartTs(); ok {
		i.VideoStartTs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatUploadedPhoto) TypeID() uint32 {
	return InputChatUploadedPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatUploadedPhoto) TypeName() string {
	return "inputChatUploadedPhoto"
}

// TypeInfo returns info about TL type.
func (i *InputChatUploadedPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatUploadedPhoto",
		ID:   InputChatUploadedPhotoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "File",
			SchemaName: "file",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "Video",
			SchemaName: "video",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "VideoStartTs",
			SchemaName: "video_start_ts",
			Null:       !i.Flags.Has(2),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChatUploadedPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatUploadedPhoto#c642724e as nil")
	}
	b.PutID(InputChatUploadedPhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatUploadedPhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatUploadedPhoto#c642724e as nil")
	}
	if !(i.File == nil) {
		i.Flags.Set(0)
	}
	if !(i.Video == nil) {
		i.Flags.Set(1)
	}
	if !(i.VideoStartTs == 0) {
		i.Flags.Set(2)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field flags: %w", err)
	}
	if i.Flags.Has(0) {
		if i.File == nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field file is nil")
		}
		if err := i.File.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field file: %w", err)
		}
	}
	if i.Flags.Has(1) {
		if i.Video == nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field video is nil")
		}
		if err := i.Video.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field video: %w", err)
		}
	}
	if i.Flags.Has(2) {
		b.PutDouble(i.VideoStartTs)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatUploadedPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatUploadedPhoto#c642724e to nil")
	}
	if err := b.ConsumeID(InputChatUploadedPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatUploadedPhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatUploadedPhoto#c642724e to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: field flags: %w", err)
		}
	}
	if i.Flags.Has(0) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: field file: %w", err)
		}
		i.File = value
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: field video: %w", err)
		}
		i.Video = value
	}
	if i.Flags.Has(2) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: field video_start_ts: %w", err)
		}
		i.VideoStartTs = value
	}
	return nil
}

// SetFile sets value of File conditional field.
func (i *InputChatUploadedPhoto) SetFile(value InputFileClass) {
	i.Flags.Set(0)
	i.File = value
}

// GetFile returns value of File conditional field and
// boolean which is true if field was set.
func (i *InputChatUploadedPhoto) GetFile() (value InputFileClass, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.File, true
}

// SetVideo sets value of Video conditional field.
func (i *InputChatUploadedPhoto) SetVideo(value InputFileClass) {
	i.Flags.Set(1)
	i.Video = value
}

// GetVideo returns value of Video conditional field and
// boolean which is true if field was set.
func (i *InputChatUploadedPhoto) GetVideo() (value InputFileClass, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Video, true
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (i *InputChatUploadedPhoto) SetVideoStartTs(value float64) {
	i.Flags.Set(2)
	i.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (i *InputChatUploadedPhoto) GetVideoStartTs() (value float64, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.VideoStartTs, true
}

// InputChatPhoto represents TL type `inputChatPhoto#8953ad37`.
// Existing photo to be set as a chat profile photo.
//
// See https://core.telegram.org/constructor/inputChatPhoto for reference.
type InputChatPhoto struct {
	// Existing photo
	ID InputPhotoClass
}

// InputChatPhotoTypeID is TL type id of InputChatPhoto.
const InputChatPhotoTypeID = 0x8953ad37

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhoto) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhoto.
var (
	_ bin.Encoder     = &InputChatPhoto{}
	_ bin.Decoder     = &InputChatPhoto{}
	_ bin.BareEncoder = &InputChatPhoto{}
	_ bin.BareDecoder = &InputChatPhoto{}

	_ InputChatPhotoClass = &InputChatPhoto{}
)

func (i *InputChatPhoto) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatPhoto) String() string {
	if i == nil {
		return "InputChatPhoto(nil)"
	}
	type Alias InputChatPhoto
	return fmt.Sprintf("InputChatPhoto%+v", Alias(*i))
}

// FillFrom fills InputChatPhoto from given interface.
func (i *InputChatPhoto) FillFrom(from interface {
	GetID() (value InputPhotoClass)
}) {
	i.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatPhoto) TypeID() uint32 {
	return InputChatPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatPhoto) TypeName() string {
	return "inputChatPhoto"
}

// TypeInfo returns info about TL type.
func (i *InputChatPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatPhoto",
		ID:   InputChatPhotoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChatPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhoto#8953ad37 as nil")
	}
	b.PutID(InputChatPhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatPhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhoto#8953ad37 as nil")
	}
	if i.ID == nil {
		return fmt.Errorf("unable to encode inputChatPhoto#8953ad37: field id is nil")
	}
	if err := i.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChatPhoto#8953ad37: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhoto#8953ad37 to nil")
	}
	if err := b.ConsumeID(InputChatPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatPhoto#8953ad37: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatPhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhoto#8953ad37 to nil")
	}
	{
		value, err := DecodeInputPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatPhoto#8953ad37: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputChatPhoto) GetID() (value InputPhotoClass) {
	return i.ID
}

// InputChatPhotoClass represents InputChatPhoto generic type.
//
// See https://core.telegram.org/type/InputChatPhoto for reference.
//
// Example:
//  g, err := tg.DecodeInputChatPhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputChatPhotoEmpty: // inputChatPhotoEmpty#1ca48f57
//  case *tg.InputChatUploadedPhoto: // inputChatUploadedPhoto#c642724e
//  case *tg.InputChatPhoto: // inputChatPhoto#8953ad37
//  default: panic(v)
//  }
type InputChatPhotoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputChatPhotoClass

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

// DecodeInputChatPhoto implements binary de-serialization for InputChatPhotoClass.
func DecodeInputChatPhoto(buf *bin.Buffer) (InputChatPhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputChatPhotoEmptyTypeID:
		// Decoding inputChatPhotoEmpty#1ca48f57.
		v := InputChatPhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	case InputChatUploadedPhotoTypeID:
		// Decoding inputChatUploadedPhoto#c642724e.
		v := InputChatUploadedPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	case InputChatPhotoTypeID:
		// Decoding inputChatPhoto#8953ad37.
		v := InputChatPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputChatPhoto boxes the InputChatPhotoClass providing a helper.
type InputChatPhotoBox struct {
	InputChatPhoto InputChatPhotoClass
}

// Decode implements bin.Decoder for InputChatPhotoBox.
func (b *InputChatPhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputChatPhotoBox to nil")
	}
	v, err := DecodeInputChatPhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputChatPhoto = v
	return nil
}

// Encode implements bin.Encode for InputChatPhotoBox.
func (b *InputChatPhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputChatPhoto == nil {
		return fmt.Errorf("unable to encode InputChatPhotoClass as nil")
	}
	return b.InputChatPhoto.Encode(buf)
}
