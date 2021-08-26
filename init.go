package main

import (
	"log"
	"math/rand"
	"time"
)

var wasInit = false

// init() is a special function. It is executed if the package is imported (or for package main, before main is called).
// This kind of function is the reason why some packages ask for anonymous import, like mysql driver.
//
// Note that you cannot call init() from the code
func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("init() function called! :)")
	wasInit = true

	rand.Seed(time.Now().UnixNano())
}
