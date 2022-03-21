package base

import (
	"fmt"
	"sort"
	"testing"
)

// sort 包
func TestSort(t *testing.T) {
	strs := []string{"c", "a", "v"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
	// sort.Float64s / Slice / Ints / Strings

}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ***自定义排序*** 自定义类型实现接口函数 Interface {
// Len() int
// Less(i, j int) bool
// Swap(i, j int)  }
//  sort.Sort([]sort.Interface)
func TestCustomeSort(t *testing.T) {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
