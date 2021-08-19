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

// HelpAppUpdateClassArray is adapter for slice of HelpAppUpdateClass.
type HelpAppUpdateClassArray []HelpAppUpdateClass

// Sort sorts slice of HelpAppUpdateClass.
func (s HelpAppUpdateClassArray) Sort(less func(a, b HelpAppUpdateClass) bool) HelpAppUpdateClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpAppUpdateClass.
func (s HelpAppUpdateClassArray) SortStable(less func(a, b HelpAppUpdateClass) bool) HelpAppUpdateClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpAppUpdateClass.
func (s HelpAppUpdateClassArray) Retain(keep func(x HelpAppUpdateClass) bool) HelpAppUpdateClassArray {
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
func (s HelpAppUpdateClassArray) First() (v HelpAppUpdateClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpAppUpdateClassArray) Last() (v HelpAppUpdateClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpAppUpdateClassArray) PopFirst() (v HelpAppUpdateClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpAppUpdateClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpAppUpdateClassArray) Pop() (v HelpAppUpdateClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpAppUpdate returns copy with only HelpAppUpdate constructors.
func (s HelpAppUpdateClassArray) AsHelpAppUpdate() (to HelpAppUpdateArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpAppUpdate)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// HelpAppUpdateArray is adapter for slice of HelpAppUpdate.
type HelpAppUpdateArray []HelpAppUpdate

// Sort sorts slice of HelpAppUpdate.
func (s HelpAppUpdateArray) Sort(less func(a, b HelpAppUpdate) bool) HelpAppUpdateArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpAppUpdate.
func (s HelpAppUpdateArray) SortStable(less func(a, b HelpAppUpdate) bool) HelpAppUpdateArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpAppUpdate.
func (s HelpAppUpdateArray) Retain(keep func(x HelpAppUpdate) bool) HelpAppUpdateArray {
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
func (s HelpAppUpdateArray) First() (v HelpAppUpdate, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpAppUpdateArray) Last() (v HelpAppUpdate, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpAppUpdateArray) PopFirst() (v HelpAppUpdate, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpAppUpdate
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpAppUpdateArray) Pop() (v HelpAppUpdate, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of HelpAppUpdate by ID.
func (s HelpAppUpdateArray) SortByID() HelpAppUpdateArray {
	return s.Sort(func(a, b HelpAppUpdate) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of HelpAppUpdate by ID.
func (s HelpAppUpdateArray) SortStableByID() HelpAppUpdateArray {
	return s.SortStable(func(a, b HelpAppUpdate) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s HelpAppUpdateArray) FillMap(to map[int]HelpAppUpdate) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s HelpAppUpdateArray) ToMap() map[int]HelpAppUpdate {
	r := make(map[int]HelpAppUpdate, len(s))
	s.FillMap(r)
	return r
}
