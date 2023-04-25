package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func takeInput(inputText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputText)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	text = strings.TrimSuffix(text, "\n")
	return text
}

func takeInputText(inputText string) string {
	fmt.Print(inputText)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()

	return choice
}
