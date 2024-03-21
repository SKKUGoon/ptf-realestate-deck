// Code generated by ent, DO NOT EDIT.

package pnu

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the pnu type in the database.
	Label = "pnu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldJibun holds the string denoting the jibun field in the database.
	FieldJibun = "jibun"
	// FieldBchk holds the string denoting the bchk field in the database.
	FieldBchk = "bchk"
	// FieldPnu holds the string denoting the pnu field in the database.
	FieldPnu = "pnu"
	// FieldColAdmSe holds the string denoting the col_adm_se field in the database.
	FieldColAdmSe = "col_adm_se"
	// FieldGeometry holds the string denoting the geometry field in the database.
	FieldGeometry = "geometry"
	// Table holds the table name of the pnu in the database.
	Table = "pnu"
)

// Columns holds all SQL columns for pnu fields.
var Columns = []string{
	FieldID,
	FieldJibun,
	FieldBchk,
	FieldPnu,
	FieldColAdmSe,
	FieldGeometry,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Pnu queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByJibun orders the results by the jibun field.
func ByJibun(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJibun, opts...).ToFunc()
}

// ByBchk orders the results by the bchk field.
func ByBchk(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBchk, opts...).ToFunc()
}

// ByPnu orders the results by the pnu field.
func ByPnu(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPnu, opts...).ToFunc()
}

// ByColAdmSe orders the results by the col_adm_se field.
func ByColAdmSe(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColAdmSe, opts...).ToFunc()
}

// ByGeometry orders the results by the geometry field.
func ByGeometry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeometry, opts...).ToFunc()
}
