package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileread() {
	dat, err := os.ReadFile("./tmp/dat")
	check(err)
	fmt.Print(string(dat))
	// 您通常会希望对文件的读取方式和内容进行更多控制。 对于这个任务，首先使用 Open 打开一个文件，以获取一个 os.File 值。

	f, err := os.Open("./tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 你也可以 Seek 到一个文件中已知的位置，并从这个位置开始读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s \n", string(b4))
	f.Close()
}

func TestFile(t *testing.T) {
	fileread()
}

func fileWrite() {

	f, err := os.Create("./tmp/write_test")
	check(err)

	defer f.Close()
	// 开始写入文本

	f.WriteString("直接写入字符串WriteString\n")

	//
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	f.Sync() // 将缓冲区的数据写入硬盘

	// 用缓冲bufbytes 写入文件内容
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("用字节缓冲器写入文本数据 Byte Write Buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes \n", n4)
	w.Flush() // 刷盘

}
func TestFilewrite(t *testing.T) {
	fileWrite()
}

// 行读取-遍历转换大小写
func lineFilter() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Printf("%s \n", ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error :", err)
		os.Exit(1)
	}
}
func TestFileLineFilter(t *testing.T) {
	lineFilter()
}

// 文件路径-
func filepathT() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	// 您应该总是使用 Join 代替手动拼接 / 和 \。 除了可移植性，Join 还会删除多余的分隔符和目录，使得路径更加规范。

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	// Dir 和 Base 可以被用于分割路径中的目录和文件。 此外，Split 可以一次调用返回上面两个函数的结果。

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	// 判断路径是否为绝对路径。

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	filename := "config.json"
	// 某些文件名包含了扩展名（文件类型）。 我们可以用 Ext 将扩展名分割出来。

	ext := filepath.Ext(filename)
	fmt.Println(ext)
	// 想获取文件名清除扩展名后的值，请使用 strings.TrmSuffix。

	fmt.Println(strings.TrimSuffix(filename, ext))
	// Rel 寻找 basepath 与 targpath 之间的相对路径。 如果相对路径不存在，则返回错误。

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

func TestFilepath(t *testing.T) {
	filepathT()
}

// 文件目录操作
func dirT() {
	_, fullFilename, _, _ := runtime.Caller(0)
	fmt.Println(fullFilename)

	full, _ := path.Split(fullFilename)
	err := os.Mkdir("subdir", 0755)
	check(err)

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)
	fmt.Printf("Listing Path: %s", "subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	fmt.Printf("Visiting Path And Run Func: %s\n", "subdir")
	err = filepath.Walk(filepath.Join(full, "subdir"), visit)
	check(err)
	defer os.RemoveAll(filepath.Join(full, "subdir"))
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println("  ", p, info.IsDir())
	return nil
}

func TestFileFir(t *testing.T) {
	dirT()
}

// 临时文件-生成会自定添加后缀，避免并发冲突
