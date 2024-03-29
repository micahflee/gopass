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

var phraseCmd = &cobra.Command{
	Use:   "phrase",
	Short: "Generate a passphrase",
	Long:  `This generates a secure and memorable passphrase. You can choose the number of words in the passphrase, the separator character, and whether or not to include a numbers and/or a symbols.`,
	Run: func(cmd *cobra.Command, args []string) {
		len, err := cmd.Flags().GetInt("length")
		if err != nil {
			fmt.Println("Error getting lenth: ", err)
			return
		}
		sep, err := cmd.Flags().GetString("separator")
		if err != nil {
			fmt.Println("Error getting separator: ", err)
			return
		}
		no_num, err := cmd.Flags().GetBool("no-number")
		if err != nil {
			fmt.Println("Error getting number: ", err)
			return
		}
		no_sym, err := cmd.Flags().GetBool("no-symbol")
		if err != nil {
			fmt.Println("Error getting symbol: ", err)
			return
		}
		no_cap, err := cmd.Flags().GetBool("no-capitalize")
		if err != nil {
			fmt.Println("Error getting capitalize: ", err)
			return
		}

		passphrase, err := generator.GeneratePassphrase(len, sep, no_num, no_sym, no_cap)
		if err != nil {
			fmt.Println("Error generating passphrase: ", err)
			return
		}
		fmt.Println(passphrase)
	},
}

func init() {
	rootCmd.AddCommand(phraseCmd)

	phraseCmd.Flags().IntP("length", "l", 6, "Number of words in the passphrase")
	phraseCmd.Flags().StringP("separator", "s", "-", "Separator character between words")
	phraseCmd.Flags().BoolP("no-number", "n", false, "Don't include a number in the passphrase")
	phraseCmd.Flags().BoolP("no-symbol", "y", false, "Don't include a symbol in the passphrase")
	phraseCmd.Flags().BoolP("no-capitalize", "c", false, "Don't capitalize a word in the passphrase")
}
