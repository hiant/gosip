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

// SIP Protocol Status Definitions
//
// http://www.iana.org/assignments/sip-parameters

package sip

import (
	"strconv"
)

type StatusCode uint16

const (
	// 1xx: Provisional -- request received, continuing to process the request.
	StatusTrying               = StatusCode(100) // Indicates server is not totally pwnd.
	StatusRinging              = StatusCode(180) // Remote phone is definitely ringing.
	StatusCallIsBeingForwarded = StatusCode(181)
	StatusQueued               = StatusCode(182)
	StatusSessionProgress      = StatusCode(183) // Establish early media (PSTN ringback)

	// 2xx: Success -- the action was successfully received, understood,
	//      and accepted;
	StatusOK             = StatusCode(200) // Call is answered
	StatusAccepted       = StatusCode(202) // [RFC3265]
	StatusNoNotification = StatusCode(204) // [RFC5839]

	// 3xx: Redirection -- further action needs to be taken in order to
	//      complete the request;
	StatusMultipleChoices    = StatusCode(300)
	StatusMovedPermanently   = StatusCode(301)
	StatusMovedTemporarily   = StatusCode(302) // Send your call there instead kthx.
	StatusUseProxy           = StatusCode(305) // You fool! Send your call there instead.
	StatusAlternativeService = StatusCode(380)

	// 4xx: Client Error -- the request contains bad syntax or cannot be
	//      fulfilled at this server;
	StatusBadRequest                   = StatusCode(400) // Missing headers, bad format, etc.
	StatusUnauthorized                 = StatusCode(401) // Resend request with auth header.
	StatusPaymentRequired              = StatusCode(402) // I am greedy.
	StatusForbidden                    = StatusCode(403) // gtfo
	StatusNotFound                     = StatusCode(404) // wat?
	StatusMethodNotAllowed             = StatusCode(405) // I don't support that type of request.
	StatusNotAcceptable                = StatusCode(406)
	StatusProxyAuthenticationRequired  = StatusCode(407)
	StatusRequestTimeout               = StatusCode(408)
	StatusConflict                     = StatusCode(409)
	StatusGone                         = StatusCode(410) // Shaniqua don't live here no more.
	StatusLengthRequired               = StatusCode(411)
	StatusConditionalRequestFailed     = StatusCode(412) // [RFC3903]
	StatusRequestEntityTooLarge        = StatusCode(413)
	StatusRequestURITooLong            = StatusCode(414)
	StatusUnsupportedMediaType         = StatusCode(415)
	StatusUnsupportedURIScheme         = StatusCode(416)
	StatusUnknownResourcePriority      = StatusCode(417)
	StatusBadExtension                 = StatusCode(420)
	StatusExtensionRequired            = StatusCode(421)
	StatusSessionIntervalTooSmall      = StatusCode(422) // [RFC4028]
	StatusIntervalTooBrief             = StatusCode(423)
	StatusUseIdentityHeader            = StatusCode(428) // [RFC4474]
	StatusProvideReferrerIdentity      = StatusCode(429) // [RFC3892]
	StatusFlowFailed                   = StatusCode(430) // [RFC5626]
	StatusAnonymityDisallowed          = StatusCode(433) // [RFC5079]
	StatusBadIdentityInfo              = StatusCode(436) // [RFC4474]
	StatusUnsupportedCertificate       = StatusCode(437) // [RFC4474]
	StatusInvalidIdentityHeader        = StatusCode(438) // [RFC4474]
	StatusFirstHopLacksOutboundSupport = StatusCode(439) // [RFC5626]
	StatusMaxBreadthExceeded           = StatusCode(440) // [RFC5393]
	StatusConsentNeeded                = StatusCode(470) // [RFC5360]
	StatusTemporarilyUnavailable       = StatusCode(480) // fast busy or soft fail
	StatusCallTransactionDoesNotExist  = StatusCode(481) // Bad news
	StatusLoopDetected                 = StatusCode(482) // Froot looping
	StatusTooManyHops                  = StatusCode(483) // Froot looping
	StatusAddressIncomplete            = StatusCode(484)
	StatusAmbiguous                    = StatusCode(485)
	StatusBusyHere                     = StatusCode(486)
	StatusRequestTerminated            = StatusCode(487)
	StatusNotAcceptableHere            = StatusCode(488)
	StatusBadEvent                     = StatusCode(489) // [RFC3265]
	StatusRequestPending               = StatusCode(491)
	StatusUndecipherable               = StatusCode(493)
	StatusSecurityAgreementRequired    = StatusCode(494) // [RFC3329]

	// 5xx: Server Error -- the server failed to fulfill an apparently
	//      valid request;
	StatusInternalServerError = StatusCode(500)
	StatusNotImplemented      = StatusCode(501)
	StatusBadGateway          = StatusCode(502)
	StatusServiceUnavailable  = StatusCode(503)
	StatusGatewayTimeout      = StatusCode(504)
	StatusVersionNotSupported = StatusCode(505)
	StatusMessageTooLarge     = StatusCode(513)
	StatusPreconditionFailure = StatusCode(580) // [RFC3312]

	// 6xx: Global Failure -- the request cannot be fulfilled at any
	//      server.
	StatusBusyEverywhere       = StatusCode(600)
	StatusDecline              = StatusCode(603)
	StatusDoesNotExistAnywhere = StatusCode(604)
	StatusNotAcceptable606     = StatusCode(606)
	StatusDialogTerminated     = StatusCode(687)
)

