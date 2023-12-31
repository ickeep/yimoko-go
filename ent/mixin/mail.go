// Package mixin mail
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/ickeep/yimoko-go/ent/ann"
	"github.com/ickeep/yimoko-go/utils"
)

// Mail mixin 邮箱
type Mail struct {
	mixin.Schema
}

// Fields _
func (Mail) Fields() []ent.Field {
	return []ent.Field{
		field.String("mail").
			Comment("邮箱").
			MaxLen(63).
			Optional().
			Default("").
			Annotations(
				ann.Field{
					PbIndex:     209,
					MaskEncrypt: utils.MaskTypeEmail,
					Query: ann.FieldQuery{
						NotEq: true,
						In:    true,
						NotIn: true,
					},
				},
			),

		field.String("mailCipher").
			MaxLen(255).
			Sensitive().
			Annotations(
				ann.Field{
					OnlyData: true,
				},
			),
	}
}

// Index _
func (Mail) Index() []ent.Index {
	return []ent.Index{
		index.Fields("mailCipher"),
	}
}
