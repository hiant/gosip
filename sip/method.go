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

// SIP Protocol Method Definitions

package sip

type RequestMethod string

const (
	MethodInvite    = RequestMethod("INVITE")    // Indicates a client is being invited to participate in a call session.
	MethodAck       = RequestMethod("ACK")       // Confirms that the client has received a final response to an INVITE request.
	MethodBye       = RequestMethod("BYE")       // Terminates a call and can be sent by either the caller or the callee.
	MethodCancel    = RequestMethod("CANCEL")    // Cancels any pending request.
	MethodOptions   = RequestMethod("OPTIONS")   // Queries the capabilities of servers.
	MethodRegister  = RequestMethod("REGISTER")  // Registers the address listed in the To header field with a SIP server.
	MethodPrack     = RequestMethod("PRACK")     // Provisional acknowledgement.
	MethodSubscribe = RequestMethod("SUBSCRIBE") // Subscribes for an Event of Notification from the Notifier.
	MethodNotify    = RequestMethod("NOTIFY")    // Notify the subscriber of a new Event.
	MethodPublish   = RequestMethod("PUBLISH")   // Publishes an event to the Server.
	MethodInfo      = RequestMethod("INFO")      // Sends mid-session information that does not modify the session state.
	MethodRefer     = RequestMethod("REFER")     // Asks recipient to issue SIP request (call transfer.)
	MethodMessage   = RequestMethod("MESSAGE")   // Transports instant messages using SIP.
	MethodUpdate    = RequestMethod("UPDATE")    // Modifies the state of a session without changing the state of the dialog.
)

func (r RequestMethod) ToString() string {
	return string(r)
}
