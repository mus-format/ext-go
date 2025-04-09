package ext

// MarshallerMUS interface wraps MarhsalMUS and SizeMUS methods.
type MarshallerMUS interface {
	MarshalMUS(bs []byte) (n int)
	SizeMUS() (size int)
}

// MarshallerMUS interface wraps the MarshalMUS and SizeMUS methods. It is
// intended for use with DTS.
type MarshallerTypedMUS interface {
	MarshalTypedMUS(bs []byte) (n int)
	SizeTypedMUS() (size int)
}
