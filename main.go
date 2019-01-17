package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	byteArrayString := strings.Join(os.Args[1:], " ")

	byteArrayString = strings.TrimSuffix(strings.TrimPrefix(byteArrayString, "["), "]")

	asList := strings.Split(byteArrayString, ", ")

	byteArray := []byte{}

	for _, s := range asList {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error converting '%s' to integer: %s\n", s, err)
			os.Exit(1)
		}
		b := byte(i)
		byteArray = append(byteArray, b)
	}

	fmt.Println(string(byteArray))
}
