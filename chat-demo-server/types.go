package main

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	net.Conn
	Name string
}

func (c *Client) Tell(msg string) {
	runes := []rune(msg)
	newline := rune('\n')

	if runes[len(runes)-1] != newline {
		runes = append(runes, newline)
	}

	msg = string(runes)
	writer := bufio.NewWriter(c)
	writer.WriteString(msg)
	writer.Flush()
}

type Mutex struct {
	channel chan interface{}
	locked  bool
	name    string
}

func (s *Mutex) Lock() {
	if s.locked {
		fmt.Printf("warning: waiting for lock on mutex %s", s.name)
	}

	s.channel <- 1
	s.locked = true
}

func (s *Mutex) Unlock() {
	if !s.locked {
		panic("tried to unlock an unlocked semaphore!")
	}

	<-s.channel
	s.locked = false
}

func NewMutex(name string) *Mutex {
	return &Mutex{name: name, channel: make(chan interface{}, 1), locked: false}
}
