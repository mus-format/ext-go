package ext

// MarshallerMUS interface wraps MarhsalMUS and SizeMUS methods.
type MarshallerMUS interface {
	MarshalMUS(bs []byte) (n int)
	SizeMUS() (size int)
}
