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

// TestGeneratePasswordLength ensures the function returns a string of the correct length
func TestGeneratePasswordLength(t *testing.T) {
	no_num := false
	no_sym := false

	password20, err := GeneratePassword(20, no_num, no_sym)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(password20) != 20 {
		t.Errorf("Expected 20 characters, got %v", len(password20))
	}

	password50, err := GeneratePassword(50, no_num, no_sym)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(password50) != 50 {
		t.Errorf("Expected 50 characters, got %v", len(password50))
	}
}

// TestGeneratePasswordTooShort ensures the function returns an error if the length is less than 3
func TestGeneratePasswordTooShort(t *testing.T) {
	length := 2
	no_num := false
	no_sym := false

	_, err := GeneratePassword(length, no_num, no_sym)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// TestGeneratePasswordNumbers ensures the function returns a string with numbers in it
func TestGeneratePasswordNumbers(t *testing.T) {
	length := 20
	no_num := false
	no_sym := true

	password, err := GeneratePassword(length, no_num, no_sym)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if matched, _ := regexp.MatchString("[0-9]", password); !matched {
		t.Errorf("No numbers in password")
	}
}

// TestGeneratePasswordSymbols ensures the function returns a string with symbols in it
func TestGeneratePasswordSymbols(t *testing.T) {
	length := 20
	no_num := true
	no_sym := false

	password, err := GeneratePassword(length, no_num, no_sym)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if matched, _ := regexp.MatchString("[!@#$%^&*()_+]", password); !matched {
		t.Errorf("No symbols in password")
	}
}

// TestGeneratePasswordComplex ensures the function returns a string with numbers and symbols in it
func TestGeneratePasswordComplex(t *testing.T) {
	length := 20
	no_num := false
	no_sym := false

	password, err := GeneratePassword(length, no_num, no_sym)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if matched, _ := regexp.MatchString("[0-9]", password); !matched {
		t.Errorf("No numbers in password")
	}

	if matched, _ := regexp.MatchString("[!@#$%^&*()_+]", password); !matched {
		t.Errorf("No symbols in password")
	}
}
