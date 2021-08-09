package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	// TODO verbose print, this is because of sign in methods are not the same type.
	var method = flag.String("type", "sha256", "Signature methods, including: sha256, sha384, sha512")
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		text := []byte(input.Text())
		if *method == "sha382" {
			sign := sha512.Sum384(text)
			fmt.Printf("%x\n", sign)
		} else if *method == "sha512" {
			sign := sha512.Sum512(text)
			fmt.Printf("%x\n", sign)
		} else {
			sign := sha256.Sum256(text)
			fmt.Printf("%x\n", sign)
		}
	}
}
