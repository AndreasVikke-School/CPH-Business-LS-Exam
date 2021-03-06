// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CheckInsColumns holds the columns for the "check_ins" table.
	CheckInsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "attendance_code", Type: field.TypeInt64},
		{Name: "student_id", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"success", "outOfTime", "notFound", "outOfRange", "error"}},
		{Name: "checkin_time", Type: field.TypeInt64},
	}
	// CheckInsTable holds the schema information for the "check_ins" table.
	CheckInsTable = &schema.Table{
		Name:       "check_ins",
		Columns:    CheckInsColumns,
		PrimaryKey: []*schema.Column{CheckInsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CheckInsTable,
	}
)

func init() {
}
