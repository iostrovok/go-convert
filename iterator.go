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

func CheckMapStringType(t any) (map[string]any, bool) {

	switch t.(type) {
	case map[string]any:
		return t.(map[string]any), true
	}

	return nil, false
}
