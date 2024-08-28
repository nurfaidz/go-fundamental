package main

import "fmt"

func float() {
	var a float32 = 3.14              // Declare a float32 variable
	var b float64 = 2.718281828459045 // Declare b float64 variable

	// Perform some basic operations with floating-point numbers
	sum := a + float32(b)        // Sum of a and b, converted to float32
	product := a * float32(b)    // Product of a and b, converted to float32
	difference := b - float64(a) // Difference of b and a, converted to float64
	quotient := b / float64(a)   // Quotient of b and a, converted to float64

	// Declare and initialize complex numbers
	var c complex64 = 1 + 2i  // 32-bit complex number
	var d complex128 = 3 + 4i // 64-bit complex number

	// Perform basic operations with complex numbers
	complexSum := c + complex64(d)     // Sum of c and d, converted to complex64
	complexProduct := c * complex64(d) // Product of c and d, converted to complex64

	// Combining floating-point numbers and complex numbers
	combinedSum := complex(float32(a), 0) + c     // Combine float32 a with complex c
	combinedProduct := complex(float64(b), 0) * d // Combine float64 b with complex d

	// Display the results
	fmt.Println(sum)
	fmt.Println(product)
	fmt.Println(difference)
	fmt.Println(quotient)
	fmt.Println(complexSum)
	fmt.Println(complexProduct)
	fmt.Println(combinedSum)
	fmt.Println(combinedProduct)
}

func main() {
	float()
}
