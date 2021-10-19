package cli

import (
	"bufio"
	"fmt"
	"os"
)

func Run(args []string) int {
	fmt.Println(args)
	// Conventionally, for os.Exit, code zero indicates success,
	// non-zero an error.
	var a app
	a.op()
	return 0
}

func getUserInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	scanner.Scan()
	return scanner.Text()
}

func outputError(explanation string, err error) {
	fmt.Println(explanation)
	if err != nil {
		fmt.Println(err)
	}
}
