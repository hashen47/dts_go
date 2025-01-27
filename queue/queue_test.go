package queue

import (
	"testing"
)

type Testcase[T comparable] struct {
	isEmpty  bool
	elements []T
}

var intTestcases *[]Testcase[int]
var floatTestcases *[]Testcase[float64]
var strTestcases *[]Testcase[string]
var byteTestcases *[]Testcase[byte]

func isSlicesEqual[T comparable](s1 *[]T, s2 *[]T) bool {
	if len(*s1) != len(*s2) {
		return false
	}
	for i := 0; i < len(*s1); i++ {
		if (*s1)[i] != (*s2)[i] {
			return false
		}
	}
	return true
}

func TestMain(m *testing.M) {
	intTestcases = &[]Testcase[int]{
		{isEmpty: true, elements: []int{}},
		{isEmpty: false, elements: []int{1}},
		{isEmpty: false, elements: []int{1, 2}},
		{isEmpty: false, elements: []int{1, 2, 3}},
		{isEmpty: false, elements: []int{1, 2, 3, 4}},
		{isEmpty: false, elements: []int{1, 2, 3, 4, 5}},
		{isEmpty: false, elements: []int{1, 2, 3, 4, 5, 6}},
		{isEmpty: false, elements: []int{1, 2, 3, 4, 5, 6, 7}},
		{isEmpty: false, elements: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{isEmpty: false, elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{isEmpty: false, elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	floatTestcases = &[]Testcase[float64]{
		{isEmpty: true, elements: []float64{}},
		{isEmpty: false, elements: []float64{1}},
		{isEmpty: false, elements: []float64{1, 2}},
		{isEmpty: false, elements: []float64{1, 2, 3}},
		{isEmpty: false, elements: []float64{1, 2, 3, 4}},
		{isEmpty: false, elements: []float64{1, 2, 3, 4, 5}},
		{isEmpty: false, elements: []float64{1, 2, 3, 4, 5, 6}},
		{isEmpty: false, elements: []float64{1, 2, 3, 4, 5, 6, 7}},
		{isEmpty: false, elements: []float64{1, 2, 3, 4, 5, 6, 7, 8}},
		{isEmpty: false, elements: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{isEmpty: false, elements: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	strTestcases = &[]Testcase[string]{
		{isEmpty: true, elements: []string{}},
		{isEmpty: false, elements: []string{"a"}},
		{isEmpty: false, elements: []string{"a", "b"}},
		{isEmpty: false, elements: []string{"a", "b", "c"}},
		{isEmpty: false, elements: []string{"a", "b", "c", "d"}},
		{isEmpty: false, elements: []string{"a", "b", "c", "d", "e"}},
		{isEmpty: false, elements: []string{"a", "b", "c", "d", "e", "f"}},
		{isEmpty: false, elements: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{isEmpty: false, elements: []string{"a", "b", "c", "d", "e", "f", "g", "h"}},
		{isEmpty: false, elements: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}},
		{isEmpty: false, elements: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}},
	}

	byteTestcases = &[]Testcase[byte]{
		{isEmpty: true, elements: []byte{}},
		{isEmpty: false, elements: []byte{'a'}},
		{isEmpty: false, elements: []byte{'a', 'b'}},
		{isEmpty: false, elements: []byte{'a', 'b', 'c'}},
		{isEmpty: false, elements: []byte{'a', 'b', 'c', 'd'}},
		{isEmpty: false, elements: []byte{'a', 'b', 'c', 'd', 'e'}},
		{isEmpty: false, elements: []byte{'a', 'b', 'c', 'd', 'e', 'f'}},
		{isEmpty: false, elements: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}},
		{isEmpty: false, elements: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}},
		{isEmpty: false, elements: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}},
		{isEmpty: false, elements: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}},
	}

	m.Run()
}

func testEnqueueHelper[T comparable](t *testing.T, testcases *[]Testcase[T]) {
	for _, testcase := range *testcases {
		q := QueueInit[T]()

		for _, element := range testcase.elements {
			q.Enqueue(element)
		}

		if len(*q.buf) != len(testcase.elements) {
			t.Fatalf("expect: %d, real: %d, elements: %v, queue: %v\n", len(testcase.elements), len(*q.buf), testcase.elements, q)
		}

		if !isSlicesEqual[T](&testcase.elements, q.buf) {
			t.Fatalf("expect: %d, real: %d, elements: %v, queue: %v\n", len(testcase.elements), len(*q.buf), testcase.elements, q)
		}
	}
}

func TestEnqueue(t *testing.T) {
	testEnqueueHelper[int](t, intTestcases)
	testEnqueueHelper[float64](t, floatTestcases)
	testEnqueueHelper[string](t, strTestcases)
	testEnqueueHelper[byte](t, byteTestcases)
}

func testPeekHelper[T comparable](t *testing.T, testcases *[]Testcase[T]) {
	for _, testcase := range *testcases {
		q := QueueInit[T]()

		for _, element := range testcase.elements {
			q.Enqueue(element)
		}

		if testcase.isEmpty {
			_, err := q.Peek()

			if err != QEmptyErr {
				t.Fatalf("expect: %v, real: %v, elements: %v, queue: %v\n", QEmptyErr, err, testcase.elements, q)
			}
		} else {
			val, err := q.Peek()

			if err != nil {
				t.Fatalf("expect: nil, real: %v, elements: %v, queue: %v\n", err, testcase.elements, q)
			}

			if val != (*q.buf)[q.i] {
				t.Fatalf("expect: %v, real: %v, elements: %v, queue: %v\n", val, (*q.buf)[q.i], testcase.elements, q)
			}
		}
	}
}

func TestPeek(t *testing.T) {
	testPeekHelper[int](t, intTestcases)
	testPeekHelper[float64](t, floatTestcases)
	testPeekHelper[string](t, strTestcases)
	testPeekHelper[byte](t, byteTestcases)
}

func testDequeueHelper[T comparable](t *testing.T, testcases *[]Testcase[T]) {
	for _, testcase := range *testcases {
		q := QueueInit[T]()

		for _, element := range testcase.elements {
			q.Enqueue(element)
		}

		if !testcase.isEmpty {
			for q.i >= 0 && q.i < len(*q.buf) {
				val, err := q.Dequeue()

				if err != nil {
					t.Fatalf("expect: nil, real: %v, elements: %v, queue: %v\n", err, testcase.elements, q)
				}

				if val != (*q.buf)[q.i-1] {
					t.Fatalf("expect: %v, real: %v, elements: %v, queue: %v\n", val, (*q.buf)[q.i-1], testcase.elements, q)
				}
			}

			_, err := q.Dequeue()

			if err != QEmptyErr {
				t.Fatalf("expect: %v, real: %v, elements: %v, queue: %v\n", QEmptyErr, err, testcase.elements, q)
			}
		}

		_, err := q.Dequeue()

		if err != QEmptyErr {
			t.Fatalf("expect: %v, real: %v, elements: %v, queue: %v\n", QEmptyErr, err, testcase.elements, q)
		}
	}
}

func TestDequeue(t *testing.T) {
	testDequeueHelper[int](t, intTestcases)
	testDequeueHelper[float64](t, floatTestcases)
	testDequeueHelper[string](t, strTestcases)
	testDequeueHelper[byte](t, byteTestcases)
}
