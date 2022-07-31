package base

import (
	"fmt"
)

func Base() {

	n := 2548
	n2 := 0x9F4

	// %x or X for hexadecimal value of 2548
	fmt.Printf("Hexadecimal value: %x \n", n)

	// %d for decimal value
	fmt.Printf("Decimal value: %d\n", n2)

	// %o for octal value
	fmt.Printf("Octal value: %o\n", n2)

	// %b for binary value
	fmt.Printf("Binary value: %b\n", n2)

	// ---- //

	/*
		NERD FACT
		'nibble' composed of 4 bits
		'byte' composed of 8 bits
		'word' composed of 16 bits
		'doubleword' composed of 32 bits
		'quadword' composed of 64 bits

	*/

	// to declare a byte

	b := make([]byte, 0)
	b = append(b, 255)
	b = append(b, 10)
	fmt.Printf("Byte: %v\n", b)

	// ---- //

	/*
		NERD FACT

		Unicode: U+0041 denoted the character 'A'
		ASCII: 1000001 denoted the character 'A'
		UTF-8: 41 denoted the character 'A"
		Runes: represents a Unicode code point.


	*/

	a := 'A'

	fmt.Printf("Unicode: %U\n", a)
	fmt.Printf("ASCII: %b\n", a)
	fmt.Printf("UTF-8: %x\n", a)

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

}
