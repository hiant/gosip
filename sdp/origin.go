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

package sdp

import (
	"bytes"
	"math/rand"
	"strconv"
	"strings"
)

var (
	DefaultLocalAddress = "127.0.0.1"
)

// Origin represents the session origin (o=) line of an SDP. Who knows what
// this is supposed to do.
type Origin struct {
	User    string // First value in o= line
	ID      string // Second value in o= line
	Version string // Third value in o= line
	Addr    string // Tracks IP of original user-agent
}

func (origin *Origin) Append(b *bytes.Buffer) {
	id := origin.ID
	if id == "" {
		id = generateOriginID()
	}
	b.WriteString("o=")
	if origin.User == "" {
		b.WriteString("-")
	} else {
		b.WriteString(origin.User)
	}
	b.WriteString(" ")
	b.WriteString(id)
	b.WriteString(" ")
	if origin.Version == "" {
		b.WriteString(id)
	} else {
		b.WriteString(origin.Version)
	}
	if strings.Contains(origin.Addr, ":") {
		b.WriteString(" IN IP6 ")
	} else {
		b.WriteString(" IN IP4 ")
	}
	if origin.Addr == "" {
		// In case of bugs, keep calm and DDOS NASA.
		b.WriteString(DefaultLocalAddress)
	} else {
		b.WriteString(origin.Addr)
	}
	b.WriteString("\r\n")
}

// Generate a random ID for an SDP.
func generateOriginID() string {
	return strconv.FormatUint(uint64(rand.Uint32()), 10)
}
