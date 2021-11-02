// Copyright 2021 The Gosip Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sip

import (
	"bytes"
	"encoding/hex"
	"math/rand"
)

const (
	hexChars = "0123456789abcdef"
)

func unhex(b byte) byte {
	switch {
	case '0' <= b && b <= '9':
		return b - '0'
	case 'a' <= b && b <= 'f':
		return b - 'a' + 10
	case 'A' <= b && b <= 'F':
		return b - 'A' + 10
	}
	return 0
}

// appendEscaped appends while URL encoding bytes that don't match the predicate..
func appendEscaped(b *bytes.Buffer, s []byte, p func(byte) bool) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if p(c) {
			b.WriteByte(c)
		} else {
			b.WriteByte('%')
			b.WriteByte(hexChars[c>>4])
			b.WriteByte(hexChars[c%16])
		}
	}
}

// appendSanitized appends stripping all characters that don't match the predicate.
func appendSanitized(b *bytes.Buffer, s []byte, p func(byte) bool) {
	for i := 0; i < len(s); i++ {
		if p(s[i]) {
			b.WriteByte(s[i])
		}
	}
}

// quote formats an address parameter value or display name.
//
// Quotation marks are only added if necessary. This implementation will
// truncate on input error.
func appendQuoted(b *bytes.Buffer, s []byte) {
	for i := 0; i < len(s); i++ {
		if !tokenc(s[i]) && s[i] != ' ' {
			appendQuoteQuoted(b, s)
			return
		}
	}
	b.Write(s)
}

// quoteQuoted formats an address parameter value or display name with quotes.
//
// This implementation will truncate on input error.
func appendQuoteQuoted(b *bytes.Buffer, s []byte) {
	b.WriteByte('"')
	for i := 0; i < len(s); i++ {
		if qdtextc(s[i]) {
			if s[i] == '\r' {
				if i+2 >= len(s) || s[i+1] != '\n' ||
					!(s[i+2] == ' ' || s[i+2] == '\t') {
					break
				}
			}
		} else {
			if !qdtextesc(s[i]) {
				break
			}
			b.WriteByte('\\')
		}
		b.WriteByte(s[i])
	}
	b.WriteByte('"')
}

// Generate a secure random number between 0 and 50,000.
func generateCSeq() int {
	return rand.Int() % 50000
}

// Generate a 48-bit secure random string like: 27c97271d363.
func generateTag() string {
	return hex.EncodeToString(randomBytes(6))
}

// Generate a SIP 2.0 Via branch ID. This is probably not suitable for use by
// stateless proxies.
func generateBranch() string {
	return "z9hG4bK-" + generateTag()
}

// Generate a secure UUID4, e.g.f47ac10b-58cc-4372-a567-0e02b2c3d479
func generateCallID() string {
	lol := randomBytes(15)
	digs := hex.EncodeToString(lol)
	uuid4 := digs[0:8] +
		"-" + digs[8:12] +
		"-4" + digs[12:15] +
		"-a" + digs[15:18] +
		"-" + digs[18:]
	return uuid4
}

func randomBytes(l int) (b []byte) {
	b = make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = byte(rand.Int())
	}
	return
}
