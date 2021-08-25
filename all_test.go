package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"sort"
	"strconv"
	"sync"
	"testing"
	"thecoolthings/closer"
	"thecoolthings/goroutines"
	"thecoolthings/lookup"
	"thecoolthings/slice_tricks"
	"thecoolthings/sorter"
	"thecoolthings/stringer"
	"thecoolthings/structs"
	"unsafe"
)

func TestInit(t *testing.T) {
	assert.Equal(t, true, wasInit)
}

func TestCloseAll(t *testing.T) {
	var s1, s2, s3 closer.Closeable

	errs := closer.CloseAll(&s1, &s2, &s3)
	if errs != nil {
		t.Error(errs)
	}
}

func TestCountMessagePerType(t *testing.T) {
	var dummyRows = []interface{}{
		5, 5, 5, 1, 1, 2, 4,
	}

	result := stringer.CountMessagePerType(dummyRows)
	for mt, count := range result {
		log.Println("There is", count, mt)
		switch mt {
		case stringer.Push:
			assert.Equal(t, count, 1)
		case stringer.Mail:
			assert.Equal(t, count, 3)
		case stringer.SMS:
			assert.Equal(t, count, 2)
		}
	}
}

func TestMyStruct_String(t *testing.T) {
	type testData struct {
		expected string
		data     stringer.MyStruct
	}
	var cases = map[uint64]testData{
		1: {
			expected: "1 received 3 messages",
			data: stringer.MyStruct{
				IdContact: 1,
				Messages:  []stringer.MessageType{stringer.Push, stringer.Push, stringer.SMS},
			},
		},
		2: {
			expected: "2 received 1 messages",
			data: stringer.MyStruct{
				IdContact: 2,
				Messages:  []stringer.MessageType{stringer.Unknown},
			},
		},
		3: {
			expected: "3 received 0 messages",
			data: stringer.MyStruct{
				IdContact: 3,
			},
		},
	}

	for id, c := range cases {
		// Simulate an output on a log via Sprintf ; this should call the String() method
		output := fmt.Sprintf("%v", c.data)
		assert.Equal(t, c.expected, output, id)
	}
}

func TestLookupSizeOf(t *testing.T) {
	// Making 2 "lookup" maps of 8 elements to show size diffs.
	var i = 0
	var mb = make(map[int]bool, 8)
	for i = 0; i < 8; i++ {
		mb[i] = true
	}
	sb := lookup.SizeOfBool(mb)
	log.Println("SizeOfBool", sb)

	i = 0
	var ms = make(map[int]struct{}, 8)
	for i = 0; i < 8; i++ {
		ms[i] = struct{}{}
	}
	i = 0
	ss := lookup.SizeOfStruct(ms)
	log.Println("SizeOfStruct", ss)

	assert.Less(t, ss, sb)
	// Showing that struct{} instanciated actually is 0 bytes.
	log.Println("Size of a struct{}{} :", int(unsafe.Sizeof(struct{}{})))
	assert.Equal(t, 0, int(unsafe.Sizeof(struct{}{})))
}

func TestLookupMaps(t *testing.T) {
	// cheat is just a value that *is* in the lookup map for the test
	myLookup, cheat := lookup.LookupMapSetup()
	// actual usage of a lookup map :
	// on "if _, exists", we don't care about the _ value, it's struct{}
	if _, exists := myLookup[cheat]; exists {
		log.Println("Found it!", cheat)
	} else {
		t.Fail()
	}

	// > 10 can't be in the map, just checking here
	if _, exists := myLookup[14]; exists {
		t.Fail()
	}
}

func TestFilterEvenNumbers(t *testing.T) {
	type sliceDesc struct {
		s      []int
		length int
	}
	sA := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 100, 200, 300,
	}
	var a = sliceDesc{
		s:      sA,
		length: len(sA),
	}

	var sExp = []int{0, 2, 4, 6, 8, 100, 200, 300}
	var expected = sliceDesc{
		s:      sExp,
		length: len(sExp),
	}

	// New value of A should be this
	var residue = []int{0, 2, 4, 6, 8, 100, 200, 300, 8, 100, 200, 300}

	b := slice_tricks.FilterEvenNumbers(a.s)

	assert.Equal(t, expected.s, b)
	assert.Equal(t, expected.length, len(b))
	assert.Equal(t, residue, sA)
}

func TestFilterSliceFunc(t *testing.T) {
	type sliceDesc struct {
		s      []int
		length int
	}
	sA := []int{
		10, 20, 21, 23, 25, 30,
	}
	var a = sliceDesc{
		s:      sA,
		length: len(sA),
	}

	var sExp = []int{10, 20, 30}
	var expected = sliceDesc{
		s:      sExp,
		length: len(sExp),
	}

	b := slice_tricks.FilterSliceFunc(a.s, func(v int) bool {
		return v%10 == 0
	})

	assert.Equal(t, expected.s, b)
}

func TestBatchWithoutAllocating(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	exp := [][]interface{}{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9},
	}

	b := slice_tricks.BatchWithoutAllocating(a, 3)
	assert.Equal(t, exp, b)
}

//Red 24, Yellow 25, Red 14, Yellow 12, Blue 16
func TestPerfectStreet(t *testing.T) {
	initial := []sorter.ColoredHouse{
		{sorter.Red, 24},
		{sorter.Yellow, 25},
		{sorter.Red, 14},
		{sorter.Yellow, 12},
		{sorter.Blue, 16},
	}
	expected := []sorter.ColoredHouse{
		{sorter.Red, 14},
		{sorter.Red, 24},
		{sorter.Blue, 16},
		{sorter.Yellow, 12},
		{sorter.Yellow, 25},
	}

	// Need to type cast here to use our custom interface explicitly
	sort.Sort(sorter.PerfectStreet(initial))
	assert.Equal(t, expected, initial)
}

func TestStructsS1Mutex(t *testing.T) {
	var s structs.S1
	var nbRoutines = 1000
	var done = make(chan bool)
	var wg sync.WaitGroup
	wg.Add(nbRoutines)
	for i := 0; i < nbRoutines; i++ {
		go func(i int) {
			wg.Wait()
			done <- assert.NotPanics(t, func() { s.AddElement(i, i) })
		}(i)

		wg.Done()
	}

	doneCount := 0
	for ret := range done {
		doneCount++
		assert.Equal(t, true, ret)
		if doneCount >= nbRoutines {
			close(done)
		}
	}

	assert.Equal(t, nbRoutines, s.Len())
}

// Slower than ranging on a map, but this just shows how to make a goroutine iterator on a map, for the syntax / general idea.
func BenchmarkIterate(b *testing.B) {
	var (
		wg    sync.WaitGroup
		N     = 10
		myMap goroutines.SafeMap
	)

	myMap.M = make(map[string]int, N)
	for i := 0; i < N; i++ {
		myMap.M[strconv.Itoa(i)] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := myMap.Iterator()
		for val := range c {
			tmpk, tmpv = val.K, val.V
		}
	}

	wg.Wait()
}

func BenchmarkRange(b *testing.B) {
	N := 10
	m := make(map[string]int, N)
	for i := 0; i < N; i++ {
		m[strconv.Itoa(i)] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k, v := range m {
			tmpk, tmpv = k, v
		}
	}
}

var tmpk, tmpv interface{}
