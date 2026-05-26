package csvutil

import (
	"reflect"
	"sync"
)

var fieldCache sync.Map // map[typeKey][]field

func cachedFields(k typeKey) fields { _ = "STUB: not implemented"; return *new(fields) }

type field struct {
	name     string
	baseType reflect.Type
	typ      reflect.Type
	tag      tag
	index    []int
}

type fields []field

func (fs fields) Len() int { _ = "STUB: not implemented"; return 0 }

func (fs fields) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (fs fields) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type typeKey struct {
	tag string
	typ reflect.Type
}

type fieldMap map[string]fields

func (m fieldMap) insert(f field) { _ = "STUB: not implemented"; return }

// insert only fields with the shortest path.

// fields that are tagged have priority.

func (m fieldMap) fields() fields { _ = "STUB: not implemented"; return *new(fields) }

func buildFields(k typeKey) fields { _ = "STUB: not implemented"; return *new(fields) }

// unexported field

// ignore embedded unexported non-struct fields.

// look for duplicate nodes on the same level. Nodes won't be
// revisited, so write all fields for the current type now.

// other nodes can have different path.

func makeIndex(index []int, v int) []int { _ = "STUB: not implemented"; return nil }
