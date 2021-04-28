package json

import (
	"reflect"
)

// This interface allows custom types to be loaded via strings in JSON.  When
// this interface is implemented, the FromString must decode into the same type
// specified in the struct, likewise the FromString must return a string used
// by Marshal to create string literals.
type CustomType interface {
	FromString(string) (interface{}, error)
	ToString() string
}

func loadCustomType(item []byte, v reflect.Value) (found bool, err error) {
	defer func() {
		recover()
	}()
	customType := reflect.TypeOf((*CustomType)(nil)).Elem()
	ret := v.Convert(customType)

	// If we get to this point, load up the value
	found = true
	var setVal interface{}
	setVal, err = ret.Interface().(CustomType).FromString(string(item))
	v.Set(reflect.ValueOf(setVal))
	return
}

//type encoderFunc func(e *encodeState, v reflect.Value, opts encOpts)

func customEncoder(t reflect.Type) (enc encoderFunc, found bool) {
	defer func() {
		recover()
	}()
	customType := reflect.TypeOf((*CustomType)(nil)).Elem()

	return func(e *encodeState, v reflect.Value, opts encOpts) {
		defer func() {
			recover()
		}()
		ret := v.Convert(customType)
		setVal := ret.Interface().(CustomType).ToString()
		if opts.quoted {
			e2 := newEncodeState()
			// Since we encode the string twice, we only need to escape HTML
			// the first time.
			e2.string(setVal, opts.escapeHTML)
			e.stringBytes(e2.Bytes(), false)
			encodeStatePool.Put(e2)
		} else {
			e.string(setVal, opts.escapeHTML)
		}
	}, t.Implements(customType)
}

func testCustomType(item []byte, v reflect.Value) (found bool) {
	defer func() {
		recover()
	}()
	custtype := reflect.TypeOf((*CustomType)(nil)).Elem()
	v.Convert(custtype)

	// If we get to this point, load up the value
	found = true
	return
}
