package main

import (
	"fmt"
)

type link struct {
	value string
	next  *link
}

func (l *link) insert(text string) {
	temp := &link{}
	temp.value = text
	temp.next = nil

	if l == nil {
		l = temp
	} else {
		currentNode := l
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = temp
	}
}
func (l *link) print() {
	var currentNode = &link{}
	currentNode = l.next
	for currentNode != nil {
		fmt.Printf("%s ->", currentNode.value)
		currentNode = currentNode.next
	}
}

func main() {
	head := new(link)

	head.insert("Monday")

	head.insert("Wednesday")
	head.insert("Tuesday")
	head.insert("Thursday")
	head.insert("Friday")

	head.print()
}
