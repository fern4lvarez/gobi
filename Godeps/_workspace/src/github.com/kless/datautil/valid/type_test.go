// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package valid

import (
	"testing"
	"time"
)

var schema = NewSchema(0)

var (
	extraBool = map[string]bool{
		"oui": true,
		"non": false,
		"o":   true,
	}

	boolTests = []struct {
		in  string
		out bool
	}{
		{"Y", true},
		{"YES", true},
		{"y", true},
		{"yes", true},

		{"N", false},
		{"NO", false},
		{"n", false},
		{"no", false},

		{"oui", true},
		{"Oui", true},
		{"non", false},
		{"NON", false},
	}
)

func TestBool(t *testing.T) {
	SetBoolStrings(extraBool)
	for _, v := range []string{"oui", "Oui", "OUI"} {
		if _, found := boolStr[v]; !found {
			t.Errorf("SetBoolStrings() want generate %q", v)
		}
	}

	for _, tt := range boolTests {
		val, err := Bool(schema, tt.in)
		if err != nil {
			t.Errorf("Bool(%q) got error %q", tt.in, err)
		} else if val != tt.out {
			t.Errorf("Bool(%q) = %v, want %v", tt.in, tt.in, tt.out)
		}
	}
}

var strString = "foo"

func TestString(t *testing.T) {
	noStrStrings := []string{
		"-1", "0", "1", // int
		"1.2", "0.1", ".1", "1.", // float
	}

	schema.flagCheck = C_StrictString

	_, err := String(schema, strString)
	if err != nil {
		t.Errorf("String(%q) got error %q", strString, err)
	}

	for _, v := range noStrStrings {
		if _, err := String(schema, v); err == nil {
			t.Errorf("String(%q) want error", v)
		} else if err.(*ValidError).Err != ErrString {
			t.Errorf("String(%q) got error %q, want %q",
				v, err.(*ValidError).Err, ErrString)
		}
	}

	flagTests := []struct {
		flag Checker
		str  string
		err  error
	}{
		{C_Alpha, "qwe1r", ErrAlpha},
		{C_Alphanumeric, ".", ErrAlphaNum},
		{C_ASCII, "qwér", ErrASCII},
	}

	for _, tt := range flagTests {
		schema.flagCheck = tt.flag

		if _, err := String(schema, tt.str); err == nil {
			t.Errorf("String(%q) want error", tt.str)
		} else if err.(*ValidError).Err != tt.err {
			t.Errorf("String(%q) got error %q, want %q",
				tt.str, err.(*ValidError).Err, tt.err)
		}
	}

	schema.flagCheck = 0
}

//func TestRawBytes(t *testing.T) {}

func TestTime(t *testing.T) {
	schema.SetTimeFmt("Mon Jan 2 15:04:05 MST 2006") // time.UnixDate
	okTime := "Mon Jan 11 11:11:11 MST 2011"

	if _, err := Time(schema, okTime); err != nil {
		t.Errorf("Time(%q) got error %q", okTime, err)
	}

	if _, err := Time(schema, time.ANSIC); err == nil {
		t.Errorf("Time(%q) want error", time.ANSIC)
	} else if testing.Verbose() {
		t.Logf("%T: %v\n", err, err)
	}
}

// == int

var intStrings = []string{"-10", "0", "10"}

func TestInt(t *testing.T) {
	for _, v := range intStrings {
		if _, err := Int(schema, v); err != nil {
			t.Errorf("Int(%q) got error %q", v, err)
		}
	}
	if _, err := Int(schema, strString); err == nil {
		t.Errorf("Int(%q) want error", strString)
	}
}

func TestInt8(t *testing.T) {
	for _, v := range intStrings {
		if _, err := Int8(schema, v); err != nil {
			t.Errorf("Int8(%q) got error %q", v, err)
		}
	}
	if _, err := Int8(schema, strString); err == nil {
		t.Errorf("Int8(%q) want error", strString)
	}
}

func TestInt16(t *testing.T) {
	for _, v := range intStrings {
		if _, err := Int16(schema, v); err != nil {
			t.Errorf("Int16(%q) got error %q", v, err)
		}
	}
	if _, err := Int16(schema, strString); err == nil {
		t.Errorf("Int16(%q) want error", strString)
	}
}

func TestInt32(t *testing.T) {
	for _, v := range intStrings {
		if _, err := Int32(schema, v); err != nil {
			t.Errorf("Int32(%q) got error %q", v, err)
		}
	}
	if _, err := Int32(schema, strString); err == nil {
		t.Errorf("Int32(%q) want error", strString)
	}
}

func TestInt64(t *testing.T) {
	for _, v := range intStrings {
		if _, err := Int64(schema, v); err != nil {
			t.Errorf("Int64(%q) got error %q", v, err)
		}
	}
	if _, err := Int64(schema, strString); err == nil {
		t.Errorf("Int64(%q) want error", strString)
	}
}

// == uint

var uintStrings = []string{"10", "0", "100"}

func TestUint(t *testing.T) {
	for _, v := range uintStrings {
		if _, err := Uint(schema, v); err != nil {
			t.Errorf("Uint(%q) got error %q", v, err)
		}
	}
	if _, err := Uint(schema, strString); err == nil {
		t.Errorf("Uint(%q) want error", strString)
	}
}

func TestUint8(t *testing.T) {
	for _, v := range uintStrings {
		if _, err := Uint8(schema, v); err != nil {
			t.Errorf("Uint8(%q) got error %q", v, err)
		}
	}
	if _, err := Uint8(schema, strString); err == nil {
		t.Errorf("Uint8(%q) want error", strString)
	}
}

