package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShouldOverrideFile(msg string) bool {
	fmt.Printf("%s (y/n): \n", msg)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return input == "y" || input == "yes"
}
