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

// SIP Via Address Library

package sip

import (
	"bytes"
	"strconv"
)

// Example: SIP/2.0/UDP 1.2.3.4:5060;branch=z9hG4bK556f77e6.
type Via struct {
	Protocol  string // should be "SIP"
	Version   string // protocol version e.g. "2.0"
	Transport string // transport type "UDP"
	Host      string // name or ip of egress interface
	Port      uint16 // network port number
	Param     *Param // param like branch, received, rport, etc.
	Next      *Via   // pointer to next via header if any
}

func (v *Via) Append(b *bytes.Buffer) {
	if v.Protocol == "" {
		b.WriteString("SIP/")
	} else {
		b.WriteString(v.Protocol)
		b.WriteString("/")
	}
	if v.Version == "" {
		b.WriteString("2.0/")
	} else {
		b.WriteString(v.Version)
		b.WriteString("/")
	}
	if v.Transport == "" {
		b.WriteString("UDP ")
	} else {
		b.WriteString(v.Transport)
		b.WriteString(" ")
	}
	b.WriteString(v.Host)
	if v.Port != 5060 {
		b.WriteString(":")
		b.WriteString(strconv.Itoa(int(v.Port)))
	}
	v.Param.Append(b)
}

// Copy returns a deep copy of via.
func (v *Via) Copy() *Via {
	if v == nil {
		return nil
	}
	res := new(Via)
	res.Protocol = v.Protocol
	res.Version = v.Version
	res.Transport = v.Transport
	res.Host = v.Host
	res.Port = v.Port
	res.Param = v.Param
	res.Next = v.Next.Copy()
	return res
}

// Branch adds a randomly generated branch ID.
func (v *Via) Branch() *Via {
	v.Param = &Param{"branch", generateBranch(), v.Param}
	return v
}

// Detach returns a shallow copy of via with Next set to nil.
func (v *Via) Detach() *Via {
	res := new(Via)
	*res = *v
	res.Next = nil
	return res
}

// Last returns pointer to last via in linked list.
func (v *Via) Last() *Via {
	if v != nil {
		for ; v.Next != nil; v = v.Next {
		}
	}
	return v
}

func (v *Via) CompareHostPort(other *Via) bool {
	if v != nil && other != nil {
		if v.Host == other.Host &&
			v.Port == other.Port {
			return true
		}
	}
	return false
}

func (v *Via) CompareBranch(other *Via) bool {
	if v != nil && other != nil {
		if b1 := v.Param.Get("branch"); b1 != nil {
			if b2 := other.Param.Get("branch"); b2 != nil {
				if b1.Value == b2.Value {
					return true
				}
			}
		}
	}
	return false
}
