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

// StorageFileUnknown represents TL type `storage.fileUnknown#aa963b05`.
// Unknown type.
//
// See https://core.telegram.org/constructor/storage.fileUnknown for reference.
type StorageFileUnknown struct {
}

// StorageFileUnknownTypeID is TL type id of StorageFileUnknown.
const StorageFileUnknownTypeID = 0xaa963b05

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileUnknown) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileUnknown.
var (
	_ bin.Encoder     = &StorageFileUnknown{}
	_ bin.Decoder     = &StorageFileUnknown{}
	_ bin.BareEncoder = &StorageFileUnknown{}
	_ bin.BareDecoder = &StorageFileUnknown{}

	_ StorageFileTypeClass = &StorageFileUnknown{}
)

func (f *StorageFileUnknown) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileUnknown) String() string {
	if f == nil {
		return "StorageFileUnknown(nil)"
	}
	type Alias StorageFileUnknown
	return fmt.Sprintf("StorageFileUnknown%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFileUnknown) TypeID() uint32 {
	return StorageFileUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFileUnknown) TypeName() string {
	return "storage.fileUnknown"
}

// TypeInfo returns info about TL type.
func (f *StorageFileUnknown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.fileUnknown",
		ID:   StorageFileUnknownTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFileUnknown) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileUnknown#aa963b05 as nil")
	}
	b.PutID(StorageFileUnknownTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFileUnknown) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileUnknown#aa963b05 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileUnknown) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileUnknown#aa963b05 to nil")
	}
	if err := b.ConsumeID(StorageFileUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileUnknown#aa963b05: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFileUnknown) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileUnknown#aa963b05 to nil")
	}
	return nil
}

// StorageFilePartial represents TL type `storage.filePartial#40bc6f52`.
// Part of a bigger file.
//
// See https://core.telegram.org/constructor/storage.filePartial for reference.
type StorageFilePartial struct {
}

// StorageFilePartialTypeID is TL type id of StorageFilePartial.
const StorageFilePartialTypeID = 0x40bc6f52

// construct implements constructor of StorageFileTypeClass.
func (f StorageFilePartial) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFilePartial.
var (
	_ bin.Encoder     = &StorageFilePartial{}
	_ bin.Decoder     = &StorageFilePartial{}
	_ bin.BareEncoder = &StorageFilePartial{}
	_ bin.BareDecoder = &StorageFilePartial{}

	_ StorageFileTypeClass = &StorageFilePartial{}
)

func (f *StorageFilePartial) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFilePartial) String() string {
	if f == nil {
		return "StorageFilePartial(nil)"
	}
	type Alias StorageFilePartial
	return fmt.Sprintf("StorageFilePartial%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFilePartial) TypeID() uint32 {
	return StorageFilePartialTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFilePartial) TypeName() string {
	return "storage.filePartial"
}

// TypeInfo returns info about TL type.
func (f *StorageFilePartial) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.filePartial",
		ID:   StorageFilePartialTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFilePartial) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePartial#40bc6f52 as nil")
	}
	b.PutID(StorageFilePartialTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFilePartial) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePartial#40bc6f52 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFilePartial) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePartial#40bc6f52 to nil")
	}
	if err := b.ConsumeID(StorageFilePartialTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.filePartial#40bc6f52: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFilePartial) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePartial#40bc6f52 to nil")
	}
	return nil
}

// StorageFileJpeg represents TL type `storage.fileJpeg#7efe0e`.
// JPEG image. MIME type: image/jpeg.
//
// See https://core.telegram.org/constructor/storage.fileJpeg for reference.
type StorageFileJpeg struct {
}

// StorageFileJpegTypeID is TL type id of StorageFileJpeg.
const StorageFileJpegTypeID = 0x7efe0e

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileJpeg) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileJpeg.
var (
	_ bin.Encoder     = &StorageFileJpeg{}
	_ bin.Decoder     = &StorageFileJpeg{}
	_ bin.BareEncoder = &StorageFileJpeg{}
	_ bin.BareDecoder = &StorageFileJpeg{}

	_ StorageFileTypeClass = &StorageFileJpeg{}
)

