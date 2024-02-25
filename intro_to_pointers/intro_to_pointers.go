package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	// fmt.Printf("To: %v\n", &m.Recipient)
	// fmt.Printf("Message: %v\n", &m.Text)
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}
