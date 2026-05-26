package csvutil

import (
	"encoding"
	"reflect"
)

var (
	textMarshaler = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	csvMarshaler  = reflect.TypeOf((*Marshaler)(nil)).Elem()
)

var (
	encodeFloat32 = encodeFloatN(32)
	encodeFloat64 = encodeFloatN(64)
)

type encodeFunc func(buf []byte, v reflect.Value, omitempty bool) ([]byte, error)

func nopEncode(buf []byte, _ reflect.Value, _ bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeFuncValue(fn marshalFunc) encodeFunc { _ = "STUB: not implemented"; return *new(encodeFunc) }

func encodeFuncValuePtr(fn marshalFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func encodeString(buf []byte, v reflect.Value, omitempty bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeInt(buf []byte, v reflect.Value, omitempty bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeUint(buf []byte, v reflect.Value, omitempty bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeFloatN(bits int) encodeFunc { _ = "STUB: not implemented"; return *new(encodeFunc) }

func encodeBool(buf []byte, v reflect.Value, omitempty bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeInterface(funcMap map[reflect.Type]marshalFunc, funcs []marshalFunc) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func encodePtrMarshaler(buf []byte, v reflect.Value, omitempty bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeTextMarshaler(buf []byte, v reflect.Value, _ bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodePtrTextMarshaler(buf []byte, v reflect.Value, omitempty bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeMarshaler(buf []byte, v reflect.Value, _ bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodePtr(typ reflect.Type, canAddr bool, funcMap map[reflect.Type]marshalFunc, funcs []marshalFunc) (encodeFunc, error) {
	_ = "STUB: not implemented"
	return *new(encodeFunc), nil
}

func encodeBytes(buf []byte, v reflect.Value, _ bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeFn(typ reflect.Type, canAddr bool, funcMap map[reflect.Type]marshalFunc, funcs []marshalFunc) (encodeFunc, error) {
	_ = "STUB: not implemented"
	return *new(encodeFunc), nil
}
