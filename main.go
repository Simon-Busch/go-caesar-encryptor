package main

import "fmt"

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain text: ", plainText)

	encryptedText := encrypt(5, plainText)
	fmt.Println("Encrypted text: ", encryptedText)

	decryptedText := decrypt(5, encryptedText)
	fmt.Println("Decrypted text: ", decryptedText)
	
}

func encrypt(key int, plainText string) (result string) {

}

func decrypt(key int, encryptedText string) (result string) {

}
