package constants

import (
	"math/rand"
	"net/http"
)

// Let's say we need a variable that requires an initialization. We could either use :
// - an init() function, but the initialization is away from the variable (and can be in another file, even!
// - an anonymous function as value of the variable.

// This kind of initialization is nice when you want a constant slice/map that requires complex logic to initialize it. Some developer will still be able to modify it.

var simpleConstant = map[string]int{
	"This is a simple thing":         2,
	"that doesn't require a closure": rand.Int(),
}

// This is done whenever the package loads (like init()), so if anything fails, it should
// panic right away.
var constant = func() map[string]int {
	// Complex init
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	_ = err
	// Read request body, do complex stuff
	_ = req
	return map[string]int{
		"Stuff from request Body": 1,
	}
}() // Call the function right away with () to make this variable the returned type
