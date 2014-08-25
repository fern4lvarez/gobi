// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package datautil implements a subset of Go types for the data management.
package datautil

// A Type represents a Go type.
type Type int

const (
	Custom Type = -1

	_ = iota
	Bool

	// Integers
	Int
	Int8
	Int16
	Int32
	Int64
	// Unsigned integers
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	// Float-point numbers
	Float32
	Float64

	String
	RawBytes

	Slice
	//Map
	//Struct
	Time
	//Nil
)

func (k Type) String() string {
	switch k {
	case Bool:
		return "bool"

	case Int:
		return "int"
	case Int8:
		return "int8"
	case Int16:
		return "int16"
	case Int32:
		return "int32"
	case Int64:
		return "int64"

	case Uint:
		return "uint"
	case Uint8:
		return "uint8"
	case Uint16:
		return "uint16"
	case Uint32:
		return "uint32"
	case Uint64:
		return "uint64"

	case Float32:
		return "float32"
	case Float64:
		return "float64"

	case String:
		return "string"
	case RawBytes:
		return "[]byte"

	case Slice:
		return "[]"
	//case Map:
	//return "map"
	//case Struct:
	//return "struct"
	case Time:
		return "time"
		//case Nil:
		//return "NULL"
	}
	panic("unimplemented")
}
