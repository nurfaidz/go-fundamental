package file

import (
	"fmt"
	"os"
)

func Read() {
	data, err := os.ReadFile("./file/output.txt")
	if err != nil {
		fmt.Println("Error Reading File:", err)
	}
	defer fmt.Println(string(data))
}