func (f *StorageFileJpeg) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileJpeg) String() string {
	if f == nil {
		return "StorageFileJpeg(nil)"
	}
	type Alias StorageFileJpeg
	return fmt.Sprintf("StorageFileJpeg%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFileJpeg) TypeID() uint32 {
	return StorageFileJpegTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFileJpeg) TypeName() string {
	return "storage.fileJpeg"
}

// TypeInfo returns info about TL type.
func (f *StorageFileJpeg) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.fileJpeg",
		ID:   StorageFileJpegTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFileJpeg) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileJpeg#7efe0e as nil")
	}
	b.PutID(StorageFileJpegTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFileJpeg) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileJpeg#7efe0e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileJpeg) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileJpeg#7efe0e to nil")
	}
	if err := b.ConsumeID(StorageFileJpegTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileJpeg#7efe0e: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFileJpeg) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileJpeg#7efe0e to nil")
	}
	return nil
}

// StorageFileGif represents TL type `storage.fileGif#cae1aadf`.
// GIF image. MIME type: image/gif.
//
// See https://core.telegram.org/constructor/storage.fileGif for reference.
type StorageFileGif struct {
}

// StorageFileGifTypeID is TL type id of StorageFileGif.
const StorageFileGifTypeID = 0xcae1aadf

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileGif) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileGif.
var (
	_ bin.Encoder     = &StorageFileGif{}
	_ bin.Decoder     = &StorageFileGif{}
	_ bin.BareEncoder = &StorageFileGif{}
	_ bin.BareDecoder = &StorageFileGif{}

	_ StorageFileTypeClass = &StorageFileGif{}
)

func (f *StorageFileGif) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileGif) String() string {
	if f == nil {
		return "StorageFileGif(nil)"
	}
	type Alias StorageFileGif
	return fmt.Sprintf("StorageFileGif%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFileGif) TypeID() uint32 {
	return StorageFileGifTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFileGif) TypeName() string {
	return "storage.fileGif"
}

// TypeInfo returns info about TL type.
func (f *StorageFileGif) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.fileGif",
		ID:   StorageFileGifTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFileGif) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileGif#cae1aadf as nil")
	}
	b.PutID(StorageFileGifTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFileGif) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileGif#cae1aadf as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileGif) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileGif#cae1aadf to nil")
	}
	if err := b.ConsumeID(StorageFileGifTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileGif#cae1aadf: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFileGif) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileGif#cae1aadf to nil")
	}
	return nil
}

// StorageFilePng represents TL type `storage.filePng#a4f63c0`.
// PNG image. MIME type: image/png.
//
// See https://core.telegram.org/constructor/storage.filePng for reference.
type StorageFilePng struct {
}

// StorageFilePngTypeID is TL type id of StorageFilePng.
const StorageFilePngTypeID = 0xa4f63c0

// construct implements constructor of StorageFileTypeClass.
func (f StorageFilePng) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFilePng.
var (
	_ bin.Encoder     = &StorageFilePng{}
	_ bin.Decoder     = &StorageFilePng{}
	_ bin.BareEncoder = &StorageFilePng{}
	_ bin.BareDecoder = &StorageFilePng{}

	_ StorageFileTypeClass = &StorageFilePng{}
)

func (f *StorageFilePng) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFilePng) String() string {
	if f == nil {
		return "StorageFilePng(nil)"
	}
	type Alias StorageFilePng
	return fmt.Sprintf("StorageFilePng%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFilePng) TypeID() uint32 {
	return StorageFilePngTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFilePng) TypeName() string {
	return "storage.filePng"
}

// TypeInfo returns info about TL type.
func (f *StorageFilePng) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.filePng",
		ID:   StorageFilePngTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFilePng) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePng#a4f63c0 as nil")
	}
	b.PutID(StorageFilePngTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFilePng) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePng#a4f63c0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFilePng) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePng#a4f63c0 to nil")
	}
	if err := b.ConsumeID(StorageFilePngTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.filePng#a4f63c0: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFilePng) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePng#a4f63c0 to nil")
	}
	return nil
}

