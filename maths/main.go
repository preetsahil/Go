package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main()  {

	var a int=2
	var b float64=4.5
    // invalid operation: a + b (mismatched types int and float64)
	fmt.Println("value of a +b is",a+int(b))

//random number
// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n) from the default Source.
//  It panics if n <= 0.
// In Golang, the rand.Seed() function is used to set a seed value to generate pseudo-random numbers.(depreciated)
// If the same seed value is used in every execution, then the same set of pseudo-random numbers is generated. 
// In order to get a different set of pseudo-random numbers, we need to update the seed value.
	// fmt.Println(rand.Intn(5))

	//random from crypto
// func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
// Int returns a uniform random value in [0, max). It panics if max <= 0.

myrandomNumber,_:=rand.Int(rand.Reader,big.NewInt(5))
fmt.Println(myrandomNumber)


	
}