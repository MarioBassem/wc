package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

const BufferSize = 16 * 1024

func countBytes(r io.Reader) (uint32, error) {
	byteCount := uint32(0)
	for {
		buff := make([]byte, BufferSize)
		n, err := r.Read(buff)
		if errors.Is(err, io.EOF) {
			byteCount += uint32(n)
			break
		}
		if err != nil {
			return 0, fmt.Errorf("failed to read from reader: %w", err)
		}

		byteCount += uint32(n)
	}

	return byteCount, nil
}

func countLines(r io.Reader) (uint32, error) {
	lineCount := uint32(0)
	for {
		buff := make([]byte, BufferSize)
		n, err := r.Read(buff)
		if errors.Is(err, io.EOF) {
			lineCount += getLinesFromBuffer(buff[:n])
			break
		}
		if err != nil {
			return 0, fmt.Errorf("failed to read from reader: %w", err)
		}

		lineCount += getLinesFromBuffer(buff[:n])
	}

	return lineCount, nil
}

func getLinesFromBuffer(buff []byte) uint32 {
	lines := uint32(0)
	for _, c := range buff {
		if c == '\n' {
			lines += 1
		}
	}

	return lines
}

func countWords(r io.Reader) (uint32, error) {
	wordCount := uint32(0)
	for {
		buff := make([]byte, BufferSize)
		n, err := r.Read(buff)
		if errors.Is(err, io.EOF) {
			wordCount += getWordsFromBuffer(buff[:n])
			break
		}
		if err != nil {
			return 0, fmt.Errorf("failed to read from buffer: %w", err)
		}

		wordCount += getWordsFromBuffer(buff[:n])
	}

	return wordCount, nil
}

func getWordsFromBuffer(buff []byte) uint32 {
	words := uint32(0)
	spliteBytes := bytes.FieldsFunc(buff, func(r rune) bool {
		return r == ' ' || r == '\n' || r == '\t'
	})

	for _, w := range spliteBytes {
		if len(bytes.Trim(w, "*-• \n\t")) > 0 {
			words += 1
		}
	}

	return words
}

func countChars(r io.Reader) (uint32, error) {
	charCount := uint32(0)
	for {
		buff := make([]byte, BufferSize)
		n, err := r.Read(buff)
		if errors.Is(err, io.EOF) {
			charCount += uint32(utf8.RuneCount(buff[:n]))
			break
		}
		if err != nil {
			return 0, fmt.Errorf("failed to read from buffer: %w", err)
		}

		charCount += uint32(utf8.RuneCount(buff[:n]))
	}

	return charCount, nil
}
