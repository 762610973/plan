package slice

import (
	"fmt"
	"testing"
)

func Test_slice1(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[9:]
	s[9] = 100
	s = nil
	fmt.Println(s1, len(s1), cap(s1))
}
