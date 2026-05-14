package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	rand "math/rand/v2"
	"net/url"
	"os"

	"github.com/ergochat/readline"
)

type Word struct {
	Source string
	Target string
}

func (w *Word) UnmarshalJSON(data []byte) error {
	var arr [2]string
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	w.Source = arr[0]
	w.Target = arr[1]
	return nil
}

const (
	cursorUp       = "\x1b[A"
	deleteLine     = "\x1b[K"
	red            = "\x1b[31m"
	green          = "\x1b[32m"
	gray           = "\x1b[90m"
	reset          = "\x1b[0m"
	deleteLastLine = cursorUp + deleteLine
	linkPrefix     = "\x1b]8;;"
	linkSuffix     = "\a"
	endLink        = linkPrefix + linkSuffix

	modeReverse = "reverse"
)

func getWords(filename string) ([]Word, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var words []Word
	err = json.Unmarshal(data, &words)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func testVocab(words []Word, mode string) (int, error) {
	readline, err := readline.New("")
	if err != nil {
		return 0, err
	}
	return testVocabReader(readline, words, mode)
}

func testVocabReader(readline *readline.Instance, words []Word, mode string) (int, error) {
	correctCount := 0
	incorrect := make([]Word, 0, len(words))
	for len(words) != 0 {
		i := rand.IntN(len(words))
		word := words[i]
		correct, err := testWord(readline, word, mode)
		if err != nil {
			return 0, err
		}
		if correct {
			correctCount++
		} else {
			incorrect = append(incorrect, word)
		}
		words[i], words[len(words)-1] = words[len(words)-1], words[i]
		words = words[:len(words)-1]
	}
	if len(incorrect) != 0 {
		ones := "ones"
		if len(incorrect) == 1 {
			ones = "one"
		}
		fmt.Printf("\nLet's try the %s you missed again.\n", ones)
		_, err := testVocabReader(readline, incorrect, mode)
		if err != nil {
			return 0, err
		}
	}
	return correctCount, nil
}

func testWord(readline *readline.Instance, word Word, mode string) (bool, error) {
	source := word.Source
	target := word.Target
	if mode == modeReverse {
		source = word.Target
		target = word.Source
	}
	fmt.Printf("What is \"%s\"?\n> ", source)
	answer, err := readline.ReadLine()
	if err != nil {
		return false, err
	}
	correct := answer == target
	color := red
	if correct {
		color = green
	}
	wiktionaryURL := fmt.Sprintf("%shttps://en.wiktionary.org/wiki/Special:Search?search=%s%s", linkPrefix, url.QueryEscape(target), linkSuffix)
	fmt.Printf("%s%s%sWhat is \"%s\"? %s%s%s%s\n", deleteLastLine, deleteLastLine,
		color, source, wiktionaryURL, answer, endLink, reset)
	if !correct {
		fmt.Printf("  %sCorrect: %s%s%s%s\n", gray, wiktionaryURL, target, endLink, reset)
	}
	return correct, nil
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("please specify a file with the vocabulary in it")
	}
	filename := os.Args[1]
	mode := ""
	if len(os.Args) >= 3 {
		mode = os.Args[2]
	}
	words, err := getWords(filename)
	if err != nil {
		return err
	}
	correct, err := testVocab(words, mode)
	if err != nil {
		return err
	}
	fmt.Printf("\nYou got %d/%d (%d%%) correct on your first go!\n", correct,
		len(words), int(math.Round((float64(correct)/float64(len(words)))*100)))
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
