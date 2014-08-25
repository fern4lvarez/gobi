// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package valid_test

import (
	"fmt"

	"github.com/kless/datautil/valid"
)

var schema = valid.NewSchema(valid.C_Required | valid.M_TrimSpace)

func ExampleString() {
	dst, err := valid.String(schema, " foo  ")
	fmt.Printf("%q : %v\n", dst, err)

	dst, err = valid.String(schema, "")
	fmt.Printf("%q : %v\n", dst, err)

	schema.Bydefault = "bar"
	dst, err = valid.String(schema, "")
	fmt.Printf("%q : %v\n", dst, err)

	schema.SetChecker(valid.M_ToUppercase)
	dst, err = valid.String(schema, " ja")
	fmt.Printf("%q : %v\n", dst, err)

	// Output:
	// "foo" : <nil>
	// "" : required string
	// "bar" : <nil>
	// " JA" : <nil>
}

func ExampleInt() {
	dst, err := valid.Int(schema, "666")
	fmt.Printf("%d : %v\n", dst, err)

	schema.SetRange(6, 66)
	dst, err = valid.Int(schema, "5")
	fmt.Printf("%d : %v\n", dst, err)
	dst, err = valid.Int(schema, "67")
	fmt.Printf("%d : %v\n", dst, err)

	// Output:
	// 666 : <nil>
	// 5 : minimum length: 6
	// 67 : maximum length: 66
}

func ExamplePattern() {
	schema.SetChecker(0)
	re, input := `[a-z]+`, "FOO"

	dst, err := valid.Pattern(schema, re, input)
	fmt.Printf("%q : %v\n", dst, err)

	schema.SetPatternFmt("letters in lowercase")
	dst, err = valid.Pattern(schema, re, input)
	fmt.Printf("%q : %v\n", dst, err)

	// Output:
	// "" : valid.Pattern: parsing "FOO": does not matches with expression: `[a-z]+`
	// "" : valid.Pattern: parsing "FOO": does not matches with format: letters in lowercase
}
