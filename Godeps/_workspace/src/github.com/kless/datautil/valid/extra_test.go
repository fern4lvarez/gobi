// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package valid

import "testing"

func init() {
	schema.flagCheck |= C_TLD | C_HTTP_FTP

	if !testing.Short() {
		schema.flagCheck |= C_DNS
	}
}

// == Algorithms

func TestLuhnChecksum(t *testing.T) {
	okTests := []string{"79927398713"}
	failTests := []string{
		"79927398710", "79927398711", "79927398712", "79927398714", "79927398715",
		"79927398716", "79927398717", "79927398718", "79927398719",
	}

	for _, v := range okTests {
		if _, err := LuhnChecksum(schema, v); err != nil {
			t.Errorf("LuhnChecksum(%q) got error: %q", v, err)
		}
	}
	for i, v := range failTests {
		if _, err := LuhnChecksum(schema, v); err == nil {
			t.Errorf("LuhnChecksum(%q) want error", v)
		} else if testing.Verbose() && i == 0 {
			t.Logf("%T: %v\n", err, err)
		}
	}
}

// == Encodings

func TestBase32(t *testing.T) {
	str := "MFZG65LOMQ======"

	if _, err := Base32(schema, str); err != nil {
		t.Errorf("Base32(%q) got error: %q", str, err)
	}
	if _, err := Base32(schema, strString); err == nil {
		t.Errorf("Base32(%q) want error", strString)
	} else if testing.Verbose() {
		t.Logf("%T: %v\n", err, err)
	}
}

func TestBase64(t *testing.T) {
	str := "YXJvdW5k"

	if _, err := Base64(schema, str); err != nil {
		t.Errorf("Base64(%q) got error: %q", str, err)
	}
	if _, err := Base64(schema, strString); err == nil {
		t.Errorf("Base64(%q) want error", strString)
	} else if testing.Verbose() {
		t.Logf("%T: %v\n", err, err)
	}
}

func TestHexadecimal(t *testing.T) {
	okTests := []string{"0xaf", "af", "256"}

	for _, v := range okTests {
		if _, err := Hexadecimal(schema, v); err != nil {
			t.Errorf("Hexadecimal(%q) got error: %q", v, err)
		}
	}
	if _, err := Hexadecimal(schema, strString); err == nil {
		t.Errorf("Hexadecimal(%q) want error", strString)
	} else if testing.Verbose() {
		t.Logf("%T: %v\n", err, err)
	}
}

// == Net

func TestEmail(t *testing.T) {
	okTests := []string{"a@foo.net", "a@site.com"}
	failTests := []string{"a@site", "a@127.0.0.1"}

	for _, v := range okTests {
		if _, err := Email(schema, v); err != nil {
			t.Errorf("Email(%q) got error: %q", v, err)
		}
	}
	for i, v := range failTests {
		if _, err := Email(schema, v); err == nil {
			t.Errorf("Email(%q) want error", v)
		} else if testing.Verbose() && i == 0 {
			t.Logf("%T: %v\n", err, err)
		}
	}
}

func TestIP(t *testing.T) {
	okTests := []string{"173.194.34.212"}
	failTests := []string{"192.168.0.1"}

	for _, v := range okTests {
		if _, err := IP(schema, v); err != nil {
			t.Errorf("IP(%q) got error: %q", v, err)
		}
	}
	for i, v := range failTests {
		if _, err := IP(schema, v); err == nil {
			t.Errorf("IP(%q) want error", v)
		} else if testing.Verbose() && i == 0 {
			t.Logf("%T: %v\n", err, err)
		}
	}
}

func TestMAC(t *testing.T) {
	okTests := []string{"01:23:45:67:89:ab"}
	failTests := []string{"01:23:45:67:89:a"}

	for _, v := range okTests {
		if _, err := MAC(schema, v); err != nil {
			t.Errorf("MAC(%q) got error: %q", v, err)
		}
	}
	for i, v := range failTests {
		if _, err := MAC(schema, v); err == nil {
			t.Errorf("MAC(%q) want error", v)
		} else if testing.Verbose() && i == 0 {
			t.Logf("%T: %v\n", err, err)
		}
	}
}

func TestURL(t *testing.T) {
	okTests := []string{"https://www.google.com"}
	failTests := []string{"foo.com", "https://foo.blogspot.co.uk"}

	for _, v := range okTests {
		if _, err := URL(schema, v); err != nil {
			t.Errorf("URL(%q) got error: %q", v, err)
		}
	}
	for i, v := range failTests {
		if _, err := URL(schema, v); err == nil {
			t.Errorf("URL(%q) want error", v)
		} else if testing.Verbose() && i == 0 {
			t.Logf("%T: %v\n", err, err)
		}
	}
}

// == Regular expressions

func TestPattern(t *testing.T) {
	re := `[a-z]+`
	schema.patternFmt = "letters in lowercase"

	_, err := Pattern(schema, re, strString)
	if err != nil {
		t.Errorf("Pattern(%q, %q) got error: %q", re, strString, err)
	}
	if _, err = Pattern(schema, re, intString); err == nil {
		t.Errorf("Pattern(%q, %q) want error", re, intString)
	} else if testing.Verbose() {
		t.Logf("%T: %v\n", err, err)
	}

	re, schema.patternFmt = `[0-9]+`, ""

	if _, err := Pattern(schema, re, intString); err != nil {
		t.Errorf("Pattern(%q, %q) got error: %q", re, intString, err)
	}
	if _, err = Pattern(schema, re, strString); err == nil {
		t.Errorf("Pattern(%q, %q) want error", re, strString)
	} else if testing.Verbose() {
		t.Logf("%T: %v\n", err, err)
	}

	// Check length of cache
	if len(reCache) != 2 {
		t.Errorf("len(reCache) got %v, want %v", len(reCache), 2)
	}
	Pattern(schema, re, strString)

	if len(reCache) != 2 {
		t.Errorf("len(reCache) got %v, want %v", len(reCache), 2)
	}

	// Bad pattern
	re = `[a-z+`

	if _, err = Pattern(schema, re, strString); err == nil {
		t.Errorf("Pattern(%q, %q) want error", re, strString)
	} else if testing.Verbose() {
		t.Logf("%T: %v\n", err, err)
	}
}
