package base

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

// 测试 正则表达式 regexp

func regeexp() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	r, _ := regexp.Compile("p([a-z]+)ch")
	// 是否符合正则匹配
	fmt.Println(r.MatchString("peach"))

	// 获取首次匹配字符下标
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	// 查找首次匹配的字符串。
	fmt.Println(r.FindString("peach punch"))
	// 获取首次匹配的字符串 和匹配到的部分差异字符
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// 完全匹配 和 匹配字符的局部位置
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// 带 All 的函数返回所有匹配项目
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// 使用字节数组 []byte
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile 模式验证字符串, 失败会返回 panic 而不是 err
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// 替换数据
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// 对匹配项目执行自定义操作
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}

func TestRegexp(t *testing.T) {
	regeexp()
}
