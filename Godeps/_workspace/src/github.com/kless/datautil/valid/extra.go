// Copyright 2014 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package valid

import (
	"encoding/base32"
	"encoding/base64"
	"errors"
	"fmt"
	"hash/adler32"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"code.google.com/p/go.net/publicsuffix"
	"github.com/kless/datautil"
)

// == Algorithms

var ErrLuhnChecksum = errors.New("invalid checksum")

// LuhnChecksum validates a variety of identification numbers, such as
// credit card numbers and IMEI numbers.
//
// https://en.wikipedia.org/wiki/Luhn_algorithm
func LuhnChecksum(s *Schema, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	odd := len(str) & 1
	sum := 0

	// From the rightmost digit, moving left
	for i := len(str) - 1; i >= 0; i-- {
		num64b, err := strconv.ParseInt(string(str[i]), 10, 0)
		if err != nil {
			return "", passError("LuhnChecksum", str, err)
		}
		num := int(num64b)

		if i&1 == odd {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}

	if (sum % 10) != 0 {
		return "", passError("LuhnChecksum", str, ErrLuhnChecksum)
	}
	return str, nil
}

// == Encodings

// Base32 checks if the string is encoded in base32 and returns the value decoded.
func Base32(s *Schema, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	data, err := base32.StdEncoding.DecodeString(str)
	if err == nil {
		return string(data), nil
	}
	if data, err = base32.HexEncoding.DecodeString(str); err == nil {
		return string(data), nil
	}

	err = passError("Base32", str, err)
	return
}

// Base64 checks if the string is encoded in base64 and returns the value decoded.
func Base64(s *Schema, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	data, err := base64.StdEncoding.DecodeString(str)
	if err == nil {
		return string(data), nil
	}
	if data, err = base64.URLEncoding.DecodeString(str); err == nil {
		return string(data), nil
	}

	err = passError("Base64", str, err)
	return
}

// Hexadecimal checks if the string is an hexadecimal value and returns the
// corresponding value.
func Hexadecimal(s *Schema, str string) (value int64, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	// Remove "0x"
	if len(str) > 2 && str[0] == '0' && str[1] == 'x' {
		str = str[2:]
	}

	if value, err = strconv.ParseInt(str, 16, 64); err != nil {
		return 0, passError("Hexadecimal", "", err)
	}
	return
}

// == Net

var (
	ErrIPDomain = errors.New("IP address in domain")
	ErrNoIP     = errors.New("no IP address")
	ErrNoScheme = errors.New(`no protocol scheme (like "http://")`)
	ErrHTTP_FTP = errors.New("no protocol schemes for http,https,ftp")
)

type ICANNError string

func (e ICANNError) Error() string {
	return `no ICANN domain: "` + string(e) + `"`
}

type URLSchemeError []string

func (e URLSchemeError) Error() string {
	return "no protocol scheme for " + strings.Join([]string(e), ",")
}

// CheckDomain uses the flag C_TLD to checking if the domain is ICANN's
// and the flag C_DNS to checking if it is mapped to an address.
func CheckDomain(flag Checker, domain string, forEmail bool) (err error) {
	if flag&C_TLD != 0 {
		if ip := net.ParseIP(domain); ip != nil {
			return ErrIPDomain
		}
		if suffix, isIcann := publicsuffix.PublicSuffix(domain); !isIcann {
			return ICANNError(suffix)
		}
	}
	if flag&C_DNS != 0 {
		if domain, err = publicsuffix.EffectiveTLDPlusOne(domain); err != nil {
			return err
		}

		if forEmail {
			_, err = net.LookupMX(domain)
		} else {
			_, err = net.LookupHost(domain)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// Email checks if the string is an email address correctly parsed and returns
// the address part.
// Uses the flag C_TLD to checking the domain and C_DNS for the DNS MX records.
func Email(s *Schema, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	addr, err := mail.ParseAddress(str)
	if err != nil {
		return "", passError("Email", str, err)
	}
	str = addr.Address

	if s.flagCheck != 0 {
		err = CheckDomain(s.flagCheck, strings.SplitAfter(str, "@")[1], true)
		if err != nil {
			return "", passError("Email", str, err)
		}
	}

	return str, nil
}

// IP checks if the string is an IP address.
// Uses the flag C_TLD to checking the domain by a reverse lookup.
func IP(s *Schema, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	if ip := net.ParseIP(str); ip == nil {
		return "", passError("IP", str, ErrNoIP)
	}
	if s.flagCheck&C_DNS != 0 {
		if _, err = net.LookupAddr(str); err != nil {
			return "", passError("IP", str, err)
		}
	}

	return str, nil
}

// MAC checks if the string is a hardware address.
func MAC(s *Schema, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	if _, err = net.ParseMAC(str); err != nil {
		return "", passError("MAC", "", err)
	}
	return str, nil
}

// URL checks if the string is an URL and returns a valid URL string.
// Uses the flags C_HTTP_FTP to checking the protocol schemes for HTTP and FTP,
// C_TLD for the domain, and C_DNS for the DNS records.
func URL(s *Schema, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	u, err := url.Parse(str)
	if err != nil {
		return "", passError("URL", str, err)
	}

	if u.Scheme == "" {
		return "", passError("URL", str, ErrNoScheme)
	}

	if s.flagCheck != 0 {
		if s.flagCheck&C_HTTP_FTP != 0 {
			switch u.Scheme {
			case "http", "https", "ftp": // found
			default:
				return "", passError("URL", str, ErrHTTP_FTP)
			}
		}

		if err = CheckDomain(s.flagCheck, u.Host, false); err != nil {
			return "", passError("URL", str, err)
		}
	}

	return u.String(), nil
}

// == Regular expressions

type PatternError struct {
	expr   string
	format string
}

func (e *PatternError) Error() string {
	if e.format != "" {
		return fmt.Sprintf("does not matches with format: %s", e.format)
	} else {
		return fmt.Sprintf("does not matches with expression: `%s`", e.expr)
	}
}

// regexpCache represents a cache of compiled regular expression.
type regexpCache map[uint32]*regexp.Regexp

var reCache regexpCache = make(map[uint32]*regexp.Regexp, 5)

// Pattern checks if the string matches a regular expression.
//
// Internally, uses a cache to store every new compiled regular expression.
func Pattern(s *Schema, expr, str string) (value string, err error) {
	if str, err = s.PreCheck(str, datautil.Custom); err != nil {
		return
	}

	crc := adler32.Checksum([]byte(expr))
	re, found := reCache[crc]
	if !found {
		re, err = regexp.Compile(expr)
		if err != nil {
			return "", passError("Pattern", str, err)
		}

		reCache[crc] = re
	}

	if ok := re.MatchString(str); !ok {
		return "", passError("Pattern", str, &PatternError{expr, s.patternFmt})
	}
	return str, nil
}
