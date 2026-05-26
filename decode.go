package csvutil

import (
	"encoding"
	"reflect"
)

var (
	textUnmarshaler = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
	csvUnmarshaler  = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
)

var intDecoders = map[int]decodeFunc{
	8:  decodeIntN(8),
	16: decodeIntN(16),
	32: decodeIntN(32),
	64: decodeIntN(64),
}

var uintDecoders = map[int]decodeFunc{
	8:  decodeUintN(8),
	16: decodeUintN(16),
	32: decodeUintN(32),
	64: decodeUintN(64),
}

var (
	decodeFloat32 = decodeFloatN(32)
	decodeFloat64 = decodeFloatN(64)
)

type decodeFunc func(s string, v reflect.Value) error

func decodeFuncValue(f func([]byte, any) error) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func decodeFuncValuePtr(f func([]byte, any) error) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func decodeString(s string, v reflect.Value) error { _ = "STUB: not implemented"; return nil }

func decodeIntN(bits int) decodeFunc { _ = "STUB: not implemented"; return *new(decodeFunc) }

func decodeUintN(bits int) decodeFunc { _ = "STUB: not implemented"; return *new(decodeFunc) }

func decodeFloatN(bits int) decodeFunc { _ = "STUB: not implemented"; return *new(decodeFunc) }

func decodeBool(s string, v reflect.Value) error { _ = "STUB: not implemented"; return nil }

func decodePtrTextUnmarshaler(s string, v reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeTextUnmarshaler(s string, v reflect.Value) error { _ = "STUB: not implemented"; return nil }

func decodePtrFieldUnmarshaler(s string, v reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeFieldUnmarshaler(s string, v reflect.Value) error { _ = "STUB: not implemented"; return nil }

func decodePtr(typ reflect.Type, funcMap map[reflect.Type]func([]byte, any) error, ifaceFuncs []ifaceDecodeFunc) (decodeFunc, error) {
	_ = "STUB: not implemented"
	return *new(decodeFunc), nil
}

func decodeInterface(funcMap map[reflect.Type]func([]byte, any) error, ifaceFuncs []ifaceDecodeFunc) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

// we may get a value receiver unmarshalers or registered funcs
// underneath the interface in which case we should call
// Unmarshal/Registered func.

func decodeBytes(s string, v reflect.Value) error { _ = "STUB: not implemented"; return nil }

func decodeFn(typ reflect.Type, funcMap map[reflect.Type]func([]byte, any) error, ifaceFuncs []ifaceDecodeFunc) (decodeFunc, error) {
	_ = "STUB: not implemented"
	return *new(decodeFunc), nil
}
