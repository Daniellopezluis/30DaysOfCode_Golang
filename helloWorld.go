package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Hello, World.\n")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}