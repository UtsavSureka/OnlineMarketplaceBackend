package swag

import "unsafe"

func StringData(s string) *byte {
	// Convert the string to a pointer to its underlying data
	return (*byte)(unsafe.Pointer(&s))
}

// hackStringBytes returns the (unsafe) underlying bytes slice of a string.
func hackStringBytes(str string) []byte {
	return unsafe.Slice(StringData(str), len(str))
}
