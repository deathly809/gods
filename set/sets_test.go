package set

import "testing"

func isSame(a ArrayCompare, N int) bool {
	for i := 0; i < N; i++ {
		if !a.Same(i) {
			return false
		}
	}
	return true
}

type ArrayCompare interface {
	Same(int) bool
}

type intCompare struct {
	a, b []int
}

type float64Compare struct {
	a, b []float64
}

func (i intCompare) Same(pos int) bool {
	if pos >= len(i.a) {
		return false
	}
	if pos >= len(i.b) {
		return false
	}

	return i.a[pos] == i.b[pos]
}

func (i float64Compare) Same(pos int) bool {
	if pos >= len(i.a) {
		return false
	}
	if pos >= len(i.b) {
		return false
	}

	return i.a[pos] == i.b[pos]
}

func testInt(t *testing.T, A, B, U, I, S []int) {
	iSet := &OrderedIntSet{A, B, nil, 0, 0}

	Union(iSet)
	if !isSame(intCompare{iSet.result, U}, len(U)) {
		t.Fail()
		t.Logf("Union(%v,%v) : Expected: %v Found: %v\n", A, B, U, iSet.result)
	}

	Intersect(iSet)
	if !isSame(intCompare{iSet.result, I}, len(I)) {
		t.Fail()
		t.Logf("Intersect(%v,%v) : Expected: %v Found: %v\n", A, B, I, iSet.result)
	}

	Subtract(iSet)
	if !isSame(intCompare{iSet.result, S}, len(S)) {
		t.Fail()
		t.Logf("Subtract(%v,%v) : Expected: %v Found: %v\n", A, B, S, iSet.result)
	}

}

func testFloat64(t *testing.T, A, B, U, I, S []float64) {
	fSet := &OrderedFloat64Set{A, B, nil, 0, 0}

	Union(fSet)
	if !isSame(float64Compare{fSet.result, U}, len(U)) {
		t.Fail()
		t.Logf("Union(%v,%v) : Expected: %v Found: %v\n", A, B, U, fSet.result)
	}

	Intersect(fSet)
	if !isSame(float64Compare{fSet.result, I}, len(I)) {
		t.Fail()
		t.Logf("Intersect(%v,%v) : Expected: %v Found: %v\n", A, B, I, fSet.result)
	}

	Subtract(fSet)
	if !isSame(float64Compare{fSet.result, S}, len(S)) {
		t.Fail()
		t.Logf("Subtract(%v,%v) : Expected: %v Found: %v\n", A, B, S, fSet.result)
	}

}

func TestIntSets(t *testing.T) {
	first := []int{1, 2, 3, 4, 5, 6}
	second := []int{5, 6, 7, 8, 9}
	third := []int{12, 13, 14, 15}

	firstSecondUnion := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstSecondIntersect := []int{5, 6}
	firstSecondSubtract := []int{1, 2, 3, 4}
	secondFirstSubtract := []int{7, 8, 9}

	firstThirdUnion := []int{1, 2, 3, 4, 5, 6, 12, 13, 14, 15}
	firstThirdIntersection := []int(nil)
	firstThirdSubtract := []int{1, 2, 3, 4, 5, 6}
	thirdFirstSubtract := []int{12, 13, 14, 15}

	secondThirdUnion := []int{5, 6, 7, 8, 9, 12, 13, 14, 15}
	secondThirdIntersection := []int(nil)
	secondThirdSubtract := []int{5, 6, 7, 8, 9}
	thirdSecondSubtract := []int{12, 13, 14, 15}

	testInt(t, first, second, firstSecondUnion, firstSecondIntersect, firstSecondSubtract)
	testInt(t, second, first, firstSecondUnion, firstSecondIntersect, secondFirstSubtract)

	testInt(t, first, third, firstThirdUnion, firstThirdIntersection, firstThirdSubtract)
	testInt(t, third, first, firstThirdUnion, firstThirdIntersection, thirdFirstSubtract)

	testInt(t, second, third, secondThirdUnion, secondThirdIntersection, secondThirdSubtract)
	testInt(t, third, second, secondThirdUnion, secondThirdIntersection, thirdSecondSubtract)
}

func TestFloat64Sets(t *testing.T) {
	first := []float64{1, 2, 3, 4, 5, 6}
	second := []float64{5, 6, 7, 8, 9}
	third := []float64{12, 13, 14, 15}

	firstSecondUnion := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstSecondIntersect := []float64{5, 6}
	firstSecondSubtract := []float64{1, 2, 3, 4}
	secondFirstSubtract := []float64{7, 8, 9}

	firstThirdUnion := []float64{1, 2, 3, 4, 5, 6, 12, 13, 14, 15}
	firstThirdIntersection := []float64(nil)
	firstThirdSubtract := []float64{1, 2, 3, 4, 5, 6}
	thirdFirstSubtract := []float64{12, 13, 14, 15}

	secondThirdUnion := []float64{5, 6, 7, 8, 9, 12, 13, 14, 15}
	secondThirdIntersection := []float64(nil)
	secondThirdSubtract := []float64{5, 6, 7, 8, 9}
	thirdSecondSubtract := []float64{12, 13, 14, 15}

	testFloat64(t, first, second, firstSecondUnion, firstSecondIntersect, firstSecondSubtract)
	testFloat64(t, second, first, firstSecondUnion, firstSecondIntersect, secondFirstSubtract)

	testFloat64(t, first, third, firstThirdUnion, firstThirdIntersection, firstThirdSubtract)
	testFloat64(t, third, first, firstThirdUnion, firstThirdIntersection, thirdFirstSubtract)

	testFloat64(t, second, third, secondThirdUnion, secondThirdIntersection, secondThirdSubtract)
	testFloat64(t, third, second, secondThirdUnion, secondThirdIntersection, thirdSecondSubtract)
}
