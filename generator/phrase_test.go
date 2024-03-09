package generator

import (
	"regexp"
	"testing"
)

// TestGeneratePassphraseLength ensures the function returns a string of the correct length
func TestGeneratePassphraseLength(t *testing.T) {
	length := 5
	sep := "-"
	num := false
	sym := false
	cap := false

	passphrase, err := GeneratePassphrase(length, sep, num, sym, cap)
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
	num := false
	sym := false
	cap := false

	passphrase, err := GeneratePassphrase(length, sep, num, sym, cap)
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
	num := true
	sym := false
	cap := false

	result, err := GeneratePassphrase(length, sep, num, sym, cap)
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
	num := false
	sym := true
	cap := false

	result, err := GeneratePassphrase(length, sep, num, sym, cap)
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
	num := false
	sym := false
	cap := true

	result, err := GeneratePassphrase(length, sep, num, sym, cap)
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
	num := true
	sym := true
	cap := true

	result, err := GeneratePassphrase(length, sep, num, sym, cap)
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
