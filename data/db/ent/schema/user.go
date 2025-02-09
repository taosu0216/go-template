package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"go-template/data/cons"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("主键"),             // 设置为主键
		field.String("name").Comment("用户名").Unique(), //.Optional(), // 设置为可空字段
		field.String("passwd").Comment("密码"),
		field.String("email").Unique(),                                     //.Unique(), // 设置为唯一字段
		field.String("phone").Unique().Optional(),                          //.Unique(), // 设置为唯一字段
		field.String("role").Comment("root / user").Default(cons.RoleUser), //.Unique(), // 设置为唯一字段
		field.Bool("is_vip").Comment("是否充值").Default(false),                //.Unique(), // 设置为唯一字段
		field.Float("balance").Default(0.0),
		field.String("create_time"),
	}
}
