package stringer

import (
	"fmt"
	"log"
)

// Stringer implemented via go:generate

//go:generate stringer -type MessageType
type MessageType int

const (
	Unknown MessageType = iota
	Mail
	Push
	SMS
)

// This would act as a conversion map from whatever input to a stable MessageType.
var dbToMessageType = map[int]MessageType{
	5: Mail,
	2: Push,
	1: SMS,
}

// Simple example of parsing interface{} data into a simple message type counter.
func CountMessagePerType(fromRows []interface{}) map[MessageType]int {
	var result = make(map[MessageType]int)

	for _, r := range fromRows {
		dbMt, ok := r.(int)
		if !ok {
			log.Println("Message type", r, "not recognized")
			continue
		}
		mt, ok := dbToMessageType[dbMt]
		result[mt]++
	}

	return result
}

// Stringer directly implemented. Usage example in all_test.go
type MyStruct struct {
	IdContact uint64
	Messages  []MessageType
}

// String() implements the Stringer interface by writing this method.
// This just prints the number of messages received by this instance,
// not the detail, assuming we don't care for debugging logs for example.
func (ms MyStruct) String() string {
	return fmt.Sprintf("%v received %v messages", ms.IdContact, len(ms.Messages))
}
