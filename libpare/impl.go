package main

import (
	// "fmt"
	"libpare/core"
	// "time"
)

type Impl struct{}

// create_task implements SysApi.
func (i Impl) create_task(req CreateTaskRequest) CreateTaskResponse {
	return CreateTaskResponse{id: 1}
	// panic("unimplemented")
}

// remove_task implements SysApi.
func (i Impl) remove_task(req RemoveTaskRequest) RemoveTaskResponse {
	return RemoveTaskResponse{success: true}
	// panic("unimplemented")
}

func init() {
	core.Meow()
	SysApiImpl = Impl{}
}