func TestUint16(t *testing.T) {
	for _, v := range uintStrings {
		if _, err := Uint16(schema, v); err != nil {
			t.Errorf("Uint16(%q) got error %q", v, err)
		}
	}
	if _, err := Uint16(schema, strString); err == nil {
		t.Errorf("Uint16(%q) want error", strString)
	}
}

func TestUint32(t *testing.T) {
	for _, v := range uintStrings {
		if _, err := Uint32(schema, v); err != nil {
			t.Errorf("Uint32(%q) got error %q", v, err)
		}
	}
	if _, err := Uint32(schema, strString); err == nil {
		t.Errorf("Uint32(%q) want error", strString)
	}
}

func TestUint64(t *testing.T) {
	for _, v := range uintStrings {
		if _, err := Uint64(schema, v); err != nil {
			t.Errorf("Uint64(%q) got error %q", v, err)
		}
	}
	if _, err := Uint64(schema, strString); err == nil {
		t.Errorf("Uint64(%q) want error", strString)
	}
}

// == float

var floatStrings = []string{"-10.1", "0.1", "10.2"}

func TestFloat32(t *testing.T) {
	for _, v := range floatStrings {
		if _, err := Float32(schema, v); err != nil {
			t.Errorf("Float32(%q) got error %q", v, err)
		}
	}
	if _, err := Float32(schema, strString); err == nil {
		t.Errorf("Float32(%q) want error", strString)
	}
}

func TestFloat64(t *testing.T) {
	for _, v := range floatStrings {
		if _, err := Float64(schema, v); err != nil {
			t.Errorf("Float64(%q) got error %q", v, err)
		}
	}
	if _, err := Float64(schema, strString); err == nil {
		t.Errorf("Float64(%q) want error", strString)
	}
}

// == Slices

func TestIntSlice(t *testing.T) {
	_, err := IntSlice(schema, intStrings)
	if err != nil {
		t.Errorf("IntSlice(%q) got error %q", intStrings, err)
	}

	if _, err = IntSlice(schema, floatStrings); err == nil {
		t.Errorf("IntSlice(%q) want error", floatStrings)
	}
}

func TestUintSlice(t *testing.T) {
	_, err := UintSlice(schema, uintStrings)
	if err != nil {
		t.Errorf("UintSlice(%q) got error %q", uintStrings, err)
	}

	if _, err = UintSlice(schema, intStrings); err == nil {
		t.Errorf("UintSlice(%q) want error", intStrings)
	}
}

func TestFloat64Slice(t *testing.T) {
	_, err := Float64Slice(schema, floatStrings)
	if err != nil {
		t.Errorf("Float64Slice(%q) got error %q", floatStrings, err)
	}

	if _, err = Float64Slice(schema, stringSlice); err == nil {
		t.Errorf("Float64Slice(%q) want error", stringSlice)
	}
}

func TestStringSlice(t *testing.T) {
	schema.flagCheck = C_StrictString

	_, err := StringSlice(schema, stringSlice)
	if err != nil {
		t.Errorf("StringSlice(%q) got error %q", stringSlice, err)
	}

	if _, err = StringSlice(schema, floatStrings); err == nil {
		t.Errorf("StringSlice(%q) want error", floatStrings)
	}
}

// Search

var (
	intSlice  = []int{-10, 0, 10}
	intString = "-10"
)

func TestContainsInt(t *testing.T) {
	found, err := ContainsInt(schema, intSlice, intString)
	if err != nil {
		t.Errorf("ContainsInt(%v, %q) got error %q", intSlice, intString, err)
	}
	if !found {
		t.Errorf("ContainsInt(%v, %q) want true", intSlice, intString)
	}

	if _, err := ContainsInt(schema, intSlice, strString); err == nil {
		t.Errorf("ContainsInt(%v, %q) want error", intSlice, strString)
	}
}

var (
	uintSlice  = []uint{10, 0, 14}
	uintString = "14"
)

func TestContainsUint(t *testing.T) {
	found, err := ContainsUint(schema, uintSlice, uintString)
	if err != nil {
		t.Errorf("ContainsUint(%v, %q) got error %q", uintSlice, uintString, err)
	}
	if !found {
		t.Errorf("ContainsUint(%v, %q) want true", uintSlice, uintString)
	}

	if _, err := ContainsUint(schema, uintSlice, strString); err == nil {
		t.Errorf("ContainsUint(%v, %q) want error", uintSlice, strString)
	}
}

var (
	floatSlice  = []float64{-10.1, 0.1, 10.2}
	floatString = "10.2"
)

func TestContainsFloat64(t *testing.T) {
	found, err := ContainsFloat64(schema, floatSlice, floatString)
	if err != nil {
		t.Errorf("ContainsFloat64(%v, %q) got error %q", floatSlice, floatString, err)
	}
	if !found {
		t.Errorf("ContainsFloat64(%v, %q) want true", floatSlice, floatString)
	}

	if _, err := ContainsFloat64(schema, floatSlice, strString); err == nil {
		t.Errorf("ContainsFloat64(%v, %q) want error", floatSlice, strString)
	}
}

var stringSlice = []string{"foo", "a", "bar"}

func TestContainsString(t *testing.T) {
	found, err := ContainsString(schema, stringSlice, strString)
	if err != nil {
		t.Errorf("ContainsString(%v, %q) got error %q", stringSlice, strString, err)
	}
	if !found {
		t.Errorf("ContainsString(%v, %q) want true", stringSlice, strString)
	}

	if _, err := ContainsString(schema, stringSlice, floatString); err == nil {
		t.Errorf("ContainsString(%v, %q) want error", stringSlice, floatString)
	}
}