// StorageFilePdf represents TL type `storage.filePdf#ae1e508d`.
// PDF document image. MIME type: application/pdf.
//
// See https://core.telegram.org/constructor/storage.filePdf for reference.
type StorageFilePdf struct {
}

// StorageFilePdfTypeID is TL type id of StorageFilePdf.
const StorageFilePdfTypeID = 0xae1e508d

// construct implements constructor of StorageFileTypeClass.
func (f StorageFilePdf) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFilePdf.
var (
	_ bin.Encoder     = &StorageFilePdf{}
	_ bin.Decoder     = &StorageFilePdf{}
	_ bin.BareEncoder = &StorageFilePdf{}
	_ bin.BareDecoder = &StorageFilePdf{}

	_ StorageFileTypeClass = &StorageFilePdf{}
)

func (f *StorageFilePdf) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFilePdf) String() string {
	if f == nil {
		return "StorageFilePdf(nil)"
	}
	type Alias StorageFilePdf
	return fmt.Sprintf("StorageFilePdf%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFilePdf) TypeID() uint32 {
	return StorageFilePdfTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFilePdf) TypeName() string {
	return "storage.filePdf"
}

// TypeInfo returns info about TL type.
func (f *StorageFilePdf) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.filePdf",
		ID:   StorageFilePdfTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFilePdf) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePdf#ae1e508d as nil")
	}
	b.PutID(StorageFilePdfTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFilePdf) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePdf#ae1e508d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFilePdf) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePdf#ae1e508d to nil")
	}
	if err := b.ConsumeID(StorageFilePdfTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.filePdf#ae1e508d: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFilePdf) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePdf#ae1e508d to nil")
	}
	return nil
}

// StorageFileMp3 represents TL type `storage.fileMp3#528a0677`.
// Mp3 audio. MIME type: audio/mpeg.
//
// See https://core.telegram.org/constructor/storage.fileMp3 for reference.
type StorageFileMp3 struct {
}

// StorageFileMp3TypeID is TL type id of StorageFileMp3.
const StorageFileMp3TypeID = 0x528a0677

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileMp3) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileMp3.
var (
	_ bin.Encoder     = &StorageFileMp3{}
	_ bin.Decoder     = &StorageFileMp3{}
	_ bin.BareEncoder = &StorageFileMp3{}
	_ bin.BareDecoder = &StorageFileMp3{}

	_ StorageFileTypeClass = &StorageFileMp3{}
)

func (f *StorageFileMp3) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileMp3) String() string {
	if f == nil {
		return "StorageFileMp3(nil)"
	}
	type Alias StorageFileMp3
	return fmt.Sprintf("StorageFileMp3%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFileMp3) TypeID() uint32 {
	return StorageFileMp3TypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFileMp3) TypeName() string {
	return "storage.fileMp3"
}

// TypeInfo returns info about TL type.
func (f *StorageFileMp3) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.fileMp3",
		ID:   StorageFileMp3TypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFileMp3) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMp3#528a0677 as nil")
	}
	b.PutID(StorageFileMp3TypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFileMp3) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMp3#528a0677 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileMp3) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMp3#528a0677 to nil")
	}
	if err := b.ConsumeID(StorageFileMp3TypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileMp3#528a0677: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFileMp3) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMp3#528a0677 to nil")
	}
	return nil
}

// StorageFileMov represents TL type `storage.fileMov#4b09ebbc`.
// Quicktime video. MIME type: video/quicktime.
//
// See https://core.telegram.org/constructor/storage.fileMov for reference.
type StorageFileMov struct {
}

