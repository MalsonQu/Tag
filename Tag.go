package Tag

import (
	"reflect"
	"strings"
)

type Tag struct {
	Tags  map[string]reflect.StructTag
	Field string
}

func (c *Tag) Parse(st interface{}) *Tag {
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

	return c
}

func (c *Tag) SetField(field string) *Tag {
	c.Field = strings.ToUpper(field)

	return c
}

func (c *Tag) Get(name string, field ...string) string {
	var f reflect.StructTag
	if len(c.Field) > 0 {
		f = c.Tags[c.Field]
	} else if len(field) > 0 {
		f = c.Tags[strings.ToUpper(field[0])]
	} else {
		panic("TAG: can not get tag content , because not set Field")
	}

	return f.Get(name)
}
