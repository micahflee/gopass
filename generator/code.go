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

// GeneratePasscode generates a numeric passcode of a given length.
// The 'length' parameter specifies the length of the passcode.
func GeneratePasscode(length int) (string, error) {
	if length < 4 {
		return "", errors.New("passcode length must be at least 4")
	}

	var numbers = "0123456789"
	var passcode string = ""

	// Generate the passcode
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
		if err != nil {
			return "", err
		}
		passcode += string(numbers[n.Int64()])
	}

	return passcode, nil
}
