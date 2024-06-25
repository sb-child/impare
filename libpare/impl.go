package main

import (
	// "fmt"
	"libpare/core"
	// "time"
)

type Demo struct{}

// demo_check implements DemoCall.
func (d Demo) demo_check(req DemoRequest) DemoResponse {
	panic("unimplemented")
}

// demo_check_async implements DemoCall.
func (d Demo) demo_check_async(req DemoRequest) DemoResponse {
	panic("unimplemented")
}

func init() {
	core.Meow()
	DemoCallImpl = Demo{}
}

// func (Demo) demo_oneway(req DemoUser) {
// 	fmt.Printf("[Go-oneway] Golang received name: %s, age: %d\n", req.name, req.age)
// }

// func (Demo) demo_check(req DemoComplicatedRequest) DemoResponse {
// 	fmt.Printf("[Go-call] Golang received req\n")
// 	fmt.Printf("[Go-call] Golang returned result\n")
// 	return DemoResponse{pass: true}
// }

// func (Demo) demo_check_async(req DemoComplicatedRequest) DemoResponse {
// 	fmt.Printf("[Go-call async] Golang received req, will sleep 1s\n")
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("[Go-call async] Golang returned result\n")
// 	return DemoResponse{pass: true}
// }

// func (Demo) demo_check_async_safe(req DemoComplicatedRequest) DemoResponse {
// 	fmt.Printf("[Go-call async drop_safe] Golang received req, will sleep 1s\n")
// 	time.Sleep(1 * time.Second)
// 	resp := DemoResponse{pass: req.balabala[0] == 1}
// 	fmt.Printf("[Go-call async drop_safe] Golang returned result, pass: %v\n", req.balabala[0] == 1)
// 	return resp
// }
