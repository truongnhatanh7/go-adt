package stack

type Stack struct {
	stack []any
}

type StackActions interface {
	Push(value any)
	Size() int
  IsEmpty() bool
	Pop() any
	Seek() any
}

func (s *Stack) Push(value any) {
	s.stack = append(s.stack, value)
}

func (s *Stack) IsEmpty() bool {
  return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) Pop() any {
	if s.Size() > 0 {
		result := (s.stack)[s.Size()-1]
		s.stack = (s.stack)[:s.Size()-1]
		return result
	}
	return nil

}

func (s *Stack) Seek() any {
  return (s.stack)[s.Size()-1]
}
