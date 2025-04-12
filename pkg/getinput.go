package pkg

import (
	"bufio"
	"os"
)

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()

	return scanner.Text()
}
