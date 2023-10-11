package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type URLMixin struct {
	mixin.Schema
}

func (URLMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").
			GoType(new(url.URL)).
			ValueScanner(field.BinaryValueScanner[*url.URL]{}),
	}
}
