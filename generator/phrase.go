/*
Copyright © 2024 Micah Lee micah@micahflee.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package generator

import (
	"crypto/rand"
	_ "embed"
	"math/big"
	"strings"
)

//go:embed eff_short_wordlist_1.txt
var wordlist string

// GeneratePassphrase generates a passphrase of a given length with certain characteristics.
// The 'length' parameter specifies the length of the passphrase.
// The 'sep' parameter specifies the separator to use between words in the passphrase.
// The 'num' parameter determines whether the passphrase should include numbers.
// The 'sym' parameter determines whether the passphrase should include symbols.
// The 'cap' parameter determines whether the passphrase should include capital letters.
func GeneratePassphrase(length int, sep string, num bool, sym bool, cap bool) (string, error) {
	allWords := strings.Split(wordlist, "\n")

	var passphrase string
	passphraseWords := make([]string, length)

	// Gather the words
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(allWords))))
		if err != nil {
			return "", err
		}
		passphraseWords[i] = allWords[n.Int64()]
	}

	// Add a number to the passphrase
	var numIndex int
	var numDigit = ""
	if num {
		// Choose a word to add a number to
		n, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return "", err
		}
		numIndex = int(n.Int64())

		// Choose a random digit
		n, err = rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		numDigit = n.String()
	}

	// Add a symbol to the passphrase
	var symIndex int
	var symChar = ""
	if sym {
		// Choose a word to add a symbol to
		n, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return "", err
		}
		symIndex = int(n.Int64())

		// Choose a random symbol
		symbols := "!@#$%^&*()_+"
		n, err = rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return "", err
		}
		symChar = string(symbols[n.Int64()])
	}

	// Capitalize a word
	var capIndex int
	if cap {
		// Choose a word to capitalize
		n, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return "", err
		}
		capIndex = int(n.Int64())
	}

	// Join the words together
	for i := 0; i < length; i++ {
		// Add the word (capitalized or not)
		if cap && i == capIndex {
			passphrase += strings.ToUpper(passphraseWords[i])
		} else {
			passphrase += passphraseWords[i]
		}

		// Add the number
		if num && i == numIndex {
			passphrase += numDigit
		}

		// Add the symbol
		if sym && i == symIndex {
			passphrase += symChar
		}

		// Add the separator
		if i < length-1 {
			passphrase += sep
		}
	}

	return passphrase, nil
}
