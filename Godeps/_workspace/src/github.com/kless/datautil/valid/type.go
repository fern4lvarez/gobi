// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package valid

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/kless/datautil"
)

// == bool

// boolStrings represents the extra strings for boolean values.
type boolStrings map[string]bool

var boolStr boolStrings = make(map[string]bool)

// SetBoolStrings sets the extra strings for parsing boolean values in strings.
// For every key, it is also added in upper and title case.
func SetBoolStrings(m map[string]bool) {
	boolStr = m

	for k, v := range m {
		boolStr[strings.ToUpper(k)] = v
		if len(k) > 1 {
			boolStr[strings.Title(k)] = v
		}
	}
}

var ErrBool = errors.New("not boolean value")

// Bool checks if the string is a boolean value.
//
// It accepts "1", "t", "true", "y", "yes" for true value,
// and "0", "f", "false", "n", "no" for false value,
// in both lower, upper and title cases.
//
// Also, checks extra values set with "SetBoolStrings()".
func Bool(s *Schema, str string) (value bool, err error) {
	if str, err = s.PreCheck(str, datautil.Bool); err != nil {
		return
	}

	if value, err = strconv.ParseBool(str); err == nil {
		return
	}
	switch str {
	case "y", "Y", "yes", "YES", "Yes":
		return true, nil
	case "n", "N", "no", "NO", "No":
		return false, nil
	}

	if v, found := boolStr[str]; found {
		return v, nil
	}
	return false, passError("Bool", str, ErrBool)
}

// == string

var (
	ErrAlpha    = errors.New("only letters are allowed (a-zA-Z)")
	ErrAlphaNum = errors.New("only letters and numbers are allowed (a-zA-Z0-9)")
	ErrASCII    = errors.New("only ASCII characters are allowed")
	ErrString   = errors.New("the value is not a string")
)

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func isLetter(char rune) bool {
	char &^= 'a' - 'A'
	return 'A' <= char && char <= 'Z'
}

// String checks if the string is not a numeric value.
// Returns an error if it is not a string.
//
// Uses the flags C_StrictString, C_ASCII, C_Alpha and C_Alphanumeric.
func String(s *Schema, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.String); err != nil {
		return
	}

	// Check if it is a numeric value.
	if s.flagCheck&C_StrictString != 0 {
		if _, err = strconv.ParseFloat(str, 64); err == nil {
			return "", passError("String", str, ErrString)
		}
	}

	if s.flagCheck&C_ASCII != 0 {
		for _, c := range str {
			if c >= 0x80 {
				return "", passError("String", str, ErrASCII)
			}
		}
	}
	if s.flagCheck&C_Alpha != 0 {
		for _, c := range str {
			if !isLetter(c) {
				return "", passError("String", str, ErrAlpha)
			}
		}
	} else if s.flagCheck&C_Alphanumeric != 0 {
		for _, c := range str {
			if !isDigit(c) && !isLetter(c) {
				return "", passError("String", str, ErrAlphaNum)
			}
		}
	}

	return str, s.PostCheck(str)
}

//func RawBytes(s *Schema, str string) (value []byte, err error) {}

// == time

// Time checks if the string is a time.
func Time(s *Schema, str string) (value time.Time, err error) {
	if str, err = s.PreCheck(str, datautil.Time); err != nil {
		return
	}

	if value, err = time.Parse(s.timeFmt, str); err != nil {
		err = passError("Time", "", err)
	}
	return
}

// == int

// Int checks if the string is an int.
func Int(s *Schema, str string) (value int, err error) {
	if str, err = s.PreCheck(str, datautil.Int); err != nil {
		return
	}

	if v, err := strconv.ParseInt(str, 10, 0); err != nil {
		return 0, err
	} else {
		value = int(v)
	}
	return value, s.PostCheck(value)
}

// Int8 checks if the string is an int8.
func Int8(s *Schema, str string) (value int8, err error) {
	if str, err = s.PreCheck(str, datautil.Int8); err != nil {
		return
	}

	if v, err := strconv.ParseInt(str, 10, 8); err != nil {
		return 0, err
	} else {
		value = int8(v)
	}
	return value, s.PostCheck(value)
}

// Int16 checks if the string is an int16.
func Int16(s *Schema, str string) (value int16, err error) {
	if str, err = s.PreCheck(str, datautil.Int16); err != nil {
		return
	}

	if v, err := strconv.ParseInt(str, 10, 16); err != nil {
		return 0, err
	} else {
		value = int16(v)
	}
	return value, s.PostCheck(value)
}

// Int32 checks if the string is an int32.
func Int32(s *Schema, str string) (value int32, err error) {
	if str, err = s.PreCheck(str, datautil.Int32); err != nil {
		return
	}

	if v, err := strconv.ParseInt(str, 10, 32); err != nil {
		return 0, err
	} else {
		value = int32(v)
	}
	return value, s.PostCheck(value)
}

// Int64 checks if the string is an int64.
func Int64(s *Schema, str string) (value int64, err error) {
	if str, err = s.PreCheck(str, datautil.Int64); err != nil {
		return
	}

	if value, err = strconv.ParseInt(str, 10, 64); err != nil {
		return
	}
	return value, s.PostCheck(value)
}

