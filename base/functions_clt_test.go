package base

import (
	"fmt"
	"strings"
	"testing"
)

// 自定义接口函数
// Func([]data,func) x

// 1.获取对应元素下标
func Index(vs []string, t string) int {
	for i, v := range vs {
		if t == v {
			return i
		}
	}
	return -1
}

// 2.判断是否包含某一个元素
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 3.是否有满足条件 func 的任意一个元素
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 4.判断是否所有元素都满足制定条件 func
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 5.取出满足条件的所有元素
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 6.对所有元素做指定操作
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func TestFuncsClt(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(Map(strs, func(s string) string {
		return strings.ToUpper(s)
	}))
	fmt.Println(Map(strs, strings.ToLower))
	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "e")
	}))
}
