package main

import "fmt"

func main() {
	for i := 60; i < 127; i++ {
		fmt.Printf("%d \t %b \t %x \t %#X \t %#o \t %q\n", i, i, i, i, i, i)
	}

}
