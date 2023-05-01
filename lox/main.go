package lox

import (
	"Golox/lox/Scanner"
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	hadError        bool
	hadRuntimeError bool
)

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	run(string(bytes))
	if hadError {
		os.Exit(65)
	}
	if hadRuntimeError {
		os.Exit(70)
	}
	return nil
}
func runPrompt() error {
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		print("> ")
		if !inputScanner.Scan() {
			break
		}
		line := inputScanner.Text()
		fmt.Println(run(line))
		hadError = false
	}
	return nil
}
func run(source string) error {
	scanner := Scanner.New(source)
	tokens, err := scanner.ScanTokens()
	if err != nil {
		return fmt.Errorf("faied to scan tokens: %w", err)
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
	return nil
}

func main() {
	var err error

	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
	} else if len(os.Args) == 2 {
		err = runFile(os.Args[1])
		if hadError {
			os.Exit(65)
		} else if hadRuntimeError {
			os.Exit(70)
		}
	} else {
		err = runPrompt()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
