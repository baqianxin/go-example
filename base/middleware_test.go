package base

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// 记录日志 中间件
func Loggin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

// 请求拦截 中间件
func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				// 非指定方式都直接返回400
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}

// 应用中间件链路
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func Test_middleware(t *testing.T) {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Loggin()))
	http.ListenAndServe(":8080", nil)
}
