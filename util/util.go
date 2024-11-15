package util

import (
	"bufio"
	"log"
	"os"
)

func CreateScannerFromFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Expected %v to be in the current folder", filename)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}
