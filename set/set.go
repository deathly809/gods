package set

// MergeSelector determines which array to use
type MergeSelector int

const (
	// First array is used
	First = MergeSelector(iota)
	// Second array is used
	Second = MergeSelector(iota)
)

// OrderedSets specifies the operations that allow for merging two sets together
type OrderedSets interface {
	Reset()
	Less(MergeSelector) bool
	Len(MergeSelector) int
	Append(MergeSelector)
	Remove(MergeSelector)
}

// Intersect performs intersection of two sets
func Intersect(m OrderedSets) {
	m.Reset()
	for m.Len(First) > 0 && m.Len(Second) > 0 {
		if m.Less(First) {
			m.Remove(First)
		} else if m.Less(Second) {
			m.Remove(Second)
		} else { // same
			m.Append(Second)
			m.Remove(First)
			m.Remove(Second)
		}
	}
}

// Union unions two sets together
func Union(m OrderedSets) {
	m.Reset()
	for m.Len(First) > 0 && m.Len(Second) > 0 {
		if m.Less(First) {
			m.Append(First)
			m.Remove(First)
		} else if m.Less(Second) {
			m.Append(Second)
			m.Remove(Second)
		} else {
			m.Append(Second)
			m.Remove(First)
			m.Remove(Second)
		}
	}

	for m.Len(First) > 0 {
		m.Append(First)
		m.Remove(First)
	}

	for m.Len(Second) > 0 {
		m.Append(Second)
		m.Remove(Second)
	}
}

// Subtract creates a new set containing everything in the first but not the second set
func Subtract(m OrderedSets) {
	m.Reset()
	for m.Len(First) > 0 && m.Len(Second) > 0 {
		if m.Less(First) {
			m.Append(First)
			m.Remove(First)
		} else if m.Less(Second) {
			m.Remove(Second)
		} else {
			m.Remove(First)
			m.Remove(Second)
		}
	}

	for m.Len(First) > 0 {
		m.Append(First)
		m.Remove(First)
	}
}
