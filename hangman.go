package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var hmArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0    |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0    |\n" +
		" |    |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0    |\n" +
		"/|    |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0    |\n" +
		"/|\\    |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0    |\n" +
		"/|\\    |\n" +
		"/     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0    |\n" +
		"/|\\    |\n" +
		"/ \\    |\n" +
		"    ===\n",
}
var wordArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER", "ZOMBIE", "ZODIAC", "FLUFF",
}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func getRandWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randWord = wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))
	return randWord

}

func showBoard() {
	fmt.Println(hmArr[len(wrongGuesses)])
	fmt.Println("Secret Word : ")
	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Print("\nIncorrect Guesses : ")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	fmt.Println()

}

func getUserLetter() string {
	reader := bufio.NewReader(os.Stdin)
	var guess string
	for true {
		fmt.Print("\n Guess a Letter")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)
		var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
		if len(guess) != 1 {
			fmt.Println("Please enter only one Letter")
		} else if !IsLetter(guess) {
			fmt.Println("Please enter a Letter")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please enter you havent guessed ")
		} else {
			return guess
		}
	}
	return guess
}

func getAllIndexes(theStr, subStr string) (indices []int) {
	if len(subStr) == 0 || (len(theStr) == 0) {
		return indices
	}
	offset := 0
	for {
		i := strings.Index(theStr[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func updateCorrectletters(letter string) {
	indexMatches := getAllIndexes(randWord, letter)
	for _, v := range indexMatches {
		correctLetters[v] = letter
	}
}
func sliceHasEmptys(theSlice []string) bool {
	for _, v := range theSlice {
		if len(v) == 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(getRandWord())

	for true {
		showBoard()
		guess := getUserLetter()
		if strings.Contains(randWord, guess) {
			updateCorrectletters(guess)
			if sliceHasEmptys(correctLetters) {
				fmt.Println("more letters to guess")
			} else {
				fmt.Println("yess the sec word is : ", randWord)
				break
			}

		} else {
			guessedLetters += guess
			wrongGuesses = append(wrongGuesses, guess)

			if len(wrongGuesses) >= 6 {
				fmt.Println("Sorry you're dead the word was  ", randWord)
				break
			}
		}
	}

}
