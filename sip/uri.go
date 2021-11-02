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

// SIP URI Library
//
// We can't use net.URL because it doesn't support SIP URIs. This is because:
// a) it doesn't support semicolon parameters; b) it doesn't extract the user
// and host information when the "//" isn't present.
//
// For example:
//
//   jart@example.test;isup-oli=29
//
// Roughly equates to:
//
//   {Scheme: "sip",
//    User: "jart",
//    Pass: "",
//    Host: "example.test",
//    Port: "",
//    Params: {"isup-oli": "29"}}
//

package sip

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	delims = ":/@;?#<>"
)

// URIParam is a linked list of ;key=values for URI parameters.
type URIParam struct {
	Name  string
	Value string
	Next  *URIParam
}

// URIHeader is a linked list of ?key=values for URI headers.
type URIHeader struct {
	Name  string
	Value string
	Next  *URIHeader
}

type URI struct {
	Scheme string     // e.g. sip, sips, tel, etc.
	User   string     // e.g. sip:USER@host
	Pass   string     // e.g. sip:user:PASS@host
	Host   string     // e.g. example.com, 1.2.3.4, etc.
	Port   uint16     // e.g. 5060, 80, etc.
	Param  *URIParam  // e.g. ;isup-oli=00;day=tuesday
	Header *URIHeader // e.g. ?subject=project%20x&lol=cat
}

//go:generate ragel -Z -G2 -o uri_parse.go uri_parse.rl

// Deep copies a URI object.
func (u *URI) Copy() *URI {
	if u == nil {
		return nil
	}
	res := new(URI)
	*res = *u
	res.Param = u.Param
	res.Header = u.Header
	return res
}

// TODO(jart): URI Comparison https://tools.ietf.org/html/rfc3261#section-19.1.4

func (u *URI) String() string {
	var b bytes.Buffer
	u.Append(&b)
	return b.String()
}

func (u *URI) Append(b *bytes.Buffer) {
	if u == nil {
		return
	}
	if u.Scheme == "" {
		b.WriteString("sip:")
	} else {
		b.WriteString(u.Scheme)
		b.WriteByte(':')
	}
	if u.User != "" {
		if u.Pass != "" {
			appendEscaped(b, []byte(u.User), userc)
			b.WriteByte(':')
			appendEscaped(b, []byte(u.Pass), passc)
		} else {
			appendEscaped(b, []byte(u.User), userc)
		}
		b.WriteByte('@')
	}
	if strings.Contains(u.Host, ":") {
		b.WriteByte('[')
		b.WriteString(u.Host)
		b.WriteByte(']')
	} else {
		b.WriteString(u.Host)
	}
	if u.Port > 0 {
		b.WriteByte(':')
		b.WriteString(strconv.FormatInt(int64(u.Port), 10))
	}
	u.Param.Append(b)
	u.Header.Append(b)
}

func (u *URI) CompareHostPort(other *URI) bool {
	if u != nil && other != nil {
		if u.Host == other.Host && u.GetPort() == other.GetPort() {
			return true
		}
	}
	return false
}

func (u *URI) GetPort() uint16 {
	if u.Port == 0 {
		if u.Scheme == "sips" {
			return 5061
		} else {
			return 5060
		}
	} else {
		return u.Port
	}
}

// Get returns an entry in O(n) time.
func (p *URIParam) Get(name string) *URIParam {
	if p == nil {
		return nil
	}
	if strings.EqualFold(p.Name, name) {
		return p
	}
	return p.Next.Get(name)
}

// Append serializes URI parameters in insertion order.
func (p *URIParam) Append(b *bytes.Buffer) {
	if p == nil {
		return
	}
	p.Next.Append(b)
	b.WriteByte(';')
	appendEscaped(b, []byte(p.Name), paramc)
	if p.Value != "" {
		b.WriteByte('=')
		appendEscaped(b, []byte(p.Value), paramc)
	}
}

// Get returns an entry in O(n) time.
func (p *URIHeader) Get(name string) *URIHeader {
	if p == nil {
		return nil
	}
	if strings.EqualFold(p.Name, name) {
		return p
	}
	return p.Next.Get(name)
}

// Append serializes URI headers in insertion order.
func (p *URIHeader) Append(b *bytes.Buffer) {
	if p == nil {
		return
	}
	if p.Next != nil {
		p.Next.Append(b)
		b.WriteByte('&')
	} else {
		b.WriteByte('?')
	}
	appendEscaped(b, []byte(p.Name), paramc)
	if p.Value != "" {
		b.WriteByte('=')
		appendEscaped(b, []byte(p.Value), paramc)
	}
}
