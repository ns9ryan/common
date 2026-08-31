package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SortMixin 通用排序字段
type SortMixin struct {
	mixin.Schema
}

// Fields 定义通用排序字段
func (SortMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("sort").
			Default(1).
			Comment("排序编号"),
	}
}
