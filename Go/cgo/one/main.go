package main

// void xl(const char* s);
import "C"

func main() {
	C.xl(C.CString("hello, world"))
}
