package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// StatusMixin 通用状态字段，适用于多状态场景
type StatusMixin struct {
	mixin.Schema
}

// Fields 定义通用状态字段
func (StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("status").
			Default(1).
			SchemaType(map[string]string{
				dialect.Postgres: "smallint",
			}).
			Comment("状态：1 正常，2 停用"),
	}
}
