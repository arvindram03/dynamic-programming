package dynamicprogramming

import "testing"

func TestFindMaxSeqReturnsOneForOneElement(t *testing.T) {
	seq := []int{0}
	res := FindMaxSeq(seq)

	AssertEquals(1, res, t)
}

func TestFindMaxSeqReturnsOneForDecreasingSeq(t *testing.T) {
	seq := []int{3, 2, 1}
	res := FindMaxSeq(seq)

	AssertEquals(1, res, t)
}

func TestFindMaxSeqReturnsLengthOfTheSeqForIncreasingSeq(t *testing.T) {
	seq := []int{0, 1, 2}
	res := FindMaxSeq(seq)

	AssertEquals(3, res, t)
}

func TestFindMaxSeqReturnsLongestSeq(t *testing.T) {
	seq := []int{0, 1, 2, 3, 1, 2, 3, 4, 5, 0}
	res := FindMaxSeq(seq)

	AssertEquals(5, res, t)
}

func AssertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Errorf("Mismatch. Expected: %+v, Actual: %+v", expected, actual)
		t.FailNow()
	}
}
