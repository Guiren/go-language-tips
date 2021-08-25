// Code generated by "stringer -type MessageType"; DO NOT EDIT.

package stringer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Mail-1]
	_ = x[Push-2]
	_ = x[SMS-3]
}

const _MessageType_name = "UnknownMailPushSMS"

var _MessageType_index = [...]uint8{0, 7, 11, 15, 18}

func (i MessageType) String() string {
	if i < 0 || i >= MessageType(len(_MessageType_index)-1) {
		return "MessageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}