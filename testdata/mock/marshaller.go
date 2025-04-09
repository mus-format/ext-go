package mock

import (
	"github.com/ymz-ncnk/mok"
)

type MarshalMUSFn func(bs []byte) (n int)
type SizeMUSFn func() (size int)

func NewMarshallerMUS() MarshallerMUS {
	return MarshallerMUS{mok.New("MarshallerMUS")}
}

type MarshallerMUS struct {
	*mok.Mock
}

func (m MarshallerMUS) RegisterMarshalMUS(fn MarshalMUSFn) MarshallerMUS {
	m.Register("MarshalMUS", fn)
	return m
}

func (m MarshallerMUS) RegisterSizeMUS(fn SizeMUSFn) MarshallerMUS {
	m.Register("SizeeMUS", fn)
	return m
}

func (m MarshallerMUS) MarshalMUS(bs []byte) (n int, err error) {
	result, err := m.Call("MarshalMUS", bs)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m MarshallerMUS) SizeMUS() (size int) {
	result, err := m.Call("SizeMUS")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
