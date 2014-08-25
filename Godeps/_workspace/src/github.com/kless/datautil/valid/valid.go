// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package valid validates strings of text.
//
// The flag Checker which is set in function NewSchema allows easily
// to check and/or to modify the input for some types.
package valid

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kless/datautil"
)

// Checker represents the check-ups and modifications for a string.
type Checker uint16

// These flags define the check-ups and modification for the input string.
const (
	C_Required Checker = 1 << iota // Required; can not be empty

	M_TrimSpace   // Remove all leading and trailing space
	M_ToLowercase // Convert to lower case
	M_ToUppercase // Convert to upper case

	c_MinLen // Check the minimum length
	c_MaxLen // Check the maximum length

	// For type String
	C_StrictString // Strict checking; no numbers in strings
	C_Alpha        // Check if the string contains only letters (a-zA-Z)
	C_Alphanumeric // Check if the string contains only letters and numbers
	C_ASCII        // Check if the string contains only ASCII characters

	C_HTTP_FTP // Check protocol scheme in type URL
	C_DNS      // Check DNS
	C_TLD      // Check Top Level Domain
)

// A Schema identifies a validation schema.
type Schema struct {
	Bydefault string // Value used by default in empty strings

	timeFmt    string // Time format
	patternFmt string // Message about the format to use into a pattern

	// Limits of length
	min, max   int
	minf, maxf float64

	dataType  datautil.Type
	flagCheck Checker

	IsSlice bool
}

// NewSchema returns a new schema with the given string Checkers and sets
// "time.RFC3339" for the time format.
func NewSchema(flag Checker) *Schema {
	return &Schema{
		flagCheck: flag,
		timeFmt:   time.RFC3339,
	}
}

// DataType returns the data type.
func (s *Schema) DataType() datautil.Type { return s.dataType }

// == Set

// SetChecker sets the Checker flags.
func (s *Schema) SetChecker(flag Checker) *Schema {
	s.flagCheck = flag
	return s
}

// SetMin sets the checking for the minimum length of a string,
// or the minimum value of a numeric type.
// The valid types for the aregument are: int, float64.
func (s *Schema) SetMin(n interface{}) *Schema {
	switch t := n.(type) {
	case int:
		s.min = t
	case float64:
		s.minf = t
	default:
		panic("the type must be int or float64")
	}

	s.flagCheck |= c_MinLen
	return s
}

// SetMax sets the checking for the maximum length of a string,
// or the maximum value of a numeric type.
// The valid types for the aregument are: int, float64.
func (s *Schema) SetMax(n interface{}) *Schema {
	switch t := n.(type) {
	case int:
		s.max = t
	case float64:
		s.maxf = t
	default:
		panic("the type must be int or float64")
	}

	s.flagCheck |= c_MaxLen
	return s
}

// SetRange sets the checking for the minimum and maximum lengths of a string,
// or the minimum and maximum values of a numeric type.
// The valid types for the areguments are: int, float64.
func (s *Schema) SetRange(min, max interface{}) *Schema {
	var firstInt, firstFloat bool

	switch t := min.(type) {
	case int:
		s.min = t
		firstInt = true
	case float64:
		s.minf = t
		firstFloat = true
	default:
		panic("min must be int or float64")
	}

	switch t := max.(type) {
	case int:
		if firstFloat {
			panic("min and max have to be of the same type")
		}
		s.max = t
	case float64:
		if firstInt {
			panic("min and max have to be of the same type")
		}
		s.maxf = t
	default:
		panic("max must be int or float64")
	}

	s.flagCheck |= c_MinLen | c_MaxLen
	return s
}

// SetPatternFmt sets the message about regular expression to show when
// the pattern is not matched.
func (s *Schema) SetPatternFmt(fmt string) *Schema {
	s.patternFmt = fmt
	return s
}

// SetTimeFmt sets the time format.
func (s *Schema) SetTimeFmt(fmt string) *Schema {
	s.timeFmt = fmt
	return s
}

// == Checking

var (
	ErrEmpty    = errors.New("empty string")
	ErrRequired = errors.New("required string")
)

