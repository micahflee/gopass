package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
)

func loadWordlist() []string {
	file, err := os.Open("eff_short_wordlist_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return words
}

func main() {
	lengthPtr := flag.Int("length", 4, "Length of passphrase to generate")
	sepPtr := flag.String("sep", "-", "Word separator")
	flag.Parse()

	words := loadWordlist()
	passphraseWords := make([]string, *lengthPtr)
	var passphrase string

	for i := 0; i < *lengthPtr; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
		if err != nil {
			log.Fatal(err)
		}
		passphraseWords[i] = words[n.Int64()]
	}

	// Join the words together
	for i := 0; i < *lengthPtr; i++ {
		passphrase += passphraseWords[i]
		if i < *lengthPtr-1 {
			passphrase += *sepPtr
		}
	}
	fmt.Println(passphrase)
}
