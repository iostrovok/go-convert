package convert

import (
	"fmt"
	"reflect"
)

// MapIt is a forward-only iterator over any map value.
type MapIt struct {
	item     reflect.Value
	keys     []reflect.Value
	i, count int
}

// MapIterator wraps i in a MapIt iterator. i must be a map; otherwise an error
// is returned. If checkLen is provided and true, an error is returned for an
// empty map as well.
func MapIterator(i any, checkLen ...bool) (*MapIt, error) {
	item := reflect.ValueOf(i)
	if item.Kind() != reflect.Map {
		return nil, fmt.Errorf("input value is not a map")
	}

	if len(checkLen) > 0 && checkLen[0] && item.Len() == 0 {
		return nil, fmt.Errorf("map is empty")
	}

	return &MapIt{
		item:  item,
		keys:  item.MapKeys(),
		count: item.Len(),
	}, nil
}

// Next advances the iterator and returns the next key-value pair.
// Returns nil, nil when all entries have been consumed.
func (it *MapIt) Next() (any, any) {

	for it.i < it.count {
		key := it.keys[it.i]
		value := it.item.MapIndex(key).Interface()
		it.i++
		return key.Interface(), value
	}

	return nil, nil
}

// HasNext reports whether there are more entries to iterate over.
func (it *MapIt) HasNext() bool {
	return it.i < it.count
}

// Len returns the total number of entries in the underlying map.
func (it *MapIt) Len() int {
	return it.count
}
