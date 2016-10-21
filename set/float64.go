package set

// OrderedFloat64Set allows for operations on sets of integers
type OrderedFloat64Set struct {
	first, second, result []float64
	posFirst, posSecond   int
}

func NewOrderedFloat64Set(a, b []float64) OrderedFloat64Set {
	return OrderedFloat64Set{first: a, second: b}
}

// Less returns the relationship first[0] < second[0] if first is true,
// otherwise return the relationship second[0] < first[0]
func (m *OrderedFloat64Set) Less(sel MergeSelector) bool {
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
func (m *OrderedFloat64Set) Append(sel MergeSelector) {
	switch sel {
	case First:
		m.result = append(m.result, m.first[m.posFirst])
	case Second:
		m.result = append(m.result, m.second[m.posSecond])
	}
}

// Len return length of first array if true, otherwise returns
// length of second array
func (m *OrderedFloat64Set) Len(sel MergeSelector) int {
	switch sel {
	case First:
		return len(m.first) - m.posFirst
	case Second:
		return len(m.second) - m.posSecond
	}
	panic("invalid selector")
}

// Remove the first element of the selected array
func (m *OrderedFloat64Set) Remove(sel MergeSelector) {
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
func (m *OrderedFloat64Set) Reset() {
	m.result = nil
	m.posFirst = 0
	m.posSecond = 0
}
