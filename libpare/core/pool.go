package core

import "sync"

type CorePool struct {
	optTable map[Option]*sync.Pool
	lock     *sync.RWMutex
}

func createPool(opt Option) *sync.Pool {
	x := sync.Pool{
		New: func() any {
			x := NewCore(opt)
			return &x
		},
	}
	return &x
}

func getFromPool(p *sync.Pool) *Core {
	return p.Get().(*Core)
}

func NewPool() CorePool {
	lock := sync.RWMutex{}
	return CorePool{lock: &lock}
}

func (x *CorePool) Get(opt Option) *Core {
	x.initTable(opt)
	return getFromPool(x.take(opt))
}

func (x *CorePool) Put(c *Core) {
	opt := c.GetOption()
	x.initTable(opt)
	x.take(opt).Put(c)
}

func (x *CorePool) Release() {
	x.lock.Lock()
	if x.optTable != nil {
		for k := range x.optTable {
			delete(x.optTable, k)
		}
	}
	x.lock.Unlock()
}

func (x *CorePool) initTable(opt Option) {
	x.lock.Lock()
	if x.optTable == nil {
		x.optTable = make(map[Option]*sync.Pool)
	}
	if _, ok := x.optTable[opt]; !ok {
		x.optTable[opt] = createPool(opt)
	}
	x.lock.Unlock()
}

// and please ensure this opt exists in the table
func (x *CorePool) take(opt Option) *sync.Pool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.optTable[opt]
}
