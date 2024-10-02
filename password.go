package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits       = "0123456789"
	Symbols      = "!@#$%&*()_+-={}[]:?,."
)

// Generates passwords by merging user info and adding random symbols
func generateRandomPasswords(targetInfo []string, desiredLength int, passwdLengthMin int, passwdLengthMax int, knownPart string) []string {
	rand.Seed(time.Now().UnixNano())

	passwords := make(map[string]struct{})
	mergedInfo := mergeUserInfo(targetInfo)

	bar := progressbar.Default(int64(desiredLength))

	var known string
	var position int

	if strings.HasPrefix(knownPart, "P") && strings.HasSuffix(knownPart, "P") {
		_, err := fmt.Sscanf(knownPart, "P%dP", &position)
		if err != nil {
			fmt.Println("Error reading position:", err)
			return nil
		}
		known = knownPart[2 : len(knownPart)-1]
	} else {
		known = knownPart
		position = 0
	}

	// Generate passwords from merged info
	for _, base := range mergedInfo {
		if len(passwords) >= desiredLength {
			break
		}
		addPasswordWithKnown(base, desiredLength, passwdLengthMin, passwdLengthMax, known, position, passwords, bar)
	}

	// If not enough passwords are generated, fill up the remaining space
	for len(passwords) < desiredLength {
		base := mergedInfo[rand.Intn(len(mergedInfo))]
		password := generatePassword(base, passwdLengthMin, passwdLengthMax, known, position)
		passwords[password] = struct{}{}
		bar.Add(1)
	}

	result := make([]string, 0, len(passwords))
	for password := range passwords {
		result = append(result, password)
	}

	return result
}

func mergeUserInfo(targetInfo []string) []string {
	var merged []string

	// Combine different fields together
	for i := 0; i < len(targetInfo); i++ {
		for j := i + 1; j < len(targetInfo); j++ {
			merged = append(merged, targetInfo[i]+targetInfo[j]) // Example: john1990
			merged = append(merged, targetInfo[j]+targetInfo[i]) // Example: 1990john
		}
	}

	// Add single fields too
	merged = append(merged, targetInfo...)

	return merged
}

// Randomly add special symbols with lower probability
func addRandomSymbols(base string) string {
	const specialProbability = 0.01 // 1% chance to add a symbol

	var result strings.Builder
	for _, char := range base {
		result.WriteRune(char)

		// Randomly decide to add a symbol
		if rand.Float64() < specialProbability {
			symbol := Symbols[rand.Intn(len(Symbols))]
			result.WriteByte(symbol)
		}
	}
	return result.String()
}

func addPasswordWithKnown(base string, desiredLength int, minLength int, maxLength int, known string, position int, passwords map[string]struct{}, bar *progressbar.ProgressBar) {
	for i := 0; i < 3 && len(passwords) < desiredLength; i++ {
		password := generatePassword(base, minLength, maxLength, known, position)
		passwords[password] = struct{}{}
		bar.Add(1)
	}
}

func generatePassword(base string, minLength int, maxLength int, known string, position int) string {
	// Adjust length to ensure there's enough space for the known part
	length := rand.Intn(maxLength-minLength+1) + minLength
	if position < 1 {
		position = 1
	}
	if position+len(known)-1 > length {
		length = position + len(known) - 1
	}

	var result strings.Builder

	// Create the prefix before the known part
	prefixLength := position - 1
	randomPrefix := addRandomSymbols(base)

	if prefixLength > len(randomPrefix) {
		prefixLength = len(randomPrefix)
	}
	result.WriteString(randomPrefix[:prefixLength]) // Add prefix

	// Insert known part at the correct position
	result.WriteString(known)

	// Fill the rest of the password with random characters
	suffixLength := length - result.Len()
	if suffixLength > 0 {
		randomSuffix := addRandomSymbols(base)
		if suffixLength > len(randomSuffix) {
			suffixLength = len(randomSuffix)
		}
		result.WriteString(randomSuffix[:suffixLength]) // Add suffix
	}

	// If the result exceeds the required length, slice it
	if result.Len() > length {
		return result.String()[:length]
	}
	return result.String()
}

/* Filter */

func filterFromWordlist(targetInfo []string, desiredLength int, minPaswdLength int, maxPaswdLength int, wordlist string) []string {

    passwords := make(map[string]struct{})

    file, err := os.Open(wordlist)
    if err != nil {
        fmt.Println("Error opening wordlist: ", err)
        return nil
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    bar := progressbar.Default(-1)

    for scanner.Scan() {
        word := scanner.Text()

        if len(word) >= minPaswdLength && len(word) <= maxPaswdLength {
            for _, target := range targetInfo {
                if strings.Contains(word, target) {
                    passwords[word] = struct{}{}
                    break
                }
            }
        }
        
        bar.Add(1)
    }

    if err != nil {
        fmt.Println("Error reading wordlist: ", err)
        return nil
    }

    result := make([]string, 0, len(passwords))
    for password := range passwords {
        result = append(result, password)
    }

    if len(result) > desiredLength {
        result = result[:desiredLength]
    }

    return result

}