// PreCheck checks and modifies the string according to the flags C_Required,
// M_TrimSpace, M_ToLowercase and M_ToUppercase.
func (s *Schema) PreCheck(str string, typ datautil.Type) (string, error) {
	if s.flagCheck&M_TrimSpace != 0 {
		str = strings.TrimSpace(str)
	}
	// Value by default
	if len(str) == 0 {
		if s.Bydefault != "" {
			return s.Bydefault, nil
		}
		if s.flagCheck&C_Required != 0 {
			return "", ErrRequired
		}
		// The slices need an empty string to know when stop of read.
		if !s.IsSlice {
			return "", ErrEmpty
		}
	}

	// Modification
	if typ == datautil.String || typ == datautil.Custom {
		if s.flagCheck&M_ToLowercase != 0 {
			str = strings.ToLower(str)
		} else if s.flagCheck&M_ToUppercase != 0 {
			str = strings.ToUpper(str)
		}
	}

	s.dataType = typ
	return str, nil
}

type minLenError struct{ i interface{} }

func (e minLenError) Error() string {
	return fmt.Sprintf("minimum length: %v", e.i)
}

type maxLenError struct{ i interface{} }

func (e maxLenError) Error() string {
	return fmt.Sprintf("maximum length: %v", e.i)
}

// PostCheck checks the length according to the flags c_MinLen and c_MaxLen.
func (s *Schema) PostCheck(value interface{}) error {
	if s.dataType == datautil.Bool ||
		(s.flagCheck&c_MinLen == 0 && s.flagCheck&c_MaxLen == 0) {
		return nil
	}
	if s.dataType == 0 {
		return fmt.Errorf("the type has not been set")
	}

	if s.dataType == datautil.String || s.dataType == datautil.Custom {
		// string
		value_str := value.(string)

		if s.flagCheck&c_MinLen != 0 && s.min != 0 && s.min > len(value_str) {
			return minLenError{s.min}
		}
		if s.flagCheck&c_MaxLen != 0 && s.max != 0 && s.max < len(value_str) {
			return maxLenError{s.max}
		}

	} else if s.dataType >= datautil.Int && s.dataType <= datautil.Int64 {
		// int
		var value_int int64
		switch value.(type) {
		case int:
			value_int = int64(value.(int))
		case int8:
			value_int = int64(value.(int8))
		case int16:
			value_int = int64(value.(int16))
		case int32:
			value_int = int64(value.(int32))
		case int64:
			value_int = value.(int64)
		}

		if s.flagCheck&c_MinLen != 0 && s.min != 0 && int64(s.min) > value_int {
			return minLenError{s.min}
		}
		if s.flagCheck&c_MaxLen != 0 && s.max != 0 && int64(s.max) < value_int {
			return maxLenError{s.max}
		}

	} else if s.dataType >= datautil.Uint && s.dataType <= datautil.Uint64 {
		// uint
		var value_uint uint64
		switch value.(type) {
		case uint:
			value_uint = uint64(value.(uint))
		case uint8:
			value_uint = uint64(value.(uint8))
		case uint16:
			value_uint = uint64(value.(uint16))
		case uint32:
			value_uint = uint64(value.(uint32))
		case uint64:
			value_uint = value.(uint64)
		}

		if s.flagCheck&c_MinLen != 0 && s.min != 0 && uint64(s.min) > value_uint {
			return minLenError{s.min}
		}
		if s.flagCheck&c_MaxLen != 0 && s.max != 0 && uint64(s.max) < value_uint {
			return maxLenError{s.max}
		}

	} else if s.dataType == datautil.Float32 || s.dataType == datautil.Float64 {
		// float
		var value_float float64
		switch value.(type) {
		case float32:
			value_float = float64(value.(float32))
		case float64:
			value_float = value.(float64)
		}

		if s.flagCheck&c_MinLen != 0 && s.minf != 0 && s.minf > value_float {
			return minLenError{s.minf}
		}
		if s.flagCheck&c_MaxLen != 0 && s.maxf != 0 && s.maxf < value_float {
			return maxLenError{s.maxf}
		}
	}

	return nil
}

// == Errors
//
// Based in code of (http://golang.org/src/pkg/strconv/atoi.go)

// A ValidError records a failed validation.
type ValidError struct {
	Func string // the failing function
	Str  string // the input
	Err  error  // the reason the validation failed
}

func (e *ValidError) Error() string {
	if e.Str != "" {
		return "valid." + e.Func + ": " + "parsing " + strconv.Quote(e.Str) +
			": " + e.Err.Error()
	}
	return "valid." + e.Func + ": " + e.Err.Error()
}

func passError(fn, str string, err error) *ValidError {
	return &ValidError{fn, str, err}
}
