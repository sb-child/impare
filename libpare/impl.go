package main

import (
	// "fmt"
	"fmt"
	"libpare/core"
	"libpare/instance"
	// "time"
)

type Impl struct{}

// create_task implements SysApi.
func (i Impl) create_task(req CreateTaskRequest) CreateTaskResponse {
	c := instance.CorePool().Get(core.Option{
		DataShards: uint32(req.data_shards),
		ParShards:  uint32(req.parity_shards),
	})
	fmt.Printf("opt: %d, %d\n", c.GetOption().DataShards, c.GetOption().ParShards)
	instance.CorePool().Put(c)
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
