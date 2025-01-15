package main

import (
    "fmt"
    "os"
)

func main() {
    // ファイルを作成する
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    } else {
        fmt.Println("File created successfully.")
    }
    // 関数が終了した際に確実にファイルを閉じる
    defer file.Close()

    // ファイルに書き込む
    _, err = file.WriteString("Hello, World!")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    // errの中身を表示
    fmt.Println("The value of err is:", err)
	

    fmt.Println("File created and written successfully.")
}