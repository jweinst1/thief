package main

//file that contains the reading function for the VM

func ReadBytes(env Env, in []byte) {
	bytes := make([]byte, 1)
	fn := byte(0)
	mode := 0
	for i:=0;i<len(in);i++ {
		switch mode {
		case 0:
			fn = byte(i)
			mode = 1
		case 1:
			if i != len(in)-1 {
				if in[i+1] == 0 {
					bytes = append(bytes, in[i])
					mode = 2
				} else {
					bytes = append(bytes, in[i])
				}
			} else {
				bytes = append(bytes, in[i])
			}
		case 2:
			OPERS[fn](env, bytes)
			mode = 0
			fn = byte(0)
			bytes = make([]byte, 1)
		}
	}
}