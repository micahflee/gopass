# Gopass

Gopass is a password, passphrase, and passcode generator that uses a [cryptographically secure random number generator]((https://pkg.go.dev/crypto/rand)). It uses the Diceware method to generate secure and memorable passphrases.

For example, if you need a passphrase, run `gopass phrase`:

```
$ gopass phrase
puma-nag2-scout!-chute-CHUNK-emu
```

If you need a password, run `gopass word`:

```
$ gopass word
KMK^XhP+n+53xBbnGeN7
```

And if you need passcode, run `gopass code`:

```
$ gopass code
060781
```

# Getting Started

I haven't put much effort into packaging. To install, you'll need to [install Go](https://go.dev/dl/) and then run:

```
go install github.com/micahflee/gopass@latest
```

Gopass will get installed in `$GOPATH/bin/gopass`.

# Usage

## Generating Passphrases

By default, passphrases contain 5 words, are separated by `-` characters, include a capitalized word, and include one number and one symbol. This is configurable using the following flags:

- `-l`, `--length int`: Number of words in the passphrase (default 6)
- `-c`, `--no-capitalize`: Don't capitalize a word in the passphrase
- `-n`, `--no-number`: Don't include a number in the passphrase
- `-y`, `--no-symbol`: Don't include a symbol in the passphrase
- `-s`, `--separator string`: Separator character between words (default `-`)

For example, if you want a 7-word passphrase separated by `.`s, and without numbers, symbols, or capitals:

```
gopass phrase --length 7 --separator . --no-number --no-symbol --no-capitalize
```

This will generate a passphrase like `lived.movie.swore.spree.sage.floss.spent`.

## Generating Passwords

By default, passwords contain 20 characters, letters, numbers, and symbols. This is configurable using the following flags:

- `-l`, `--length int`: Number of characters in the password (default 20)
- `-n`, `--no-numbers`: Don't include numbers in the password
- `-y`, `--no-symbols`: Don't include symbols in the password

For example, if you want a 16 character password with letters and numbers but no symbols:

```
gopass word --length 16 --no-symbols
```

This will generate a password like `QKiyeYcyd4B6ndq5`.

## Generating Passcodes

By default, passcodes are 6 digits long. This is configurable:

- `-l`, `--length int`: Number of digits in the passcode (default 6)

For example, if you want a 4 digit passcode:

```
gopass code --length 4
```

This will generate a passcode like `4320`.