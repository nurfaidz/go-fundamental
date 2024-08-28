package write

import (
	"fmt"
	"os"
)

func Write() {
	file, err := os.Create("./file/output.txt")
	if err != nil {
		fmt.Println("Error Creating File:", err)
	}
	defer file.Close()

	data := []byte("Hello, Golang!")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error Writing to file:", err)
	}
}
