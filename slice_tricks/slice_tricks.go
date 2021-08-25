package slice_tricks

// Some cool stuff from https://github.com/golang/go/wiki/SliceTricks
// Those are the most useful in my opinion. Opinions are great if you want to make a burger.

// **** AppendVector ****
func AppendSliceToSliceBad(a, b []interface{}) (result []interface{}) {
	// Setting capacity instead of len() to be able to use append everywhere
	result = make([]interface{}, 0, len(a)+len(b))

	for _, val := range a {
		result = append(result, val)
	}
	for _, val := range b {
		result = append(result, val)
	}

	return
}

func AppendSliceToSliceGood(a, b []interface{}) (result []interface{}) {
	// Append returns a pointer either to slice A's memory space, or a new one if the cap() wasn't big enough.
	result = append(a, b...)

	return
}

// **** Filtering without allocating ****
/* NOTE : Filtering this way uses the original slice's memory space, thus rendering that slice "useless". Example :
Let's say we have this slice :
	Slice A 		: [0, 1, 2, 3, 4, 5] (len 6, cap 6)
We want a slice that gets every even number and returns a slice of them. So we want :
	Slice Result 	: [0, 2, 4] (len 3, cap N) - we don't care about the capacity.

If we don't care about Slice A after this filter, we can actually re-use the same memory space and trick Go, using different len() via a new slice pointer.
The end result of this filtering would be 2 slices : a slice Result, with what we want, and the original slice A, modified/overwritten to accomodate the filter, making it unusable. Same example, result after filtering :
	Slice A : [0, 2, 4, 3, 4, 5] (len 6) // Notice how slice Result is at the start, but what comes after is unchanged.
	Slice B : [0, 2, 4] (len 3)
*/
func FilterEvenNumbers(a []int) (result []int) {
	// result and A point to the same memory, but result has a len of 0 (and same capacity as A)
	result = a[:0]
	for _, value := range a {
		// If we want to keep this value, add it to result
		if value%2 == 0 {
			result = append(result, value)
		}
	}

	return
}

// We could filter with a better function signature, asking for filter logic as well :
func FilterSliceFunc(a []int, filt func(int) bool) (result []int) {
	result = a[:0]
	for _, value := range a {
		if filt(value) {
			result = append(result, value)
		}
	}

	return
}

// BatchWithoutAllocating is very useful in manual SQL queries when using a lot of placeholders.
// For a query like `SELECT * FROM table WHERE id IN (?,?,?,...,?)`, there cannot be more than 65535 `?` in default
// MySQL configurations. Plus, performances are better if the number of placeholders stays low.
//
// This prepares batches so looping/querying over them is trivial, and with almost no memory overhead.
func BatchWithoutAllocating(a []interface{}, batchSize int) (b [][]interface{}) {
	b = make([][]interface{}, 0, (len(a)+batchSize-1)/batchSize)

	for batchSize < len(a) {
		a, b = a[batchSize:], append(b, a[0:batchSize:batchSize])
	}
	b = append(b, a)

	return
}
