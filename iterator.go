package convert

import (
	"fmt"
	"reflect"
)

// It is a forward-only iterator over any slice value.
type It struct {
	items    reflect.Value
	i, count int
}

// Iterator wraps i in an It iterator. i must be a slice; otherwise an error is
// returned. If checkLen is provided and true, an error is returned for an empty
// slice as well.
func Iterator(i any, checkLen ...bool) (*It, error) {
	items := reflect.ValueOf(i)
	if items.Kind() != reflect.Slice {
		return nil, fmt.Errorf("data is not a list/slice")
	}

	if len(checkLen) > 0 && checkLen[0] && items.Len() == 0 {
		return nil, fmt.Errorf("list/slice is empty")
	}

	return &It{
		items: items,
		count: items.Len(),
	}, nil
}

// NextNotNil advances the iterator and returns the next non-nil element.
// Returns nil when no more non-nil elements remain.
func (it *It) NextNotNil() any {

	for it.i < it.count {
		out := it.items.Index(it.i).Interface()
		it.i++
		if out != nil {
			return out
		}
	}

	return nil
}

// NextNotNilMapString advances the iterator and returns the next element that
// is a non-nil map[string]any. Elements of other types are skipped.
// Returns nil, false when the iterator is exhausted.
func (it *It) NextNotNilMapString() (map[string]any, bool) {
	for {
		i := it.NextNotNil()
		if i == nil {
			return nil, false
		}

		out, find := CheckMapStringType(i)
		if find {
			return out, true
		}
	}

	return nil, false
}

// NextNotEmptyString advances the iterator and returns the next element that
// is a non-empty string. Non-string and empty-string elements are skipped.
// Returns "" when the iterator is exhausted.
func (it *It) NextNotEmptyString() string {
	out := ""
	for {
		s := it.NextNotNil()
		if s == nil {
			return ""
		}

		switch s.(type) {
		case string:
			out = s.(string)
		}

		// don't return empty string
		if out != "" {
			return out
		}
	}

	return ""
}

// Len returns the total number of elements in the underlying slice.
func (it *It) Len() int {
	return it.count
}

// CheckMapStringType asserts that t is of type map[string]any.
// Returns the map and true on success, or nil and false otherwise.
func CheckMapStringType(t any) (map[string]any, bool) {

	switch t.(type) {
	case map[string]any:
		return t.(map[string]any), true
	}

	return nil, false
}
