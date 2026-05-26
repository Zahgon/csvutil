package csvutil

import (
	"reflect"
)

type tag struct {
	name      string
	prefix    string
	empty     bool
	omitEmpty bool
	ignore    bool
	inline    bool
}

func parseTag(tagname string, field reflect.StructField) (t tag) {
	_ = "STUB: not implemented"
	return *new(tag)
}
