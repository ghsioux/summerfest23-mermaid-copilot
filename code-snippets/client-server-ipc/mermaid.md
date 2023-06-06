Prompt: packagpackage main

import (
    "fmt"
    "time"
)

func main() {
    msgChan := make(chan string)

    go client(msgChan)
    go server(msgChan)

    time.Sleep(2 * time.Second)
}

func client(msgChan chan<- string) {
    fmt.Println("Client: Hello from client!")
    msgChan <- "Hello from client!"

    response := <-msgChan
    fmt.Println("Client received:", response)

    for i := 0; i < 3; i++ {
        fmt.Println("Client: Send me some data please :)")
        msgChan <- "Send me some data please :)"

        response := <-msgChan
        fmt.Println("Client received:", response)
    }

    fmt.Println("Client: Thank you, goodbye from client!")
    msgChan <- "Thank you, goodbye from client!"

    response := <-msgChan
    fmt.Println("Client received:", response)
}

func server(msgChan chan string) {
    response := <-msgChan
    fmt.Println("Server received:", response)

    fmt.Println("Server: Hello from server!")
    msgChan <- "Hello from server!"

    for i := 0; i < 3; i++ {
        response := <-msgChan
        fmt.Println("Server received:", response)

        fmt.Println("Server: Here's the data")
        msgChan <- "Here's the data"
    }

    response = <-msgChan
    fmt.Println("Server received:", response)

    fmt.Println("Server: Goodbye from server!")
    msgChan <- "Goodbye from server!"
}