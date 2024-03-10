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
	"errors"
	"math/big"
)

func containsNumbers(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func containsSymbols(s string) bool {
	for _, c := range s {
		if (c >= '!' && c <= '/') || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= '~') {
			return true
		}
	}
	return false
}

// GeneratePassword generates a password of a given length with certain characteristics.
// The 'length' parameter specifies the length of the password.
// The 'num' parameter determines whether the password should include numbers.
// The 'sym' parameter determines whether the password should include symbols.
func GeneratePassword(length int, no_num bool, no_sym bool) (string, error) {
	if length < 3 {
		return "", errors.New("password length must be at least 3")
	}

	var lowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"
	var upperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numbers = "0123456789"
	var symbols = "!@#$%^&*()_+"

	var characters string
	characters += lowerCaseLetters
	characters += upperCaseLetters
	if !no_num {
		characters += numbers
	}
	if !no_sym {
		characters += symbols
	}

	// Generate the password
	var password string
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		password += string(characters[n.Int64()])
	}

	// Ensure it meets the requirements
	if !no_num {
		if !containsNumbers(password) {
			return GeneratePassword(length, no_num, no_sym)
		}
	}
	if !no_sym {
		if !containsSymbols(password) {
			return GeneratePassword(length, no_num, no_sym)
		}
	}

	return password, nil
}
