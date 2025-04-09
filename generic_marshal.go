package ext

// MarshalMUS creates and returns a byte slice filled with the serialized data.
func MarshalMUS(v MarshallerMUS) (bs []byte) {
	bs = make([]byte, v.SizeMUS())
	v.MarshalMUS(bs)
	return
}
