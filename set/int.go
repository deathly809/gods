package set

// OrderedIntSet allows for operations on sets of integers
type OrderedIntSet struct {
	first, second, result []int
	posFirst, posSecond   int
}

func NewOrderedIntSet(a, b []int) OrderedIntSet {
	return OrderedIntSet{first: a, second: b}
}

// Less returns the relationship first[0] < second[0] if first is true,
// otherwise return the relationship second[0] < first[0]
func (m *OrderedIntSet) Less(sel MergeSelector) bool {
	switch sel {
	case First:
		return m.first[m.posFirst] < m.second[m.posSecond]
	case Second:
		return m.first[m.posFirst] > m.second[m.posSecond]
	}
	panic("invalid selector")
}

// Append appends the and removes first element of the first array if true,
//	otherwise the first element of the second array.
func (m *OrderedIntSet) Append(sel MergeSelector) {
	switch sel {
	case First:
		m.result = append(m.result, m.first[m.posFirst])
	case Second:
		m.result = append(m.result, m.second[m.posSecond])
	}
}

// Len return length of first array if true, otherwise returns
// length of second array
func (m *OrderedIntSet) Len(sel MergeSelector) int {
	switch sel {
	case First:
		return len(m.first) - m.posFirst
	case Second:
		return len(m.second) - m.posSecond
	}
	panic("invalid selector")
}

// Remove the first element of the selected array
func (m *OrderedIntSet) Remove(sel MergeSelector) {
	switch sel {
	case First:
		m.posFirst++
	case Second:
		m.posSecond++
	default:
		panic("unknown selector")

	}
}

// Clear the result
func (m *OrderedIntSet) Reset() {
	m.result = nil
	m.posFirst = 0
	m.posSecond = 0
}
