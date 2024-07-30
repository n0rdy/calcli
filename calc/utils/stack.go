package utils

const (
	errEmptyStack = "stack is empty"
)

type Stack struct {
	elements []float64
}

func (s *Stack) Push(e float64) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() float64 {
	n := len(s.elements)
	if n == 0 {
		panic(errEmptyStack)
	}

	e := s.elements[n-1]
	s.elements = s.elements[:n-1]
	return e
}
