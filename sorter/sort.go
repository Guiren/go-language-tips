package sorter

// Like the Stringer example, the Sortable interface is very interesting to implement, and very easy.
// Sort requires 3 functions : Less, Len, and

type sortableData struct {
	idContact            uint64
	messageReceivedCount uint
}

// Custom type alias is required since we need to write methods, but any slice type
// "[]something" isn't useable in a method's definition.
type dataSortedByCount []sortableData

// Len should return the total length of the dataSortedByCount.
func (s dataSortedByCount) Len() int {
	return len(s)
}

// Less uses custom logic to point out "what makes s[i] less than s[j]".
// Here we use the number of messages received, but it could be sorted by IdContact, or something else!
func (s dataSortedByCount) Less(i, j int) bool {
	return s[i].messageReceivedCount < s[j].messageReceivedCount
}

// Swap describes how one would swap 2 elements with indexes i and j, in place.
// This example is a regular swap implementation, but it could use custom logic.
func (s dataSortedByCount) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// We're sorting per count, but we could also make another implementation and type alias to sort by idContact! You can try to implement this below :
type dataSortedByIdContact []sortableData

func (d dataSortedByIdContact) Len() int {
	panic("implement me")
}

func (d dataSortedByIdContact) Less(i, j int) bool {
	panic("implement me")
}

func (d dataSortedByIdContact) Swap(i, j int) {
	panic("implement me")
}

// We could have a more complex sorting method, with 2 criterias.
// Here we'll have a street of potential houses, that we'd like to sort by TColor
// and people's age, so a starting slice like :
//
//  Red 24, Yellow 25, Red 14, Yellow 12, Blue 16
//
// would become :
//
//  Red 14, Red 24, Blue 16, Yellow 12, Yellow 25
//
// Red, then blue, then yellow (TColor's order), AND ages going up.
type TColor int

const (
	Red TColor = iota
	Blue
	Green
	Yellow
	Black
)

type ColoredHouse struct {
	Color         TColor
	InhabitantAge int
}

type PerfectStreet []ColoredHouse

func (p PerfectStreet) Len() int {
	return len(p)
}

func (p PerfectStreet) Less(i, j int) bool {
	colorI, ageI := p[i].Color, p[i].InhabitantAge
	colorJ, ageJ := p[j].Color, p[j].InhabitantAge
	if colorI > colorJ {
		return false
	}
	if colorI < colorJ {
		return true
	}

	return ageI < ageJ
}

// Here the swap is simple, but we could still imagine a more complex swap, moving only
// some attributes inside each structure around. Of course, Less and Swap need to be correlated so that the Sort() function actually finishes.
func (p PerfectStreet) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
