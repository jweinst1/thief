package main

func Byte_to_i32(in []byte) int32 {
    //must be length 4
    return  (int32(in[0]) << 24 | int32(in[1]) << 16 | int32(in[2]) << 8 | int32(in[3]))
}

func Byte_to_i(in []byte) int {
    //must be length 4
    return  int((int32(in[0]) << 24 | int32(in[1]) << 16 | int32(in[2]) << 8 | int32(in[3])))
}