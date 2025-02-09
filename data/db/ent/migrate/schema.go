// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "passwd", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "role", Type: field.TypeString, Default: "user"},
		{Name: "is_vip", Type: field.TypeBool, Default: false},
		{Name: "balance", Type: field.TypeFloat64, Default: 0},
		{Name: "create_time", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
