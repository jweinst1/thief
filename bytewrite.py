#python file for writing bytes to test thief

btest = open("test.thief", "w")
bytes_to_write = bytearray([1, 44, 56, 55, 55, 0])

btest.write(bytes_to_write)

btest.close()
