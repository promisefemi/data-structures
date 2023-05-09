package main

import (
	"fmt"
	"log"
	"os"
)

type linkedList struct {
	// value can be whatever datatype fits your particular scenario
	value string
	next  *linkedList
}

func (node *linkedList) insert(text string) {
	// first we create a temporary value for our new node
	tempNode := &linkedList{value: text, next: nil}
	// remember that the next pointer in the last node is always nil until it points to another node
	if node == nil {
		// if this is a head node and it is nill, make the current tempNode the head
		node = tempNode
	} else {
		currentNode := node
		// go through every node until we get to the last node,
		for currentNode.next != nil {
			// check if it is the last node, if not move to the next
			currentNode = currentNode.next
		}
		// when we get to the last node, place the new node we are creating
		currentNode.next = tempNode
	}
}

func (node *linkedList) print() {
	// start from the first node
	currentNode := node
	// loop through each node, making sure it is not nil
	for currentNode != nil {
		//print the value out
		fmt.Printf("%s -> ", currentNode.value)
		//move to the next node
		currentNode = currentNode.next
	}
	fmt.Println("")
}

func (node *linkedList) deleteFirstNode() *linkedList {
	node = node.next
	return node
}

func (node *linkedList) deleteLastNode() {
	currentNode := node
	nextNode := node.next
	for nextNode.next != nil {
		nextNode = nextNode.next
		currentNode = currentNode.next
	}
	currentNode.next = nil
	/*
		A different way of solving the problem
		p := node
		for p.next.next != nil {
			p = p.next
		}
		p.next = nil*/
}

func (node *linkedList) insertNodeAtSpecificPosition(value string, position int) {
	currentNode := node
	fmt.Println(value, position)
	newNode := &linkedList{value: value}
	for i := 0; i < position; i++ {
		currentNode = currentNode.next
	}
	newNode.next = currentNode.next
	currentNode.next = newNode
}
func (node *linkedList) insertAtBeginning(value string) *linkedList {
	newNode := &linkedList{value: value}
	tempNode := node
	newNode.next = tempNode
	return newNode
}

func (node *linkedList) deleteSpecificNode(index int) {
	currentNode := node
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
}

func main() {
	list := new(linkedList)

	for {
		var selected string
		fmt.Println("(1) enter a new item")
		fmt.Println("(2) print list")
		fmt.Println("(3) delete from head/beginning")
		fmt.Println("(4) delete from last")
		fmt.Println("(5) delete from position")
		fmt.Println("(6) insert into position")
		fmt.Println("(7) insert at head")
		fmt.Println("(8) exit")
		_, err := fmt.Scanln(&selected)
		if err != nil {
			log.Fatalln(err)
		}

		switch selected {
		case "1":
			fmt.Println("Enter new value")
			var item string
			_, err := fmt.Scanln(&item)
			if err != nil {
				log.Fatalln(err)
			}
			list.insert(item)
		case "2":
			list.print()
		case "3":
			list = list.deleteFirstNode()
		case "4":
			list.deleteLastNode()
		case "5":
			fmt.Println("Delete from a position")
			var i int
			_, err = fmt.Scanln(&i)
			if err != nil {
				log.Fatalln(err)
			}
			list.deleteSpecificNode(i)

		case "6":
			fmt.Println("Enter new value")
			var item string
			_, err := fmt.Scanln(&item)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Enter new item position")
			var i int
			_, err = fmt.Scanln(&i)
			if err != nil {
				log.Fatalln(err)
			}

			list.insertNodeAtSpecificPosition(item, i)
		case "7":
			fmt.Println("Enter new value")
			var item string
			_, err := fmt.Scanln(&item)
			if err != nil {
				log.Fatalln(err)
			}
			list = list.insertAtBeginning(item)
		default:
			os.Exit(0)

		}
	}
}
