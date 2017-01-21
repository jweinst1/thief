package main

func Byte_to_i32(in []byte) int32 {
    //must be length 4
    return  (int32(in[0]) << 24 | int32(in[1]) << 16 | int32(in[2]) << 8 | int32(in[3]))
}

func Byte_to_i(in []byte) int {
    //must be length 4
    return  int((int32(in[0]) << 24 | int32(in[1]) << 16 | int32(in[2]) << 8 | int32(in[3])))
}

func Byte_to_s(in []byte) string {
	return string(in)
}

func Byte_to_b(in []byte) bool {
	switch in[0] {
	case 0:
		return false
	case 1:
		return true
	default:
		panic("Invalid Bool value")
	}
}

//struct that acts as an iterator for a byte stream
type ByteStream struct {
	bytes []byte
	i int
	done bool
}

func NewByteStream(lst []byte) ByteStream {
	return ByteStream{lst, 0, false}
}

func (self *ByteStream) next() (byte, bool) {
	if self.done {
		return (byte(0), true)
	} else {
		new_b := (self.bytes[self.i++], false)
		if self.i == len(self.bytes) {self.done = true}
		return new_b
	}
}
