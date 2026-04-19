package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("[COMPUTER]")
	num := runtime.NumCPU()
	versionn := runtime.Version()
	fmt.Println("CPU cores", num)
	fmt.Println(versionn)
}
