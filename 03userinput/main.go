package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcom := "welcom to user input"
	fmt.Println(welcom) //bufio,os pkg.go.dev

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for pizza:")

	//comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thankgs for rating ,", input)
	fmt.Printf("type of this rating %T ", input)
}
