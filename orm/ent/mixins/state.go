package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// StateMixin 通用状态字段
type StateMixin struct {
	mixin.Schema
}

// Fields 定义通用状态字段
func (StateMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("state").
			Default(true).
			Optional().
			Comment("状态：true 正常，false 禁用"),
	}
}
