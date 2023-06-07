package main

import (
    "fmt"
)

func client(server chan<- string, client <-chan string) {
    server <- "Hello from client!"
    fmt.Println(<-client)
    for i := 0; i < 3; i++ {
        server <- "Send me some data please :)"
        fmt.Println(<-client)
    }
    server <- "Thank you, goodbye from client!"
    fmt.Println(<-client)
}

func server(client chan<- string, server <-chan string) {
    fmt.Println(<-server)
    client <- "Hello from server!"
    for i := 0; i < 3; i++ {
        fmt.Println(<-server)
        client <- "Here's the data"
    }
    fmt.Println(<-server)
    client <- "Goodbye from server!"
}

func main() {
    clientToServer := make(chan string)
    serverToClient := make(chan string)

    go client(clientToServer, serverToClient)
    go server(clientToServer, serverToClient)

    // Wait for the processes to finish
    var input string
    fmt.Scanln(&input)
}