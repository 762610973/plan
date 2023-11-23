package main

import (
	"log/slog"
	"net/http"
)

func main() {
	handleFunc()
}

func handleFunc() {
	// 没有精细匹配到的路由, 默认会命中到这里, 具体如何处理, 要看自己的逻辑
	// /xl,/get,/post等都会命中到这里
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(r.URL.String()))
	})
	// http.handleFunc是对单一路由进行定义处理方法
	// 只有/handleFunc1会命中到这里
	http.HandleFunc("/handleFunc1", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("this is handleFunc1"))
	})
	// 只有/handleFunc2会命中到这里进行处理
	http.HandleFunc("/handleFunc2", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("this is handleFunc2"))
	})
	// 此处需要填写nil
	err := http.ListenAndServe(":1112", nil)
	if err != nil {
		slog.Error("listen failed", err)
	}
}

type handler struct {
	addr string
}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(r.URL.String()))
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("response success", r.Host, r.Method, r.URL.String())
}
func handle() {
	h := handler{
		addr: ":1112",
	}
	// http.Handle 在声明时无法对路由进行清晰的展示, 只能在handler.ServeHTTP中进行处理
	http.Handle("/", h)
	// 此处需要填写h, 只会调用h.ServeHTTP
	err := http.ListenAndServe(h.addr, h)
	if err != nil {
		slog.Error("listen failed", err)
	}
	server := http.Server{}
	_ = server.Close()
}