// == uint

// Uint checks if the string is an uint.
func Uint(s *Schema, str string) (value uint, err error) {
	if str, err = s.PreCheck(str, datautil.Uint); err != nil {
		return
	}

	if v, err := strconv.ParseUint(str, 10, 0); err != nil {
		return 0, err
	} else {
		value = uint(v)
	}
	return value, s.PostCheck(value)
}

// Uint8 checks if the string is an uint8.
func Uint8(s *Schema, str string) (value uint8, err error) {
	if str, err = s.PreCheck(str, datautil.Uint8); err != nil {
		return
	}

	if v, err := strconv.ParseUint(str, 10, 8); err != nil {
		return 0, err
	} else {
		value = uint8(v)
	}
	return value, s.PostCheck(value)
}

// Uint16 checks if the string is an uint16.
func Uint16(s *Schema, str string) (value uint16, err error) {
	if str, err = s.PreCheck(str, datautil.Uint16); err != nil {
		return
	}

	if v, err := strconv.ParseUint(str, 10, 16); err != nil {
		return 0, err
	} else {
		value = uint16(v)
	}
	return value, s.PostCheck(value)
}

// Uint32 checks if the string is an uint32.
func Uint32(s *Schema, str string) (value uint32, err error) {
	if str, err = s.PreCheck(str, datautil.Uint32); err != nil {
		return
	}

	if v, err := strconv.ParseUint(str, 10, 32); err != nil {
		return 0, err
	} else {
		value = uint32(v)
	}
	return value, s.PostCheck(value)
}

// Uint64 checks if the string is an uint64.
func Uint64(s *Schema, str string) (value uint64, err error) {
	if str, err = s.PreCheck(str, datautil.Uint64); err != nil {
		return
	}

	if value, err = strconv.ParseUint(str, 10, 64); err != nil {
		return
	}
	return value, s.PostCheck(value)
}

// == float

// Float32 checks if the string is a float32.
func Float32(s *Schema, str string) (value float32, err error) {
	if str, err = s.PreCheck(str, datautil.Float32); err != nil {
		return
	}

	if v, err := strconv.ParseFloat(str, 32); err != nil {
		return 0, err
	} else {
		value = float32(v)
	}
	return value, s.PostCheck(value)
}

// Float64 checks if the string is a float64.
func Float64(s *Schema, str string) (value float64, err error) {
	if str, err = s.PreCheck(str, datautil.Float64); err != nil {
		return
	}

	if value, err = strconv.ParseFloat(str, 64); err != nil {
		return
	}
	return value, s.PostCheck(value)
}

// == Slices

// IntSlice checks if each array element is an int.
func IntSlice(s *Schema, array []string) (values []int, err error) {
	s.IsSlice = true

	for _, elem := range array {
		if i, err := Int(s, elem); err != nil {
			s.IsSlice = false
			return nil, err
		} else {
			values = append(values, i)
		}
	}

	s.IsSlice = false
	return
}

// UintSlice checks if each array element is an uint.
func UintSlice(s *Schema, array []string) (values []uint, err error) {
	s.IsSlice = true

	for _, elem := range array {
		if u, err := Uint(s, elem); err != nil {
			s.IsSlice = false
			return nil, err
		} else {
			values = append(values, u)
		}
	}

	s.IsSlice = false
	return
}

// Float64Slice checks if each array element is a float64.
func Float64Slice(s *Schema, array []string) (values []float64, err error) {
	s.IsSlice = true

	for _, elem := range array {
		if f64, err := Float64(s, elem); err != nil {
			s.IsSlice = false
			return nil, err
		} else {
			values = append(values, f64)
		}
	}

	s.IsSlice = false
	return
}

// StringSlice checks if each array element is a string.
func StringSlice(s *Schema, array []string) (values []string, err error) {
	s.IsSlice = true

	for _, elem := range array {
		if str, err := String(s, elem); err != nil {
			s.IsSlice = false
			return nil, err
		} else {
			values = append(values, str)
		}
	}

	s.IsSlice = false
	return
}

// Search

// ContainsInt checks if the string str as int is within array.
func ContainsInt(s *Schema, array []int, str string) (found bool, err error) {
	value, err := Int(s, str)
	if err != nil {
		return
	}

	for _, v := range array {
		if v == value {
			return true, nil
		}
	}
	return false, nil
}

// ContainsUint checks if the string str as uint is within array.
func ContainsUint(s *Schema, array []uint, str string) (found bool, err error) {
	value, err := Uint(s, str)
	if err != nil {
		return
	}

	for _, v := range array {
		if v == value {
			return true, nil
		}
	}
	return false, nil
}

// ContainsFloat64 checks if the string str as float64 is within array.
func ContainsFloat64(s *Schema, array []float64, str string) (found bool, err error) {
	value, err := Float64(s, str)
	if err != nil {
		return
	}

	for _, v := range array {
		if v == value {
			return true, nil
		}
	}
	return false, nil
}

// ContainsString checks if the string str is within array.
func ContainsString(s *Schema, array []string, str string) (found bool, err error) {
	if _, err = String(s, str); err != nil {
		return
	}

	for _, v := range array {
		if v == str {
			return true, nil
		}
	}
	return false, nil
}
