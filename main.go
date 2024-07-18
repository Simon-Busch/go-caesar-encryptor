package main

import (
	"fmt"
	"strings"
)

const originalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"


func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain text: ", plainText)

	encryptedText := encrypt(5, plainText)
	fmt.Println("Encrypted text: ", encryptedText)

	decryptedText := decrypt(5, encryptedText)
	fmt.Println("Decrypted text: ", decryptedText)

}

func encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetters)

	var hashedString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(originalLetters, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetters)) % len(originalLetters)
			hashedString += string(hashLetter[letterPosition])
		}
		return r
	}

	strings.Map(findOne, plainText)
	return hashedString
}

func decrypt(key int, encryptedText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetters)
	var hashedString = ""
	findOne := func(r rune)rune {
		pos := strings.Index(hashLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalLetters)) % len(originalLetters)
			hashedString += string(originalLetters[letterPosition])
			return r
		}
		return r
	}


	strings.Map(findOne, encryptedText)
	return hashedString
}

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOverLetter := string(runes[0 : len(letter)-key])
	return fmt.Sprintf("%s%s", lastLetterKey, leftOverLetter)
}
