package base

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// var addr = flag.String("server addr", ":8090", "server address")
// var server http.Server

func Test_Context(t *testing.T) {
	// handler := http.HandlerFunc(hello)
	// server = http.Server{
	// 	Addr:    *addr,
	// 	Handler: handler,
	// }
	// server.ListenAndServe()
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context() //, cancel := context.WithTimeout(req.Context(), 3*time.Second)
	// ctx.Done()
	fmt.Println("server: hello handler started.")
	defer fmt.Println("server: bye handler ended.")
	// defer cancel()
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server: ", err)
		internalErr := http.StatusInternalServerError
		http.Error(w, err.Error(), internalErr)

		// server.Shutdown(ctx)
	}
}
