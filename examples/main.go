package main

import (
	"fmt"
	rest "go-rest-server"
	model "go-rest-server/model"
)

type SimpleJSON struct {
	Message string `json:"message"`
}

func main() {
	rest.CreateRestServer(8081).SetHandlers(
		model.MakeRequestHandler(model.POST, "/test1").SetIncome(test1HandlerBefore1, test1HandlerBefore2).SetHandler(test1Handler).SetOutcome(test1HandlerAfter1),
		model.MakeRequestHandler(model.GET, "/test2").SetHandler(test2Handler),
	).Start()
}

func test1HandlerBefore1(context *model.HTTPContext) {
	fmt.Println("# 1 Before test1")
}

func test1HandlerBefore2(context *model.HTTPContext) {
	fmt.Println("# 2 Before test1")
}

func test1HandlerAfter1(context *model.HTTPContext) {
	fmt.Println("# 1 After test1")
}

func test1Handler(context *model.HTTPContext) *model.ResponseEntity {
	return &model.ResponseEntity{Body: SimpleJSON{"And you!"}, Headers: nil}
}

func test2Handler(context *model.HTTPContext) *model.ResponseEntity {
	return &model.ResponseEntity{Body: SimpleJSON{"And you!"}, Headers: nil}
}
