package core

import "sync"

type CorePool struct {
	optTable map[Option]*sync.Pool
}

func createPool(opt Option) *sync.Pool {
	return &sync.Pool{
		New: func() any {
			x := NewCore(opt)
			return &x
		},
	}
}

func getFromPool(p *sync.Pool) *Core {
	return p.Get().(*Core)
}

func NewPool() CorePool {
	return CorePool{}
}

func (x CorePool) Get(opt Option) *Core {
	x.initTable(opt)
	return getFromPool(x.take(opt))
}

func (x CorePool) Put(c *Core) {
	opt := c.GetOption()
	x.initTable(opt)
	x.take(opt).Put(c)
}

func (x CorePool) initTable(opt Option) {
	if x.optTable == nil {
		x.optTable = make(map[Option]*sync.Pool)
	}
	if _, ok := x.optTable[opt]; !ok {
		x.optTable[opt] = createPool(opt)
	}
}

// and please ensure this opt exists in the table
func (x CorePool) take(opt Option) *sync.Pool {
	return x.optTable[opt]
}
