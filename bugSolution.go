package main

import (

        "fmt"
        "os"
)

func main() {
        data, err := os.ReadFile("my_file.txt")
        if err != nil {
                fmt.Println("Error reading file:", err)
                return
        }
        fmt.Println("File content:", string(data))
}