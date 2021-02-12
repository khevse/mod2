package mod2

import "github.com/khevse/mod3/v2"

type Impl struct {
}

func New() *Impl {
	return &Impl{}
}

func (i Impl) String() string {
	return mod3.New().String()
}
