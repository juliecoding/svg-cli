package cli

import (
	"bufio"
	"fmt"
	"os"
)

func Run() int {
	var a = app{
		filters: initFilters(),
		config:  getConfig(),
	}
	return a.op()
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
