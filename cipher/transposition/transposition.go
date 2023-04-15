// A Transposition cipher implementation.
// Is an encryption method in which the positions occupied by units of plaintext
// (usually being characters or groups of characters) are shifted according to a regular system,
// so that the ciphertext is a permutation of the plaintext.
// Reference: https://en.wikipedia.org/wiki/Transposition_cipher
package transposition

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

const placeholder = ' '

var (
	ErrNoTextToEncrypt = errors.New("no text to encrypt")
	ErrKeyMissing      = errors.New("missing Key")
)

func getIndex(wordSet []rune, subString rune) int {
	n := len(wordSet)
	for i := 0; i < n; i++ {
		if wordSet[i] == subString {
			return i
		}
	}

	return 0
}

func getKey(keyWord string) []int {
	keyWord = strings.ToLower(keyWord)
	word := []rune(keyWord)
	sortedWord := make([]rune, len(word))
	usedLettersMap := make(map[rune]int)
	wordLength := len(word)
	resultKey := make([]int, wordLength)
	copy(sortedWord, word)

	sort.Slice(sortedWord, func(i, j int) bool { return sortedWord[i] < sortedWord[j] })

	for i := 0; i < wordLength; i++ {
		char := word[i]
		numberOfUsage := usedLettersMap[char]
		resultKey[i] = getIndex(sortedWord, char) + numberOfUsage + 1 // +1 - so that the indexing does not start from 0
		numberOfUsage++
		usedLettersMap[char] = numberOfUsage
	}

	return resultKey
}

func Encrypt(text []rune, keyWord string) (result []rune, err error) {
	key := getKey(keyWord)
	keyLength := len(key)
	textLength := len(text)

	if keyLength <= 0 {
		return nil, ErrKeyMissing
	}
	if textLength <= 0 {
		return nil, ErrNoTextToEncrypt
	}
	if text[len(text)-1] == placeholder {
		return nil, fmt.Errorf("%w: cannot encrypt a text, %q, ending with the placeholder char %q", ErrNoTextToEncrypt, text, placeholder)
	}

	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text = append(text, placeholder)
	}

	textLength = len(text)
	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[key[j]-1] = text[i+j]
		}
		result = append(result, transposition...)
	}
	return
}

func Decrypt(text []rune, keyWord string) (result []rune, err error) {
	key := getKey(keyWord)
	textLength := len(text)

	if textLength <= 0 {
		return nil, ErrNoTextToEncrypt
	}

	keyLength := len(key)
	if keyLength <= 0 {
		return nil, ErrKeyMissing
	}

	n := textLength % keyLength
	for i := 0; i < keyLength-n; i++ {
		text = append(text, placeholder)
	}

	for i := 0; i < textLength; i += keyLength {
		transposition := make([]rune, keyLength)
		for j := 0; j < keyLength; j++ {
			transposition[j] = text[i+key[j]-1]
		}
		result = append(result, transposition...)
	}

	return []rune(strings.TrimRight(string(result), string(placeholder))), nil
}
