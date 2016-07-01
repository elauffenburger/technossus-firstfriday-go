package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "net"
    "os"
)

var port = flag.String("port", "1337", "specifies the port to connect on")

func main() {
    conn, err := net.Dial("tcp", ":"+*port)

    if err != nil {
        panic("couldn't connect!")
    }

    go func(conn net.Conn) {
        reader := bufio.NewReader(os.Stdin)
        writer := bufio.NewWriter(conn)

        for {
            msg, err := reader.ReadString('\n')
            if err != nil {
                panic("couldn't understand input")
            }

            writer.WriteString(msg)
            writer.Flush()
        }
    }(conn)

    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')

        if err != nil {
            panic(fmt.Sprintf("error reading from server: %v", err))
        }

        io.WriteString(os.Stdout, "received: "+msg)
    }
}
