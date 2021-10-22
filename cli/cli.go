package cli

import (
	"bufio"
	"fmt"
	"os"
)

func Run() int {
	// Conventionally, for os.Exit, code zero indicates success,
	// non-zero an error.
	var a = app{
		filters: initFilters(),
		config: getConfig(),
	}
	a.op()
	return 0
}

func getUserInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	scanner.Scan()
	return scanner.Text()
}

func output(explanation string, err error) {
	fmt.Println(explanation)
	if err != nil {
		fmt.Println(err)
	}
}
