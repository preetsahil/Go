package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Package strconv implements conversions to and from string representations of basic data types.
// Package strings implements simple functions to manipulate UTF-8 encoded strings.
func main()  {

	fmt.Println("Hello Welcome to GoLang")

	reader:=bufio.NewReader(os.Stdin)

	input,_:=reader.ReadString('\n')
//ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.
// If the base argument is 0, the true base is implied by the string's prefix following the sign (if present): 2 for "0b", 8 for "0" or "0o", 16 for "0x", and 10 otherwise.
//  Also, for argument base 0 only, underscore characters are permitted as defined by the Go syntax for integer literals.

// if input is 9_0 answer will be 91 with base 0
// 0o_3 this is 3 with base 8 answer 4
// 0x_3 this is 3 with base 16 answer 4
	increase,err:=strconv.ParseInt(strings.TrimSpace(input),10,64)

// 	Hello Welcome to GoLang
// 0x_15  to convert to base 10
// 22

	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println(increase+1)
	}

	
}