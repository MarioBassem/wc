package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	output, err := Run()
	if err != nil {
		log.Fatalf("wc: error: %s", err.Error())
	}

	fmt.Print(output)
}

func Run() (string, error) {
	args := os.Args[1:]
	if len(args) == 1 {
		f, err := os.Open(args[0])
		if err != nil {
			return "", fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		wordCount, err := countWords(f)
		if err != nil {
			return "", fmt.Errorf("failed to count file words: %w", err)
		}

		if _, err := f.Seek(0, 0); err != nil {
			return "", fmt.Errorf("failed to seek file: %w", err)
		}

		byteCount, err := countBytes(f)
		if err != nil {
			return "", fmt.Errorf("failed to count file bytes: %w", err)
		}

		if _, err := f.Seek(0, 0); err != nil {
			return "", fmt.Errorf("failed to seek file: %w", err)
		}

		lineCount, err := countLines(f)
		if err != nil {
			return "", fmt.Errorf("failed to count file lines: %w", err)
		}

		return fmt.Sprintf("%d %d %d %s\n", lineCount, wordCount, byteCount, args[0]), nil
	}

	if len(args) == 2 {
		// flag and file are provided
		f, err := os.Open(args[1])
		if err != nil {
			return "", fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		switch args[0] {
		case "-c":
			byteCount, err := countBytes(f)
			if err != nil {
				return "", fmt.Errorf("failed to count bytes: %w", err)
			}
			return fmt.Sprintf("%d %s\n", byteCount, args[1]), nil
		case "-l":
			// count the number of lines in file
			lineCount, err := countLines(f)
			if err != nil {
				return "", fmt.Errorf("failed to count lines: %w", err)
			}
			return fmt.Sprintf("%d %s\n", lineCount, args[1]), nil
		case "-w":
			// count the number of words in file
			count, err := countWords(f)
			if err != nil {
				return "", fmt.Errorf("failed to count words: %w", err)
			}
			return fmt.Sprintf("%d %s\n", count, args[1]), nil
		case "-m":
			// count the number of characters in file
			count, err := countChars(f)
			if err != nil {
				return "", fmt.Errorf("faeild to count lines: %w", err)
			}
			return fmt.Sprintf("%d %s\n", count, args[1]), nil
		default:
			return "", fmt.Errorf(`available options are:
									-c		print the byte count
									-l		print the line count
									-m		print the characters count
									-w		print the word count`)
		}
	}

	return "", fmt.Errorf("Usage: wc [OPTION] FILE")

}
