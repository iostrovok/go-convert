package convert

/*
	Iterator over slice of interfaces
*/

import (
	"fmt"
	"reflect"
)

type It struct {
	items    reflect.Value
	i, count int
}

func Iterator(i interface{}, checkLen ...bool) (*It, error) {
	items := reflect.ValueOf(i)
	if items.Kind() != reflect.Slice {
		return nil, fmt.Errorf("Data is not a list/slice")
	}

	if len(checkLen) > 0 && checkLen[0] && items.Len() == 0 {
		return nil, fmt.Errorf("list/slice is empty")
	}

	return &It{
		items: items,
		count: items.Len(),
	}, nil
}

func (it *It) NextNotNil() interface{} {

	for it.i < it.count {
		out := it.items.Index(it.i).Interface()
		it.i++
		if out != nil {
			return out
		}
	}

	return nil
}

func (it *It) NextNotNilMapString() map[string]interface{} {

	for {
		i := it.NextNotNil()
		if i == nil {
			return nil
		}
		out, find := CheckMapStringType(i)
		if find {
			return out
		}
	}

	return nil
}

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

func (it *It) Len() int {
	return it.count
}

func CheckMapStringType(t interface{}) (map[string]interface{}, bool) {

	switch t.(type) {
	case map[string]interface{}:
		return t.(map[string]interface{}), true
	}

	return nil, false
}
