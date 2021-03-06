// Code generated by nero, DO NOT EDIT.
package repository

import (
	"time"

	"github.com/sf9v/nero/comparison"
)

// PredFunc is a predicate function
type PredFunc func(*comparison.Predicates)

// IDEq is a "equal" operator on "id" column
func IDEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Arg: id,
		})
	}
}

// IDNotEq is a "not equal" operator on "id" column
func IDNotEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Arg: id,
		})
	}
}

// IDGt is a "greater than" operator on "id" column
func IDGt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Arg: id,
		})
	}
}

// IDGtOrEq is a "greater than or equal" operator on "id" column
func IDGtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Arg: id,
		})
	}
}

// IDLt is a "less than" operator on "id" column
func IDLt(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Arg: id,
		})
	}
}

// IDLtOrEq is a "less than or equal" operator on "id" column
func IDLtOrEq(id int64) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Arg: id,
		})
	}
}

// IDIn is a "in" operator on "id" column
func IDIn(ids ...int64) PredFunc {
	args := []interface{}{}
	for _, v := range ids {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// IDNotIn is a "not in" operator on "id" column
func IDNotIn(ids ...int64) PredFunc {
	args := []interface{}{}
	for _, v := range ids {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}

// NameEq is a "equal" operator on "name" column
func NameEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Eq,
			Arg: name,
		})
	}
}

// NameNotEq is a "not equal" operator on "name" column
func NameNotEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.NotEq,
			Arg: name,
		})
	}
}

// NameGt is a "greater than" operator on "name" column
func NameGt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Gt,
			Arg: name,
		})
	}
}

// NameGtOrEq is a "greater than or equal" operator on "name" column
func NameGtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.GtOrEq,
			Arg: name,
		})
	}
}

// NameLt is a "less than" operator on "name" column
func NameLt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Lt,
			Arg: name,
		})
	}
}

// NameLtOrEq is a "less than or equal" operator on "name" column
func NameLtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.LtOrEq,
			Arg: name,
		})
	}
}

// NameIn is a "in" operator on "name" column
func NameIn(names ...string) PredFunc {
	args := []interface{}{}
	for _, v := range names {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// NameNotIn is a "not in" operator on "name" column
func NameNotIn(names ...string) PredFunc {
	args := []interface{}{}
	for _, v := range names {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}

// CreatedAtEq is a "equal" operator on "created_at" column
func CreatedAtEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Eq,
			Arg: createdAt,
		})
	}
}

// CreatedAtNotEq is a "not equal" operator on "created_at" column
func CreatedAtNotEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.NotEq,
			Arg: createdAt,
		})
	}
}

// CreatedAtGt is a "greater than" operator on "created_at" column
func CreatedAtGt(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Gt,
			Arg: createdAt,
		})
	}
}

// CreatedAtGtOrEq is a "greater than or equal" operator on "created_at" column
func CreatedAtGtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.GtOrEq,
			Arg: createdAt,
		})
	}
}

// CreatedAtLt is a "less than" operator on "created_at" column
func CreatedAtLt(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.Lt,
			Arg: createdAt,
		})
	}
}

// CreatedAtLtOrEq is a "less than or equal" operator on "created_at" column
func CreatedAtLtOrEq(createdAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.LtOrEq,
			Arg: createdAt,
		})
	}
}

// CreatedAtIn is a "in" operator on "created_at" column
func CreatedAtIn(createdAts ...*time.Time) PredFunc {
	args := []interface{}{}
	for _, v := range createdAts {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// CreatedAtNotIn is a "not in" operator on "created_at" column
func CreatedAtNotIn(createdAts ...*time.Time) PredFunc {
	args := []interface{}{}
	for _, v := range createdAts {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "created_at",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}

// UpdatedAtEq is a "equal" operator on "updated_at" column
func UpdatedAtEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Eq,
			Arg: updatedAt,
		})
	}
}

// UpdatedAtNotEq is a "not equal" operator on "updated_at" column
func UpdatedAtNotEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.NotEq,
			Arg: updatedAt,
		})
	}
}

// UpdatedAtGt is a "greater than" operator on "updated_at" column
func UpdatedAtGt(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Gt,
			Arg: updatedAt,
		})
	}
}

// UpdatedAtGtOrEq is a "greater than or equal" operator on "updated_at" column
func UpdatedAtGtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.GtOrEq,
			Arg: updatedAt,
		})
	}
}

// UpdatedAtLt is a "less than" operator on "updated_at" column
func UpdatedAtLt(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.Lt,
			Arg: updatedAt,
		})
	}
}

// UpdatedAtLtOrEq is a "less than or equal" operator on "updated_at" column
func UpdatedAtLtOrEq(updatedAt *time.Time) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.LtOrEq,
			Arg: updatedAt,
		})
	}
}

// UpdatedAtIn is a "in" operator on "updated_at" column
func UpdatedAtIn(updatedAts ...*time.Time) PredFunc {
	args := []interface{}{}
	for _, v := range updatedAts {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// UpdatedAtNotIn is a "not in" operator on "updated_at" column
func UpdatedAtNotIn(updatedAts ...*time.Time) PredFunc {
	args := []interface{}{}
	for _, v := range updatedAts {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "updated_at",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}
