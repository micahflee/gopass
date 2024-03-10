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
	"regexp"
	"testing"
)

// TestGeneratePassphraseLength ensures the function returns a string of the correct length
func TestGeneratePassphraseLength(t *testing.T) {
	length := 5
	sep := "-"
	no_num := true
	no_sym := true
	no_cap := true

	passphrase, err := GeneratePassphrase(length, sep, no_num, no_sym, no_cap)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	words := regexp.MustCompile(sep).Split(passphrase, -1)
	if len(words) != length {
		t.Errorf("Expected %v words, got %v", length, len(words))
	}
}

// TestGeneratePassphraseSeparator ensures the function returns a string with the correct separator
func TestGeneratePassphraseSeparator(t *testing.T) {
	length := 5
	sep := "ASDF"
	no_num := true
	no_sym := true
	no_cap := true

	passphrase, err := GeneratePassphrase(length, sep, no_num, no_sym, no_cap)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Make sure the separator is in the passphrase length-1 times
	if matched, _ := regexp.MatchString(sep, passphrase); !matched {
		t.Errorf("Separator not in passphrase")
	}

	words := regexp.MustCompile(sep).Split(passphrase, -1)
	if len(words) != length {
		t.Errorf("Expected %v words, got %v", length, len(words))
	}
}

// TestGeneratePassphraseNumbers ensures the function returns a string with a number
func TestGeneratePassphraseNumbers(t *testing.T) {
	length := 5
	sep := "-"
	no_num := false
	no_sym := true
	no_cap := true

	result, err := GeneratePassphrase(length, sep, no_num, no_sym, no_cap)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Check if a number is in the passphrase
	if matched, _ := regexp.MatchString("[0-9]", result); !matched {
		t.Errorf("No number in the passphrase")
	}
}

// TestGeneratePassphraseSymbols ensures the function returns a string with a symbol
func TestGeneratePassphraseSymbols(t *testing.T) {
	length := 5
	sep := "-"
	no_num := true
	no_sym := false
	no_cap := true

	result, err := GeneratePassphrase(length, sep, no_num, no_sym, no_cap)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Check if a symbol is in the passphrase
	if matched, _ := regexp.MatchString("[!@#$%^&*()_+]", result); !matched {
		t.Errorf("No symbol in the passphrase")
	}
}

// TestGeneratePassphraseCapitalize ensures the function returns a string with a capitalized word
func TestGeneratePassphraseCapitalize(t *testing.T) {
	length := 5
	sep := "-"
	no_num := true
	no_sym := true
	no_cap := false

	result, err := GeneratePassphrase(length, sep, no_num, no_sym, no_cap)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Check if a word is capitalized in the passphrase
	if matched, _ := regexp.MatchString("[A-Z]", result); !matched {
		t.Errorf("No capitalized word in the passphrase")
	}
}

// TestGeneratePassphraseComplex ensures the function returns a string with a number, symbol, and capitalized word
func TestGeneratePassphraseComplex(t *testing.T) {
	length := 5
	sep := "-"
	no_num := false
	no_sym := false
	no_cap := false

	result, err := GeneratePassphrase(length, sep, no_num, no_sym, no_cap)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// Check if a number is in the passphrase
	if matched, _ := regexp.MatchString("[0-9]", result); !matched {
		t.Errorf("No number in the passphrase")
	}

	// Check if a symbol is in the passphrase
	if matched, _ := regexp.MatchString("[!@#$%^&*()_+]", result); !matched {
		t.Errorf("No symbol in the passphrase")
	}

	// Check if a word is capitalized in the passphrase
	if matched, _ := regexp.MatchString("[A-Z]", result); !matched {
		t.Errorf("No capitalized word in the passphrase")
	}
}
