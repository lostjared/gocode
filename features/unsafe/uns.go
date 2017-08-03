
/*
using some of the features of unsafe package
Sizeof, Pointer
*/

package main

import (
	"fmt"
	"os"
	"unsafe"
)

func FloatBits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func Word(v uint32) uint16 {
	return *(*uint16)(unsafe.Pointer(&v))
}

func main() {
	fmt.Println("Unsafe package")
	s := unsafe.Sizeof(int64(0))
	fmt.Println("size of int64: ", s)
	s = unsafe.Sizeof(float32(0))
	fmt.Println("size of float32: ", s)
	s = unsafe.Sizeof(rune(0))
	fmt.Println("size of rune: ", s)
	fmt.Printf("%#016x\n", FloatBits(1.0))
	var x uint32 = 0xFFFFFFFF
	val := Word(x)
	fmt.Println("val: ", val)
	first_byte := *(*uint8)(unsafe.Pointer(&x))
	fmt.Println("First byte: ", first_byte)
	os.Exit(0)
}