package main

import (
	"bufio"
	"fmt"
	"os"
)
//Package bufio implements buffered I/O.
// It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) 
//that also implements the interface but provides buffering and some help for textual I/O.

//Package os provides a platform-independent interface to operating system functionality.
// The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers.
// Often, more information is available within the error. For example, if a call that takes a file name fails, such as Open or Stat, 
// the error will include the failing file name when printed and will be of type *PathError, which may be unpacked for more information.

func main()  {

	fmt.Println("Hello Welcome to GoLang")

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Fav jersey number in football: ")

	input,_:=reader.ReadString('\n')

	fmt.Println("Your Fav jersey number is: ",input)
	fmt.Printf("Your Fav jersey number is: %T\n",input)
}