package main

import (
	"fmt"
	model "go-rest-server/model"
	"net/http"
)

func main() {
	CreateRestServer(8081).Handle(
		model.CreateHTTPHandler(model.POST, "/test1").HandleOps(test1HandlerBefore1, test1HandlerBefore2, test1Handler, test1HandlerAfter1),
		model.CreateHTTPHandler(model.POST, "/test2").HandleOp(test2Handler),
	).Start("")
}

func test1HandlerBefore1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# 1 Before test1")
}

func test1HandlerBefore2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# 2 Before test1")
}

func test1HandlerAfter1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# 1 After test1")
}

func test1Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Just test1"))
}

func test2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Another test2"))
}
