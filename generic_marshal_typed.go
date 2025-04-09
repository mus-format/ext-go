package ext

// MarshalTypedMUS creates and returns a byte slice filled with the serialized
// data.
func MarshalTypedMUS(v MarshallerTypedMUS) (bs []byte) {
	bs = make([]byte, v.SizeTypedMUS())
	v.MarshalTypedMUS(bs)
	return
}
