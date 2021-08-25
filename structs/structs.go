package structs

import (
	"sync"
)

// Scenario : S1 has a map that will be accessed by multiple goroutines, and need to be
// protected via a Mutex. You could either name that mutex inside the structure, or let
// it be anonymous. In some cases, it might be clearer to Lock the structure itself.
// It's all semantic preference.
//
// Note that you cannot do this if the anonymous member is a collection (slice/map)
type S1 struct {
	sync.Mutex

	protectedMap map[int]int
}

func (s *S1) Len() int {
	s.Lock()
	defer s.Unlock()
	return len(s.protectedMap)
}

func (s *S1) AddElement(key, val int) {
	s.Lock()
	if s.protectedMap == nil {
		s.protectedMap = make(map[int]int)
	}
	s.protectedMap[key] = val
	s.Unlock()
}

// Random useful stuff

func (_ *S1) DoSomething() {
	// Impossible to access self here. Pretty specific, but can be useful to show that
	// this function doesn't actually do anything with the structure itself.
}

type S2 struct {
	attr string
}

func (s S2) ReadOnlyMethod() {
	// This code is read-only on `s` : it's not passed by pointer, so s is a copy of whatever called it.
	// Performance-wise, unless the structure is HUGE (see http.Request for example), it doesn't really matter and makes the code clearer.
	s.attr = "changed" // this changes the copy, not the original structure
}

func (s *S2) WriteMethod() {
	// Passed by pointer. When reading this signature, I'm usually assuming this modifies
	// the structure somehow.
}

func (s S1) MutexWarning() {
	// Specifically for this kind of case : sync.Mutex needs to be passed by address so that any locker will try to lock the same mutex. This method is NOT by address (no pointer on s1 in the signature), so locking the mutex here is very useless.
	// Another way to handle this could be to make the mutex a pointer in the structure itself.
}

type s3 struct {
	*sync.Mutex // Nice, but careful on initialization!
	a           int
}

// let's make a constructor for s3
func NewS3() *s3 {
	return &s3{
		Mutex: new(sync.Mutex),
		a:     1,
	}
}

func MapWithStruct() {
	var m1 = make(map[string]s3)
	var m2 = make(map[string]*s3)

	m1["Bad"] = s3{}
	//m1["Bad"].a++ // <- Error! You have to do some very ugly stuff to change this value

	tmp := m1["Bad"] // Very
	tmp.a++          // Very
	m1["Bad"] = tmp  // Ugly!

	m2["Good"] = &s3{} // Using a pointer instead, no overhead
	m2["Good"].a++     // Works! Because it's a pointer, everything's fine =)
}