// StorageFileMovTypeID is TL type id of StorageFileMov.
const StorageFileMovTypeID = 0x4b09ebbc

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileMov) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileMov.
var (
	_ bin.Encoder     = &StorageFileMov{}
	_ bin.Decoder     = &StorageFileMov{}
	_ bin.BareEncoder = &StorageFileMov{}
	_ bin.BareDecoder = &StorageFileMov{}

	_ StorageFileTypeClass = &StorageFileMov{}
)

func (f *StorageFileMov) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileMov) String() string {
	if f == nil {
		return "StorageFileMov(nil)"
	}
	type Alias StorageFileMov
	return fmt.Sprintf("StorageFileMov%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFileMov) TypeID() uint32 {
	return StorageFileMovTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFileMov) TypeName() string {
	return "storage.fileMov"
}

// TypeInfo returns info about TL type.
func (f *StorageFileMov) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.fileMov",
		ID:   StorageFileMovTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFileMov) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMov#4b09ebbc as nil")
	}
	b.PutID(StorageFileMovTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFileMov) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMov#4b09ebbc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileMov) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMov#4b09ebbc to nil")
	}
	if err := b.ConsumeID(StorageFileMovTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileMov#4b09ebbc: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFileMov) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMov#4b09ebbc to nil")
	}
	return nil
}

// StorageFileMp4 represents TL type `storage.fileMp4#b3cea0e4`.
// MPEG-4 video. MIME type: video/mp4.
//
// See https://core.telegram.org/constructor/storage.fileMp4 for reference.
type StorageFileMp4 struct {
}

// StorageFileMp4TypeID is TL type id of StorageFileMp4.
const StorageFileMp4TypeID = 0xb3cea0e4

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileMp4) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileMp4.
var (
	_ bin.Encoder     = &StorageFileMp4{}
	_ bin.Decoder     = &StorageFileMp4{}
	_ bin.BareEncoder = &StorageFileMp4{}
	_ bin.BareDecoder = &StorageFileMp4{}

	_ StorageFileTypeClass = &StorageFileMp4{}
)

func (f *StorageFileMp4) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileMp4) String() string {
	if f == nil {
		return "StorageFileMp4(nil)"
	}
	type Alias StorageFileMp4
	return fmt.Sprintf("StorageFileMp4%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFileMp4) TypeID() uint32 {
	return StorageFileMp4TypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFileMp4) TypeName() string {
	return "storage.fileMp4"
}

// TypeInfo returns info about TL type.
func (f *StorageFileMp4) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.fileMp4",
		ID:   StorageFileMp4TypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFileMp4) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMp4#b3cea0e4 as nil")
	}
	b.PutID(StorageFileMp4TypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFileMp4) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMp4#b3cea0e4 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileMp4) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMp4#b3cea0e4 to nil")
	}
	if err := b.ConsumeID(StorageFileMp4TypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileMp4#b3cea0e4: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFileMp4) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMp4#b3cea0e4 to nil")
	}
	return nil
}

// StorageFileWebp represents TL type `storage.fileWebp#1081464c`.
// WEBP image. MIME type: image/webp.
//
// See https://core.telegram.org/constructor/storage.fileWebp for reference.
type StorageFileWebp struct {
}

// StorageFileWebpTypeID is TL type id of StorageFileWebp.
const StorageFileWebpTypeID = 0x1081464c

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileWebp) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileWebp.
var (
	_ bin.Encoder     = &StorageFileWebp{}
	_ bin.Decoder     = &StorageFileWebp{}
	_ bin.BareEncoder = &StorageFileWebp{}
	_ bin.BareDecoder = &StorageFileWebp{}

	_ StorageFileTypeClass = &StorageFileWebp{}
)

