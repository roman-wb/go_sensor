package generator

import (
	"fmt"
)

type Stack struct {
	items []*Element
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.items)
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(item *Element) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() *Element {
	if !s.Empty() {
		item := s.items[s.Len()-1]
		s.items = s.items[:s.Len()-1]
		return item
	}
	return nil
}

func (s *Stack) Peek() *Element {
	if !s.Empty() {
		return s.items[s.Len()-1]
	}
	return nil
}
