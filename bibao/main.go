package main

func main(){
	i:=incr()
	println(i()) // 1
	println(i()) // 2
	println(i()) // 3
	println(incr()()) // 1
	println(incr()()) // 1
	println(incr()()) // 1
}

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

//-----------------------------

