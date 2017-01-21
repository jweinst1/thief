package main

//file that contains the reading function for the VM

func ReadBytes(env Env, in []byte) {
	OPERS[in[0]](env, in[1:len(in)-1])
}