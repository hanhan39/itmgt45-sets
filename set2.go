package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Shift letter:")
	fmt.Println(shiftLetter("Z", 1))

	fmt.Println("Caesar cipher:")
	fmt.Println(caesarCipher("HELLO HI", 3))

	fmt.Println("Shift by letter:")
	fmt.Println(shiftByLetter("A", "D"))

	fmt.Println("Vigenere cipher:")
	fmt.Println(vigenereCipher("LONGTEXT", "KEY"))

	fmt.Println("Scytale cipher:")
	fmt.Println(scytaleCipher("INFORMATION_AGE", 3))

	fmt.Println("Scytale decipher:")
	fmt.Println(scytaleDecipher("IMNNA_FTAOIGROE", 3))
}

func shiftLetter(letter string, shift int) string {
	// create alphabet array
	alphabet := make([]string, 26)
	// for every character in array:
	for i := 0; i < 26; i++ {
		// add i to value of A
		alphabet[i] = string('A' + i)
	}

	// if character found in alphabet:
	if slices.Contains(alphabet, letter) == true {
		// add shift to index of alphabet character
		shift = shift + slices.Index(alphabet, letter)
		// wrap around w/ modulo
		shift = shift % 26
		// return shifted character from alphabet
		return alphabet[shift]
	} else {
		// return space if otherwise
		return " "
	}
}

func caesarCipher(message string, shift int) string {
	// create empty array for new message
	newMessage := make([]string, len(message))
	// same as before
	alphabet := make([]string, 26)
	for i := 0; i < 26; i++ {
		alphabet[i] = string('A' + i)
	}

	// for each letter in message:
	for i, letter := range message {
		// if character found in alphabet:
		if slices.Contains(alphabet, string(letter)) == true {
			// add shift to index value of letter
			index := shift + slices.Index(alphabet, string(letter))
			// wrap around w/ modulo
			index = index % 26
			// append character at index to new message
			newMessage[i] = alphabet[index]
		} else {
			// append space if otherwise
			newMessage[i] = " "
		}
	}
	// return constructed message
	return strings.Join(newMessage, "")
}

func shiftByLetter(letter string, letterShift string) string {
	// same as before
	alphabet := make([]string, 26)
	for i := 0; i < 26; i++ {
		alphabet[i] = string('A' + i)
	}

	// if both characters found in alphabet:
	if slices.Contains(alphabet, letter) == true && slices.Contains(alphabet, letterShift) == true {
		// get index of letterShift
		shift := slices.Index(alphabet, letterShift)
		// add shift to index of letter in alphabet
		shift = shift + slices.Index(alphabet, letter)
		// wrap around w/ modulo
		shift = shift % 26
		// return shifted character from alphabet
		return alphabet[shift]
	} else {
		// return space if otherwise
		return " "
	}
}

func vigenereCipher(message string, key string) string {
	// same as before
	newMessage := make([]string, len(message))
	alphabet := make([]string, 26)
	for i := 0; i < 26; i++ {
		alphabet[i] = string('A' + i)
	}

	// if length of key & message not equal:
	if len(key) != len(message) {
		// create new key string array
		newKey := make([]string, len(message))
		// construct repeating pattern from original key w/ modulo until message length reached
		for i := 0; i < len(message); i++ {
			newKey[i] = string(key[i%len(key)])
		}
		// set key to new key
		key = strings.Join(newKey, "")
	}

	// for each letter in message:
	for i, letter := range message {
		// if letter and current key character found in alphabet:
		if slices.Contains(alphabet, string(letter)) == true && slices.Contains(alphabet, string(key[i])) == true {
			// set shift value to index of current key character in alphabet
			shift := slices.Index(alphabet, string(key[i]))
			// add shift value to index of letter in alphabet
			index := shift + slices.Index(alphabet, string(letter))
			// wrap around w/ modulo
			index = index % 26
			// append character at index to new message
			newMessage[i] = alphabet[index]
		} else {
			// append space if otherwise
			newMessage[i] = " "
		}
	}
	// return constructed message
	return strings.Join(newMessage, "")
}

func scytaleCipher(message string, shift int) string {
	// if message length is not multiple of shift:
	if len(message)%shift != 0 {
		// append underscores to message until it is
		for len(message)%shift != 0 {
			message += "_"
		}
	}
	// create new message string array
	newMessage := make([]string, len(message))

	// until message length is reached:
	for i := 0; i < len(message); i++ {
		// perform operation specified in guidelines
		scytaleIndex := int(i/shift + (len(message)/shift)*(i%shift))
		// get character found at index
		index := message[scytaleIndex]
		// append character at index to new message
		newMessage[i] = string(index)
	}
	// return constructed message
	return strings.Join(newMessage, "")
}

func scytaleDecipher(message string, shift int) string {
	// create new message string array
	newMessage := make([]string, len(message))

	// mostly identical, except for one line
	for i := 0; i < len(message); i++ {
		// flip around the values in original line, replace i/shift with i/(len(message)/shift)
		scytaleIndex := int(i/(len(message)/shift) + i%(len(message)/shift)*shift)
		index := message[scytaleIndex]
		newMessage[i] = string(index)
	}
	// return deciphered message
	return strings.Join(newMessage, "")
}
