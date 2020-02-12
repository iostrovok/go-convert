package convert

/*
	Iterator over map of interfaces
*/

import (
	"fmt"
	"reflect"
)

type MapIt struct {
	item     reflect.Value
	keys     []reflect.Value
	i, count int
}

func MapIterator(i interface{}, checkLen ...bool) (*MapIt, error) {
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

func (it *MapIt) Next() (interface{}, interface{}) {

	for it.i < it.count {
		key := it.keys[it.i]
		value := it.item.MapIndex(key).Interface()
		it.i++
		return key.Interface(), value
	}

	return nil, nil
}

func (it *MapIt) HasNext() bool {
	return it.i < it.count
}

func (it *MapIt) Len() int {
	return it.count
}
