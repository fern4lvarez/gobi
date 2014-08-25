// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package valid

import (
	"testing"

	"github.com/kless/datautil"
)

func TestPreCheck(t *testing.T) {
	flagTests := []struct {
		flag Checker
		in   string
		out  string
	}{
		{M_TrimSpace, "  foo ", "foo"},
		{M_ToLowercase, " Foo ", " foo "},
		{M_ToUppercase, "bar", "BAR"},
	}

	for _, tt := range flagTests {
		schema.flagCheck = tt.flag

		out, err := schema.PreCheck(tt.in, datautil.String)
		if err != nil {
			t.Errorf("PreCheck(%q) got error: %q", tt.in, err)
		} else if out != tt.out {
			t.Errorf("PreCheck(%q) got %s, want %s", tt.in, out, tt.out)
		}
	}

	emptyTests := []struct {
		flag Checker
		err  error
		def  string
	}{
		{0, ErrEmpty, ""},
		{C_Required, ErrRequired, ""},
		{0, nil, "valueBydefault"},
	}

	in := ""
	for i, tt := range emptyTests {
		schema.flagCheck = tt.flag
		if tt.def != "" {
			schema.Bydefault = tt.def
		}

		out, err := schema.PreCheck(in, datautil.String)
		if err == nil && tt.def == "" {
			t.Errorf("[%d] PreCheck(%q) want error", i, in)
		} else if err != tt.err {
			t.Errorf("[%d] PreCheck(%q) got error %q, want %q", i, in, err, tt.err)
		}
		if tt.def != "" && out != tt.def {
			t.Errorf("[%d] PreCheck(%q) got %q, want %q", i, in, out, tt.def)
		}
	}
}

func TestPostCheck(t *testing.T) {
	schema.flagCheck = c_MinLen | c_MaxLen

	// Data used for argument in PostCheck
	value_str := "four"
	value_int := 33
	value_uint := uint(22)
	value_float := 11.22

	flagTests := []struct {
		dataType datautil.Type
		checker  Checker
		value    interface{}
		hasError bool
	}{
		// string
		{datautil.String, c_MinLen, len(value_str) - 1, false},
		{datautil.String, c_MinLen, len(value_str), false},
		{datautil.String, c_MinLen, len(value_str) + 1, true},

		{datautil.String, c_MaxLen, len(value_str) - 1, true},
		{datautil.String, c_MaxLen, len(value_str), false},
		{datautil.String, c_MaxLen, len(value_str) + 1, false},

		// int
		{datautil.Int8, c_MinLen, value_int - 1, false},
		{datautil.Int8, c_MinLen, value_int, false},
		{datautil.Int8, c_MinLen, value_int + 1, true},

		{datautil.Int8, c_MaxLen, value_int - 1, true},
		{datautil.Int8, c_MaxLen, value_int, false},
		{datautil.Int8, c_MaxLen, value_int + 1, false},

		// uint
		{datautil.Uint8, c_MinLen, value_uint - uint(1), false},
		{datautil.Uint8, c_MinLen, value_uint, false},
		{datautil.Uint8, c_MinLen, value_uint + uint(1), true},

		{datautil.Uint8, c_MaxLen, value_uint - uint(1), true},
		{datautil.Uint8, c_MaxLen, value_uint, false},
		{datautil.Uint8, c_MaxLen, value_uint + uint(1), false},

		// float
		{datautil.Float64, c_MinLen, value_float - 1.0, false},
		{datautil.Float64, c_MinLen, value_float, false},
		{datautil.Float64, c_MinLen, value_float + 1.0, true},

		{datautil.Float64, c_MaxLen, value_float - 1.0, true},
		{datautil.Float64, c_MaxLen, value_float, false},
		{datautil.Float64, c_MaxLen, value_float + 1.0, false},
	}

	var errExpected error
	var value interface{}

	for i, tt := range flagTests {
		schema.dataType = tt.dataType
		if !tt.hasError {
			errExpected = nil
		}

		switch tt.checker {
		case c_MinLen:
			schema.max, schema.maxf = 0, 0
			if tt.hasError {
				errExpected = minLenError{tt.value}
			}
			switch tt.value.(type) {
			case int:
				schema.min = tt.value.(int)
			case uint:
				schema.min = int(tt.value.(uint))
			case float64:
				schema.minf = tt.value.(float64)
			}
		case c_MaxLen:
			schema.min, schema.minf = 0, 0
			if tt.hasError {
				errExpected = maxLenError{tt.value}
			}
			switch tt.value.(type) {
			case int:
				schema.max = tt.value.(int)
			case uint:
				schema.max = int(tt.value.(uint))
			case float64:
				schema.maxf = tt.value.(float64)
			}
		}

		switch tt.dataType {
		case datautil.String:
			value = value_str
		case datautil.Int8:
			value = value_int
		case datautil.Uint8:
			value = value_uint
		case datautil.Float64:
			value = value_float
		}

		err := schema.PostCheck(value)

		if tt.hasError {
			if err.Error() != errExpected.Error() {
				t.Errorf("[%d] PostCheck(%v) got error \"%v\", want %q",
					i, value, err, errExpected)
			}
		} else if err != nil {
			t.Errorf("[%d] PostCheck(%v) got error %q", i, value, err)
		}
	}
}
