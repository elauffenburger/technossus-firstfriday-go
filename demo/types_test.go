package main

import (
    "bufio"
    "github.com/elauffenburger/technossus-firstfriday-go/demo/types"
    "net"
    "testing"
)

func TestInterfaces(t *testing.T) {
    name := "Bob"

    person := &types.Person{Name: name}
    takesHasName := func(hn types.HasName) string {
        return hn.GetName()
    }

    result := takesHasName(person)

    if result != name {
        t.Errorf("Expected result to be '%s', got '%s'", name, result)
    }
}

func TestTypeSwitch(t *testing.T) {
    var person interface{} = &types.Person{}
    var num interface{} = int(1)
    var num32 interface{} = int32(1)

    if _, ok := person.(*types.Person); !ok {
        t.Error("Expected type assertion to succeed for *types.Person -> *types.Person")
    }

    if _, ok := person.(types.Person); ok {
        t.Error("Expected type assertion to fail for *types.Person -> types.Person")
    }

    if _, ok := num.(int); !ok {
        t.Error("Expected type assertion to succeed to int -> int")
    }

    if _, ok := num32.(int); ok {
        t.Error("Expected type assertion to fail for int32 -> int")
    }
}

func TestNetAndAdHocPolymorphismBecauseWhyNot(t *testing.T) {
    port := "1543"
    message := "i'm a string!"

    type stringReader interface {
        ReadString(delim byte) (string, error)
    }

    readString := func(reader stringReader) {
        msg, _ := reader.ReadString(0)

        if msg != message {
            t.Errorf("Expected msg to be '%s', received '%s'", message, msg)
            return
        }
    }

    go func() {
        // just to make this as insane as possible
        listener, _ := net.Listen("tcp", ":"+port)

        for {
            connection, err := listener.Accept()

            if err != nil {
                break
            }

            go func(conn net.Conn) {
                defer func() {
                    conn.Close()
                }()

                writer := bufio.NewWriter(conn)

                writer.WriteString(message)
                writer.Flush()
            }(connection)
        }
    }()

    conn, _ := net.Dial("tcp", "localhost:"+port)
    reader := bufio.NewReader(conn)

    readString(reader)
}

type Doer struct{}

func (d *Doer) DoSomething() bool {
    return true
}

func TestJustAdHocPolymorphism(t *testing.T) {
    type doesStuff interface {
        DoSomething() bool
    }

    var obj interface{} = &Doer{}

    if _, ok := obj.(doesStuff); !ok {
        t.Errorf("Expected *Doer -> doesStuff assertion to be to succeed")
    }
}
