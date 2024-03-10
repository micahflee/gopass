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
package cmd

import (
	"fmt"

	"github.com/micahflee/gopass/generator"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Generate a numeric passcode",
	Long:  `This generates a secure numeric passcode of a given length.`,
	Run: func(cmd *cobra.Command, args []string) {
		len, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Println("Error getting length: ", err)
			return
		}

		// Ensure password has a minimum length of 4
		if len < 4 {
			fmt.Println("Error: Passcode length must be at least 4")
			return
		}

		passcode, err := generator.GeneratePasscode(len)
		if err != nil {
			fmt.Println("Error generating passcode: ", err)
			return
		}
		fmt.Println(passcode)
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	codeCmd.Flags().IntP("length", "l", 6, "Number of digits in the passcode")
}
