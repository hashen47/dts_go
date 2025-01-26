package stack

import (
	"testing"
)

type TestCase[T any] struct {
	isEmpty  bool
	elements []T
}

var intTestcases []TestCase[int]
var strTestcases []TestCase[string]
var floatTestcases []TestCase[float64]
var byteTestcases []TestCase[byte]

func TestMain(m *testing.M) {
	intTestcases = []TestCase[int]{
		{isEmpty: true, elements: []int{}},
		{elements: []int{1}},
		{elements: []int{1, 2}},
		{elements: []int{1, 2, 3}},
		{elements: []int{1, 2, 3, 4}},
		{elements: []int{1, 2, 3, 4, 5}},
		{elements: []int{1, 2, 3, 4, 5, 6}},
		{elements: []int{1, 2, 3, 4, 5, 6, 7}},
		{elements: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	floatTestcases = []TestCase[float64]{
		{isEmpty: true, elements: []float64{}},
		{elements: []float64{1}},
		{elements: []float64{1, 2}},
		{elements: []float64{1, 2, 3}},
		{elements: []float64{1, 2, 3, 4}},
		{elements: []float64{1, 2, 3, 4, 5}},
		{elements: []float64{1, 2, 3, 4, 5, 6}},
		{elements: []float64{1, 2, 3, 4, 5, 6, 7}},
		{elements: []float64{1, 2, 3, 4, 5, 6, 7, 8}},
		{elements: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{elements: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	strTestcases = []TestCase[string]{
		{isEmpty: true, elements: []string{}},
		{elements: []string{"a"}},
		{elements: []string{"a", "b"}},
		{elements: []string{"a", "b", "c"}},
		{elements: []string{"a", "b", "c", "d"}},
		{elements: []string{"a", "b", "c", "d", "e"}},
		{elements: []string{"a", "b", "c", "d", "e", "f"}},
		{elements: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	byteTestcases = []TestCase[byte]{
		{isEmpty: true, elements: []byte{}},
		{elements: []byte{'a'}},
		{elements: []byte{'a', 'b'}},
		{elements: []byte{'a', 'b', 'c'}},
		{elements: []byte{'a', 'b', 'c', 'd'}},
		{elements: []byte{'a', 'b', 'c', 'd', 'e'}},
		{elements: []byte{'a', 'b', 'c', 'd', 'e', 'f'}},
		{elements: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}},
	}

	m.Run()
}

func isSlicesEqual[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func testPushHelper[T comparable](t *testing.T, testcase *TestCase[T], s *Stack[T]) {
	for _, el := range testcase.elements {
		s.Push(el)
	}

	if !isSlicesEqual(testcase.elements, s.Buf) {
		t.Fatalf("expect: %v, real: %v\n", testcase.elements, s.Buf)
	}
}

func TestPush(t *testing.T) {
	for _, testcase := range intTestcases {
		s := StackInit[int]()
		testPushHelper(t, &testcase, s)
	}

	for _, testcase := range floatTestcases {
		s := StackInit[float64]()
		testPushHelper(t, &testcase, s)
	}

	for _, testcase := range strTestcases {
		s := StackInit[string]()
		testPushHelper(t, &testcase, s)
	}

	for _, testcase := range byteTestcases {
		s := StackInit[byte]()
		testPushHelper(t, &testcase, s)
	}
}

func testIsEmptyHelper[T comparable](t *testing.T, testcase *TestCase[T], s *Stack[T]) {
	for _, el := range testcase.elements {
		s.Push(el)
	}

	if (s.IsEmpty() && len(testcase.elements) != 0) ||
		(!s.IsEmpty() && len(testcase.elements) == 0) {
		t.Fatalf("expect: %v, real: %v\n, stack: %v", testcase.isEmpty, s.IsEmpty(), s)
	}
}

func TestIsEmpty(t *testing.T) {
	for _, testcase := range intTestcases {
		s := StackInit[int]()
		testIsEmptyHelper(t, &testcase, s)
	}

	for _, testcase := range floatTestcases {
		s := StackInit[float64]()
		testIsEmptyHelper(t, &testcase, s)
	}

	for _, testcase := range strTestcases {
		s := StackInit[string]()
		testIsEmptyHelper(t, &testcase, s)
	}

	for _, testcase := range byteTestcases {
		s := StackInit[byte]()
		testIsEmptyHelper(t, &testcase, s)
	}
}

func testPopHelper[T comparable](t *testing.T, testcase *TestCase[T], s *Stack[T]) {
	for _, el := range (*testcase).elements {
		s.Push(el)
	}

	if len(s.Buf) == 0 {
		_, err := s.Pop()

		if err == nil {
			t.Fatalf("expect: <error>, real: %v, stack: %v\n", err, s)
		}
	} else {
		for {
			if s.I == -1 {
				break
			}

			val, err := s.Pop()

			if err != nil {
				t.Fatalf("expect: nil, real: %v, stack: %v\n", err, s)
			}

			expectVal := (*testcase).elements[s.I+1]

			if expectVal != val {
				t.Fatalf("expect: %v, real: %v, stack: %v\n", expectVal, val, s)
			}
		}
	}
}

func TestPop(t *testing.T) {
	for _, testcase := range intTestcases {
		s := StackInit[int]()
		testPopHelper(t, &testcase, s)
	}

	for _, testcase := range floatTestcases {
		s := StackInit[float64]()
		testPopHelper(t, &testcase, s)
	}

	for _, testcase := range strTestcases {
		s := StackInit[string]()
		testPopHelper(t, &testcase, s)
	}

	for _, testcase := range byteTestcases {
		s := StackInit[byte]()
		testPopHelper(t, &testcase, s)
	}
}