func (s StatusCode) Phrase() string {
	if phrase, ok := phrases[s]; ok {
		return phrase
	}
	return "Unknown Status Code"
}

func (s StatusCode) ToString() string {
	return strconv.Itoa(int(s))
}

func Phrase(status int) string {
	if phrase, ok := phrases[StatusCode(status)]; ok {
		return phrase
	}
	return "Unknown Status Code"
}

var phrases = map[StatusCode]string{
	StatusTrying:                       "Trying",
	StatusRinging:                      "Ringing",
	StatusCallIsBeingForwarded:         "Call Is Being Forwarded",
	StatusQueued:                       "Queued",
	StatusSessionProgress:              "Session Progress",
	StatusOK:                           "OK",
	StatusAccepted:                     "Accepted",
	StatusNoNotification:               "No Notification",
	StatusMultipleChoices:              "Multiple Choices",
	StatusMovedPermanently:             "Moved Permanently",
	StatusMovedTemporarily:             "Moved Temporarily",
	StatusUseProxy:                     "Use Proxy",
	StatusAlternativeService:           "Alternative Service",
	StatusBadRequest:                   "Bad Request",
	StatusUnauthorized:                 "Unauthorized",
	StatusPaymentRequired:              "Payment Required",
	StatusForbidden:                    "Forbidden",
	StatusNotFound:                     "Not Found",
	StatusMethodNotAllowed:             "Method Not Allowed",
	StatusNotAcceptable:                "Not Acceptable",
	StatusProxyAuthenticationRequired:  "Proxy Authentication Required",
	StatusRequestTimeout:               "Request Timeout",
	StatusConflict:                     "Conflict",
	StatusGone:                         "Gone",
	StatusLengthRequired:               "Length Required",
	StatusConditionalRequestFailed:     "Conditional Request Failed",
	StatusRequestEntityTooLarge:        "Request Entity Too Large",
	StatusRequestURITooLong:            "Request-URI Too Long",
	StatusUnsupportedMediaType:         "Unsupported Media Type",
	StatusUnsupportedURIScheme:         "Unsupported URI Scheme",
	StatusUnknownResourcePriority:      "Unknown Resource-Priority",
	StatusBadExtension:                 "Bad Extension",
	StatusExtensionRequired:            "Extension Required",
	StatusSessionIntervalTooSmall:      "Session Interval Too Small",
	StatusIntervalTooBrief:             "Interval Too Brief",
	StatusUseIdentityHeader:            "Use Identity Header",
	StatusProvideReferrerIdentity:      "Provide Referrer Identity",
	StatusFlowFailed:                   "Flow Failed",
	StatusAnonymityDisallowed:          "Anonymity Disallowed",
	StatusBadIdentityInfo:              "Bad Identity-Info",
	StatusUnsupportedCertificate:       "Unsupported Certificate",
	StatusInvalidIdentityHeader:        "Invalid Identity Header",
	StatusFirstHopLacksOutboundSupport: "First Hop Lacks Outbound Support",
	StatusMaxBreadthExceeded:           "Max-Breadth Exceeded",
	StatusConsentNeeded:                "Consent Needed",
	StatusTemporarilyUnavailable:       "Temporarily Unavailable",
	StatusCallTransactionDoesNotExist:  "Call/Transaction Does Not Exist",
	StatusLoopDetected:                 "Loop Detected",
	StatusTooManyHops:                  "Too Many Hops",
	StatusAddressIncomplete:            "Address StatusIncomplete",
	StatusAmbiguous:                    "Ambiguous",
	StatusBusyHere:                     "Busy Here",
	StatusRequestTerminated:            "Request Terminated",
	StatusNotAcceptableHere:            "Not Acceptable Here",
	StatusBadEvent:                     "Bad Event",
	StatusRequestPending:               "Request Pending",
	StatusUndecipherable:               "Undecipherable",
	StatusSecurityAgreementRequired:    "Security Agreement Required",
	StatusInternalServerError:          "Internal Server Error",
	StatusNotImplemented:               "Not Implemented",
	StatusBadGateway:                   "Bad Gateway",
	StatusServiceUnavailable:           "Service Unavailable",
	StatusGatewayTimeout:               "Gateway Time-out",
	StatusVersionNotSupported:          "Version Not Supported",
	StatusMessageTooLarge:              "Message Too Large",
	StatusPreconditionFailure:          "Precondition Failure",
	StatusBusyEverywhere:               "Busy Everywhere",
	StatusDecline:                      "Decline",
	StatusDoesNotExistAnywhere:         "Does Not Exist Anywhere",
	StatusNotAcceptable606:             "Not Acceptable",
	StatusDialogTerminated:             "Dialog Terminated",
}
