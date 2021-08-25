package lookup

import (
	"math/rand"
	"unsafe"
)

func SizeOfBool(m map[int]bool) interface{} {
	var dummyInt int
	var dummyBool bool
	return (len(m) * 8) + (len(m) * 8 * int(unsafe.Sizeof(dummyInt))) + (len(m) * 8 * int(unsafe.Sizeof(dummyBool)))
}

func SizeOfStruct(m map[int]struct{}) interface{} {
	var dummyInt int
	var dummyStruct struct{}

	return (len(m) * 8) + (len(m) * 8 * int(unsafe.Sizeof(dummyInt))) + (len(m) * 8 * int(unsafe.Sizeof(dummyStruct)))
}

// Example of a lookup map usage. Lookup maps need to answer the question :
// "Is this element in here?" - and that's it. They could have interesting values associated
// with the key that was just found, but mainly lookup is important
func LookupMapSetup() (map[int]struct{}, int) {
	var myLookup = make(map[int]struct{}, 10)
	var cheat int

	// init
	for i := 0; i < 10; i++ {
		random := (rand.Int() % 10) + 1
		if cheat == 0 {
			cheat = random
		}
		myLookup[random] = struct{}{}
	}

	return myLookup, cheat
}
