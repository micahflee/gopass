# Gopass

Gopass is a secure password and passphrase generator that uses the Diceware method to generate secure and memorable passphrases.

Passwords and passphrases are generated using a [cryptographically secure random number generator](https://pkg.go.dev/crypto/rand).

For example, if you need a passphrase, run `gopass phrase`:

```
$ gopass phrase
puma-nag2-scout!-chute-CHUNK-emu
```

And if you need a password, run `gopass word`:

```
$ gopass word
KMK^XhP+n+53xBbnGeN7
```

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
