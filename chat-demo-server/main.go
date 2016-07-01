package main

import (
    "bufio"
    "flag"
    "fmt"
    "net"
    "os"
)

// flags
var port = flag.String("port", "1337", "Set the port to listen on")

// listeners
var onJoin = make(chan *Client)
var onMessage = make(chan string)

// queues
var clients = make([]*Client, 0)

// locks
var clientsMutex = NewMutex("clients")

func addClient(c *Client) {
    clientsMutex.Lock()

    // do some work
    clients = append(clients, c)

    clientsMutex.Unlock()

    c.Tell("welcome to the party!")

    fmt.Printf("accepted new client: %v\n", c.Name)
    broadcast(fmt.Sprintf("new user: %v", c.Name))
}

func dropClient(c *Client) {
    fmt.Printf("dropping client %v\n", c.RemoteAddr())

    for i, cl := range clients {
        if c == cl {
            clients = append(clients[:i], clients[i+1:]...)

            c.Close()
            return
        }
    }

    panic(fmt.Sprintf("failed to remove client %v", c.Name))
}

func handleClient(c *Client) {
    reader := bufio.NewReader(c)

    for {
        // wait to read from buffer
        msg, err := reader.ReadString('\n')

        if err != nil {
            dropClient(c)
            return
        }

        // let everyone know we got a message
        onMessage <- fmt.Sprintf("message from %v: %s", c.Name, msg)
    }
}

func broadcast(str string) {
    for _, c := range clients {
        c.Tell(str)
    }
}

func main() {
    listener, err := net.Listen("tcp", ":"+*port)

    if err != nil {
        panic(fmt.Sprintf("couldn't start server on port %s", *port))
    }

    defer func() {
        listener.Close()
    }()

    go func() {
        for {
            select {
            case c, ok := <-onJoin:
                // handle someone joining

                if !ok {
                    break
                }

                addClient(c)
            case msg, ok := <-onMessage:
                // handle a new message
                fmt.Fprint(os.Stdout, msg)

                if !ok {
                    break
                }

                broadcast(msg)
            }
        }
    }()

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("problem accepting client %v", err)
        }

        c := &Client{Conn: conn, Name: conn.RemoteAddr().String()}
        onJoin <- c

        go handleClient(c)
    }
}
