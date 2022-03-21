package base

import (
	"fmt"
	"os"
	"testing"
)

// 格式化输出

type point struct {
	x, y int
}

func printf() {
	p := point{1, 2}
	// 结构 %v 实例的值
	fmt.Printf("struct1:%v\n", p)
	// 结构 %+v 实例值+结构字段
	fmt.Printf("struct2:%+v\n", p)
	// 结构 %#v 根据语法输出值，产生该值的源码片段
	fmt.Printf("struct3:%#v\n", p)
	// 打印类型 %T
	fmt.Printf("type:%T\n", p)
	// 格式化bool值 %t
	fmt.Printf("bool:%t\n", true)
	// 十进制整型 %d
	fmt.Printf("int: %d\n", 112)
	// 二进制表示 %b
	fmt.Printf("bin: %b\n", 112)
	// 字符转换
	fmt.Printf("char: %c\n", 33)
	// 十六进制编码
	fmt.Printf("hex: %x\n", 456)
	// 十进制浮点类型
	fmt.Printf("float: %f\n", 45.6)
	// 十进制浮点科学计数法
	fmt.Printf("float-e: %e\n", 11120000.0)
	fmt.Printf("float-E: %E\n", 11120000.0)
	// base-16字符编码
	fmt.Printf("str1: %s\n", "\"string\"")
	fmt.Printf("str2: %q\n", "\"string\"")
	fmt.Printf("str3: %x\n", "hex this")
	// 指针类型
	fmt.Printf("point: %p\n", &p)
	// 格式化长度
	fmt.Printf("width1: |%8d|%8d|\n", 12, 345)
	fmt.Printf("width1: |%8.2f|%8.2f|\n", 12.1222211, 345.4555)
	// 格式化长度 左对齐
	fmt.Printf("width1: |%-8.2f|%-8.2f|\n", 12.1222211, 345.4555)
	// 格式化长度也可用于字符串
	// 格式化字符串返回值
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	// 指定输出目标
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}

func TestPrintF(t *testing.T) {
	printf()
}
