package Tag

import (
	"reflect"
)

type Tag struct {
	Tags  map[string]reflect.StructTag
	Field string
}

func (c *Tag) Parse(st interface{}) {
	t := reflect.TypeOf(st)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("TAG: param is not struct")
	}

	c.Tags = make(map[string]reflect.StructTag, t.NumField())

	for i := 0; i != t.NumField(); i++ {
		c.Tags[t.Field(i).Name] = t.Field(i).Tag
	}
}

func (c *Tag) SetField(field string) {
	c.Field = field
}

func (c *Tag) Get(name string, field ...string) string {
	var f reflect.StructTag
	if len(c.Field) > 0 {
		f = c.Tags[c.Field]
	} else if len(field) > 0 {
		f = c.Tags[field[0]]
	} else {
		panic("TAG: can not get tag content , because not set Field")
	}

	return f.Get(name)
}
