package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// IDMixin 通用主键及时间字段
type IDMixin struct {
	mixin.Schema
}

// Fields 定义通用主键、创建时间和更新时间字段
func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Comment("主键 ID"),

		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("创建时间"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}
