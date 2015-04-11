package sip

import (
	"bytes"
)

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
