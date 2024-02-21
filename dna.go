package dna

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Compress(source string) (compressed string, err error) {

	if len(source) == 0 {
		return "", fmt.Errorf("source string is empty")
	}

	if len(source) == 1 {
		return source + "1", nil
	}

	var prevRune rune
	var runeCounter int = 1
	var runeIndex = 0

	builder := strings.Builder{}
	for _, currentRune := range source {

		if runeIndex == 0 {
			prevRune = currentRune
			runeIndex++
			continue
		}

		if prevRune == currentRune {
			runeCounter++
		} else {
			_, err = builder.WriteString(string(prevRune))
			if err != nil {
				return "", err
			}
			_, err = builder.WriteString(strconv.Itoa(runeCounter))
			if err != nil {
				return "", err
			}
			runeCounter = 1
		}

		prevRune = currentRune
		runeIndex++
	}

	if runeCounter > 1 {
		_, err = builder.WriteString(string(prevRune))
		if err != nil {
			return "", err
		}
		_, err = builder.WriteString(strconv.Itoa(runeCounter))
		if err != nil {
			return "", err
		}

	}

	return builder.String(), nil
}

func Decompress(compressed string) (source string, err error) {

	if utf8.RuneCountInString(compressed)%2 != 0 {
		return "", fmt.Errorf("error is not properly compressed, expected length should be even")
	}

	sb := strings.Builder{}
	var currentRune rune
	var runeIndex = 0
	for _, compressedRune := range compressed {

		if runeIndex%2 == 0 {
			currentRune = compressedRune
		} else {

			currentRunCount, err := strconv.Atoi(string(compressedRune))
			if err != nil {
				return "", fmt.Errorf("unable to parse current rune count")
			}
			for i := 0; i < currentRunCount; i++ {

				_, err = sb.WriteRune(currentRune)
				if err != nil {
					return "", fmt.Errorf("unable to write rune")
				}

			}
		}

		runeIndex++

	}

	return sb.String(), nil
}
