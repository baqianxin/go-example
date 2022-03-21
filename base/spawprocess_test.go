package base

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

// 测试 命令行使用： 生成外部进程
func TestSpawProcess(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep cmddddd\ngoodbye grep"))
	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -l -a -h")
	lsOut, _ := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -alh")
	fmt.Println(string(lsOut))

}

// 测试执行外部进程：环境变量，系统程序，运行参数

// 这里是真正的 syscall.Exec 调用。
// 如果这个调用成功，那么我们的进程将在这里结束
// ，并被 /bin/ls -a -l -h 进程代替。
//  如果存在错误，那么我们将会得到一个返回值。
func TestRunSpawproces(t *testing.T) {
	binary, lookErr := exec.LookPath("ls")

	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// 注意 Go 没有提供 Unix 经典的 fork 函数。
// 一般来说，这没有问题，因为启动协程、生成进程和执行进程，
// 已经涵盖了 fork 的大多数使用场景。
