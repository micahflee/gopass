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

// TestGeneratePasscodeLength ensures the function returns a string of the correct length
func TestGeneratePasscodeLength(t *testing.T) {
	passcode4, err := GeneratePasscode(4)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(passcode4) != 4 {
		t.Errorf("Expected 4 characters, got %v", len(passcode4))
	}

	passcode10, err := GeneratePasscode(10)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(passcode10) != 10 {
		t.Errorf("Expected 10 characters, got %v", len(passcode10))
	}
}

// TestGeneratePasscodeTooShort ensures the function returns an error if the length is less than 4
func TestGeneratePasscodeTooShort(t *testing.T) {
	length := 3

	_, err := GeneratePasscode(length)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// TestGeneratePasscodeRegex ensures the function returns a string that only contains numbers
func TestGeneratePasscodeRegex(t *testing.T) {
	passcode, err := GeneratePasscode(10)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if match, _ := regexp.MatchString("^[0-9]+$", passcode); !match {
		t.Errorf("Expected only numbers, got %v", passcode)
	}
}
