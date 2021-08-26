// Code generated by "stringer -type=Protocol -linecomment"; DO NOT EDIT.

package easyproxy

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ProtocolHTTP-0]
	_ = x[ProtocolHTTPS-1]
	_ = x[ProtocolSOCKS5-2]
}

const _Protocol_name = "httphttpssocks5"

var _Protocol_index = [...]uint8{0, 4, 9, 15}

func (i Protocol) String() string {
	if i >= Protocol(len(_Protocol_index)-1) {
		return "Protocol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Protocol_name[_Protocol_index[i]:_Protocol_index[i+1]]
}