func (f *StorageFileWebp) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileWebp) String() string {
	if f == nil {
		return "StorageFileWebp(nil)"
	}
	type Alias StorageFileWebp
	return fmt.Sprintf("StorageFileWebp%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StorageFileWebp) TypeID() uint32 {
	return StorageFileWebpTypeID
}

// TypeName returns name of type in TL schema.
func (*StorageFileWebp) TypeName() string {
	return "storage.fileWebp"
}

// TypeInfo returns info about TL type.
func (f *StorageFileWebp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storage.fileWebp",
		ID:   StorageFileWebpTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *StorageFileWebp) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileWebp#1081464c as nil")
	}
	b.PutID(StorageFileWebpTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *StorageFileWebp) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileWebp#1081464c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileWebp) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileWebp#1081464c to nil")
	}
	if err := b.ConsumeID(StorageFileWebpTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileWebp#1081464c: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *StorageFileWebp) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileWebp#1081464c to nil")
	}
	return nil
}

// StorageFileTypeClass represents storage.FileType generic type.
//
// See https://core.telegram.org/type/storage.FileType for reference.
//
// Example:
//  g, err := tg.DecodeStorageFileType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.StorageFileUnknown: // storage.fileUnknown#aa963b05
//  case *tg.StorageFilePartial: // storage.filePartial#40bc6f52
//  case *tg.StorageFileJpeg: // storage.fileJpeg#7efe0e
//  case *tg.StorageFileGif: // storage.fileGif#cae1aadf
//  case *tg.StorageFilePng: // storage.filePng#a4f63c0
//  case *tg.StorageFilePdf: // storage.filePdf#ae1e508d
//  case *tg.StorageFileMp3: // storage.fileMp3#528a0677
//  case *tg.StorageFileMov: // storage.fileMov#4b09ebbc
//  case *tg.StorageFileMp4: // storage.fileMp4#b3cea0e4
//  case *tg.StorageFileWebp: // storage.fileWebp#1081464c
//  default: panic(v)
//  }
type StorageFileTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StorageFileTypeClass

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

// DecodeStorageFileType implements binary de-serialization for StorageFileTypeClass.
func DecodeStorageFileType(buf *bin.Buffer) (StorageFileTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StorageFileUnknownTypeID:
		// Decoding storage.fileUnknown#aa963b05.
		v := StorageFileUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFilePartialTypeID:
		// Decoding storage.filePartial#40bc6f52.
		v := StorageFilePartial{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileJpegTypeID:
		// Decoding storage.fileJpeg#7efe0e.
		v := StorageFileJpeg{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileGifTypeID:
		// Decoding storage.fileGif#cae1aadf.
		v := StorageFileGif{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFilePngTypeID:
		// Decoding storage.filePng#a4f63c0.
		v := StorageFilePng{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFilePdfTypeID:
		// Decoding storage.filePdf#ae1e508d.
		v := StorageFilePdf{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileMp3TypeID:
		// Decoding storage.fileMp3#528a0677.
		v := StorageFileMp3{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileMovTypeID:
		// Decoding storage.fileMov#4b09ebbc.
		v := StorageFileMov{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileMp4TypeID:
		// Decoding storage.fileMp4#b3cea0e4.
		v := StorageFileMp4{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileWebpTypeID:
		// Decoding storage.fileWebp#1081464c.
		v := StorageFileWebp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// StorageFileType boxes the StorageFileTypeClass providing a helper.
type StorageFileTypeBox struct {
	FileType StorageFileTypeClass
}

// Decode implements bin.Decoder for StorageFileTypeBox.
func (b *StorageFileTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StorageFileTypeBox to nil")
	}
	v, err := DecodeStorageFileType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FileType = v
	return nil
}

// Encode implements bin.Encode for StorageFileTypeBox.
func (b *StorageFileTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FileType == nil {
		return fmt.Errorf("unable to encode StorageFileTypeClass as nil")
	}
	return b.FileType.Encode(buf)
}
