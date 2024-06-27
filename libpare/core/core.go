package core

// import "github.com/klauspost/reedsolomon"

type Core struct {
	opt Option
}

func NewCore(opt Option) Core {
	return Core{
		opt: opt,
	}
}

func (x Core) GetOption() Option {
	return x.opt
}

func (x Core) Meow() {

}

func Meow() {
	// enc, err := reedsolomon.New(*dataShards, *parShards)
}
