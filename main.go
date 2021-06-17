package main

import (
	"crypto/rand"
	"fmt"
	"io"
)
var nums = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func main() {
	x := 5 // change this value for the number of digits to generate for the random no
	randomNo := GenerateRandomNo(x)
	fmt.Println(randomNo)	
}

func GenerateRandomNo(x int) string {
	buffer := make([]byte, x)
	n, err := io.ReadAtLeast(rand.Reader, buffer, x)
		
	if  err != nil || n != x {
		return GenerateRandomNo(x)
	}
	
	for i := 0; i < len(buffer); i++ {
		buffer [i] = nums[int(buffer[i])%len(nums)]
	}
	return string(buffer)
}
