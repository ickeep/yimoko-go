// Package mixin appid
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/ickeep/yimoko-go/ent/ann"
)

// Appid mixin 通常用于 sass 场景中标识不同的 app
type Appid struct {
	mixin.Schema
}

// Fields _
func (Appid) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("appid").
			Min(1000).
			Optional().
			Comment("appid").
			Immutable().
			Annotations(ann.Field{
				PbIndex:   205,
				SassField: true,
				OnlyMeta:  true,
			}),
	}
}

// Index _
func (Appid) Index() []ent.Index {
	return []ent.Index{
		index.Fields("appid"),
	}
}
