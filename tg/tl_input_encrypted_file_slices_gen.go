//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// InputEncryptedFileClassArray is adapter for slice of InputEncryptedFileClass.
type InputEncryptedFileClassArray []InputEncryptedFileClass

// Sort sorts slice of InputEncryptedFileClass.
func (s InputEncryptedFileClassArray) Sort(less func(a, b InputEncryptedFileClass) bool) InputEncryptedFileClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputEncryptedFileClass.
func (s InputEncryptedFileClassArray) SortStable(less func(a, b InputEncryptedFileClass) bool) InputEncryptedFileClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputEncryptedFileClass.
func (s InputEncryptedFileClassArray) Retain(keep func(x InputEncryptedFileClass) bool) InputEncryptedFileClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputEncryptedFileClassArray) First() (v InputEncryptedFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputEncryptedFileClassArray) Last() (v InputEncryptedFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputEncryptedFileClassArray) PopFirst() (v InputEncryptedFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputEncryptedFileClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputEncryptedFileClassArray) Pop() (v InputEncryptedFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputEncryptedFileUploaded returns copy with only InputEncryptedFileUploaded constructors.
func (s InputEncryptedFileClassArray) AsInputEncryptedFileUploaded() (to InputEncryptedFileUploadedArray) {
	for _, elem := range s {
		value, ok := elem.(*InputEncryptedFileUploaded)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputEncryptedFile returns copy with only InputEncryptedFile constructors.
func (s InputEncryptedFileClassArray) AsInputEncryptedFile() (to InputEncryptedFileArray) {
	for _, elem := range s {
		value, ok := elem.(*InputEncryptedFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputEncryptedFileBigUploaded returns copy with only InputEncryptedFileBigUploaded constructors.
func (s InputEncryptedFileClassArray) AsInputEncryptedFileBigUploaded() (to InputEncryptedFileBigUploadedArray) {
	for _, elem := range s {
		value, ok := elem.(*InputEncryptedFileBigUploaded)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s InputEncryptedFileClassArray) AppendOnlyNotEmpty(to []NotEmptyInputEncryptedFile) []NotEmptyInputEncryptedFile {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s InputEncryptedFileClassArray) AsNotEmpty() (to []NotEmptyInputEncryptedFile) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s InputEncryptedFileClassArray) FirstAsNotEmpty() (v NotEmptyInputEncryptedFile, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s InputEncryptedFileClassArray) LastAsNotEmpty() (v NotEmptyInputEncryptedFile, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *InputEncryptedFileClassArray) PopFirstAsNotEmpty() (v NotEmptyInputEncryptedFile, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *InputEncryptedFileClassArray) PopAsNotEmpty() (v NotEmptyInputEncryptedFile, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// InputEncryptedFileUploadedArray is adapter for slice of InputEncryptedFileUploaded.
type InputEncryptedFileUploadedArray []InputEncryptedFileUploaded

// Sort sorts slice of InputEncryptedFileUploaded.
func (s InputEncryptedFileUploadedArray) Sort(less func(a, b InputEncryptedFileUploaded) bool) InputEncryptedFileUploadedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputEncryptedFileUploaded.
func (s InputEncryptedFileUploadedArray) SortStable(less func(a, b InputEncryptedFileUploaded) bool) InputEncryptedFileUploadedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputEncryptedFileUploaded.
func (s InputEncryptedFileUploadedArray) Retain(keep func(x InputEncryptedFileUploaded) bool) InputEncryptedFileUploadedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputEncryptedFileUploadedArray) First() (v InputEncryptedFileUploaded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputEncryptedFileUploadedArray) Last() (v InputEncryptedFileUploaded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputEncryptedFileUploadedArray) PopFirst() (v InputEncryptedFileUploaded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputEncryptedFileUploaded
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputEncryptedFileUploadedArray) Pop() (v InputEncryptedFileUploaded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputEncryptedFileArray is adapter for slice of InputEncryptedFile.
type InputEncryptedFileArray []InputEncryptedFile

// Sort sorts slice of InputEncryptedFile.
func (s InputEncryptedFileArray) Sort(less func(a, b InputEncryptedFile) bool) InputEncryptedFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputEncryptedFile.
func (s InputEncryptedFileArray) SortStable(less func(a, b InputEncryptedFile) bool) InputEncryptedFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputEncryptedFile.
func (s InputEncryptedFileArray) Retain(keep func(x InputEncryptedFile) bool) InputEncryptedFileArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputEncryptedFileArray) First() (v InputEncryptedFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputEncryptedFileArray) Last() (v InputEncryptedFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputEncryptedFileArray) PopFirst() (v InputEncryptedFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputEncryptedFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputEncryptedFileArray) Pop() (v InputEncryptedFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputEncryptedFileBigUploadedArray is adapter for slice of InputEncryptedFileBigUploaded.
type InputEncryptedFileBigUploadedArray []InputEncryptedFileBigUploaded

// Sort sorts slice of InputEncryptedFileBigUploaded.
func (s InputEncryptedFileBigUploadedArray) Sort(less func(a, b InputEncryptedFileBigUploaded) bool) InputEncryptedFileBigUploadedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputEncryptedFileBigUploaded.
func (s InputEncryptedFileBigUploadedArray) SortStable(less func(a, b InputEncryptedFileBigUploaded) bool) InputEncryptedFileBigUploadedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputEncryptedFileBigUploaded.
func (s InputEncryptedFileBigUploadedArray) Retain(keep func(x InputEncryptedFileBigUploaded) bool) InputEncryptedFileBigUploadedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputEncryptedFileBigUploadedArray) First() (v InputEncryptedFileBigUploaded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputEncryptedFileBigUploadedArray) Last() (v InputEncryptedFileBigUploaded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputEncryptedFileBigUploadedArray) PopFirst() (v InputEncryptedFileBigUploaded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputEncryptedFileBigUploaded
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputEncryptedFileBigUploadedArray) Pop() (v InputEncryptedFileBigUploaded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
