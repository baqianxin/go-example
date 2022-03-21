package base

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// 系统信号与协程通道-chan & signal

func TestSignal(t *testing.T) {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	select {
	case <-done:
		fmt.Println("exiting")
	case <-time.After(10 * time.Second):
		fmt.Println("timeout!")
	}

}

// os.Exit(1)
func TestExit(t *testing.T) {
	defer fmt.Println("exit!")
	os.Exit(3) // os.Exit 之后不会掉用 defer
}
