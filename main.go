package main

import "fmt"
import "gostudy"
import (
	"bufio"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World,您好，世界!")
	gostudy.Variable()

	reader := bufio.NewReader(os.Stdin)
	key, _, _ := reader.ReadLine()
	log.Println("command", key)
}
