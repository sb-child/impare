package instance

import "libpare/core"

// singleton

var corePool *core.CorePool

// init functions

func initCorePool() {
	x := core.NewPool()
	corePool = &x
}

// get functions

func CorePool() *core.CorePool {
	return corePool
}

// init

func init() {
	initCorePool()
}
