// Strings

// What is a String?
// A string is a slice of bytes in Go. Strings can be created by enclosing a set of characters inside double quotes ""

package main

// func main() {
// 	name := "Hello World"
// 	fmt.Println(name)
// }

// Strings in Go are Univode compliant and are UTF-8 Encoded

// Since a string is a slice of byptes, its possible to access each byte of a string.

// func printBytes(s string) {
// 	fmt.Printf("Bytes: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// func main() {
// 	name := "Hello World"
// 	fmt.Printf("String: %s\n", name)
// 	printBytes(name)
// }

// %x is the format specifer for hexadecimal

// Unicode character set maps each character in the wolrd to a unique number. This ensures that there are no collisions between alphabets of different languages. These numebrs are platform independent/

// UTF-8 encoding is a variable sized encoding scheme to represent unicode code points in memory. Variable sized encoding means the code points are represented using 1, 2, 3 or 4 bytes depending on their size.

// func printBytes(s string) {
// 	fmt.Printf("Bytes: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// %x is hexadecimal notation (for floating-point and complex constitents)

// func printChars(s string) {
// 	fmt.Printf("Characters: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%c ", s[i])
// 	}
// }

// %c is the character represented by the corresponding Unicode code point.
// %c is used to print the characters of the string

// func main() {
// 	name := "Hello World"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)
// }

// %s is the uninterpreted bytes of the string or slice

// serious bug:

// import (
// 	"fmt"
// )

// func printBytes(s string) {
// 	fmt.Printf("Bytes: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// func printChars(s string) {
// 	fmt.Printf("Characters: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%c ", s[i])
// 	}
// }

// func main() {
// 	name := "Hello World"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)
// 	fmt.Printf("\n\n")
// 	name = "Señor"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)
// }

// In UTF-8 encoding a code point can occupy more than 1 byte. SO how do we solve this. This is where rune saves us.

// Rune

// A rune is a builtin type in Go and it's the alias of int32. Rune represents a Unicode code point in Go. It doesn't matter how many bytes the code point occupies, it can be represented by a rune. Let's modify the above program to print characters using a rune.

// import (
// 	"fmt"
// )

// func printBytes(s string) {
// 	fmt.Printf("Bytes: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// func printChars(s string) {
// 	fmt.Printf("Characters: ")
// 	runes := []rune(s)
// 	for i := 0; i < len(runes); i++ {
// 		fmt.Printf("%c ", runes[i])
// 	}
// }

// func main() {
// 	name := "Hello World"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)
// 	fmt.Printf("\n\n")
// 	name = "Señor"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)
// }

// Accessing individual runes using foor range loop

//  Go offers us a much easier way to iterate over the individual runes of a string using the `for range` loop

// import (
// 	"fmt"
// )

// func charsAndBytePosition(s string) {
// 	for index, rune := range s {
// 		fmt.Printf("%c starts at byte %d\n", rune, index)
// 	}
// }

// func main() {
// 	name := "Señor"
// 	charsAndBytePosition(name)
// }

// func main() {
// 	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
// 	str := string(byteSlice)
// 	fmt.Println(str)
// }

// The `RuneCountInString(s string) (n int)` function of the utf8 package can be used to find the length of the string. This method takes a string as an argument and retuns the number of runes in it.

// `len(s)` is used to find the number of bytes in the string and it doesn't return the string length. As we already discussed, some Unicode characters have code points that occupy more than 1 byte. Using `len` to find out the length of those strings will return incorrect string length.

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	word1 := "Señor"
// 	fmt.Printf("String: %s\n", word1)
// 	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
// 	fmt.Printf("Number of bytes: %d\n", len(word1))
// }

// String comparison

// The `==` operator is used to compare two strings for equality. If both the strings are equal, then the result is `true` else its `false`

import (
	"fmt"
)

// func compareStrings(str1 string, str2 string) {
// 	if str1 == str2 {
// 		fmt.Printf("%s and %s are equal\n", str1, str2)
// 		return
// 	}
// 	fmt.Printf("%s and %s are not equal\n", str1, str2)
// }

// func main() {
// 	string1 := "Go"
// 	string2 := "Go"
// 	compareStrings(string1, string2)

// 	string3 := "hello"
// 	string4 := "world"
// 	compareStrings(string3, string4)
// }

// Strings are immutable in Go. Once a string is created it's not possible to change it

// To work around string immutablility, strings are converted to a slice of runes. Thien that slice is mutated with whatever changes are needed and converted back to a new string.

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
